// Code generated by go-swagger; DO NOT EDIT.

package road

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new road API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for road API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
RoadDisruptedStreets gets a list of disrupted streets if no date filters are provided current disruptions are returned
*/
func (a *Client) RoadDisruptedStreets(params *RoadDisruptedStreetsParams) (*RoadDisruptedStreetsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRoadDisruptedStreetsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Road_DisruptedStreets",
		Method:             "GET",
		PathPattern:        "/Road/all/Street/Disruption",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/json", "text/xml"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RoadDisruptedStreetsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RoadDisruptedStreetsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Road_DisruptedStreets: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RoadDisruption gets active disruptions filtered by road ids
*/
func (a *Client) RoadDisruption(params *RoadDisruptionParams) (*RoadDisruptionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRoadDisruptionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Road_Disruption",
		Method:             "GET",
		PathPattern:        "/Road/{ids}/Disruption",
		ProducesMediaTypes: []string{"application/geo+json", "application/json", "application/xml", "text/json", "text/xml"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RoadDisruptionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RoadDisruptionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Road_Disruption: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RoadDisruptionByID gets a list of active disruptions filtered by disruption ids
*/
func (a *Client) RoadDisruptionByID(params *RoadDisruptionByIDParams) (*RoadDisruptionByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRoadDisruptionByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Road_DisruptionById",
		Method:             "GET",
		PathPattern:        "/Road/all/Disruption/{disruptionIds}",
		ProducesMediaTypes: []string{"application/geo+json", "application/json", "application/xml", "text/json", "text/xml"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RoadDisruptionByIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RoadDisruptionByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Road_DisruptionById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RoadGet gets the road with the specified id e g a1
*/
func (a *Client) RoadGet(params *RoadGetParams) (*RoadGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRoadGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Road_Get",
		Method:             "GET",
		PathPattern:        "/Road/{ids}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/json", "text/xml"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RoadGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RoadGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Road_Get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RoadMetaCategories gets a list of valid road disruption categories
*/
func (a *Client) RoadMetaCategories(params *RoadMetaCategoriesParams) (*RoadMetaCategoriesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRoadMetaCategoriesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Road_MetaCategories",
		Method:             "GET",
		PathPattern:        "/Road/Meta/Categories",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/json", "text/xml"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RoadMetaCategoriesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RoadMetaCategoriesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Road_MetaCategories: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RoadMetaSeverities gets a list of valid road disruption severity codes
*/
func (a *Client) RoadMetaSeverities(params *RoadMetaSeveritiesParams) (*RoadMetaSeveritiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRoadMetaSeveritiesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Road_MetaSeverities",
		Method:             "GET",
		PathPattern:        "/Road/Meta/Severities",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/json", "text/xml"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RoadMetaSeveritiesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RoadMetaSeveritiesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Road_MetaSeverities: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RoadStatus gets the specified roads with the status aggregated over the date range specified or now until the end of today if no dates are passed
*/
func (a *Client) RoadStatus(params *RoadStatusParams) (*RoadStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRoadStatusParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Road_Status",
		Method:             "GET",
		PathPattern:        "/Road/{ids}/Status",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/json", "text/xml"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RoadStatusReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RoadStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Road_Status: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
