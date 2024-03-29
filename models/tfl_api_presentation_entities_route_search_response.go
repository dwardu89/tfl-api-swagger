// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TflAPIPresentationEntitiesRouteSearchResponse tfl Api presentation entities route search response
// swagger:model Tfl.Api.Presentation.Entities.RouteSearchResponse
type TflAPIPresentationEntitiesRouteSearchResponse struct {

	// input
	Input string `json:"input,omitempty"`

	// search matches
	SearchMatches []*TflAPIPresentationEntitiesRouteSearchMatch `json:"searchMatches"`
}

// Validate validates this tfl Api presentation entities route search response
func (m *TflAPIPresentationEntitiesRouteSearchResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSearchMatches(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TflAPIPresentationEntitiesRouteSearchResponse) validateSearchMatches(formats strfmt.Registry) error {

	if swag.IsZero(m.SearchMatches) { // not required
		return nil
	}

	for i := 0; i < len(m.SearchMatches); i++ {
		if swag.IsZero(m.SearchMatches[i]) { // not required
			continue
		}

		if m.SearchMatches[i] != nil {
			if err := m.SearchMatches[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("searchMatches" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesRouteSearchResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesRouteSearchResponse) UnmarshalBinary(b []byte) error {
	var res TflAPIPresentationEntitiesRouteSearchResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
