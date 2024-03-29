// Code generated by go-swagger; DO NOT EDIT.

package road

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

// NewRoadDisruptionByIDParams creates a new RoadDisruptionByIDParams object
// with the default values initialized.
func NewRoadDisruptionByIDParams() *RoadDisruptionByIDParams {
	var ()
	return &RoadDisruptionByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRoadDisruptionByIDParamsWithTimeout creates a new RoadDisruptionByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRoadDisruptionByIDParamsWithTimeout(timeout time.Duration) *RoadDisruptionByIDParams {
	var ()
	return &RoadDisruptionByIDParams{

		timeout: timeout,
	}
}

// NewRoadDisruptionByIDParamsWithContext creates a new RoadDisruptionByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewRoadDisruptionByIDParamsWithContext(ctx context.Context) *RoadDisruptionByIDParams {
	var ()
	return &RoadDisruptionByIDParams{

		Context: ctx,
	}
}

// NewRoadDisruptionByIDParamsWithHTTPClient creates a new RoadDisruptionByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRoadDisruptionByIDParamsWithHTTPClient(client *http.Client) *RoadDisruptionByIDParams {
	var ()
	return &RoadDisruptionByIDParams{
		HTTPClient: client,
	}
}

/*RoadDisruptionByIDParams contains all the parameters to send to the API endpoint
for the road disruption by Id operation typically these are written to a http.Request
*/
type RoadDisruptionByIDParams struct {

	/*DisruptionIds
	  Comma-separated list of disruption identifiers to filter by.

	*/
	DisruptionIds []string
	/*StripContent
	  Optional, defaults to false. When true, removes every property/node except for id, point, severity, severityDescription, startDate, endDate, corridor details, location and comments.

	*/
	StripContent *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the road disruption by Id params
func (o *RoadDisruptionByIDParams) WithTimeout(timeout time.Duration) *RoadDisruptionByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the road disruption by Id params
func (o *RoadDisruptionByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the road disruption by Id params
func (o *RoadDisruptionByIDParams) WithContext(ctx context.Context) *RoadDisruptionByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the road disruption by Id params
func (o *RoadDisruptionByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the road disruption by Id params
func (o *RoadDisruptionByIDParams) WithHTTPClient(client *http.Client) *RoadDisruptionByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the road disruption by Id params
func (o *RoadDisruptionByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDisruptionIds adds the disruptionIds to the road disruption by Id params
func (o *RoadDisruptionByIDParams) WithDisruptionIds(disruptionIds []string) *RoadDisruptionByIDParams {
	o.SetDisruptionIds(disruptionIds)
	return o
}

// SetDisruptionIds adds the disruptionIds to the road disruption by Id params
func (o *RoadDisruptionByIDParams) SetDisruptionIds(disruptionIds []string) {
	o.DisruptionIds = disruptionIds
}

// WithStripContent adds the stripContent to the road disruption by Id params
func (o *RoadDisruptionByIDParams) WithStripContent(stripContent *bool) *RoadDisruptionByIDParams {
	o.SetStripContent(stripContent)
	return o
}

// SetStripContent adds the stripContent to the road disruption by Id params
func (o *RoadDisruptionByIDParams) SetStripContent(stripContent *bool) {
	o.StripContent = stripContent
}

// WriteToRequest writes these params to a swagger request
func (o *RoadDisruptionByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesDisruptionIds := o.DisruptionIds

	joinedDisruptionIds := swag.JoinByFormat(valuesDisruptionIds, "")
	// path array param disruptionIds
	// SetPathParam does not support variadric arguments, since we used JoinByFormat
	// we can send the first item in the array as it's all the items of the previous
	// array joined together
	if len(joinedDisruptionIds) > 0 {
		if err := r.SetPathParam("disruptionIds", joinedDisruptionIds[0]); err != nil {
			return err
		}
	}

	if o.StripContent != nil {

		// query param stripContent
		var qrStripContent bool
		if o.StripContent != nil {
			qrStripContent = *o.StripContent
		}
		qStripContent := swag.FormatBool(qrStripContent)
		if qStripContent != "" {
			if err := r.SetQueryParam("stripContent", qStripContent); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
