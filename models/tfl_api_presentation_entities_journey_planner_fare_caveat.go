// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TflAPIPresentationEntitiesJourneyPlannerFareCaveat tfl Api presentation entities journey planner fare caveat
// swagger:model Tfl.Api.Presentation.Entities.JourneyPlanner.FareCaveat
type TflAPIPresentationEntitiesJourneyPlannerFareCaveat struct {

	// text
	Text string `json:"text,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this tfl Api presentation entities journey planner fare caveat
func (m *TflAPIPresentationEntitiesJourneyPlannerFareCaveat) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesJourneyPlannerFareCaveat) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesJourneyPlannerFareCaveat) UnmarshalBinary(b []byte) error {
	var res TflAPIPresentationEntitiesJourneyPlannerFareCaveat
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
