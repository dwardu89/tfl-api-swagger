// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TflAPIPresentationEntitiesTimetablesDisambiguationOption tfl Api presentation entities timetables disambiguation option
// swagger:model Tfl.Api.Presentation.Entities.Timetables.DisambiguationOption
type TflAPIPresentationEntitiesTimetablesDisambiguationOption struct {

	// description
	Description string `json:"description,omitempty"`

	// uri
	URI string `json:"uri,omitempty"`
}

// Validate validates this tfl Api presentation entities timetables disambiguation option
func (m *TflAPIPresentationEntitiesTimetablesDisambiguationOption) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesTimetablesDisambiguationOption) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesTimetablesDisambiguationOption) UnmarshalBinary(b []byte) error {
	var res TflAPIPresentationEntitiesTimetablesDisambiguationOption
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
