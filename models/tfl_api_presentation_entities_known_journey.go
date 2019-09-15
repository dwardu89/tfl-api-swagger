// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TflAPIPresentationEntitiesKnownJourney tfl Api presentation entities known journey
// swagger:model Tfl.Api.Presentation.Entities.KnownJourney
type TflAPIPresentationEntitiesKnownJourney struct {

	// hour
	Hour string `json:"hour,omitempty"`

	// interval Id
	IntervalID int32 `json:"intervalId,omitempty"`

	// minute
	Minute string `json:"minute,omitempty"`
}

// Validate validates this tfl Api presentation entities known journey
func (m *TflAPIPresentationEntitiesKnownJourney) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesKnownJourney) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesKnownJourney) UnmarshalBinary(b []byte) error {
	var res TflAPIPresentationEntitiesKnownJourney
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}