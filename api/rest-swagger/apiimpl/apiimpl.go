package apiimpl

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/rs/cors"

	"github.com/Smart-Purveyance-Tracker/backend/api/rest-swagger/models"
	"github.com/Smart-Purveyance-Tracker/backend/api/rest-swagger/restapi/operations"
	"github.com/Smart-Purveyance-Tracker/backend/entity"
	"github.com/Smart-Purveyance-Tracker/backend/service"
	"github.com/Smart-Purveyance-Tracker/backend/service/auth"
)

type Impl struct {
	userSvc service.User
	authSvc auth.Service
}

func NewImpl(userSvc service.User, authSvc auth.Service) *Impl {
	return &Impl{
		userSvc: userSvc,
		authSvc: authSvc,
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

	api.ScanProductsHandler = operations.ScanProductsHandlerFunc(func(params operations.ScanProductsParams) middleware.Responder {
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

	api.ScanCheckHandler = operations.ScanCheckHandlerFunc(func(params operations.ScanCheckParams) middleware.Responder {
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

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

func getStrPtr(s string) *string {
	return &s
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
