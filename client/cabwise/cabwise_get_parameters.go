// Code generated by go-swagger; DO NOT EDIT.

package cabwise

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

// NewCabwiseGetParams creates a new CabwiseGetParams object
// with the default values initialized.
func NewCabwiseGetParams() *CabwiseGetParams {
	var ()
	return &CabwiseGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCabwiseGetParamsWithTimeout creates a new CabwiseGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCabwiseGetParamsWithTimeout(timeout time.Duration) *CabwiseGetParams {
	var ()
	return &CabwiseGetParams{

		timeout: timeout,
	}
}

// NewCabwiseGetParamsWithContext creates a new CabwiseGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewCabwiseGetParamsWithContext(ctx context.Context) *CabwiseGetParams {
	var ()
	return &CabwiseGetParams{

		Context: ctx,
	}
}

// NewCabwiseGetParamsWithHTTPClient creates a new CabwiseGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCabwiseGetParamsWithHTTPClient(client *http.Client) *CabwiseGetParams {
	var ()
	return &CabwiseGetParams{
		HTTPClient: client,
	}
}

/*CabwiseGetParams contains all the parameters to send to the API endpoint
for the cabwise get operation typically these are written to a http.Request
*/
type CabwiseGetParams struct {

	/*ForceXML
	  Force Xml

	*/
	ForceXML *bool
	/*Lat
	  Latitude

	*/
	Lat float64
	/*LegacyFormat
	  Legacy Format

	*/
	LegacyFormat *bool
	/*Lon
	  Longitude

	*/
	Lon float64
	/*MaxResults
	  An optional parameter to limit the number of results return. Default and maximum is 20.

	*/
	MaxResults *int32
	/*Name
	  Trading name of operating company

	*/
	Name *string
	/*Optype
	  Operator Type e.g Minicab, Executive, Limousine

	*/
	Optype *string
	/*Radius
	  The radius of the bounding circle in metres

	*/
	Radius *float64
	/*TwentyFourSevenOnly
	  Twenty Four Seven Only

	*/
	TwentyFourSevenOnly *bool
	/*Wc
	  Wheelchair accessible

	*/
	Wc *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the cabwise get params
func (o *CabwiseGetParams) WithTimeout(timeout time.Duration) *CabwiseGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cabwise get params
func (o *CabwiseGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cabwise get params
func (o *CabwiseGetParams) WithContext(ctx context.Context) *CabwiseGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cabwise get params
func (o *CabwiseGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cabwise get params
func (o *CabwiseGetParams) WithHTTPClient(client *http.Client) *CabwiseGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cabwise get params
func (o *CabwiseGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithForceXML adds the forceXML to the cabwise get params
func (o *CabwiseGetParams) WithForceXML(forceXML *bool) *CabwiseGetParams {
	o.SetForceXML(forceXML)
	return o
}

// SetForceXML adds the forceXml to the cabwise get params
func (o *CabwiseGetParams) SetForceXML(forceXML *bool) {
	o.ForceXML = forceXML
}

// WithLat adds the lat to the cabwise get params
func (o *CabwiseGetParams) WithLat(lat float64) *CabwiseGetParams {
	o.SetLat(lat)
	return o
}

// SetLat adds the lat to the cabwise get params
func (o *CabwiseGetParams) SetLat(lat float64) {
	o.Lat = lat
}

// WithLegacyFormat adds the legacyFormat to the cabwise get params
func (o *CabwiseGetParams) WithLegacyFormat(legacyFormat *bool) *CabwiseGetParams {
	o.SetLegacyFormat(legacyFormat)
	return o
}

// SetLegacyFormat adds the legacyFormat to the cabwise get params
func (o *CabwiseGetParams) SetLegacyFormat(legacyFormat *bool) {
	o.LegacyFormat = legacyFormat
}

// WithLon adds the lon to the cabwise get params
func (o *CabwiseGetParams) WithLon(lon float64) *CabwiseGetParams {
	o.SetLon(lon)
	return o
}

// SetLon adds the lon to the cabwise get params
func (o *CabwiseGetParams) SetLon(lon float64) {
	o.Lon = lon
}

// WithMaxResults adds the maxResults to the cabwise get params
func (o *CabwiseGetParams) WithMaxResults(maxResults *int32) *CabwiseGetParams {
	o.SetMaxResults(maxResults)
	return o
}

// SetMaxResults adds the maxResults to the cabwise get params
func (o *CabwiseGetParams) SetMaxResults(maxResults *int32) {
	o.MaxResults = maxResults
}

// WithName adds the name to the cabwise get params
func (o *CabwiseGetParams) WithName(name *string) *CabwiseGetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the cabwise get params
func (o *CabwiseGetParams) SetName(name *string) {
	o.Name = name
}

// WithOptype adds the optype to the cabwise get params
func (o *CabwiseGetParams) WithOptype(optype *string) *CabwiseGetParams {
	o.SetOptype(optype)
	return o
}

// SetOptype adds the optype to the cabwise get params
func (o *CabwiseGetParams) SetOptype(optype *string) {
	o.Optype = optype
}

// WithRadius adds the radius to the cabwise get params
func (o *CabwiseGetParams) WithRadius(radius *float64) *CabwiseGetParams {
	o.SetRadius(radius)
	return o
}

// SetRadius adds the radius to the cabwise get params
func (o *CabwiseGetParams) SetRadius(radius *float64) {
	o.Radius = radius
}

// WithTwentyFourSevenOnly adds the twentyFourSevenOnly to the cabwise get params
func (o *CabwiseGetParams) WithTwentyFourSevenOnly(twentyFourSevenOnly *bool) *CabwiseGetParams {
	o.SetTwentyFourSevenOnly(twentyFourSevenOnly)
	return o
}

// SetTwentyFourSevenOnly adds the twentyFourSevenOnly to the cabwise get params
func (o *CabwiseGetParams) SetTwentyFourSevenOnly(twentyFourSevenOnly *bool) {
	o.TwentyFourSevenOnly = twentyFourSevenOnly
}

// WithWc adds the wc to the cabwise get params
func (o *CabwiseGetParams) WithWc(wc *string) *CabwiseGetParams {
	o.SetWc(wc)
	return o
}

// SetWc adds the wc to the cabwise get params
func (o *CabwiseGetParams) SetWc(wc *string) {
	o.Wc = wc
}

// WriteToRequest writes these params to a swagger request
func (o *CabwiseGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ForceXML != nil {

		// query param forceXml
		var qrForceXML bool
		if o.ForceXML != nil {
			qrForceXML = *o.ForceXML
		}
		qForceXML := swag.FormatBool(qrForceXML)
		if qForceXML != "" {
			if err := r.SetQueryParam("forceXml", qForceXML); err != nil {
				return err
			}
		}

	}

	// query param lat
	qrLat := o.Lat
	qLat := swag.FormatFloat64(qrLat)
	if qLat != "" {
		if err := r.SetQueryParam("lat", qLat); err != nil {
			return err
		}
	}

	if o.LegacyFormat != nil {

		// query param legacyFormat
		var qrLegacyFormat bool
		if o.LegacyFormat != nil {
			qrLegacyFormat = *o.LegacyFormat
		}
		qLegacyFormat := swag.FormatBool(qrLegacyFormat)
		if qLegacyFormat != "" {
			if err := r.SetQueryParam("legacyFormat", qLegacyFormat); err != nil {
				return err
			}
		}

	}

	// query param lon
	qrLon := o.Lon
	qLon := swag.FormatFloat64(qrLon)
	if qLon != "" {
		if err := r.SetQueryParam("lon", qLon); err != nil {
			return err
		}
	}

	if o.MaxResults != nil {

		// query param maxResults
		var qrMaxResults int32
		if o.MaxResults != nil {
			qrMaxResults = *o.MaxResults
		}
		qMaxResults := swag.FormatInt32(qrMaxResults)
		if qMaxResults != "" {
			if err := r.SetQueryParam("maxResults", qMaxResults); err != nil {
				return err
			}
		}

	}

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}

	}

	if o.Optype != nil {

		// query param optype
		var qrOptype string
		if o.Optype != nil {
			qrOptype = *o.Optype
		}
		qOptype := qrOptype
		if qOptype != "" {
			if err := r.SetQueryParam("optype", qOptype); err != nil {
				return err
			}
		}

	}

	if o.Radius != nil {

		// query param radius
		var qrRadius float64
		if o.Radius != nil {
			qrRadius = *o.Radius
		}
		qRadius := swag.FormatFloat64(qrRadius)
		if qRadius != "" {
			if err := r.SetQueryParam("radius", qRadius); err != nil {
				return err
			}
		}

	}

	if o.TwentyFourSevenOnly != nil {

		// query param twentyFourSevenOnly
		var qrTwentyFourSevenOnly bool
		if o.TwentyFourSevenOnly != nil {
			qrTwentyFourSevenOnly = *o.TwentyFourSevenOnly
		}
		qTwentyFourSevenOnly := swag.FormatBool(qrTwentyFourSevenOnly)
		if qTwentyFourSevenOnly != "" {
			if err := r.SetQueryParam("twentyFourSevenOnly", qTwentyFourSevenOnly); err != nil {
				return err
			}
		}

	}

	if o.Wc != nil {

		// query param wc
		var qrWc string
		if o.Wc != nil {
			qrWc = *o.Wc
		}
		qWc := qrWc
		if qWc != "" {
			if err := r.SetQueryParam("wc", qWc); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
