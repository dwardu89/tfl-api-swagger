// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TflAPIPresentationEntitiesTrainLoading tfl Api presentation entities train loading
// swagger:model Tfl.Api.Presentation.Entities.TrainLoading
type TflAPIPresentationEntitiesTrainLoading struct {

	// Direction in regards to Journey Planner i.e. inbound or outbound
	Direction string `json:"direction,omitempty"`

	// The Line Name e.g. "Victoria"
	Line string `json:"line,omitempty"`

	// Direction of the Line e.g. NB, SB, WB etc.
	LineDirection string `json:"lineDirection,omitempty"`

	// Naptan of the adjacent station
	NaptanTo string `json:"naptanTo,omitempty"`

	// Direction displayed on the platform e.g. NB, SB, WB etc.
	PlatformDirection string `json:"platformDirection,omitempty"`

	// Time in 24hr format with 15 minute intervals e.g. 0500-0515, 0515-0530 etc.
	TimeSlice string `json:"timeSlice,omitempty"`

	// Scale between 1-6,
	//              1 = Very quiet, 2 = Quiet, 3 = Fairly busy, 4 = Busy, 5 = Very busy, 6 = Exceptionally busy
	Value int32 `json:"value,omitempty"`
}

// Validate validates this tfl Api presentation entities train loading
func (m *TflAPIPresentationEntitiesTrainLoading) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesTrainLoading) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesTrainLoading) UnmarshalBinary(b []byte) error {
	var res TflAPIPresentationEntitiesTrainLoading
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}