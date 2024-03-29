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

// NewPlaceGetOverlayParams creates a new PlaceGetOverlayParams object
// with the default values initialized.
func NewPlaceGetOverlayParams() *PlaceGetOverlayParams {
	var ()
	return &PlaceGetOverlayParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPlaceGetOverlayParamsWithTimeout creates a new PlaceGetOverlayParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPlaceGetOverlayParamsWithTimeout(timeout time.Duration) *PlaceGetOverlayParams {
	var ()
	return &PlaceGetOverlayParams{

		timeout: timeout,
	}
}

// NewPlaceGetOverlayParamsWithContext creates a new PlaceGetOverlayParams object
// with the default values initialized, and the ability to set a context for a request
func NewPlaceGetOverlayParamsWithContext(ctx context.Context) *PlaceGetOverlayParams {
	var ()
	return &PlaceGetOverlayParams{

		Context: ctx,
	}
}

// NewPlaceGetOverlayParamsWithHTTPClient creates a new PlaceGetOverlayParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPlaceGetOverlayParamsWithHTTPClient(client *http.Client) *PlaceGetOverlayParams {
	var ()
	return &PlaceGetOverlayParams{
		HTTPClient: client,
	}
}

/*PlaceGetOverlayParams contains all the parameters to send to the API endpoint
for the place get overlay operation typically these are written to a http.Request
*/
type PlaceGetOverlayParams struct {

	/*Height
	  The height of the requested overlay.

	*/
	Height int32
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
	/*Width
	  The width of the requested overlay.

	*/
	Width int32
	/*Z
	  The zoom level

	*/
	Z int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the place get overlay params
func (o *PlaceGetOverlayParams) WithTimeout(timeout time.Duration) *PlaceGetOverlayParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the place get overlay params
func (o *PlaceGetOverlayParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the place get overlay params
func (o *PlaceGetOverlayParams) WithContext(ctx context.Context) *PlaceGetOverlayParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the place get overlay params
func (o *PlaceGetOverlayParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the place get overlay params
func (o *PlaceGetOverlayParams) WithHTTPClient(client *http.Client) *PlaceGetOverlayParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the place get overlay params
func (o *PlaceGetOverlayParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHeight adds the height to the place get overlay params
func (o *PlaceGetOverlayParams) WithHeight(height int32) *PlaceGetOverlayParams {
	o.SetHeight(height)
	return o
}

// SetHeight adds the height to the place get overlay params
func (o *PlaceGetOverlayParams) SetHeight(height int32) {
	o.Height = height
}

// WithLat adds the lat to the place get overlay params
func (o *PlaceGetOverlayParams) WithLat(lat string) *PlaceGetOverlayParams {
	o.SetLat(lat)
	return o
}

// SetLat adds the lat to the place get overlay params
func (o *PlaceGetOverlayParams) SetLat(lat string) {
	o.Lat = lat
}

// WithLocationLat adds the locationLat to the place get overlay params
func (o *PlaceGetOverlayParams) WithLocationLat(locationLat float64) *PlaceGetOverlayParams {
	o.SetLocationLat(locationLat)
	return o
}

// SetLocationLat adds the locationLat to the place get overlay params
func (o *PlaceGetOverlayParams) SetLocationLat(locationLat float64) {
	o.LocationLat = locationLat
}

// WithLocationLon adds the locationLon to the place get overlay params
func (o *PlaceGetOverlayParams) WithLocationLon(locationLon float64) *PlaceGetOverlayParams {
	o.SetLocationLon(locationLon)
	return o
}

// SetLocationLon adds the locationLon to the place get overlay params
func (o *PlaceGetOverlayParams) SetLocationLon(locationLon float64) {
	o.LocationLon = locationLon
}

// WithLon adds the lon to the place get overlay params
func (o *PlaceGetOverlayParams) WithLon(lon string) *PlaceGetOverlayParams {
	o.SetLon(lon)
	return o
}

// SetLon adds the lon to the place get overlay params
func (o *PlaceGetOverlayParams) SetLon(lon string) {
	o.Lon = lon
}

// WithType adds the typeVar to the place get overlay params
func (o *PlaceGetOverlayParams) WithType(typeVar []string) *PlaceGetOverlayParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the place get overlay params
func (o *PlaceGetOverlayParams) SetType(typeVar []string) {
	o.Type = typeVar
}

// WithWidth adds the width to the place get overlay params
func (o *PlaceGetOverlayParams) WithWidth(width int32) *PlaceGetOverlayParams {
	o.SetWidth(width)
	return o
}

// SetWidth adds the width to the place get overlay params
func (o *PlaceGetOverlayParams) SetWidth(width int32) {
	o.Width = width
}

// WithZ adds the z to the place get overlay params
func (o *PlaceGetOverlayParams) WithZ(z int32) *PlaceGetOverlayParams {
	o.SetZ(z)
	return o
}

// SetZ adds the z to the place get overlay params
func (o *PlaceGetOverlayParams) SetZ(z int32) {
	o.Z = z
}

// WriteToRequest writes these params to a swagger request
func (o *PlaceGetOverlayParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param height
	if err := r.SetPathParam("height", swag.FormatInt32(o.Height)); err != nil {
		return err
	}

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

	// path param width
	if err := r.SetPathParam("width", swag.FormatInt32(o.Width)); err != nil {
		return err
	}

	// path param z
	if err := r.SetPathParam("z", swag.FormatInt32(o.Z)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
