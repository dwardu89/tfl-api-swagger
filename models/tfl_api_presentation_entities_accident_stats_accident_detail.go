// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TflAPIPresentationEntitiesAccidentStatsAccidentDetail tfl Api presentation entities accident stats accident detail
// swagger:model Tfl.Api.Presentation.Entities.AccidentStats.AccidentDetail
type TflAPIPresentationEntitiesAccidentStatsAccidentDetail struct {

	// borough
	Borough string `json:"borough,omitempty"`

	// casualties
	Casualties []*TflAPIPresentationEntitiesAccidentStatsCasualty `json:"casualties"`

	// date
	// Format: date-time
	Date strfmt.DateTime `json:"date,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// lat
	Lat float64 `json:"lat,omitempty"`

	// location
	Location string `json:"location,omitempty"`

	// lon
	Lon float64 `json:"lon,omitempty"`

	// severity
	Severity string `json:"severity,omitempty"`

	// vehicles
	Vehicles []*TflAPIPresentationEntitiesAccidentStatsVehicle `json:"vehicles"`
}

// Validate validates this tfl Api presentation entities accident stats accident detail
func (m *TflAPIPresentationEntitiesAccidentStatsAccidentDetail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCasualties(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVehicles(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TflAPIPresentationEntitiesAccidentStatsAccidentDetail) validateCasualties(formats strfmt.Registry) error {

	if swag.IsZero(m.Casualties) { // not required
		return nil
	}

	for i := 0; i < len(m.Casualties); i++ {
		if swag.IsZero(m.Casualties[i]) { // not required
			continue
		}

		if m.Casualties[i] != nil {
			if err := m.Casualties[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("casualties" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TflAPIPresentationEntitiesAccidentStatsAccidentDetail) validateDate(formats strfmt.Registry) error {

	if swag.IsZero(m.Date) { // not required
		return nil
	}

	if err := validate.FormatOf("date", "body", "date-time", m.Date.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TflAPIPresentationEntitiesAccidentStatsAccidentDetail) validateVehicles(formats strfmt.Registry) error {

	if swag.IsZero(m.Vehicles) { // not required
		return nil
	}

	for i := 0; i < len(m.Vehicles); i++ {
		if swag.IsZero(m.Vehicles[i]) { // not required
			continue
		}

		if m.Vehicles[i] != nil {
			if err := m.Vehicles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vehicles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesAccidentStatsAccidentDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesAccidentStatsAccidentDetail) UnmarshalBinary(b []byte) error {
	var res TflAPIPresentationEntitiesAccidentStatsAccidentDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}