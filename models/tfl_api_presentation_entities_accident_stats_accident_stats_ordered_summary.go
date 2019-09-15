// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TflAPIPresentationEntitiesAccidentStatsAccidentStatsOrderedSummary tfl Api presentation entities accident stats accident stats ordered summary
// swagger:model Tfl.Api.Presentation.Entities.AccidentStats.AccidentStatsOrderedSummary
type TflAPIPresentationEntitiesAccidentStatsAccidentStatsOrderedSummary struct {

	// accidents
	Accidents int32 `json:"accidents,omitempty"`

	// borough
	Borough string `json:"borough,omitempty"`

	// year
	Year int32 `json:"year,omitempty"`
}

// Validate validates this tfl Api presentation entities accident stats accident stats ordered summary
func (m *TflAPIPresentationEntitiesAccidentStatsAccidentStatsOrderedSummary) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesAccidentStatsAccidentStatsOrderedSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesAccidentStatsAccidentStatsOrderedSummary) UnmarshalBinary(b []byte) error {
	var res TflAPIPresentationEntitiesAccidentStatsAccidentStatsOrderedSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}