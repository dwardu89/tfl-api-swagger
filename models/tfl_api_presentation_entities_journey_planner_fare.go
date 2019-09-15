// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TflAPIPresentationEntitiesJourneyPlannerFare tfl Api presentation entities journey planner fare
// swagger:model Tfl.Api.Presentation.Entities.JourneyPlanner.Fare
type TflAPIPresentationEntitiesJourneyPlannerFare struct {

	// charge level
	ChargeLevel string `json:"chargeLevel,omitempty"`

	// charge profile name
	ChargeProfileName string `json:"chargeProfileName,omitempty"`

	// cost
	Cost int32 `json:"cost,omitempty"`

	// high zone
	HighZone int32 `json:"highZone,omitempty"`

	// is hopper fare
	IsHopperFare bool `json:"isHopperFare,omitempty"`

	// low zone
	LowZone int32 `json:"lowZone,omitempty"`

	// off peak
	OffPeak int32 `json:"offPeak,omitempty"`

	// peak
	Peak int32 `json:"peak,omitempty"`

	// taps
	Taps []*TflAPIPresentationEntitiesJourneyPlannerFareTap `json:"taps"`
}

// Validate validates this tfl Api presentation entities journey planner fare
func (m *TflAPIPresentationEntitiesJourneyPlannerFare) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTaps(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TflAPIPresentationEntitiesJourneyPlannerFare) validateTaps(formats strfmt.Registry) error {

	if swag.IsZero(m.Taps) { // not required
		return nil
	}

	for i := 0; i < len(m.Taps); i++ {
		if swag.IsZero(m.Taps[i]) { // not required
			continue
		}

		if m.Taps[i] != nil {
			if err := m.Taps[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("taps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesJourneyPlannerFare) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesJourneyPlannerFare) UnmarshalBinary(b []byte) error {
	var res TflAPIPresentationEntitiesJourneyPlannerFare
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
