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

// TflAPIPresentationEntitiesAdditionalProperties tfl Api presentation entities additional properties
// swagger:model Tfl.Api.Presentation.Entities.AdditionalProperties
type TflAPIPresentationEntitiesAdditionalProperties struct {

	// category
	Category string `json:"category,omitempty"`

	// key
	Key string `json:"key,omitempty"`

	// modified
	// Format: date-time
	Modified strfmt.DateTime `json:"modified,omitempty"`

	// source system key
	SourceSystemKey string `json:"sourceSystemKey,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this tfl Api presentation entities additional properties
func (m *TflAPIPresentationEntitiesAdditionalProperties) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateModified(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TflAPIPresentationEntitiesAdditionalProperties) validateModified(formats strfmt.Registry) error {

	if swag.IsZero(m.Modified) { // not required
		return nil
	}

	if err := validate.FormatOf("modified", "body", "date-time", m.Modified.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesAdditionalProperties) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesAdditionalProperties) UnmarshalBinary(b []byte) error {
	var res TflAPIPresentationEntitiesAdditionalProperties
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
