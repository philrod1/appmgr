// Code generated by go-swagger; DO NOT EDIT.

package health

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetHealthAliveOKCode is the HTTP code returned for type GetHealthAliveOK
const GetHealthAliveOKCode int = 200

/*GetHealthAliveOK Status of xApp Manager is ok

swagger:response getHealthAliveOK
*/
type GetHealthAliveOK struct {
}

// NewGetHealthAliveOK creates GetHealthAliveOK with default headers values
func NewGetHealthAliveOK() *GetHealthAliveOK {

	return &GetHealthAliveOK{}
}

// WriteResponse to the client
func (o *GetHealthAliveOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}
