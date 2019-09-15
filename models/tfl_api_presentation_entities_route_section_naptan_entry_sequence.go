// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TflAPIPresentationEntitiesRouteSectionNaptanEntrySequence tfl Api presentation entities route section naptan entry sequence
// swagger:model Tfl.Api.Presentation.Entities.RouteSectionNaptanEntrySequence
type TflAPIPresentationEntitiesRouteSectionNaptanEntrySequence struct {

	// ordinal
	Ordinal int32 `json:"ordinal,omitempty"`

	// stop point
	StopPoint *TflAPIPresentationEntitiesStopPoint `json:"stopPoint,omitempty"`
}

// Validate validates this tfl Api presentation entities route section naptan entry sequence
func (m *TflAPIPresentationEntitiesRouteSectionNaptanEntrySequence) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStopPoint(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TflAPIPresentationEntitiesRouteSectionNaptanEntrySequence) validateStopPoint(formats strfmt.Registry) error {

	if swag.IsZero(m.StopPoint) { // not required
		return nil
	}

	if m.StopPoint != nil {
		if err := m.StopPoint.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stopPoint")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesRouteSectionNaptanEntrySequence) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesRouteSectionNaptanEntrySequence) UnmarshalBinary(b []byte) error {
	var res TflAPIPresentationEntitiesRouteSectionNaptanEntrySequence
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}