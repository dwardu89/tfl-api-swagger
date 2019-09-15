// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TflAPIPresentationEntitiesAccidentStatsVehicle tfl Api presentation entities accident stats vehicle
// swagger:model Tfl.Api.Presentation.Entities.AccidentStats.Vehicle
type TflAPIPresentationEntitiesAccidentStatsVehicle struct {

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this tfl Api presentation entities accident stats vehicle
func (m *TflAPIPresentationEntitiesAccidentStatsVehicle) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesAccidentStatsVehicle) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesAccidentStatsVehicle) UnmarshalBinary(b []byte) error {
	var res TflAPIPresentationEntitiesAccidentStatsVehicle
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}