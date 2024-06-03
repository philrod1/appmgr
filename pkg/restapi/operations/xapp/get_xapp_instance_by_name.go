// Code generated by go-swagger; DO NOT EDIT.

package xapp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetXappInstanceByNameHandlerFunc turns a function with the right signature into a get xapp instance by name handler
type GetXappInstanceByNameHandlerFunc func(GetXappInstanceByNameParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetXappInstanceByNameHandlerFunc) Handle(params GetXappInstanceByNameParams) middleware.Responder {
	return fn(params)
}

// GetXappInstanceByNameHandler interface for that can handle valid get xapp instance by name params
type GetXappInstanceByNameHandler interface {
	Handle(GetXappInstanceByNameParams) middleware.Responder
}

// NewGetXappInstanceByName creates a new http.Handler for the get xapp instance by name operation
func NewGetXappInstanceByName(ctx *middleware.Context, handler GetXappInstanceByNameHandler) *GetXappInstanceByName {
	return &GetXappInstanceByName{Context: ctx, Handler: handler}
}

/*GetXappInstanceByName swagger:route GET /xapps/{xAppName}/instances/{xAppInstanceName} xapp getXappInstanceByName

Returns the status of a given xapp

*/
type GetXappInstanceByName struct {
	Context *middleware.Context
	Handler GetXappInstanceByNameHandler
}

func (o *GetXappInstanceByName) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetXappInstanceByNameParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
