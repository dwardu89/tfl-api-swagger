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

// NewStopPointCrowdingParams creates a new StopPointCrowdingParams object
// with the default values initialized.
func NewStopPointCrowdingParams() *StopPointCrowdingParams {
	var ()
	return &StopPointCrowdingParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStopPointCrowdingParamsWithTimeout creates a new StopPointCrowdingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStopPointCrowdingParamsWithTimeout(timeout time.Duration) *StopPointCrowdingParams {
	var ()
	return &StopPointCrowdingParams{

		timeout: timeout,
	}
}

// NewStopPointCrowdingParamsWithContext creates a new StopPointCrowdingParams object
// with the default values initialized, and the ability to set a context for a request
func NewStopPointCrowdingParamsWithContext(ctx context.Context) *StopPointCrowdingParams {
	var ()
	return &StopPointCrowdingParams{

		Context: ctx,
	}
}

// NewStopPointCrowdingParamsWithHTTPClient creates a new StopPointCrowdingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStopPointCrowdingParamsWithHTTPClient(client *http.Client) *StopPointCrowdingParams {
	var ()
	return &StopPointCrowdingParams{
		HTTPClient: client,
	}
}

/*StopPointCrowdingParams contains all the parameters to send to the API endpoint
for the stop point crowding operation typically these are written to a http.Request
*/
type StopPointCrowdingParams struct {

	/*Direction
	  The direction of travel. Can be inbound or outbound.

	*/
	Direction string
	/*ID
	  The Naptan id of the stop

	*/
	ID string
	/*Line
	  A particular line e.g. victoria, circle, northern etc.

	*/
	Line string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the stop point crowding params
func (o *StopPointCrowdingParams) WithTimeout(timeout time.Duration) *StopPointCrowdingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stop point crowding params
func (o *StopPointCrowdingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stop point crowding params
func (o *StopPointCrowdingParams) WithContext(ctx context.Context) *StopPointCrowdingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stop point crowding params
func (o *StopPointCrowdingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stop point crowding params
func (o *StopPointCrowdingParams) WithHTTPClient(client *http.Client) *StopPointCrowdingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stop point crowding params
func (o *StopPointCrowdingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDirection adds the direction to the stop point crowding params
func (o *StopPointCrowdingParams) WithDirection(direction string) *StopPointCrowdingParams {
	o.SetDirection(direction)
	return o
}

// SetDirection adds the direction to the stop point crowding params
func (o *StopPointCrowdingParams) SetDirection(direction string) {
	o.Direction = direction
}

// WithID adds the id to the stop point crowding params
func (o *StopPointCrowdingParams) WithID(id string) *StopPointCrowdingParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the stop point crowding params
func (o *StopPointCrowdingParams) SetID(id string) {
	o.ID = id
}

// WithLine adds the line to the stop point crowding params
func (o *StopPointCrowdingParams) WithLine(line string) *StopPointCrowdingParams {
	o.SetLine(line)
	return o
}

// SetLine adds the line to the stop point crowding params
func (o *StopPointCrowdingParams) SetLine(line string) {
	o.Line = line
}

// WriteToRequest writes these params to a swagger request
func (o *StopPointCrowdingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param direction
	qrDirection := o.Direction
	qDirection := qrDirection
	if qDirection != "" {
		if err := r.SetQueryParam("direction", qDirection); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param line
	if err := r.SetPathParam("line", o.Line); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
