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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPlaceGetAtParams creates a new PlaceGetAtParams object
// with the default values initialized.
func NewPlaceGetAtParams() *PlaceGetAtParams {
	var ()
	return &PlaceGetAtParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPlaceGetAtParamsWithTimeout creates a new PlaceGetAtParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPlaceGetAtParamsWithTimeout(timeout time.Duration) *PlaceGetAtParams {
	var ()
	return &PlaceGetAtParams{

		timeout: timeout,
	}
}

// NewPlaceGetAtParamsWithContext creates a new PlaceGetAtParams object
// with the default values initialized, and the ability to set a context for a request
func NewPlaceGetAtParamsWithContext(ctx context.Context) *PlaceGetAtParams {
	var ()
	return &PlaceGetAtParams{

		Context: ctx,
	}
}

// NewPlaceGetAtParamsWithHTTPClient creates a new PlaceGetAtParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPlaceGetAtParamsWithHTTPClient(client *http.Client) *PlaceGetAtParams {
	var ()
	return &PlaceGetAtParams{
		HTTPClient: client,
	}
}

/*PlaceGetAtParams contains all the parameters to send to the API endpoint
for the place get at operation typically these are written to a http.Request
*/
type PlaceGetAtParams struct {

	/*Lat*/
	Lat string
	/*LocationLat*/
	LocationLat float64
	/*LocationLon*/
	LocationLon float64
	/*Lon*/
	Lon string
	/*Type
	  The place type (a valid list of place types can be obtained from the /Place/Meta/placeTypes endpoint)

	*/
	Type []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the place get at params
func (o *PlaceGetAtParams) WithTimeout(timeout time.Duration) *PlaceGetAtParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the place get at params
func (o *PlaceGetAtParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the place get at params
func (o *PlaceGetAtParams) WithContext(ctx context.Context) *PlaceGetAtParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the place get at params
func (o *PlaceGetAtParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the place get at params
func (o *PlaceGetAtParams) WithHTTPClient(client *http.Client) *PlaceGetAtParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the place get at params
func (o *PlaceGetAtParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLat adds the lat to the place get at params
func (o *PlaceGetAtParams) WithLat(lat string) *PlaceGetAtParams {
	o.SetLat(lat)
	return o
}

// SetLat adds the lat to the place get at params
func (o *PlaceGetAtParams) SetLat(lat string) {
	o.Lat = lat
}

// WithLocationLat adds the locationLat to the place get at params
func (o *PlaceGetAtParams) WithLocationLat(locationLat float64) *PlaceGetAtParams {
	o.SetLocationLat(locationLat)
	return o
}

// SetLocationLat adds the locationLat to the place get at params
func (o *PlaceGetAtParams) SetLocationLat(locationLat float64) {
	o.LocationLat = locationLat
}

// WithLocationLon adds the locationLon to the place get at params
func (o *PlaceGetAtParams) WithLocationLon(locationLon float64) *PlaceGetAtParams {
	o.SetLocationLon(locationLon)
	return o
}

// SetLocationLon adds the locationLon to the place get at params
func (o *PlaceGetAtParams) SetLocationLon(locationLon float64) {
	o.LocationLon = locationLon
}

// WithLon adds the lon to the place get at params
func (o *PlaceGetAtParams) WithLon(lon string) *PlaceGetAtParams {
	o.SetLon(lon)
	return o
}

// SetLon adds the lon to the place get at params
func (o *PlaceGetAtParams) SetLon(lon string) {
	o.Lon = lon
}

// WithType adds the typeVar to the place get at params
func (o *PlaceGetAtParams) WithType(typeVar []string) *PlaceGetAtParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the place get at params
func (o *PlaceGetAtParams) SetType(typeVar []string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *PlaceGetAtParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param lat
	qrLat := o.Lat
	qLat := qrLat
	if qLat != "" {
		if err := r.SetQueryParam("lat", qLat); err != nil {
			return err
		}
	}

	// query param location.lat
	qrLocationLat := o.LocationLat
	qLocationLat := swag.FormatFloat64(qrLocationLat)
	if qLocationLat != "" {
		if err := r.SetQueryParam("location.lat", qLocationLat); err != nil {
			return err
		}
	}

	// query param location.lon
	qrLocationLon := o.LocationLon
	qLocationLon := swag.FormatFloat64(qrLocationLon)
	if qLocationLon != "" {
		if err := r.SetQueryParam("location.lon", qLocationLon); err != nil {
			return err
		}
	}

	// query param lon
	qrLon := o.Lon
	qLon := qrLon
	if qLon != "" {
		if err := r.SetQueryParam("lon", qLon); err != nil {
			return err
		}
	}

	valuesType := o.Type

	joinedType := swag.JoinByFormat(valuesType, "")
	// path array param type
	// SetPathParam does not support variadric arguments, since we used JoinByFormat
	// we can send the first item in the array as it's all the items of the previous
	// array joined together
	if len(joinedType) > 0 {
		if err := r.SetPathParam("type", joinedType[0]); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
