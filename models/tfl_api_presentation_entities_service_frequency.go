// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TflAPIPresentationEntitiesServiceFrequency tfl Api presentation entities service frequency
// swagger:model Tfl.Api.Presentation.Entities.ServiceFrequency
type TflAPIPresentationEntitiesServiceFrequency struct {

	// highest frequency
	HighestFrequency float64 `json:"highestFrequency,omitempty"`

	// lowest frequency
	LowestFrequency float64 `json:"lowestFrequency,omitempty"`
}

// Validate validates this tfl Api presentation entities service frequency
func (m *TflAPIPresentationEntitiesServiceFrequency) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesServiceFrequency) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesServiceFrequency) UnmarshalBinary(b []byte) error {
	var res TflAPIPresentationEntitiesServiceFrequency
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}