// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TflAPIPresentationEntitiesFaresTicketType tfl Api presentation entities fares ticket type
// swagger:model Tfl.Api.Presentation.Entities.Fares.TicketType
type TflAPIPresentationEntitiesFaresTicketType struct {

	// description
	Description string `json:"description,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this tfl Api presentation entities fares ticket type
func (m *TflAPIPresentationEntitiesFaresTicketType) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesFaresTicketType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesFaresTicketType) UnmarshalBinary(b []byte) error {
	var res TflAPIPresentationEntitiesFaresTicketType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
