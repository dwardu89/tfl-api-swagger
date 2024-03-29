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

// TflAPIPresentationEntitiesSearchResponse tfl Api presentation entities search response
// swagger:model Tfl.Api.Presentation.Entities.SearchResponse
type TflAPIPresentationEntitiesSearchResponse struct {

	// from
	From int32 `json:"from,omitempty"`

	// matches
	Matches []*TflAPIPresentationEntitiesSearchMatch `json:"matches"`

	// max score
	MaxScore float64 `json:"maxScore,omitempty"`

	// page
	Page int32 `json:"page,omitempty"`

	// page size
	PageSize int32 `json:"pageSize,omitempty"`

	// provider
	Provider string `json:"provider,omitempty"`

	// query
	Query string `json:"query,omitempty"`

	// total
	Total int32 `json:"total,omitempty"`
}

// Validate validates this tfl Api presentation entities search response
func (m *TflAPIPresentationEntitiesSearchResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMatches(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TflAPIPresentationEntitiesSearchResponse) validateMatches(formats strfmt.Registry) error {

	if swag.IsZero(m.Matches) { // not required
		return nil
	}

	for i := 0; i < len(m.Matches); i++ {
		if swag.IsZero(m.Matches[i]) { // not required
			continue
		}

		if m.Matches[i] != nil {
			if err := m.Matches[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("matches" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesSearchResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesSearchResponse) UnmarshalBinary(b []byte) error {
	var res TflAPIPresentationEntitiesSearchResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
