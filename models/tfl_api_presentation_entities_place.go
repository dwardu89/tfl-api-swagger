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

// TflAPIPresentationEntitiesPlace tfl Api presentation entities place
// swagger:model Tfl.Api.Presentation.Entities.Place
type TflAPIPresentationEntitiesPlace struct {

	// A bag of additional key/value pairs with extra information about this place.
	AdditionalProperties []*TflAPIPresentationEntitiesAdditionalProperties `json:"additionalProperties"`

	// children
	Children []*TflAPIPresentationEntitiesPlace `json:"children"`

	// children urls
	ChildrenUrls []string `json:"childrenUrls"`

	// A human readable name.
	CommonName string `json:"commonName,omitempty"`

	// The distance of the place from its search point, if this is the result
	//             of a geographical search, otherwise zero.
	Distance float64 `json:"distance,omitempty"`

	// A unique identifier.
	ID string `json:"id,omitempty"`

	// WGS84 latitude of the location.
	Lat float64 `json:"lat,omitempty"`

	// WGS84 longitude of the location.
	Lon float64 `json:"lon,omitempty"`

	// The type of Place. See /Place/Meta/placeTypes for possible values.
	PlaceType string `json:"placeType,omitempty"`

	// The unique location of this resource.
	URL string `json:"url,omitempty"`
}

// Validate validates this tfl Api presentation entities place
func (m *TflAPIPresentationEntitiesPlace) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdditionalProperties(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChildren(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TflAPIPresentationEntitiesPlace) validateAdditionalProperties(formats strfmt.Registry) error {

	if swag.IsZero(m.AdditionalProperties) { // not required
		return nil
	}

	for i := 0; i < len(m.AdditionalProperties); i++ {
		if swag.IsZero(m.AdditionalProperties[i]) { // not required
			continue
		}

		if m.AdditionalProperties[i] != nil {
			if err := m.AdditionalProperties[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("additionalProperties" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TflAPIPresentationEntitiesPlace) validateChildren(formats strfmt.Registry) error {

	if swag.IsZero(m.Children) { // not required
		return nil
	}

	for i := 0; i < len(m.Children); i++ {
		if swag.IsZero(m.Children[i]) { // not required
			continue
		}

		if m.Children[i] != nil {
			if err := m.Children[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("children" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesPlace) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesPlace) UnmarshalBinary(b []byte) error {
	var res TflAPIPresentationEntitiesPlace
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}