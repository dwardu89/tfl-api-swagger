// Code generated by go-swagger; DO NOT EDIT.

package occupancy

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

// NewOccupancyGetChargeConnectorStatusParams creates a new OccupancyGetChargeConnectorStatusParams object
// with the default values initialized.
func NewOccupancyGetChargeConnectorStatusParams() *OccupancyGetChargeConnectorStatusParams {
	var ()
	return &OccupancyGetChargeConnectorStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOccupancyGetChargeConnectorStatusParamsWithTimeout creates a new OccupancyGetChargeConnectorStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOccupancyGetChargeConnectorStatusParamsWithTimeout(timeout time.Duration) *OccupancyGetChargeConnectorStatusParams {
	var ()
	return &OccupancyGetChargeConnectorStatusParams{

		timeout: timeout,
	}
}

// NewOccupancyGetChargeConnectorStatusParamsWithContext creates a new OccupancyGetChargeConnectorStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewOccupancyGetChargeConnectorStatusParamsWithContext(ctx context.Context) *OccupancyGetChargeConnectorStatusParams {
	var ()
	return &OccupancyGetChargeConnectorStatusParams{

		Context: ctx,
	}
}

// NewOccupancyGetChargeConnectorStatusParamsWithHTTPClient creates a new OccupancyGetChargeConnectorStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOccupancyGetChargeConnectorStatusParamsWithHTTPClient(client *http.Client) *OccupancyGetChargeConnectorStatusParams {
	var ()
	return &OccupancyGetChargeConnectorStatusParams{
		HTTPClient: client,
	}
}

/*OccupancyGetChargeConnectorStatusParams contains all the parameters to send to the API endpoint
for the occupancy get charge connector status operation typically these are written to a http.Request
*/
type OccupancyGetChargeConnectorStatusParams struct {

	/*Ids*/
	Ids []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the occupancy get charge connector status params
func (o *OccupancyGetChargeConnectorStatusParams) WithTimeout(timeout time.Duration) *OccupancyGetChargeConnectorStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the occupancy get charge connector status params
func (o *OccupancyGetChargeConnectorStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the occupancy get charge connector status params
func (o *OccupancyGetChargeConnectorStatusParams) WithContext(ctx context.Context) *OccupancyGetChargeConnectorStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the occupancy get charge connector status params
func (o *OccupancyGetChargeConnectorStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the occupancy get charge connector status params
func (o *OccupancyGetChargeConnectorStatusParams) WithHTTPClient(client *http.Client) *OccupancyGetChargeConnectorStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the occupancy get charge connector status params
func (o *OccupancyGetChargeConnectorStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIds adds the ids to the occupancy get charge connector status params
func (o *OccupancyGetChargeConnectorStatusParams) WithIds(ids []string) *OccupancyGetChargeConnectorStatusParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the occupancy get charge connector status params
func (o *OccupancyGetChargeConnectorStatusParams) SetIds(ids []string) {
	o.Ids = ids
}

// WriteToRequest writes these params to a swagger request
func (o *OccupancyGetChargeConnectorStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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