// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TflAPIPresentationEntitiesFaresPassengerType tfl Api presentation entities fares passenger type
// swagger:model Tfl.Api.Presentation.Entities.Fares.PassengerType
type TflAPIPresentationEntitiesFaresPassengerType struct {

	// description
	Description string `json:"description,omitempty"`

	// display name
	DisplayName string `json:"displayName,omitempty"`

	// display order
	DisplayOrder int32 `json:"displayOrder,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this tfl Api presentation entities fares passenger type
func (m *TflAPIPresentationEntitiesFaresPassengerType) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesFaresPassengerType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesFaresPassengerType) UnmarshalBinary(b []byte) error {
	var res TflAPIPresentationEntitiesFaresPassengerType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}