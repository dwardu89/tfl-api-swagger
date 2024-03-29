// Code generated by go-swagger; DO NOT EDIT.

package line

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new line API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for line API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
LineArrivals gets the list of arrival predictions for given line ids based at the given stop
*/
func (a *Client) LineArrivals(params *LineArrivalsParams) (*LineArrivalsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLineArrivalsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Line_Arrivals",
		Method:             "GET",
		PathPattern:        "/Line/{ids}/Arrivals/{stopPointId}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/json", "text/xml"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LineArrivalsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LineArrivalsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Line_Arrivals: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
LineDisruption gets disruptions for the given line ids
*/
func (a *Client) LineDisruption(params *LineDisruptionParams) (*LineDisruptionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLineDisruptionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Line_Disruption",
		Method:             "GET",
		PathPattern:        "/Line/{ids}/Disruption",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/json", "text/xml"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LineDisruptionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LineDisruptionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Line_Disruption: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
LineDisruptionByMode gets disruptions for all lines of the given modes
*/
func (a *Client) LineDisruptionByMode(params *LineDisruptionByModeParams) (*LineDisruptionByModeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLineDisruptionByModeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Line_DisruptionByMode",
		Method:             "GET",
		PathPattern:        "/Line/Mode/{modes}/Disruption",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/json", "text/xml"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LineDisruptionByModeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LineDisruptionByModeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Line_DisruptionByMode: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
LineGet gets lines that match the specified line ids
*/
func (a *Client) LineGet(params *LineGetParams) (*LineGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLineGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Line_Get",
		Method:             "GET",
		PathPattern:        "/Line/{ids}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/json", "text/xml"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LineGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LineGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Line_Get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
LineGetByMode gets lines that serve the given modes
*/
func (a *Client) LineGetByMode(params *LineGetByModeParams) (*LineGetByModeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLineGetByModeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Line_GetByMode",
		Method:             "GET",
		PathPattern:        "/Line/Mode/{modes}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/json", "text/xml"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LineGetByModeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LineGetByModeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Line_GetByMode: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
LineLineRoutesByIds gets all valid routes for given line ids including the name and id of the originating and terminating stops for each route
*/
func (a *Client) LineLineRoutesByIds(params *LineLineRoutesByIdsParams) (*LineLineRoutesByIdsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLineLineRoutesByIdsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Line_LineRoutesByIds",
		Method:             "GET",
		PathPattern:        "/Line/{ids}/Route",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/json", "text/xml"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LineLineRoutesByIdsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LineLineRoutesByIdsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Line_LineRoutesByIds: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
LineMetaDisruptionCategories gets a list of valid disruption categories
*/
func (a *Client) LineMetaDisruptionCategories(params *LineMetaDisruptionCategoriesParams) (*LineMetaDisruptionCategoriesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLineMetaDisruptionCategoriesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Line_MetaDisruptionCategories",
		Method:             "GET",
		PathPattern:        "/Line/Meta/DisruptionCategories",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/json", "text/xml"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LineMetaDisruptionCategoriesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LineMetaDisruptionCategoriesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Line_MetaDisruptionCategories: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
LineMetaModes gets a list of valid modes
*/
func (a *Client) LineMetaModes(params *LineMetaModesParams) (*LineMetaModesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLineMetaModesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Line_MetaModes",
		Method:             "GET",
		PathPattern:        "/Line/Meta/Modes",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/json", "text/xml"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LineMetaModesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LineMetaModesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Line_MetaModes: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
LineMetaServiceTypes gets a list of valid service types to filter on
*/
func (a *Client) LineMetaServiceTypes(params *LineMetaServiceTypesParams) (*LineMetaServiceTypesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLineMetaServiceTypesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Line_MetaServiceTypes",
		Method:             "GET",
		PathPattern:        "/Line/Meta/ServiceTypes",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/json", "text/xml"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LineMetaServiceTypesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LineMetaServiceTypesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Line_MetaServiceTypes: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
LineMetaSeverity gets a list of valid severity codes
*/
func (a *Client) LineMetaSeverity(params *LineMetaSeverityParams) (*LineMetaSeverityOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLineMetaSeverityParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Line_MetaSeverity",
		Method:             "GET",
		PathPattern:        "/Line/Meta/Severity",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/json", "text/xml"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LineMetaSeverityReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LineMetaSeverityOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Line_MetaSeverity: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
LineRoute gets all valid routes for all lines including the name and id of the originating and terminating stops for each route
*/
func (a *Client) LineRoute(params *LineRouteParams) (*LineRouteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLineRouteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Line_Route",
		Method:             "GET",
		PathPattern:        "/Line/Route",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/json", "text/xml"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LineRouteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LineRouteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Line_Route: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
LineRouteByMode gets all lines and their valid routes for given modes including the name and id of the originating and terminating stops for each route
*/
func (a *Client) LineRouteByMode(params *LineRouteByModeParams) (*LineRouteByModeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLineRouteByModeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Line_RouteByMode",
		Method:             "GET",
		PathPattern:        "/Line/Mode/{modes}/Route",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/json", "text/xml"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LineRouteByModeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LineRouteByModeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Line_RouteByMode: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
LineRouteSequence gets all valid routes for given line id including the sequence of stops on each route
*/
func (a *Client) LineRouteSequence(params *LineRouteSequenceParams) (*LineRouteSequenceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLineRouteSequenceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Line_RouteSequence",
		Method:             "GET",
		PathPattern:        "/Line/{id}/Route/Sequence/{direction}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/json", "text/xml"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LineRouteSequenceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LineRouteSequenceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Line_RouteSequence: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
LineSearch searches for lines or routes matching the query string
*/
func (a *Client) LineSearch(params *LineSearchParams) (*LineSearchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLineSearchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Line_Search",
		Method:             "GET",
		PathPattern:        "/Line/Search/{query}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/json", "text/xml"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LineSearchReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LineSearchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Line_Search: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
LineStatus gets the line status for given line ids during the provided dates e g minor delays
*/
func (a *Client) LineStatus(params *LineStatusParams) (*LineStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLineStatusParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Line_Status",
		Method:             "GET",
		PathPattern:        "/Line/{ids}/Status/{StartDate}/to/{EndDate}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/json", "text/xml"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LineStatusReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LineStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Line_Status: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
LineStatusByIds gets the line status of for given line ids e g minor delays
*/
func (a *Client) LineStatusByIds(params *LineStatusByIdsParams) (*LineStatusByIdsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLineStatusByIdsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Line_StatusByIds",
		Method:             "GET",
		PathPattern:        "/Line/{ids}/Status",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/json", "text/xml"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LineStatusByIdsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LineStatusByIdsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Line_StatusByIds: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
LineStatusByMode gets the line status of for all lines for the given modes
*/
func (a *Client) LineStatusByMode(params *LineStatusByModeParams) (*LineStatusByModeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLineStatusByModeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Line_StatusByMode",
		Method:             "GET",
		PathPattern:        "/Line/Mode/{modes}/Status",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/json", "text/xml"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LineStatusByModeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LineStatusByModeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Line_StatusByMode: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
LineStatusBySeverity gets the line status for all lines with a given severity a list of valid severity codes can be obtained from a call to line meta severity
*/
func (a *Client) LineStatusBySeverity(params *LineStatusBySeverityParams) (*LineStatusBySeverityOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLineStatusBySeverityParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Line_StatusBySeverity",
		Method:             "GET",
		PathPattern:        "/Line/Status/{severity}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/json", "text/xml"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LineStatusBySeverityReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LineStatusBySeverityOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Line_StatusBySeverity: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
LineStopPoints gets a list of the stations that serve the given line id
*/
func (a *Client) LineStopPoints(params *LineStopPointsParams) (*LineStopPointsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLineStopPointsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Line_StopPoints",
		Method:             "GET",
		PathPattern:        "/Line/{id}/StopPoints",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/json", "text/xml"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LineStopPointsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LineStopPointsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Line_StopPoints: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
LineTimetable gets the timetable for a specified station on the give line
*/
func (a *Client) LineTimetable(params *LineTimetableParams) (*LineTimetableOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLineTimetableParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Line_Timetable",
		Method:             "GET",
		PathPattern:        "/Line/{id}/Timetable/{fromStopPointId}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/json", "text/xml"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LineTimetableReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LineTimetableOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Line_Timetable: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
LineTimetableTo gets the timetable for a specified station on the give line with specified destination
*/
func (a *Client) LineTimetableTo(params *LineTimetableToParams) (*LineTimetableToOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLineTimetableToParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Line_TimetableTo",
		Method:             "GET",
		PathPattern:        "/Line/{id}/Timetable/{fromStopPointId}/to/{toStopPointId}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/json", "text/xml"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LineTimetableToReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LineTimetableToOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Line_TimetableTo: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
