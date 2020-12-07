package service

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
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
}

type ProductImpl struct {
	productRepo repository.Product
	productScan ScanAdapter
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

func NewProductScanAdapter(uri string) *ProductScanAdapter {
	return &ProductScanAdapter{
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
	var products []productResp
	err = json.NewDecoder(resp.Body).Decode(&products)
	if err != nil {
		return ProductScanResponse{}, err
	}
	return toProducts(products), nil
}

func (p *ProductImpl) ScanProducts(args ScanProductsArgs) (ProductScanResponse, error) {
	return p.productScan.Scan(args.Image)
}

func toProducts(resp []productResp) ProductScanResponse {
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
