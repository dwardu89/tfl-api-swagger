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

// TflAPIPresentationEntitiesValidityPeriod Represents a period for which a planned works is valid.
// swagger:model Tfl.Api.Presentation.Entities.ValidityPeriod
type TflAPIPresentationEntitiesValidityPeriod struct {

	// Gets or sets the start date.
	// Format: date-time
	FromDate strfmt.DateTime `json:"fromDate,omitempty"`

	// If true is a realtime status rather than planned or info
	IsNow bool `json:"isNow,omitempty"`

	// Gets or sets the end date.
	// Format: date-time
	ToDate strfmt.DateTime `json:"toDate,omitempty"`
}

// Validate validates this tfl Api presentation entities validity period
func (m *TflAPIPresentationEntitiesValidityPeriod) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFromDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateToDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TflAPIPresentationEntitiesValidityPeriod) validateFromDate(formats strfmt.Registry) error {

	if swag.IsZero(m.FromDate) { // not required
		return nil
	}

	if err := validate.FormatOf("fromDate", "body", "date-time", m.FromDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TflAPIPresentationEntitiesValidityPeriod) validateToDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ToDate) { // not required
		return nil
	}

	if err := validate.FormatOf("toDate", "body", "date-time", m.ToDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesValidityPeriod) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesValidityPeriod) UnmarshalBinary(b []byte) error {
	var res TflAPIPresentationEntitiesValidityPeriod
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
