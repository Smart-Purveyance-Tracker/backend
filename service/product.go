package service

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/Smart-Purveyance-Tracker/backend/entity"
	"github.com/Smart-Purveyance-Tracker/backend/repository"
)

type Product interface {
	ByID(id string) (entity.Product, error)
	Create(product entity.Product) (entity.Product, error)
	List(args repository.ProductListArgs) ([]entity.Product, error)
	Update(product entity.Product) (entity.Product, error)
	ScanProducts(args ScanProductsArgs) (ProductScanResponse, error)
	ScanCheck(args ScanProductsArgs) (ProductScanResponse, error)
}

type ProductImpl struct {
	productRepo repository.Product
	productScan ScanAdapter
	checkScan   ScanAdapter
}

func NewProductImpl(product repository.Product, productScan ScanAdapter) *ProductImpl {
	return &ProductImpl{
		productRepo: product,
		productScan: productScan,
	}
}

func (p *ProductImpl) ByID(id string) (entity.Product, error) {
	return p.productRepo.Find(id)
}

func (p *ProductImpl) Create(product entity.Product) (entity.Product, error) {
	return p.productRepo.Insert(product)
}

func (p *ProductImpl) Update(product entity.Product) (entity.Product, error) {
	return p.productRepo.Update(product)
}

func (p *ProductImpl) List(args repository.ProductListArgs) ([]entity.Product, error) {
	return p.productRepo.List(args)
}

type ScanProductsArgs struct {
	BoughtAt time.Time
	Type     string
	Image    io.Reader
}

type ProductCount struct {
	Product entity.Product
	Count   uint64
}

type productResp struct {
	Classification string    `json:"cls"`
	Confidence     float64   `json:"conf"`
	PixelLocation  []float64 `json:"xyxy"`
}

type checkScanResp struct {
	Date     string `json:"date"`
	Products []struct {
		Category string  `json:"category"`
		FullName string  `json:"full_name"`
		Price    float64 `json:"price"`
	} `json:"products"`
	Shop string `json:"shop"`
}

type ProductScanResponse struct {
	Products      []entity.Product
	ProductCounts []ProductCount
}

type ScanAdapter interface {
	Scan(image io.Reader) (ProductScanResponse, error)
}

type ProductScanAdapter struct {
	client *http.Client
	uri    string
}

type CheckScanAdapter struct {
	client *http.Client
	uri    string
}

func NewProductScanAdapter(uri string) *ProductScanAdapter {
	return &ProductScanAdapter{
		uri: uri,
		client: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

func NewCheckScanAdapter(uri string) *CheckScanAdapter {
	return &CheckScanAdapter{
		uri: uri,
		client: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

func newfileUploadRequest(uri string, image io.Reader, paramName string) (*http.Request, error) {
	fileContents, err := ioutil.ReadAll(image)
	if err != nil {
		return nil, err
	}
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, paramName)
	if err != nil {
		return nil, err
	}
	part.Write(fileContents)

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", writer.FormDataContentType())
	return req, nil
}

func (a *ProductScanAdapter) Scan(image io.Reader) (ProductScanResponse, error) {
	req, err := newfileUploadRequest(a.uri+"/process_image", image, "image")
	if err != nil {
		return ProductScanResponse{}, err
	}
	resp, err := a.client.Do(req)
	if err != nil {
		return ProductScanResponse{}, err
	}
	defer resp.Body.Close()
	bb, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ProductScanResponse{}, err
	}
	log.Print(string(bb))
	var products []productResp
	err = json.Unmarshal(bb, &products)
	if err != nil {
		return ProductScanResponse{}, err
	}
	return a.toProducts(products), nil
}

func (a *CheckScanAdapter) Scan(image io.Reader) (ProductScanResponse, error) {
	req, err := newfileUploadRequest(a.uri+"/process_image", image, "image")
	if err != nil {
		return ProductScanResponse{}, err
	}
	resp, err := a.client.Do(req)
	if err != nil {
		return ProductScanResponse{}, err
	}
	defer resp.Body.Close()
	bb, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ProductScanResponse{}, err
	}
	log.Print(string(bb))
	var checkResp checkScanResp
	err = json.Unmarshal(bb, &checkResp)
	if err != nil {
		return ProductScanResponse{}, err
	}
	return a.toProducts(checkResp), nil
}

func (a *CheckScanAdapter) toProducts(resp checkScanResp) ProductScanResponse {
	scanResp := ProductScanResponse{
		Products:      make([]entity.Product, 0),
		ProductCounts: make([]ProductCount, 0),
	}
	productToCount := make(map[string]int)

	for _, p := range resp.Products {
		product := entity.Product{
			Name: p.FullName,
			Type: p.Category,
		}
		scanResp.Products = append(scanResp.Products, product)
		productToCount[p.Category]++
	}
	for product, count := range productToCount {
		scanResp.ProductCounts = append(scanResp.ProductCounts, ProductCount{
			Product: entity.Product{
				Type: product,
			},
			Count: uint64(count),
		})
	}
	return scanResp
}

func (p *ProductImpl) ScanProducts(args ScanProductsArgs) (ProductScanResponse, error) {
	resp, err := p.productScan.Scan(args.Image)
	if err != nil {
		return ProductScanResponse{}, err
	}
	for i := range resp.Products {
		resp.Products[i].BoughtAt = args.BoughtAt
		resp.Products[i], err = p.Create(resp.Products[i])
		if err != nil {
			return ProductScanResponse{}, err
		}
	}
	return resp, nil
}

func (p *ProductImpl) ScanCheck(args ScanProductsArgs) (ProductScanResponse, error) {
	resp, err := p.checkScan.Scan(args.Image)
	if err != nil {
		return ProductScanResponse{}, err
	}
	for i := range resp.Products {
		resp.Products[i].BoughtAt = args.BoughtAt
		resp.Products[i], err = p.Create(resp.Products[i])
		if err != nil {
			return ProductScanResponse{}, err
		}
	}
	return resp, nil
}

func (ProductScanAdapter) toProducts(resp []productResp) ProductScanResponse {
	scanResp := ProductScanResponse{
		Products:      make([]entity.Product, 0, len(resp)),
		ProductCounts: make([]ProductCount, 0, len(resp)),
	}
	productToCount := make(map[string]int)
	for _, p := range resp {
		scanResp.Products = append(scanResp.Products, entity.Product{
			Type: p.Classification,
		})
		productToCount[p.Classification]++
	}
	for product, count := range productToCount {
		scanResp.ProductCounts = append(scanResp.ProductCounts, ProductCount{
			Product: entity.Product{
				Type: product,
			},
			Count: uint64(count),
		})
	}
	return scanResp
}
