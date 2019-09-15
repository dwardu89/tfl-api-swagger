// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TflAPIPresentationEntitiesMode tfl Api presentation entities mode
// swagger:model Tfl.Api.Presentation.Entities.Mode
type TflAPIPresentationEntitiesMode struct {

	// is fare paying
	IsFarePaying bool `json:"isFarePaying,omitempty"`

	// is scheduled service
	IsScheduledService bool `json:"isScheduledService,omitempty"`

	// is tfl service
	IsTflService bool `json:"isTflService,omitempty"`

	// mode name
	ModeName string `json:"modeName,omitempty"`
}

// Validate validates this tfl Api presentation entities mode
func (m *TflAPIPresentationEntitiesMode) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesMode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesMode) UnmarshalBinary(b []byte) error {
	var res TflAPIPresentationEntitiesMode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
