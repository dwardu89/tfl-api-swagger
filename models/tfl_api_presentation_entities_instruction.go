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

// TflAPIPresentationEntitiesInstruction tfl Api presentation entities instruction
// swagger:model Tfl.Api.Presentation.Entities.Instruction
type TflAPIPresentationEntitiesInstruction struct {

	// detailed
	Detailed string `json:"detailed,omitempty"`

	// steps
	Steps []*TflAPIPresentationEntitiesInstructionStep `json:"steps"`

	// summary
	Summary string `json:"summary,omitempty"`
}

// Validate validates this tfl Api presentation entities instruction
func (m *TflAPIPresentationEntitiesInstruction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSteps(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TflAPIPresentationEntitiesInstruction) validateSteps(formats strfmt.Registry) error {

	if swag.IsZero(m.Steps) { // not required
		return nil
	}

	for i := 0; i < len(m.Steps); i++ {
		if swag.IsZero(m.Steps[i]) { // not required
			continue
		}

		if m.Steps[i] != nil {
			if err := m.Steps[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("steps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesInstruction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesInstruction) UnmarshalBinary(b []byte) error {
	var res TflAPIPresentationEntitiesInstruction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
