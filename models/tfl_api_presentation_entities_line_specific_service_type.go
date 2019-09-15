// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TflAPIPresentationEntitiesLineSpecificServiceType tfl Api presentation entities line specific service type
// swagger:model Tfl.Api.Presentation.Entities.LineSpecificServiceType
type TflAPIPresentationEntitiesLineSpecificServiceType struct {

	// service type
	ServiceType *TflAPIPresentationEntitiesLineServiceTypeInfo `json:"serviceType,omitempty"`

	// stop serves service type
	StopServesServiceType bool `json:"stopServesServiceType,omitempty"`
}

// Validate validates this tfl Api presentation entities line specific service type
func (m *TflAPIPresentationEntitiesLineSpecificServiceType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateServiceType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TflAPIPresentationEntitiesLineSpecificServiceType) validateServiceType(formats strfmt.Registry) error {

	if swag.IsZero(m.ServiceType) { // not required
		return nil
	}

	if m.ServiceType != nil {
		if err := m.ServiceType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("serviceType")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesLineSpecificServiceType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesLineSpecificServiceType) UnmarshalBinary(b []byte) error {
	var res TflAPIPresentationEntitiesLineSpecificServiceType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
