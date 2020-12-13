// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewSwaggerAPI creates a new Swagger instance
func NewSwaggerAPI(spec *loads.Document) *SwaggerAPI {
	return &SwaggerAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		CreateProductHandler: CreateProductHandlerFunc(func(params CreateProductParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation CreateProduct has not yet been implemented")
		}),
		DeleteProductHandler: DeleteProductHandlerFunc(func(params DeleteProductParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation DeleteProduct has not yet been implemented")
		}),
		GetProductHandler: GetProductHandlerFunc(func(params GetProductParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation GetProduct has not yet been implemented")
		}),
		GetStatusHandler: GetStatusHandlerFunc(func(params GetStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation GetStatus has not yet been implemented")
		}),
		LoginHandler: LoginHandlerFunc(func(params LoginParams) middleware.Responder {
			return middleware.NotImplemented("operation Login has not yet been implemented")
		}),
		ProductListHandler: ProductListHandlerFunc(func(params ProductListParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation ProductList has not yet been implemented")
		}),
		ScanCheckHandler: ScanCheckHandlerFunc(func(params ScanCheckParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation ScanCheck has not yet been implemented")
		}),
		ScanProductsHandler: ScanProductsHandlerFunc(func(params ScanProductsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation ScanProducts has not yet been implemented")
		}),
		SignupHandler: SignupHandlerFunc(func(params SignupParams) middleware.Responder {
			return middleware.NotImplemented("operation Signup has not yet been implemented")
		}),
		UpdateProductHandler: UpdateProductHandlerFunc(func(params UpdateProductParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation UpdateProduct has not yet been implemented")
		}),

		// Applies when the "Authorization" header is set
		BearerAuth: func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (Bearer) Authorization from header param [Authorization] has not yet been implemented")
		},
		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*SwaggerAPI test */
type SwaggerAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// BearerAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key Authorization provided in the header
	BearerAuth func(string) (interface{}, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// CreateProductHandler sets the operation handler for the create product operation
	CreateProductHandler CreateProductHandler
	// DeleteProductHandler sets the operation handler for the delete product operation
	DeleteProductHandler DeleteProductHandler
	// GetProductHandler sets the operation handler for the get product operation
	GetProductHandler GetProductHandler
	// GetStatusHandler sets the operation handler for the get status operation
	GetStatusHandler GetStatusHandler
	// LoginHandler sets the operation handler for the login operation
	LoginHandler LoginHandler
	// ProductListHandler sets the operation handler for the product list operation
	ProductListHandler ProductListHandler
	// ScanCheckHandler sets the operation handler for the scan check operation
	ScanCheckHandler ScanCheckHandler
	// ScanProductsHandler sets the operation handler for the scan products operation
	ScanProductsHandler ScanProductsHandler
	// SignupHandler sets the operation handler for the signup operation
	SignupHandler SignupHandler
	// UpdateProductHandler sets the operation handler for the update product operation
	UpdateProductHandler UpdateProductHandler
	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *SwaggerAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *SwaggerAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *SwaggerAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *SwaggerAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *SwaggerAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *SwaggerAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *SwaggerAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *SwaggerAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *SwaggerAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the SwaggerAPI
func (o *SwaggerAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.BearerAuth == nil {
		unregistered = append(unregistered, "AuthorizationAuth")
	}

	if o.CreateProductHandler == nil {
		unregistered = append(unregistered, "CreateProductHandler")
	}
	if o.DeleteProductHandler == nil {
		unregistered = append(unregistered, "DeleteProductHandler")
	}
	if o.GetProductHandler == nil {
		unregistered = append(unregistered, "GetProductHandler")
	}
	if o.GetStatusHandler == nil {
		unregistered = append(unregistered, "GetStatusHandler")
	}
	if o.LoginHandler == nil {
		unregistered = append(unregistered, "LoginHandler")
	}
	if o.ProductListHandler == nil {
		unregistered = append(unregistered, "ProductListHandler")
	}
	if o.ScanCheckHandler == nil {
		unregistered = append(unregistered, "ScanCheckHandler")
	}
	if o.ScanProductsHandler == nil {
		unregistered = append(unregistered, "ScanProductsHandler")
	}
	if o.SignupHandler == nil {
		unregistered = append(unregistered, "SignupHandler")
	}
	if o.UpdateProductHandler == nil {
		unregistered = append(unregistered, "UpdateProductHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *SwaggerAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *SwaggerAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	result := make(map[string]runtime.Authenticator)
	for name := range schemes {
		switch name {
		case "Bearer":
			scheme := schemes[name]
			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, o.BearerAuth)

		}
	}
	return result
}

// Authorizer returns the registered authorizer
func (o *SwaggerAPI) Authorizer() runtime.Authorizer {
	return o.APIAuthorizer
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *SwaggerAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *SwaggerAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *SwaggerAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the swagger API
func (o *SwaggerAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *SwaggerAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/product/{productID}"] = NewCreateProduct(o.context, o.CreateProductHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/product/{productID}"] = NewDeleteProduct(o.context, o.DeleteProductHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/product/{productID}"] = NewGetProduct(o.context, o.GetProductHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/status"] = NewGetStatus(o.context, o.GetStatusHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/login"] = NewLogin(o.context, o.LoginHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/product/list"] = NewProductList(o.context, o.ProductListHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/scanCheck"] = NewScanCheck(o.context, o.ScanCheckHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/scanProducts"] = NewScanProducts(o.context, o.ScanProductsHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/signup"] = NewSignup(o.context, o.SignupHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/product/{productID}"] = NewUpdateProduct(o.context, o.UpdateProductHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *SwaggerAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *SwaggerAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *SwaggerAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *SwaggerAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *SwaggerAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
