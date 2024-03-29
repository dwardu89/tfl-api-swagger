// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TflAPIPresentationEntitiesPeriod tfl Api presentation entities period
// swagger:model Tfl.Api.Presentation.Entities.Period
type TflAPIPresentationEntitiesPeriod struct {

	// frequency
	Frequency *TflAPIPresentationEntitiesServiceFrequency `json:"frequency,omitempty"`

	// from time
	FromTime *TflAPIPresentationEntitiesTwentyFourHourClockTime `json:"fromTime,omitempty"`

	// to time
	ToTime *TflAPIPresentationEntitiesTwentyFourHourClockTime `json:"toTime,omitempty"`

	// type
	// Enum: [Normal FrequencyHours FrequencyMinutes Unknown]
	Type string `json:"type,omitempty"`
}

// Validate validates this tfl Api presentation entities period
func (m *TflAPIPresentationEntitiesPeriod) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFrequency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFromTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateToTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TflAPIPresentationEntitiesPeriod) validateFrequency(formats strfmt.Registry) error {

	if swag.IsZero(m.Frequency) { // not required
		return nil
	}

	if m.Frequency != nil {
		if err := m.Frequency.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("frequency")
			}
			return err
		}
	}

	return nil
}

func (m *TflAPIPresentationEntitiesPeriod) validateFromTime(formats strfmt.Registry) error {

	if swag.IsZero(m.FromTime) { // not required
		return nil
	}

	if m.FromTime != nil {
		if err := m.FromTime.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fromTime")
			}
			return err
		}
	}

	return nil
}

func (m *TflAPIPresentationEntitiesPeriod) validateToTime(formats strfmt.Registry) error {

	if swag.IsZero(m.ToTime) { // not required
		return nil
	}

	if m.ToTime != nil {
		if err := m.ToTime.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("toTime")
			}
			return err
		}
	}

	return nil
}

var tflApiPresentationEntitiesPeriodTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Normal","FrequencyHours","FrequencyMinutes","Unknown"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tflApiPresentationEntitiesPeriodTypeTypePropEnum = append(tflApiPresentationEntitiesPeriodTypeTypePropEnum, v)
	}
}

const (

	// TflAPIPresentationEntitiesPeriodTypeNormal captures enum value "Normal"
	TflAPIPresentationEntitiesPeriodTypeNormal string = "Normal"

	// TflAPIPresentationEntitiesPeriodTypeFrequencyHours captures enum value "FrequencyHours"
	TflAPIPresentationEntitiesPeriodTypeFrequencyHours string = "FrequencyHours"

	// TflAPIPresentationEntitiesPeriodTypeFrequencyMinutes captures enum value "FrequencyMinutes"
	TflAPIPresentationEntitiesPeriodTypeFrequencyMinutes string = "FrequencyMinutes"

	// TflAPIPresentationEntitiesPeriodTypeUnknown captures enum value "Unknown"
	TflAPIPresentationEntitiesPeriodTypeUnknown string = "Unknown"
)

// prop value enum
func (m *TflAPIPresentationEntitiesPeriod) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, tflApiPresentationEntitiesPeriodTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TflAPIPresentationEntitiesPeriod) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesPeriod) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesPeriod) UnmarshalBinary(b []byte) error {
	var res TflAPIPresentationEntitiesPeriod
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
