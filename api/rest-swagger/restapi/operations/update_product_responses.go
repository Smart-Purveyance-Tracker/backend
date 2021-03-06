// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Smart-Purveyance-Tracker/backend/api/rest-swagger/models"
)

// UpdateProductOKCode is the HTTP code returned for type UpdateProductOK
const UpdateProductOKCode int = 200

/*UpdateProductOK update product

swagger:response updateProductOK
*/
type UpdateProductOK struct {

	/*
	  In: Body
	*/
	Payload *models.Product `json:"body,omitempty"`
}

// NewUpdateProductOK creates UpdateProductOK with default headers values
func NewUpdateProductOK() *UpdateProductOK {

	return &UpdateProductOK{}
}

// WithPayload adds the payload to the update product o k response
func (o *UpdateProductOK) WithPayload(payload *models.Product) *UpdateProductOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update product o k response
func (o *UpdateProductOK) SetPayload(payload *models.Product) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateProductOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*UpdateProductDefault error

swagger:response updateProductDefault
*/
type UpdateProductDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateProductDefault creates UpdateProductDefault with default headers values
func NewUpdateProductDefault(code int) *UpdateProductDefault {
	if code <= 0 {
		code = 500
	}

	return &UpdateProductDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the update product default response
func (o *UpdateProductDefault) WithStatusCode(code int) *UpdateProductDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the update product default response
func (o *UpdateProductDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the update product default response
func (o *UpdateProductDefault) WithPayload(payload *models.Error) *UpdateProductDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update product default response
func (o *UpdateProductDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateProductDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
