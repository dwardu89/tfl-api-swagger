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

// TflAPIPresentationEntitiesLineServiceType tfl Api presentation entities line service type
// swagger:model Tfl.Api.Presentation.Entities.LineServiceType
type TflAPIPresentationEntitiesLineServiceType struct {

	// line name
	LineName string `json:"lineName,omitempty"`

	// line specific service types
	LineSpecificServiceTypes []*TflAPIPresentationEntitiesLineSpecificServiceType `json:"lineSpecificServiceTypes"`
}

// Validate validates this tfl Api presentation entities line service type
func (m *TflAPIPresentationEntitiesLineServiceType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLineSpecificServiceTypes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TflAPIPresentationEntitiesLineServiceType) validateLineSpecificServiceTypes(formats strfmt.Registry) error {

	if swag.IsZero(m.LineSpecificServiceTypes) { // not required
		return nil
	}

	for i := 0; i < len(m.LineSpecificServiceTypes); i++ {
		if swag.IsZero(m.LineSpecificServiceTypes[i]) { // not required
			continue
		}

		if m.LineSpecificServiceTypes[i] != nil {
			if err := m.LineSpecificServiceTypes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("lineSpecificServiceTypes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesLineServiceType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesLineServiceType) UnmarshalBinary(b []byte) error {
	var res TflAPIPresentationEntitiesLineServiceType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
