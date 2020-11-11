// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/rs/cors"

	"github.com/Smart-Purveyance-Tracker/backend/entity"
	"github.com/Smart-Purveyance-Tracker/backend/models"
	"github.com/Smart-Purveyance-Tracker/backend/repository"
	"github.com/Smart-Purveyance-Tracker/backend/restapi/operations"
	"github.com/Smart-Purveyance-Tracker/backend/service"
	"github.com/Smart-Purveyance-Tracker/backend/service/auth"
)

//go:generate swagger generate server --target ../../backend --name Swagger --spec ../swagger-api/swagger.yml --principal interface{}

func configureFlags(api *operations.SwaggerAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.SwaggerAPI) http.Handler {
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

	userRepo := repository.NewUserInMem()
	userSvc := service.NewUserImpl(userRepo)
	jwtSvc := auth.NewJWTService([]byte("todo"))
	api.SignupHandler = operations.SignupHandlerFunc(func(params operations.SignupParams) middleware.Responder {
		user, err := userSvc.Create(entity.User{
			Email:    *params.UserInfo.Email,
			Password: *params.UserInfo.Password,
		})
		if err != nil {
			return operations.NewSignupDefault(http.StatusInternalServerError).WithPayload(&models.Error{
				Message: getStrPtr(err.Error()),
			})
		}

		token, err := jwtSvc.GenerateToken(user.ID)
		if err != nil {
			return operations.NewSignupDefault(http.StatusInternalServerError).WithPayload(&models.Error{
				Message: getStrPtr(err.Error()),
			})
		}

		return operations.NewSignupOK().WithPayload(&models.User{
			ID:    int64(user.ID),
			Email: user.Email,
		}).WithAuthenthication("Bearer " + token)
	})

	api.LoginHandler = operations.LoginHandlerFunc(func(params operations.LoginParams) middleware.Responder {
		user, err := userSvc.Login(*params.UserInfo.Email, *params.UserInfo.Password)
		if err != nil {
			return operations.NewLoginDefault(http.StatusInternalServerError).WithPayload(&models.Error{
				Message: getStrPtr(err.Error()),
			})
		}
		token, err := jwtSvc.GenerateToken(user.ID)
		if err != nil {
			return operations.NewLoginDefault(http.StatusInternalServerError).WithPayload(&models.Error{
				Message: getStrPtr(err.Error()),
			})
		}

		return operations.NewLoginOK().WithPayload(&models.User{
			ID:    int64(user.ID),
			Email: user.Email,
		}).WithAuthenthication("Bearer " + token)
	})

	api.ScanProductsHandler = operations.ScanProductsHandlerFunc(func(params operations.ScanProductsParams) middleware.Responder {
		return operations.NewScanProductsOK().WithPayload([]*models.Product{
			{
				ID:   1,
				Name: "овощ",
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

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	handleCORS := cors.Default().Handler

	return handleCORS(handler)
}
