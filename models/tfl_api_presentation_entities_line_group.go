// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TflAPIPresentationEntitiesLineGroup tfl Api presentation entities line group
// swagger:model Tfl.Api.Presentation.Entities.LineGroup
type TflAPIPresentationEntitiesLineGroup struct {

	// line identifier
	LineIdentifier []string `json:"lineIdentifier"`

	// naptan Id reference
	NaptanIDReference string `json:"naptanIdReference,omitempty"`

	// station atco code
	StationAtcoCode string `json:"stationAtcoCode,omitempty"`
}

// Validate validates this tfl Api presentation entities line group
func (m *TflAPIPresentationEntitiesLineGroup) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesLineGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesLineGroup) UnmarshalBinary(b []byte) error {
	var res TflAPIPresentationEntitiesLineGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}