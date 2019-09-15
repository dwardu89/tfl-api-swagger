// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TflAPIPresentationEntitiesBay tfl Api presentation entities bay
// swagger:model Tfl.Api.Presentation.Entities.Bay
type TflAPIPresentationEntitiesBay struct {

	// bay count
	BayCount int32 `json:"bayCount,omitempty"`

	// bay type
	BayType string `json:"bayType,omitempty"`

	// free
	Free int32 `json:"free,omitempty"`

	// occupied
	Occupied int32 `json:"occupied,omitempty"`
}

// Validate validates this tfl Api presentation entities bay
func (m *TflAPIPresentationEntitiesBay) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesBay) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesBay) UnmarshalBinary(b []byte) error {
	var res TflAPIPresentationEntitiesBay
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}