// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TflAPIPresentationEntitiesInstructionStep tfl Api presentation entities instruction step
// swagger:model Tfl.Api.Presentation.Entities.InstructionStep
type TflAPIPresentationEntitiesInstructionStep struct {

	// cumulative distance
	CumulativeDistance int32 `json:"cumulativeDistance,omitempty"`

	// cumulative travel time
	CumulativeTravelTime int32 `json:"cumulativeTravelTime,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// description heading
	DescriptionHeading string `json:"descriptionHeading,omitempty"`

	// distance
	Distance int32 `json:"distance,omitempty"`

	// latitude
	Latitude float64 `json:"latitude,omitempty"`

	// longitude
	Longitude float64 `json:"longitude,omitempty"`

	// path attribute
	PathAttribute *TflAPIPresentationEntitiesPathAttribute `json:"pathAttribute,omitempty"`

	// sky direction
	SkyDirection int32 `json:"skyDirection,omitempty"`

	// sky direction description
	// Enum: [North NorthEast East SouthEast South SouthWest West NorthWest]
	SkyDirectionDescription string `json:"skyDirectionDescription,omitempty"`

	// street name
	StreetName string `json:"streetName,omitempty"`

	// track type
	// Enum: [CycleSuperHighway CanalTowpath QuietRoad ProvisionForCyclists BusyRoads None PushBike Quietway]
	TrackType string `json:"trackType,omitempty"`

	// turn direction
	TurnDirection string `json:"turnDirection,omitempty"`
}

// Validate validates this tfl Api presentation entities instruction step
func (m *TflAPIPresentationEntitiesInstructionStep) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePathAttribute(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSkyDirectionDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrackType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TflAPIPresentationEntitiesInstructionStep) validatePathAttribute(formats strfmt.Registry) error {

	if swag.IsZero(m.PathAttribute) { // not required
		return nil
	}

	if m.PathAttribute != nil {
		if err := m.PathAttribute.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pathAttribute")
			}
			return err
		}
	}

	return nil
}

var tflApiPresentationEntitiesInstructionStepTypeSkyDirectionDescriptionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["North","NorthEast","East","SouthEast","South","SouthWest","West","NorthWest"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tflApiPresentationEntitiesInstructionStepTypeSkyDirectionDescriptionPropEnum = append(tflApiPresentationEntitiesInstructionStepTypeSkyDirectionDescriptionPropEnum, v)
	}
}

const (

	// TflAPIPresentationEntitiesInstructionStepSkyDirectionDescriptionNorth captures enum value "North"
	TflAPIPresentationEntitiesInstructionStepSkyDirectionDescriptionNorth string = "North"

	// TflAPIPresentationEntitiesInstructionStepSkyDirectionDescriptionNorthEast captures enum value "NorthEast"
	TflAPIPresentationEntitiesInstructionStepSkyDirectionDescriptionNorthEast string = "NorthEast"

	// TflAPIPresentationEntitiesInstructionStepSkyDirectionDescriptionEast captures enum value "East"
	TflAPIPresentationEntitiesInstructionStepSkyDirectionDescriptionEast string = "East"

	// TflAPIPresentationEntitiesInstructionStepSkyDirectionDescriptionSouthEast captures enum value "SouthEast"
	TflAPIPresentationEntitiesInstructionStepSkyDirectionDescriptionSouthEast string = "SouthEast"

	// TflAPIPresentationEntitiesInstructionStepSkyDirectionDescriptionSouth captures enum value "South"
	TflAPIPresentationEntitiesInstructionStepSkyDirectionDescriptionSouth string = "South"

	// TflAPIPresentationEntitiesInstructionStepSkyDirectionDescriptionSouthWest captures enum value "SouthWest"
	TflAPIPresentationEntitiesInstructionStepSkyDirectionDescriptionSouthWest string = "SouthWest"

	// TflAPIPresentationEntitiesInstructionStepSkyDirectionDescriptionWest captures enum value "West"
	TflAPIPresentationEntitiesInstructionStepSkyDirectionDescriptionWest string = "West"

	// TflAPIPresentationEntitiesInstructionStepSkyDirectionDescriptionNorthWest captures enum value "NorthWest"
	TflAPIPresentationEntitiesInstructionStepSkyDirectionDescriptionNorthWest string = "NorthWest"
)

// prop value enum
func (m *TflAPIPresentationEntitiesInstructionStep) validateSkyDirectionDescriptionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, tflApiPresentationEntitiesInstructionStepTypeSkyDirectionDescriptionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TflAPIPresentationEntitiesInstructionStep) validateSkyDirectionDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.SkyDirectionDescription) { // not required
		return nil
	}

	// value enum
	if err := m.validateSkyDirectionDescriptionEnum("skyDirectionDescription", "body", m.SkyDirectionDescription); err != nil {
		return err
	}

	return nil
}

var tflApiPresentationEntitiesInstructionStepTypeTrackTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CycleSuperHighway","CanalTowpath","QuietRoad","ProvisionForCyclists","BusyRoads","None","PushBike","Quietway"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tflApiPresentationEntitiesInstructionStepTypeTrackTypePropEnum = append(tflApiPresentationEntitiesInstructionStepTypeTrackTypePropEnum, v)
	}
}

const (

	// TflAPIPresentationEntitiesInstructionStepTrackTypeCycleSuperHighway captures enum value "CycleSuperHighway"
	TflAPIPresentationEntitiesInstructionStepTrackTypeCycleSuperHighway string = "CycleSuperHighway"

	// TflAPIPresentationEntitiesInstructionStepTrackTypeCanalTowpath captures enum value "CanalTowpath"
	TflAPIPresentationEntitiesInstructionStepTrackTypeCanalTowpath string = "CanalTowpath"

	// TflAPIPresentationEntitiesInstructionStepTrackTypeQuietRoad captures enum value "QuietRoad"
	TflAPIPresentationEntitiesInstructionStepTrackTypeQuietRoad string = "QuietRoad"

	// TflAPIPresentationEntitiesInstructionStepTrackTypeProvisionForCyclists captures enum value "ProvisionForCyclists"
	TflAPIPresentationEntitiesInstructionStepTrackTypeProvisionForCyclists string = "ProvisionForCyclists"

	// TflAPIPresentationEntitiesInstructionStepTrackTypeBusyRoads captures enum value "BusyRoads"
	TflAPIPresentationEntitiesInstructionStepTrackTypeBusyRoads string = "BusyRoads"

	// TflAPIPresentationEntitiesInstructionStepTrackTypeNone captures enum value "None"
	TflAPIPresentationEntitiesInstructionStepTrackTypeNone string = "None"

	// TflAPIPresentationEntitiesInstructionStepTrackTypePushBike captures enum value "PushBike"
	TflAPIPresentationEntitiesInstructionStepTrackTypePushBike string = "PushBike"

	// TflAPIPresentationEntitiesInstructionStepTrackTypeQuietway captures enum value "Quietway"
	TflAPIPresentationEntitiesInstructionStepTrackTypeQuietway string = "Quietway"
)

// prop value enum
func (m *TflAPIPresentationEntitiesInstructionStep) validateTrackTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, tflApiPresentationEntitiesInstructionStepTypeTrackTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TflAPIPresentationEntitiesInstructionStep) validateTrackType(formats strfmt.Registry) error {

	if swag.IsZero(m.TrackType) { // not required
		return nil
	}

	// value enum
	if err := m.validateTrackTypeEnum("trackType", "body", m.TrackType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesInstructionStep) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesInstructionStep) UnmarshalBinary(b []byte) error {
	var res TflAPIPresentationEntitiesInstructionStep
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
