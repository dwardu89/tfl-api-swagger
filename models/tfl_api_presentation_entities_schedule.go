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

// TflAPIPresentationEntitiesSchedule tfl Api presentation entities schedule
// swagger:model Tfl.Api.Presentation.Entities.Schedule
type TflAPIPresentationEntitiesSchedule struct {

	// first journey
	FirstJourney *TflAPIPresentationEntitiesKnownJourney `json:"firstJourney,omitempty"`

	// known journeys
	KnownJourneys []*TflAPIPresentationEntitiesKnownJourney `json:"knownJourneys"`

	// last journey
	LastJourney *TflAPIPresentationEntitiesKnownJourney `json:"lastJourney,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// periods
	Periods []*TflAPIPresentationEntitiesPeriod `json:"periods"`
}

// Validate validates this tfl Api presentation entities schedule
func (m *TflAPIPresentationEntitiesSchedule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFirstJourney(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKnownJourneys(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastJourney(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePeriods(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TflAPIPresentationEntitiesSchedule) validateFirstJourney(formats strfmt.Registry) error {

	if swag.IsZero(m.FirstJourney) { // not required
		return nil
	}

	if m.FirstJourney != nil {
		if err := m.FirstJourney.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("firstJourney")
			}
			return err
		}
	}

	return nil
}

func (m *TflAPIPresentationEntitiesSchedule) validateKnownJourneys(formats strfmt.Registry) error {

	if swag.IsZero(m.KnownJourneys) { // not required
		return nil
	}

	for i := 0; i < len(m.KnownJourneys); i++ {
		if swag.IsZero(m.KnownJourneys[i]) { // not required
			continue
		}

		if m.KnownJourneys[i] != nil {
			if err := m.KnownJourneys[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("knownJourneys" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TflAPIPresentationEntitiesSchedule) validateLastJourney(formats strfmt.Registry) error {

	if swag.IsZero(m.LastJourney) { // not required
		return nil
	}

	if m.LastJourney != nil {
		if err := m.LastJourney.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastJourney")
			}
			return err
		}
	}

	return nil
}

func (m *TflAPIPresentationEntitiesSchedule) validatePeriods(formats strfmt.Registry) error {

	if swag.IsZero(m.Periods) { // not required
		return nil
	}

	for i := 0; i < len(m.Periods); i++ {
		if swag.IsZero(m.Periods[i]) { // not required
			continue
		}

		if m.Periods[i] != nil {
			if err := m.Periods[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("periods" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesSchedule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesSchedule) UnmarshalBinary(b []byte) error {
	var res TflAPIPresentationEntitiesSchedule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
