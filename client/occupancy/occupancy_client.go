// Code generated by go-swagger; DO NOT EDIT.

package occupancy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new occupancy API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for occupancy API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
OccupancyGet gets the occupancy for a car park with a given id
*/
func (a *Client) OccupancyGet(params *OccupancyGetParams) (*OccupancyGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOccupancyGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Occupancy_Get",
		Method:             "GET",
		PathPattern:        "/Occupancy/CarPark/{id}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/json", "text/xml"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &OccupancyGetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*OccupancyGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Occupancy_Get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
OccupancyGetAllChargeConnectorStatus gets the occupancy for all charge connectors
*/
func (a *Client) OccupancyGetAllChargeConnectorStatus(params *OccupancyGetAllChargeConnectorStatusParams) (*OccupancyGetAllChargeConnectorStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOccupancyGetAllChargeConnectorStatusParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Occupancy_GetAllChargeConnectorStatus",
		Method:             "GET",
		PathPattern:        "/Occupancy/ChargeConnector",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/json", "text/xml"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &OccupancyGetAllChargeConnectorStatusReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*OccupancyGetAllChargeConnectorStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Occupancy_GetAllChargeConnectorStatus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
OccupancyGetBikePointsOccupancies gets the occupancy for bike points
*/
func (a *Client) OccupancyGetBikePointsOccupancies(params *OccupancyGetBikePointsOccupanciesParams) (*OccupancyGetBikePointsOccupanciesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOccupancyGetBikePointsOccupanciesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Occupancy_GetBikePointsOccupancies",
		Method:             "GET",
		PathPattern:        "/Occupancy/BikePoints/{ids}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/json", "text/xml"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &OccupancyGetBikePointsOccupanciesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*OccupancyGetBikePointsOccupanciesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Occupancy_GetBikePointsOccupancies: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
OccupancyGetChargeConnectorStatus gets the occupancy for a charge connectors with a given id source system place Id
*/
func (a *Client) OccupancyGetChargeConnectorStatus(params *OccupancyGetChargeConnectorStatusParams) (*OccupancyGetChargeConnectorStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOccupancyGetChargeConnectorStatusParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Occupancy_GetChargeConnectorStatus",
		Method:             "GET",
		PathPattern:        "/Occupancy/ChargeConnector/{ids}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/json", "text/xml"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &OccupancyGetChargeConnectorStatusReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*OccupancyGetChargeConnectorStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Occupancy_GetChargeConnectorStatus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
