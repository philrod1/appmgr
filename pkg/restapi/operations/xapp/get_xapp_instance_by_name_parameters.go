// Code generated by go-swagger; DO NOT EDIT.

package xapp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetXappInstanceByNameParams creates a new GetXappInstanceByNameParams object
// no default values defined in spec.
func NewGetXappInstanceByNameParams() GetXappInstanceByNameParams {

	return GetXappInstanceByNameParams{}
}

// GetXappInstanceByNameParams contains all the bound params for the get xapp instance by name operation
// typically these are obtained from a http.Request
//
// swagger:parameters getXappInstanceByName
type GetXappInstanceByNameParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Name of xApp instance to get information
	  Required: true
	  In: path
	*/
	XAppInstanceName string
	/*Name of xApp
	  Required: true
	  In: path
	*/
	XAppName string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetXappInstanceByNameParams() beforehand.
func (o *GetXappInstanceByNameParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rXAppInstanceName, rhkXAppInstanceName, _ := route.Params.GetOK("xAppInstanceName")
	if err := o.bindXAppInstanceName(rXAppInstanceName, rhkXAppInstanceName, route.Formats); err != nil {
		res = append(res, err)
	}

	rXAppName, rhkXAppName, _ := route.Params.GetOK("xAppName")
	if err := o.bindXAppName(rXAppName, rhkXAppName, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindXAppInstanceName binds and validates parameter XAppInstanceName from path.
func (o *GetXappInstanceByNameParams) bindXAppInstanceName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.XAppInstanceName = raw

	return nil
}

// bindXAppName binds and validates parameter XAppName from path.
func (o *GetXappInstanceByNameParams) bindXAppName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.XAppName = raw

	return nil
}
