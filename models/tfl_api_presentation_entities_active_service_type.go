// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TflAPIPresentationEntitiesActiveServiceType tfl Api presentation entities active service type
// swagger:model Tfl.Api.Presentation.Entities.ActiveServiceType
type TflAPIPresentationEntitiesActiveServiceType struct {

	// mode
	Mode string `json:"mode,omitempty"`

	// service type
	ServiceType string `json:"serviceType,omitempty"`
}

// Validate validates this tfl Api presentation entities active service type
func (m *TflAPIPresentationEntitiesActiveServiceType) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesActiveServiceType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesActiveServiceType) UnmarshalBinary(b []byte) error {
	var res TflAPIPresentationEntitiesActiveServiceType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
