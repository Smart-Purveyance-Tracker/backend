// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewProductListParams creates a new ProductListParams object
// no default values defined in spec.
func NewProductListParams() ProductListParams {

	return ProductListParams{}
}

// ProductListParams contains all the bound params for the product list operation
// typically these are obtained from a http.Request
//
// swagger:parameters productList
type ProductListParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: query
	*/
	Date *strfmt.Date
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewProductListParams() beforehand.
func (o *ProductListParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qDate, qhkDate, _ := qs.GetOK("date")
	if err := o.bindDate(qDate, qhkDate, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindDate binds and validates parameter Date from query.
func (o *ProductListParams) bindDate(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	// Format: date
	value, err := formats.Parse("date", raw)
	if err != nil {
		return errors.InvalidType("date", "query", "strfmt.Date", raw)
	}
	o.Date = (value.(*strfmt.Date))

	if err := o.validateDate(formats); err != nil {
		return err
	}

	return nil
}

// validateDate carries on validations for parameter Date
func (o *ProductListParams) validateDate(formats strfmt.Registry) error {

	if err := validate.FormatOf("date", "query", "date", o.Date.String(), formats); err != nil {
		return err
	}
	return nil
}
