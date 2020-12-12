package apiimpl

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/rs/cors"

	"github.com/Smart-Purveyance-Tracker/backend/api/rest-swagger/models"
	"github.com/Smart-Purveyance-Tracker/backend/api/rest-swagger/restapi/operations"
	"github.com/Smart-Purveyance-Tracker/backend/entity"
	"github.com/Smart-Purveyance-Tracker/backend/repository"
	"github.com/Smart-Purveyance-Tracker/backend/service"
	"github.com/Smart-Purveyance-Tracker/backend/service/auth"
)

type Server struct {
	userSvc    service.User
	productSvc service.Product
	authSvc    auth.Service
}

func NewServer(userSvc service.User, authSvc auth.Service, productSvc service.Product) *Server {
	return &Server{
		userSvc:    userSvc,
		authSvc:    authSvc,
		productSvc: productSvc,
	}
}

func (s *Server) login(params operations.LoginParams) middleware.Responder {
	user, err := s.userSvc.Login(*params.UserInfo.Email, *params.UserInfo.Password)
	if err != nil {
		return operations.NewLoginDefault(http.StatusInternalServerError).WithPayload(newAPIErr(err.Error()))
	}
	token, err := s.authSvc.GenerateToken(user.ID)
	if err != nil {
		return operations.NewLoginDefault(http.StatusInternalServerError).WithPayload(newAPIErr(err.Error()))
	}

	return operations.NewLoginOK().WithPayload(&models.User{
		ID:    user.ID,
		Email: user.Email,
	}).WithAuthenthication("Bearer " + token)
}

func (s *Server) signup(params operations.SignupParams) middleware.Responder {
	user, err := s.userSvc.Create(entity.User{
		Email:    *params.UserInfo.Email,
		Password: *params.UserInfo.Password,
	})
	if err != nil {
		return operations.NewSignupDefault(http.StatusInternalServerError).WithPayload(newAPIErr(err.Error()))
	}

	token, err := s.authSvc.GenerateToken(user.ID)
	if err != nil {
		return operations.NewSignupDefault(http.StatusInternalServerError).WithPayload(newAPIErr(err.Error()))
	}

	return operations.NewSignupOK().WithPayload(&models.User{
		ID:    user.ID,
		Email: user.Email,
	}).WithAuthenthication("Bearer " + token)
}

// todo: add check for user id
func (s *Server) getProduct(params operations.GetProductParams) middleware.Responder {
	product, err := s.productSvc.ByID(params.ProductID)
	if err != nil {
		return operations.NewGetProductDefault(http.StatusInternalServerError).WithPayload(newAPIErr(err.Error()))
	}

	return operations.NewGetProductOK().WithPayload(toModelProduct(product))
}

func (s *Server) updateProduct(params operations.UpdateProductParams, userID string) middleware.Responder {
	params.Product.ID = params.ProductID
	product, err := s.productSvc.Update(toEntityProduct(params.Product, userID))
	if err != nil {
		return operations.NewUpdateProductDefault(http.StatusInternalServerError).WithPayload(newAPIErr(err.Error()))
	}
	return operations.NewCreateProductOK().WithPayload(toModelProduct(product))
}

func (s *Server) getProductList(params operations.ProductListParams, userID string) middleware.Responder {
	products, err := s.productSvc.List(repository.ProductListArgs{
		UserID: &userID,
		Date:   (*time.Time)(params.Date),
	})
	if err != nil {
		return operations.NewProductListDefault(http.StatusInternalServerError).WithPayload(newAPIErr(err.Error()))
	}
	return operations.NewProductListOK().WithPayload(toModelProducts(products))
}

func (s *Server) createProduct(params operations.CreateProductParams, userID string) middleware.Responder {
	product, err := s.productSvc.Create(toEntityProduct(params.Product, userID))
	if err != nil {
		return operations.NewCreateProductDefault(http.StatusInternalServerError).WithPayload(newAPIErr(err.Error()))
	}
	return operations.NewCreateProductOK().WithPayload(toModelProduct(product))
}

func newAPIErr(msg string) *models.Error {
	return &models.Error{
		Message: &msg,
	}
}

func ConfigureAPI(api *operations.SwaggerAPI, impl *Server) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.BearerAuth = func(s string) (interface{}, error) {
		ss := strings.Split(s, " ")
		if len(ss) != 2 {
			return nil, fmt.Errorf("can't parse token")
		}
		userID, err := impl.authSvc.UserIDFromToken(ss[1])
		if err != nil {
			return nil, err
		}
		return userID, nil
	}

	api.GetStatusHandler = operations.GetStatusHandlerFunc(func(params operations.GetStatusParams) middleware.Responder {
		return operations.NewGetStatusOK().WithPayload(&operations.GetStatusOKBody{
			Status: "OK",
		})
	})

	api.SignupHandler = operations.SignupHandlerFunc(func(params operations.SignupParams) middleware.Responder {
		return impl.signup(params)
	})

	api.LoginHandler = operations.LoginHandlerFunc(func(params operations.LoginParams) middleware.Responder {
		return impl.login(params)
	})

	api.ScanProductsHandler = operations.ScanProductsHandlerFunc(func(params operations.ScanProductsParams, _ interface{}) middleware.Responder {
		mForm := params.HTTPRequest.MultipartForm
		for name, files := range mForm.File {
			log.Print(name)
			for _, f := range files {
				log.Print(f.Header)
				log.Print(f.Size)
				log.Print(f.Filename)
			}
		}
		boughtAt := time.Now()
		if params.ScanDate != nil {
			boughtAt = time.Time(*params.ScanDate)
		}
		defer params.Image.Close()
		resp, err := impl.productSvc.ScanProducts(service.ScanProductsArgs{
			BoughtAt: boughtAt,
			Image:    params.Image,
		})
		if err != nil {
			return operations.NewScanProductsDefault(http.StatusInternalServerError).WithPayload(newAPIErr(err.Error()))
		}
		return operations.NewScanProductsOK().WithPayload(toScanResponse(resp))
	})

	//api.ScanCheckHandler = operations.ScanCheckHandlerFunc(func(params operations.ScanCheckParams, _ interface{}) middleware.Responder {
	//	boughtAt := time.Now()
	//	if params.ScanDate != nil {
	//		boughtAt = time.Time(*params.ScanDate)
	//	}
	//	counts, err := impl.productSvc.ScanProducts(service.ScanProductsArgs{BoughtAt: boughtAt})
	//	if err != nil {
	//		return operations.NewScanCheckDefault(http.StatusInternalServerError).WithPayload(newAPIErr(err.Error()))
	//	}
	//	return operations.NewScanProductsOK().WithPayload(toModelProductCount(counts))
	//})

	api.GetProductHandler = operations.GetProductHandlerFunc(func(params operations.GetProductParams, _ interface{}) middleware.Responder {
		return impl.getProduct(params)
	})

	api.UpdateProductHandler = operations.UpdateProductHandlerFunc(func(params operations.UpdateProductParams, id interface{}) middleware.Responder {
		return impl.updateProduct(params, id.(string))
	})

	api.CreateProductHandler = operations.CreateProductHandlerFunc(func(params operations.CreateProductParams, id interface{}) middleware.Responder {
		uID := id.(string)
		return impl.createProduct(params, uID)
	})

	api.ProductListHandler = operations.ProductListHandlerFunc(func(params operations.ProductListParams, id interface{}) middleware.Responder {
		uID := id.(string)
		return impl.getProductList(params, uID)
	})

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

func toModelProduct(product entity.Product) *models.Product {
	var t *strfmt.DateTime
	if !product.BoughtAt.IsZero() {
		tmp := strfmt.DateTime(product.BoughtAt)
		t = &tmp
	}
	return &models.Product{
		ID:       product.ID,
		BoughtAt: t,
		InStock:  product.InStock,
		Name:     product.Name,
		Type:     product.Type,
	}
}

func toEntityProduct(product *models.Product, userID string) entity.Product {
	//boughAt := *time.Time(product.BoughtAt)
	var t time.Time
	if product.BoughtAt != nil {
		t = time.Time(*product.BoughtAt)
	}
	return entity.Product{
		ID:       product.ID,
		Name:     product.Name,
		Type:     product.Type,
		BoughtAt: t,
		UserID:   userID,
		InStock:  product.InStock,
	}
}

func toScanResponse(resp service.ProductScanResponse) *models.ScanResponse {
	var res models.ScanResponse
	res.Products = toModelProducts(resp.Products)

	for i := range resp.ProductCounts {
		res.ProductCounts = append(res.ProductCounts, &models.ProductCount{
			Product: &models.Product{
				Type: resp.ProductCounts[i].Product.Type,
			},
			Count: int64(resp.ProductCounts[i].Count),
		})
	}
	return &res
}

func toModelProducts(pp []entity.Product) []*models.Product {
	res := make([]*models.Product, 0, len(pp))
	for _, p := range pp {
		res = append(res, toModelProduct(p))
	}
	return res
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	handleCORS := cors.Default().Handler

	return handleCORS(handler)
}

func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}
