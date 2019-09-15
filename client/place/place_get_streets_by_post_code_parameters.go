// Code generated by go-swagger; DO NOT EDIT.

package place

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

// NewPlaceGetStreetsByPostCodeParams creates a new PlaceGetStreetsByPostCodeParams object
// with the default values initialized.
func NewPlaceGetStreetsByPostCodeParams() *PlaceGetStreetsByPostCodeParams {
	var ()
	return &PlaceGetStreetsByPostCodeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPlaceGetStreetsByPostCodeParamsWithTimeout creates a new PlaceGetStreetsByPostCodeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPlaceGetStreetsByPostCodeParamsWithTimeout(timeout time.Duration) *PlaceGetStreetsByPostCodeParams {
	var ()
	return &PlaceGetStreetsByPostCodeParams{

		timeout: timeout,
	}
}

// NewPlaceGetStreetsByPostCodeParamsWithContext creates a new PlaceGetStreetsByPostCodeParams object
// with the default values initialized, and the ability to set a context for a request
func NewPlaceGetStreetsByPostCodeParamsWithContext(ctx context.Context) *PlaceGetStreetsByPostCodeParams {
	var ()
	return &PlaceGetStreetsByPostCodeParams{

		Context: ctx,
	}
}

// NewPlaceGetStreetsByPostCodeParamsWithHTTPClient creates a new PlaceGetStreetsByPostCodeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPlaceGetStreetsByPostCodeParamsWithHTTPClient(client *http.Client) *PlaceGetStreetsByPostCodeParams {
	var ()
	return &PlaceGetStreetsByPostCodeParams{
		HTTPClient: client,
	}
}

/*PlaceGetStreetsByPostCodeParams contains all the parameters to send to the API endpoint
for the place get streets by post code operation typically these are written to a http.Request
*/
type PlaceGetStreetsByPostCodeParams struct {

	/*Postcode*/
	Postcode string
	/*PostcodeInputPostcode*/
	PostcodeInputPostcode *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the place get streets by post code params
func (o *PlaceGetStreetsByPostCodeParams) WithTimeout(timeout time.Duration) *PlaceGetStreetsByPostCodeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the place get streets by post code params
func (o *PlaceGetStreetsByPostCodeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the place get streets by post code params
func (o *PlaceGetStreetsByPostCodeParams) WithContext(ctx context.Context) *PlaceGetStreetsByPostCodeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the place get streets by post code params
func (o *PlaceGetStreetsByPostCodeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the place get streets by post code params
func (o *PlaceGetStreetsByPostCodeParams) WithHTTPClient(client *http.Client) *PlaceGetStreetsByPostCodeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the place get streets by post code params
func (o *PlaceGetStreetsByPostCodeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPostcode adds the postcode to the place get streets by post code params
func (o *PlaceGetStreetsByPostCodeParams) WithPostcode(postcode string) *PlaceGetStreetsByPostCodeParams {
	o.SetPostcode(postcode)
	return o
}

// SetPostcode adds the postcode to the place get streets by post code params
func (o *PlaceGetStreetsByPostCodeParams) SetPostcode(postcode string) {
	o.Postcode = postcode
}

// WithPostcodeInputPostcode adds the postcodeInputPostcode to the place get streets by post code params
func (o *PlaceGetStreetsByPostCodeParams) WithPostcodeInputPostcode(postcodeInputPostcode *string) *PlaceGetStreetsByPostCodeParams {
	o.SetPostcodeInputPostcode(postcodeInputPostcode)
	return o
}

// SetPostcodeInputPostcode adds the postcodeInputPostcode to the place get streets by post code params
func (o *PlaceGetStreetsByPostCodeParams) SetPostcodeInputPostcode(postcodeInputPostcode *string) {
	o.PostcodeInputPostcode = postcodeInputPostcode
}

// WriteToRequest writes these params to a swagger request
func (o *PlaceGetStreetsByPostCodeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param postcode
	qrPostcode := o.Postcode
	qPostcode := qrPostcode
	if qPostcode != "" {
		if err := r.SetQueryParam("postcode", qPostcode); err != nil {
			return err
		}
	}

	if o.PostcodeInputPostcode != nil {

		// query param postcodeInput.postcode
		var qrPostcodeInputPostcode string
		if o.PostcodeInputPostcode != nil {
			qrPostcodeInputPostcode = *o.PostcodeInputPostcode
		}
		qPostcodeInputPostcode := qrPostcodeInputPostcode
		if qPostcodeInputPostcode != "" {
			if err := r.SetQueryParam("postcodeInput.postcode", qPostcodeInputPostcode); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
