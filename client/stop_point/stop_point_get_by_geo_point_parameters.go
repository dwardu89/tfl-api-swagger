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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewStopPointGetByGeoPointParams creates a new StopPointGetByGeoPointParams object
// with the default values initialized.
func NewStopPointGetByGeoPointParams() *StopPointGetByGeoPointParams {
	var ()
	return &StopPointGetByGeoPointParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStopPointGetByGeoPointParamsWithTimeout creates a new StopPointGetByGeoPointParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStopPointGetByGeoPointParamsWithTimeout(timeout time.Duration) *StopPointGetByGeoPointParams {
	var ()
	return &StopPointGetByGeoPointParams{

		timeout: timeout,
	}
}

// NewStopPointGetByGeoPointParamsWithContext creates a new StopPointGetByGeoPointParams object
// with the default values initialized, and the ability to set a context for a request
func NewStopPointGetByGeoPointParamsWithContext(ctx context.Context) *StopPointGetByGeoPointParams {
	var ()
	return &StopPointGetByGeoPointParams{

		Context: ctx,
	}
}

// NewStopPointGetByGeoPointParamsWithHTTPClient creates a new StopPointGetByGeoPointParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStopPointGetByGeoPointParamsWithHTTPClient(client *http.Client) *StopPointGetByGeoPointParams {
	var ()
	return &StopPointGetByGeoPointParams{
		HTTPClient: client,
	}
}

/*StopPointGetByGeoPointParams contains all the parameters to send to the API endpoint
for the stop point get by geo point operation typically these are written to a http.Request
*/
type StopPointGetByGeoPointParams struct {

	/*Categories
	  an optional list of comma separated property categories to return in the StopPoint's property bag. If null or empty, all categories of property are returned. Pass the keyword "none" to return no properties (a valid list of categories can be obtained from the /StopPoint/Meta/categories endpoint)

	*/
	Categories []string
	/*LocationLat*/
	LocationLat float64
	/*LocationLon*/
	LocationLon float64
	/*Modes
	  the list of modes to search (comma separated mode names e.g. tube,dlr)

	*/
	Modes []string
	/*Radius
	  the radius of the bounding circle in metres (default : 200)

	*/
	Radius *int32
	/*ReturnLines
	  true to return the lines that each stop point serves as a nested resource

	*/
	ReturnLines *bool
	/*StopTypes
	  a list of stopTypes that should be returned (a list of valid stop types can be obtained from the StopPoint/meta/stoptypes endpoint)

	*/
	StopTypes []string
	/*UseStopPointHierarchy
	  Re-arrange the output into a parent/child hierarchy

	*/
	UseStopPointHierarchy *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the stop point get by geo point params
func (o *StopPointGetByGeoPointParams) WithTimeout(timeout time.Duration) *StopPointGetByGeoPointParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stop point get by geo point params
func (o *StopPointGetByGeoPointParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stop point get by geo point params
func (o *StopPointGetByGeoPointParams) WithContext(ctx context.Context) *StopPointGetByGeoPointParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stop point get by geo point params
func (o *StopPointGetByGeoPointParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stop point get by geo point params
func (o *StopPointGetByGeoPointParams) WithHTTPClient(client *http.Client) *StopPointGetByGeoPointParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stop point get by geo point params
func (o *StopPointGetByGeoPointParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCategories adds the categories to the stop point get by geo point params
func (o *StopPointGetByGeoPointParams) WithCategories(categories []string) *StopPointGetByGeoPointParams {
	o.SetCategories(categories)
	return o
}

// SetCategories adds the categories to the stop point get by geo point params
func (o *StopPointGetByGeoPointParams) SetCategories(categories []string) {
	o.Categories = categories
}

// WithLocationLat adds the locationLat to the stop point get by geo point params
func (o *StopPointGetByGeoPointParams) WithLocationLat(locationLat float64) *StopPointGetByGeoPointParams {
	o.SetLocationLat(locationLat)
	return o
}

// SetLocationLat adds the locationLat to the stop point get by geo point params
func (o *StopPointGetByGeoPointParams) SetLocationLat(locationLat float64) {
	o.LocationLat = locationLat
}

// WithLocationLon adds the locationLon to the stop point get by geo point params
func (o *StopPointGetByGeoPointParams) WithLocationLon(locationLon float64) *StopPointGetByGeoPointParams {
	o.SetLocationLon(locationLon)
	return o
}

// SetLocationLon adds the locationLon to the stop point get by geo point params
func (o *StopPointGetByGeoPointParams) SetLocationLon(locationLon float64) {
	o.LocationLon = locationLon
}

// WithModes adds the modes to the stop point get by geo point params
func (o *StopPointGetByGeoPointParams) WithModes(modes []string) *StopPointGetByGeoPointParams {
	o.SetModes(modes)
	return o
}

// SetModes adds the modes to the stop point get by geo point params
func (o *StopPointGetByGeoPointParams) SetModes(modes []string) {
	o.Modes = modes
}

// WithRadius adds the radius to the stop point get by geo point params
func (o *StopPointGetByGeoPointParams) WithRadius(radius *int32) *StopPointGetByGeoPointParams {
	o.SetRadius(radius)
	return o
}

// SetRadius adds the radius to the stop point get by geo point params
func (o *StopPointGetByGeoPointParams) SetRadius(radius *int32) {
	o.Radius = radius
}

// WithReturnLines adds the returnLines to the stop point get by geo point params
func (o *StopPointGetByGeoPointParams) WithReturnLines(returnLines *bool) *StopPointGetByGeoPointParams {
	o.SetReturnLines(returnLines)
	return o
}

// SetReturnLines adds the returnLines to the stop point get by geo point params
func (o *StopPointGetByGeoPointParams) SetReturnLines(returnLines *bool) {
	o.ReturnLines = returnLines
}

// WithStopTypes adds the stopTypes to the stop point get by geo point params
func (o *StopPointGetByGeoPointParams) WithStopTypes(stopTypes []string) *StopPointGetByGeoPointParams {
	o.SetStopTypes(stopTypes)
	return o
}

// SetStopTypes adds the stopTypes to the stop point get by geo point params
func (o *StopPointGetByGeoPointParams) SetStopTypes(stopTypes []string) {
	o.StopTypes = stopTypes
}

// WithUseStopPointHierarchy adds the useStopPointHierarchy to the stop point get by geo point params
func (o *StopPointGetByGeoPointParams) WithUseStopPointHierarchy(useStopPointHierarchy *bool) *StopPointGetByGeoPointParams {
	o.SetUseStopPointHierarchy(useStopPointHierarchy)
	return o
}

// SetUseStopPointHierarchy adds the useStopPointHierarchy to the stop point get by geo point params
func (o *StopPointGetByGeoPointParams) SetUseStopPointHierarchy(useStopPointHierarchy *bool) {
	o.UseStopPointHierarchy = useStopPointHierarchy
}

// WriteToRequest writes these params to a swagger request
func (o *StopPointGetByGeoPointParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesCategories := o.Categories

	joinedCategories := swag.JoinByFormat(valuesCategories, "multi")
	// query array param categories
	if err := r.SetQueryParam("categories", joinedCategories...); err != nil {
		return err
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

	valuesModes := o.Modes

	joinedModes := swag.JoinByFormat(valuesModes, "multi")
	// query array param modes
	if err := r.SetQueryParam("modes", joinedModes...); err != nil {
		return err
	}

	if o.Radius != nil {

		// query param radius
		var qrRadius int32
		if o.Radius != nil {
			qrRadius = *o.Radius
		}
		qRadius := swag.FormatInt32(qrRadius)
		if qRadius != "" {
			if err := r.SetQueryParam("radius", qRadius); err != nil {
				return err
			}
		}

	}

	if o.ReturnLines != nil {

		// query param returnLines
		var qrReturnLines bool
		if o.ReturnLines != nil {
			qrReturnLines = *o.ReturnLines
		}
		qReturnLines := swag.FormatBool(qrReturnLines)
		if qReturnLines != "" {
			if err := r.SetQueryParam("returnLines", qReturnLines); err != nil {
				return err
			}
		}

	}

	valuesStopTypes := o.StopTypes

	joinedStopTypes := swag.JoinByFormat(valuesStopTypes, "multi")
	// query array param stopTypes
	if err := r.SetQueryParam("stopTypes", joinedStopTypes...); err != nil {
		return err
	}

	if o.UseStopPointHierarchy != nil {

		// query param useStopPointHierarchy
		var qrUseStopPointHierarchy bool
		if o.UseStopPointHierarchy != nil {
			qrUseStopPointHierarchy = *o.UseStopPointHierarchy
		}
		qUseStopPointHierarchy := swag.FormatBool(qrUseStopPointHierarchy)
		if qUseStopPointHierarchy != "" {
			if err := r.SetQueryParam("useStopPointHierarchy", qUseStopPointHierarchy); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}