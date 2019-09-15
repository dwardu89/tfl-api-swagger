// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TflAPIPresentationEntitiesStopPointSequence tfl Api presentation entities stop point sequence
// swagger:model Tfl.Api.Presentation.Entities.StopPointSequence
type TflAPIPresentationEntitiesStopPointSequence struct {

	// The id of this branch.
	BranchID int32 `json:"branchId,omitempty"`

	// direction
	Direction string `json:"direction,omitempty"`

	// line Id
	LineID string `json:"lineId,omitempty"`

	// line name
	LineName string `json:"lineName,omitempty"`

	// The ids of the next branch(es) in the sequence. Note that the next and previous branch id can be
	//             identical in the case of a looped route e.g. the Circle line.
	NextBranchIds []int32 `json:"nextBranchIds"`

	// The ids of the previous branch(es) in the sequence. Note that the next and previous branch id can be
	//             identical in the case of a looped route e.g. the Circle line.
	PrevBranchIds []int32 `json:"prevBranchIds"`

	// service type
	// Enum: [Regular Night]
	ServiceType string `json:"serviceType,omitempty"`

	// stop point
	StopPoint []*TflAPIPresentationEntitiesMatchedStop `json:"stopPoint"`
}

// Validate validates this tfl Api presentation entities stop point sequence
func (m *TflAPIPresentationEntitiesStopPointSequence) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateServiceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStopPoint(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var tflApiPresentationEntitiesStopPointSequenceTypeServiceTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Regular","Night"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tflApiPresentationEntitiesStopPointSequenceTypeServiceTypePropEnum = append(tflApiPresentationEntitiesStopPointSequenceTypeServiceTypePropEnum, v)
	}
}

const (

	// TflAPIPresentationEntitiesStopPointSequenceServiceTypeRegular captures enum value "Regular"
	TflAPIPresentationEntitiesStopPointSequenceServiceTypeRegular string = "Regular"

	// TflAPIPresentationEntitiesStopPointSequenceServiceTypeNight captures enum value "Night"
	TflAPIPresentationEntitiesStopPointSequenceServiceTypeNight string = "Night"
)

// prop value enum
func (m *TflAPIPresentationEntitiesStopPointSequence) validateServiceTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, tflApiPresentationEntitiesStopPointSequenceTypeServiceTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TflAPIPresentationEntitiesStopPointSequence) validateServiceType(formats strfmt.Registry) error {

	if swag.IsZero(m.ServiceType) { // not required
		return nil
	}

	// value enum
	if err := m.validateServiceTypeEnum("serviceType", "body", m.ServiceType); err != nil {
		return err
	}

	return nil
}

func (m *TflAPIPresentationEntitiesStopPointSequence) validateStopPoint(formats strfmt.Registry) error {

	if swag.IsZero(m.StopPoint) { // not required
		return nil
	}

	for i := 0; i < len(m.StopPoint); i++ {
		if swag.IsZero(m.StopPoint[i]) { // not required
			continue
		}

		if m.StopPoint[i] != nil {
			if err := m.StopPoint[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("stopPoint" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesStopPointSequence) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesStopPointSequence) UnmarshalBinary(b []byte) error {
	var res TflAPIPresentationEntitiesStopPointSequence
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
