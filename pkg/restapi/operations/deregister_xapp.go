// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeregisterXappHandlerFunc turns a function with the right signature into a deregister xapp handler
type DeregisterXappHandlerFunc func(DeregisterXappParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeregisterXappHandlerFunc) Handle(params DeregisterXappParams) middleware.Responder {
	return fn(params)
}

// DeregisterXappHandler interface for that can handle valid deregister xapp params
type DeregisterXappHandler interface {
	Handle(DeregisterXappParams) middleware.Responder
}

// NewDeregisterXapp creates a new http.Handler for the deregister xapp operation
func NewDeregisterXapp(ctx *middleware.Context, handler DeregisterXappHandler) *DeregisterXapp {
	return &DeregisterXapp{Context: ctx, Handler: handler}
}

/*DeregisterXapp swagger:route POST /deregister xapp registration deregisterXapp

Deregister an existing xApp

*/
type DeregisterXapp struct {
	Context *middleware.Context
	Handler DeregisterXappHandler
}

func (o *DeregisterXapp) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeregisterXappParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
