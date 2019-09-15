// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TflAPIPresentationEntitiesStreetSegment tfl Api presentation entities street segment
// swagger:model Tfl.Api.Presentation.Entities.StreetSegment
type TflAPIPresentationEntitiesStreetSegment struct {

	// geoJSON formatted LineString containing two latitude/longitude (WGS84) pairs that identify the start and end points of the street segment.
	LineString string `json:"lineString,omitempty"`

	// The ID from the source system of the disruption that this street belongs to.
	SourceSystemID int64 `json:"sourceSystemId,omitempty"`

	// The key of the source system of the disruption that this street belongs to.
	SourceSystemKey string `json:"sourceSystemKey,omitempty"`

	// A 16 digit unique integer identifying a OS ITN (Ordnance Survey Integrated Transport Network) road link.
	Toid string `json:"toid,omitempty"`
}

// Validate validates this tfl Api presentation entities street segment
func (m *TflAPIPresentationEntitiesStreetSegment) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesStreetSegment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesStreetSegment) UnmarshalBinary(b []byte) error {
	var res TflAPIPresentationEntitiesStreetSegment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}