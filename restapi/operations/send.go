// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// SendHandlerFunc turns a function with the right signature into a send handler
type SendHandlerFunc func(SendParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn SendHandlerFunc) Handle(params SendParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// SendHandler interface for that can handle valid send params
type SendHandler interface {
	Handle(SendParams, interface{}) middleware.Responder
}

// NewSend creates a new http.Handler for the send operation
func NewSend(ctx *middleware.Context, handler SendHandler) *Send {
	return &Send{Context: ctx, Handler: handler}
}

/*Send swagger:route GET /send send

send some crypto

send some crypto

*/
type Send struct {
	Context *middleware.Context
	Handler SendHandler
}

func (o *Send) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewSendParams()

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
