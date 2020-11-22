package apiimpl

import (
	"fmt"
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

type Impl struct {
	userSvc    service.User
	productSvc service.Product
	authSvc    auth.Service
}

func NewImpl(userSvc service.User, authSvc auth.Service, productSvc service.Product) *Impl {
	return &Impl{
		userSvc:    userSvc,
		authSvc:    authSvc,
		productSvc: productSvc,
	}
}

func ConfigureAPI(api *operations.SwaggerAPI, impl *Impl) http.Handler {
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
		user, err := impl.userSvc.Create(entity.User{
			Email:    *params.UserInfo.Email,
			Password: *params.UserInfo.Password,
		})
		if err != nil {
			return operations.NewSignupDefault(http.StatusInternalServerError).WithPayload(&models.Error{
				Message: getStrPtr(err.Error()),
			})
		}

		token, err := impl.authSvc.GenerateToken(user.ID)
		if err != nil {
			return operations.NewSignupDefault(http.StatusInternalServerError).WithPayload(&models.Error{
				Message: getStrPtr(err.Error()),
			})
		}

		return operations.NewSignupOK().WithPayload(&models.User{
			ID:    user.ID,
			Email: user.Email,
		}).WithAuthenthication("Bearer " + token)
	})

	api.LoginHandler = operations.LoginHandlerFunc(func(params operations.LoginParams) middleware.Responder {
		user, err := impl.userSvc.Login(*params.UserInfo.Email, *params.UserInfo.Password)
		if err != nil {
			return operations.NewLoginDefault(http.StatusInternalServerError).WithPayload(&models.Error{
				Message: getStrPtr(err.Error()),
			})
		}
		token, err := impl.authSvc.GenerateToken(user.ID)
		if err != nil {
			return operations.NewLoginDefault(http.StatusInternalServerError).WithPayload(&models.Error{
				Message: getStrPtr(err.Error()),
			})
		}

		return operations.NewLoginOK().WithPayload(&models.User{
			ID:    user.ID,
			Email: user.Email,
		}).WithAuthenthication("Bearer " + token)
	})

	api.ScanProductsHandler = operations.ScanProductsHandlerFunc(func(params operations.ScanProductsParams, _ interface{}) middleware.Responder {
		return operations.NewScanProductsOK().WithPayload([]*models.ProductCount{
			{
				Count: 1,
				Product: &models.Product{
					ID:   "1",
					Name: "ОВОЩ",
				},
			},
		})
	})

	api.ScanCheckHandler = operations.ScanCheckHandlerFunc(func(params operations.ScanCheckParams, _ interface{}) middleware.Responder {
		return operations.NewScanCheckOK().WithPayload([]*models.ProductCount{
			{
				Count: 1,
				Product: &models.Product{
					ID:   "1",
					Name: "ОВОЩ",
				},
			},
		})
	})

	api.GetProductHandler = operations.GetProductHandlerFunc(func(params operations.GetProductParams, _ interface{}) middleware.Responder {
		product, err := impl.productSvc.ByID(params.ProductID)
		if err != nil {
			return operations.NewGetProductDefault(http.StatusInternalServerError).WithPayload(&models.Error{
				Message: getStrPtr(err.Error()),
			})
		}

		return operations.NewGetProductOK().WithPayload(toModelProduct(product))
	})

	api.UpdateProductHandler = operations.UpdateProductHandlerFunc(func(params operations.UpdateProductParams, id interface{}) middleware.Responder {
		uID := id.(string)
		params.Product.ID = params.ProductID
		product, err := impl.productSvc.Update(toEntity(params.Product, uID))
		if err != nil {
			return operations.NewUpdateProductDefault(http.StatusInternalServerError).WithPayload(&models.Error{
				Message: getStrPtr(err.Error()),
			})
		}
		return operations.NewCreateProductOK().WithPayload(toModelProduct(product))
	})

	api.CreateProductHandler = operations.CreateProductHandlerFunc(func(params operations.CreateProductParams, id interface{}) middleware.Responder {
		uID := id.(string)
		product, err := impl.productSvc.Create(toEntity(params.Product, uID))
		if err != nil {
			return operations.NewCreateProductDefault(http.StatusInternalServerError).WithPayload(&models.Error{
				Message: getStrPtr(err.Error()),
			})
		}
		return operations.NewCreateProductOK().WithPayload(toModelProduct(product))
	})

	api.ProductListHandler = operations.ProductListHandlerFunc(func(params operations.ProductListParams, id interface{}) middleware.Responder {
		uID := id.(string)
		products, err := impl.productSvc.List(repository.ProductListArgs{
			UserID: &uID,
			Date:   (*time.Time)(params.Date),
		})
		if err != nil {
			return operations.NewProductListDefault(http.StatusInternalServerError).WithPayload(&models.Error{
				Message: getStrPtr(err.Error()),
			})
		}
		return operations.NewProductListOK().WithPayload(toModelProducts(products))
	})

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

func getStrPtr(s string) *string {
	return &s
}

func toModelProduct(product entity.Product) *models.Product {
	return &models.Product{
		ID:       product.ID,
		BoughtAt: strfmt.DateTime(product.BoughtAt),
		InStock:  product.InStock,
		Name:     product.Name,
		Type:     product.Type,
	}
}

func toEntity(product *models.Product, userID string) entity.Product {
	boughAt := time.Time(product.BoughtAt)
	return entity.Product{
		ID:       product.ID,
		Name:     product.Name,
		Type:     product.Type,
		BoughtAt: boughAt,
		UserID:   userID,
		InStock:  product.InStock,
	}
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
