// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Magicking/dimple/models"
)

// SendOKCode is the HTTP code returned for type SendOK
const SendOKCode int = 200

/*SendOK return txid

swagger:response sendOK
*/
type SendOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewSendOK creates SendOK with default headers values
func NewSendOK() *SendOK {
	return &SendOK{}
}

// WithPayload adds the payload to the send o k response
func (o *SendOK) WithPayload(payload string) *SendOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the send o k response
func (o *SendOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SendOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*SendDefault Unexpected error

swagger:response sendDefault
*/
type SendDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewSendDefault creates SendDefault with default headers values
func NewSendDefault(code int) *SendDefault {
	if code <= 0 {
		code = 500
	}

	return &SendDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the send default response
func (o *SendDefault) WithStatusCode(code int) *SendDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the send default response
func (o *SendDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the send default response
func (o *SendDefault) WithPayload(payload *models.Error) *SendDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the send default response
func (o *SendDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SendDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
