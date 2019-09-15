// Code generated by go-swagger; DO NOT EDIT.

package stop_point

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewStopPointGetTaxiRanksByIdsParams creates a new StopPointGetTaxiRanksByIdsParams object
// with the default values initialized.
func NewStopPointGetTaxiRanksByIdsParams() *StopPointGetTaxiRanksByIdsParams {
	var ()
	return &StopPointGetTaxiRanksByIdsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStopPointGetTaxiRanksByIdsParamsWithTimeout creates a new StopPointGetTaxiRanksByIdsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStopPointGetTaxiRanksByIdsParamsWithTimeout(timeout time.Duration) *StopPointGetTaxiRanksByIdsParams {
	var ()
	return &StopPointGetTaxiRanksByIdsParams{

		timeout: timeout,
	}
}

// NewStopPointGetTaxiRanksByIdsParamsWithContext creates a new StopPointGetTaxiRanksByIdsParams object
// with the default values initialized, and the ability to set a context for a request
func NewStopPointGetTaxiRanksByIdsParamsWithContext(ctx context.Context) *StopPointGetTaxiRanksByIdsParams {
	var ()
	return &StopPointGetTaxiRanksByIdsParams{

		Context: ctx,
	}
}

// NewStopPointGetTaxiRanksByIdsParamsWithHTTPClient creates a new StopPointGetTaxiRanksByIdsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStopPointGetTaxiRanksByIdsParamsWithHTTPClient(client *http.Client) *StopPointGetTaxiRanksByIdsParams {
	var ()
	return &StopPointGetTaxiRanksByIdsParams{
		HTTPClient: client,
	}
}

/*StopPointGetTaxiRanksByIdsParams contains all the parameters to send to the API endpoint
for the stop point get taxi ranks by ids operation typically these are written to a http.Request
*/
type StopPointGetTaxiRanksByIdsParams struct {

	/*StopPointID
	  stopPointId is required to get the taxi ranks.

	*/
	StopPointID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the stop point get taxi ranks by ids params
func (o *StopPointGetTaxiRanksByIdsParams) WithTimeout(timeout time.Duration) *StopPointGetTaxiRanksByIdsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stop point get taxi ranks by ids params
func (o *StopPointGetTaxiRanksByIdsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stop point get taxi ranks by ids params
func (o *StopPointGetTaxiRanksByIdsParams) WithContext(ctx context.Context) *StopPointGetTaxiRanksByIdsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stop point get taxi ranks by ids params
func (o *StopPointGetTaxiRanksByIdsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stop point get taxi ranks by ids params
func (o *StopPointGetTaxiRanksByIdsParams) WithHTTPClient(client *http.Client) *StopPointGetTaxiRanksByIdsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stop point get taxi ranks by ids params
func (o *StopPointGetTaxiRanksByIdsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithStopPointID adds the stopPointID to the stop point get taxi ranks by ids params
func (o *StopPointGetTaxiRanksByIdsParams) WithStopPointID(stopPointID string) *StopPointGetTaxiRanksByIdsParams {
	o.SetStopPointID(stopPointID)
	return o
}

// SetStopPointID adds the stopPointId to the stop point get taxi ranks by ids params
func (o *StopPointGetTaxiRanksByIdsParams) SetStopPointID(stopPointID string) {
	o.StopPointID = stopPointID
}

// WriteToRequest writes these params to a swagger request
func (o *StopPointGetTaxiRanksByIdsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param stopPointId
	if err := r.SetPathParam("stopPointId", o.StopPointID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
