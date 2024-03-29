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

// TflAPIPresentationEntitiesStreet tfl Api presentation entities street
// swagger:model Tfl.Api.Presentation.Entities.Street
type TflAPIPresentationEntitiesStreet struct {

	// Type of road closure. Some example values:
	//             Open = road is open, not blocked, not closed, not restricted. It maybe that the disruption has been moved out of the carriageway.
	//             Partial Closure = road is partially blocked, closed or restricted.
	//             Full Closure = road is fully blocked or closed.
	Closure string `json:"closure,omitempty"`

	// The direction of the disruption on the street. Some example values:
	//             All Directions
	//             All Approaches
	//             Clockwise
	//             Anti-Clockwise
	//             Northbound
	//             Eastbound
	//             Southbound
	//             Westbound
	//             Both Directions
	Directions string `json:"directions,omitempty"`

	// Street name
	Name string `json:"name,omitempty"`

	// Geographic description of the sections of this street that are affected.
	Segments []*TflAPIPresentationEntitiesStreetSegment `json:"segments"`

	// The ID from the source system of the disruption that this street belongs to.
	SourceSystemID int64 `json:"sourceSystemId,omitempty"`

	// The key of the source system of the disruption that this street belongs to.
	SourceSystemKey string `json:"sourceSystemKey,omitempty"`
}

// Validate validates this tfl Api presentation entities street
func (m *TflAPIPresentationEntitiesStreet) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSegments(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TflAPIPresentationEntitiesStreet) validateSegments(formats strfmt.Registry) error {

	if swag.IsZero(m.Segments) { // not required
		return nil
	}

	for i := 0; i < len(m.Segments); i++ {
		if swag.IsZero(m.Segments[i]) { // not required
			continue
		}

		if m.Segments[i] != nil {
			if err := m.Segments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("segments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesStreet) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesStreet) UnmarshalBinary(b []byte) error {
	var res TflAPIPresentationEntitiesStreet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
