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

// TflAPIPresentationEntitiesLineStatus tfl Api presentation entities line status
// swagger:model Tfl.Api.Presentation.Entities.LineStatus
type TflAPIPresentationEntitiesLineStatus struct {

	// created
	// Format: date-time
	Created strfmt.DateTime `json:"created,omitempty"`

	// disruption
	Disruption *TflAPIPresentationEntitiesDisruption `json:"disruption,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// line Id
	LineID string `json:"lineId,omitempty"`

	// modified
	// Format: date-time
	Modified strfmt.DateTime `json:"modified,omitempty"`

	// reason
	Reason string `json:"reason,omitempty"`

	// status severity
	StatusSeverity int32 `json:"statusSeverity,omitempty"`

	// status severity description
	StatusSeverityDescription string `json:"statusSeverityDescription,omitempty"`

	// validity periods
	ValidityPeriods []*TflAPIPresentationEntitiesValidityPeriod `json:"validityPeriods"`
}

// Validate validates this tfl Api presentation entities line status
func (m *TflAPIPresentationEntitiesLineStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisruption(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidityPeriods(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TflAPIPresentationEntitiesLineStatus) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TflAPIPresentationEntitiesLineStatus) validateDisruption(formats strfmt.Registry) error {

	if swag.IsZero(m.Disruption) { // not required
		return nil
	}

	if m.Disruption != nil {
		if err := m.Disruption.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("disruption")
			}
			return err
		}
	}

	return nil
}

func (m *TflAPIPresentationEntitiesLineStatus) validateModified(formats strfmt.Registry) error {

	if swag.IsZero(m.Modified) { // not required
		return nil
	}

	if err := validate.FormatOf("modified", "body", "date-time", m.Modified.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TflAPIPresentationEntitiesLineStatus) validateValidityPeriods(formats strfmt.Registry) error {

	if swag.IsZero(m.ValidityPeriods) { // not required
		return nil
	}

	for i := 0; i < len(m.ValidityPeriods); i++ {
		if swag.IsZero(m.ValidityPeriods[i]) { // not required
			continue
		}

		if m.ValidityPeriods[i] != nil {
			if err := m.ValidityPeriods[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("validityPeriods" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesLineStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesLineStatus) UnmarshalBinary(b []byte) error {
	var res TflAPIPresentationEntitiesLineStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}