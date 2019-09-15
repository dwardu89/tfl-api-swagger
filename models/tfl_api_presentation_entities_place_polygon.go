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

// TflAPIPresentationEntitiesPlacePolygon tfl Api presentation entities place polygon
// swagger:model Tfl.Api.Presentation.Entities.PlacePolygon
type TflAPIPresentationEntitiesPlacePolygon struct {

	// common name
	CommonName string `json:"commonName,omitempty"`

	// geo points
	GeoPoints []*TflAPICommonGeoPoint `json:"geoPoints"`
}

// Validate validates this tfl Api presentation entities place polygon
func (m *TflAPIPresentationEntitiesPlacePolygon) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGeoPoints(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TflAPIPresentationEntitiesPlacePolygon) validateGeoPoints(formats strfmt.Registry) error {

	if swag.IsZero(m.GeoPoints) { // not required
		return nil
	}

	for i := 0; i < len(m.GeoPoints); i++ {
		if swag.IsZero(m.GeoPoints[i]) { // not required
			continue
		}

		if m.GeoPoints[i] != nil {
			if err := m.GeoPoints[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("geoPoints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesPlacePolygon) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesPlacePolygon) UnmarshalBinary(b []byte) error {
	var res TflAPIPresentationEntitiesPlacePolygon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
