// Code generated by go-swagger; DO NOT EDIT.

package vehicle

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewVehicleGetParams creates a new VehicleGetParams object
// with the default values initialized.
func NewVehicleGetParams() *VehicleGetParams {
	var ()
	return &VehicleGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVehicleGetParamsWithTimeout creates a new VehicleGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVehicleGetParamsWithTimeout(timeout time.Duration) *VehicleGetParams {
	var ()
	return &VehicleGetParams{

		timeout: timeout,
	}
}

// NewVehicleGetParamsWithContext creates a new VehicleGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewVehicleGetParamsWithContext(ctx context.Context) *VehicleGetParams {
	var ()
	return &VehicleGetParams{

		Context: ctx,
	}
}

// NewVehicleGetParamsWithHTTPClient creates a new VehicleGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewVehicleGetParamsWithHTTPClient(client *http.Client) *VehicleGetParams {
	var ()
	return &VehicleGetParams{
		HTTPClient: client,
	}
}

/*VehicleGetParams contains all the parameters to send to the API endpoint
for the vehicle get operation typically these are written to a http.Request
*/
type VehicleGetParams struct {

	/*Ids
	  A comma-separated list of vehicle ids e.g. LX58CFV,LX11AZB,LX58CFE. Max approx. 25 ids.

	*/
	Ids []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the vehicle get params
func (o *VehicleGetParams) WithTimeout(timeout time.Duration) *VehicleGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the vehicle get params
func (o *VehicleGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the vehicle get params
func (o *VehicleGetParams) WithContext(ctx context.Context) *VehicleGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the vehicle get params
func (o *VehicleGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the vehicle get params
func (o *VehicleGetParams) WithHTTPClient(client *http.Client) *VehicleGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the vehicle get params
func (o *VehicleGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIds adds the ids to the vehicle get params
func (o *VehicleGetParams) WithIds(ids []string) *VehicleGetParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the vehicle get params
func (o *VehicleGetParams) SetIds(ids []string) {
	o.Ids = ids
}

// WriteToRequest writes these params to a swagger request
func (o *VehicleGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesIds := o.Ids

	joinedIds := swag.JoinByFormat(valuesIds, "")
	// path array param ids
	// SetPathParam does not support variadric arguments, since we used JoinByFormat
	// we can send the first item in the array as it's all the items of the previous
	// array joined together
	if len(joinedIds) > 0 {
		if err := r.SetPathParam("ids", joinedIds[0]); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}