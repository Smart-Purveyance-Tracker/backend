// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ScanCheckHandlerFunc turns a function with the right signature into a scan check handler
type ScanCheckHandlerFunc func(ScanCheckParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ScanCheckHandlerFunc) Handle(params ScanCheckParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ScanCheckHandler interface for that can handle valid scan check params
type ScanCheckHandler interface {
	Handle(ScanCheckParams, interface{}) middleware.Responder
}

// NewScanCheck creates a new http.Handler for the scan check operation
func NewScanCheck(ctx *middleware.Context, handler ScanCheckHandler) *ScanCheck {
	return &ScanCheck{Context: ctx, Handler: handler}
}

/*ScanCheck swagger:route POST /scanCheck scanCheck

Uploads a file.

*/
type ScanCheck struct {
	Context *middleware.Context
	Handler ScanCheckHandler
}

func (o *ScanCheck) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewScanCheckParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// ScanCheckBody scan check body
//
// swagger:model ScanCheckBody
type ScanCheckBody struct {

	// body
	// Required: true
	Body *string `json:"body"`
}

// Validate validates this scan check body
func (o *ScanCheckBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateBody(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ScanCheckBody) validateBody(formats strfmt.Registry) error {

	if err := validate.Required("image"+"."+"body", "body", o.Body); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this scan check body based on context it is used
func (o *ScanCheckBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ScanCheckBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ScanCheckBody) UnmarshalBinary(b []byte) error {
	var res ScanCheckBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
