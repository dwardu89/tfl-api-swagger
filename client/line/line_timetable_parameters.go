// Code generated by go-swagger; DO NOT EDIT.

package line

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

// NewLineTimetableParams creates a new LineTimetableParams object
// with the default values initialized.
func NewLineTimetableParams() *LineTimetableParams {
	var ()
	return &LineTimetableParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLineTimetableParamsWithTimeout creates a new LineTimetableParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLineTimetableParamsWithTimeout(timeout time.Duration) *LineTimetableParams {
	var ()
	return &LineTimetableParams{

		timeout: timeout,
	}
}

// NewLineTimetableParamsWithContext creates a new LineTimetableParams object
// with the default values initialized, and the ability to set a context for a request
func NewLineTimetableParamsWithContext(ctx context.Context) *LineTimetableParams {
	var ()
	return &LineTimetableParams{

		Context: ctx,
	}
}

// NewLineTimetableParamsWithHTTPClient creates a new LineTimetableParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLineTimetableParamsWithHTTPClient(client *http.Client) *LineTimetableParams {
	var ()
	return &LineTimetableParams{
		HTTPClient: client,
	}
}

/*LineTimetableParams contains all the parameters to send to the API endpoint
for the line timetable operation typically these are written to a http.Request
*/
type LineTimetableParams struct {

	/*FromStopPointID
	  The originating station's stop point id (station naptan code e.g. 940GZZLUASL, you can use /StopPoint/Search/{query} endpoint to find a stop point id from a station name)

	*/
	FromStopPointID string
	/*ID
	  A single line id e.g. victoria

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the line timetable params
func (o *LineTimetableParams) WithTimeout(timeout time.Duration) *LineTimetableParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the line timetable params
func (o *LineTimetableParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the line timetable params
func (o *LineTimetableParams) WithContext(ctx context.Context) *LineTimetableParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the line timetable params
func (o *LineTimetableParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the line timetable params
func (o *LineTimetableParams) WithHTTPClient(client *http.Client) *LineTimetableParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the line timetable params
func (o *LineTimetableParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFromStopPointID adds the fromStopPointID to the line timetable params
func (o *LineTimetableParams) WithFromStopPointID(fromStopPointID string) *LineTimetableParams {
	o.SetFromStopPointID(fromStopPointID)
	return o
}

// SetFromStopPointID adds the fromStopPointId to the line timetable params
func (o *LineTimetableParams) SetFromStopPointID(fromStopPointID string) {
	o.FromStopPointID = fromStopPointID
}

// WithID adds the id to the line timetable params
func (o *LineTimetableParams) WithID(id string) *LineTimetableParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the line timetable params
func (o *LineTimetableParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *LineTimetableParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param fromStopPointId
	if err := r.SetPathParam("fromStopPointId", o.FromStopPointID); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
