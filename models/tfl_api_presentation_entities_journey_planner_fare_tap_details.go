// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TflAPIPresentationEntitiesJourneyPlannerFareTapDetails tfl Api presentation entities journey planner fare tap details
// swagger:model Tfl.Api.Presentation.Entities.JourneyPlanner.FareTapDetails
type TflAPIPresentationEntitiesJourneyPlannerFareTapDetails struct {

	// bus route Id
	BusRouteID string `json:"busRouteId,omitempty"`

	// host device type
	HostDeviceType string `json:"hostDeviceType,omitempty"`

	// mode type
	ModeType string `json:"modeType,omitempty"`

	// national location code
	NationalLocationCode int32 `json:"nationalLocationCode,omitempty"`

	// tap timestamp
	// Format: date-time
	TapTimestamp strfmt.DateTime `json:"tapTimestamp,omitempty"`

	// validation type
	ValidationType string `json:"validationType,omitempty"`
}

// Validate validates this tfl Api presentation entities journey planner fare tap details
func (m *TflAPIPresentationEntitiesJourneyPlannerFareTapDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTapTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TflAPIPresentationEntitiesJourneyPlannerFareTapDetails) validateTapTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(m.TapTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("tapTimestamp", "body", "date-time", m.TapTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesJourneyPlannerFareTapDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesJourneyPlannerFareTapDetails) UnmarshalBinary(b []byte) error {
	var res TflAPIPresentationEntitiesJourneyPlannerFareTapDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
