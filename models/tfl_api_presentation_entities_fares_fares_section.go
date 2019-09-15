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

// TflAPIPresentationEntitiesFaresFaresSection tfl Api presentation entities fares fares section
// swagger:model Tfl.Api.Presentation.Entities.Fares.FaresSection
type TflAPIPresentationEntitiesFaresFaresSection struct {

	// header
	Header string `json:"header,omitempty"`

	// index
	Index int32 `json:"index,omitempty"`

	// journey
	Journey *TflAPIPresentationEntitiesFaresJourney `json:"journey,omitempty"`

	// messages
	Messages []*TflAPIPresentationEntitiesMessage `json:"messages"`

	// rows
	Rows []*TflAPIPresentationEntitiesFaresFareDetails `json:"rows"`
}

// Validate validates this tfl Api presentation entities fares fares section
func (m *TflAPIPresentationEntitiesFaresFaresSection) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateJourney(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRows(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TflAPIPresentationEntitiesFaresFaresSection) validateJourney(formats strfmt.Registry) error {

	if swag.IsZero(m.Journey) { // not required
		return nil
	}

	if m.Journey != nil {
		if err := m.Journey.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("journey")
			}
			return err
		}
	}

	return nil
}

func (m *TflAPIPresentationEntitiesFaresFaresSection) validateMessages(formats strfmt.Registry) error {

	if swag.IsZero(m.Messages) { // not required
		return nil
	}

	for i := 0; i < len(m.Messages); i++ {
		if swag.IsZero(m.Messages[i]) { // not required
			continue
		}

		if m.Messages[i] != nil {
			if err := m.Messages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("messages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TflAPIPresentationEntitiesFaresFaresSection) validateRows(formats strfmt.Registry) error {

	if swag.IsZero(m.Rows) { // not required
		return nil
	}

	for i := 0; i < len(m.Rows); i++ {
		if swag.IsZero(m.Rows[i]) { // not required
			continue
		}

		if m.Rows[i] != nil {
			if err := m.Rows[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rows" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesFaresFaresSection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesFaresFaresSection) UnmarshalBinary(b []byte) error {
	var res TflAPIPresentationEntitiesFaresFaresSection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
