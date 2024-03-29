// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TflAPIPresentationEntitiesJourneyPlannerLeg tfl Api presentation entities journey planner leg
// swagger:model Tfl.Api.Presentation.Entities.JourneyPlanner.Leg
type TflAPIPresentationEntitiesJourneyPlannerLeg struct {

	// arrival point
	ArrivalPoint *TflAPIPresentationEntitiesPoint `json:"arrivalPoint,omitempty"`

	// arrival time
	// Format: date-time
	ArrivalTime strfmt.DateTime `json:"arrivalTime,omitempty"`

	// departure point
	DeparturePoint *TflAPIPresentationEntitiesPoint `json:"departurePoint,omitempty"`

	// departure time
	// Format: date-time
	DepartureTime strfmt.DateTime `json:"departureTime,omitempty"`

	// disruptions
	Disruptions []*TflAPIPresentationEntitiesDisruption `json:"disruptions"`

	// distance
	Distance float64 `json:"distance,omitempty"`

	// duration
	Duration int32 `json:"duration,omitempty"`

	// has fixed locations
	// Read Only: true
	HasFixedLocations *bool `json:"hasFixedLocations,omitempty"`

	// Describes the action the user need to take for this section, E.g. "walk to the
	//             district line"
	Instruction *TflAPIPresentationEntitiesInstruction `json:"instruction,omitempty"`

	// is disrupted
	// Read Only: true
	IsDisrupted *bool `json:"isDisrupted,omitempty"`

	// mode
	Mode *TflAPIPresentationEntitiesIdentifier `json:"mode,omitempty"`

	// obstacles
	Obstacles []*TflAPIPresentationEntitiesJourneyPlannerObstacle `json:"obstacles"`

	// path
	Path *TflAPIPresentationEntitiesJourneyPlannerPath `json:"path,omitempty"`

	// planned works
	PlannedWorks []*TflAPIPresentationEntitiesJourneyPlannerPlannedWork `json:"plannedWorks"`

	// route options
	RouteOptions []*TflAPIPresentationEntitiesJourneyPlannerRouteOption `json:"routeOptions"`

	// speed
	Speed string `json:"speed,omitempty"`
}

// Validate validates this tfl Api presentation entities journey planner leg
func (m *TflAPIPresentationEntitiesJourneyPlannerLeg) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArrivalPoint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateArrivalTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeparturePoint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDepartureTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisruptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstruction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObstacles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlannedWorks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRouteOptions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TflAPIPresentationEntitiesJourneyPlannerLeg) validateArrivalPoint(formats strfmt.Registry) error {

	if swag.IsZero(m.ArrivalPoint) { // not required
		return nil
	}

	if m.ArrivalPoint != nil {
		if err := m.ArrivalPoint.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("arrivalPoint")
			}
			return err
		}
	}

	return nil
}

func (m *TflAPIPresentationEntitiesJourneyPlannerLeg) validateArrivalTime(formats strfmt.Registry) error {

	if swag.IsZero(m.ArrivalTime) { // not required
		return nil
	}

	if err := validate.FormatOf("arrivalTime", "body", "date-time", m.ArrivalTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TflAPIPresentationEntitiesJourneyPlannerLeg) validateDeparturePoint(formats strfmt.Registry) error {

	if swag.IsZero(m.DeparturePoint) { // not required
		return nil
	}

	if m.DeparturePoint != nil {
		if err := m.DeparturePoint.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("departurePoint")
			}
			return err
		}
	}

	return nil
}

func (m *TflAPIPresentationEntitiesJourneyPlannerLeg) validateDepartureTime(formats strfmt.Registry) error {

	if swag.IsZero(m.DepartureTime) { // not required
		return nil
	}

	if err := validate.FormatOf("departureTime", "body", "date-time", m.DepartureTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TflAPIPresentationEntitiesJourneyPlannerLeg) validateDisruptions(formats strfmt.Registry) error {

	if swag.IsZero(m.Disruptions) { // not required
		return nil
	}

	for i := 0; i < len(m.Disruptions); i++ {
		if swag.IsZero(m.Disruptions[i]) { // not required
			continue
		}

		if m.Disruptions[i] != nil {
			if err := m.Disruptions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("disruptions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TflAPIPresentationEntitiesJourneyPlannerLeg) validateInstruction(formats strfmt.Registry) error {

	if swag.IsZero(m.Instruction) { // not required
		return nil
	}

	if m.Instruction != nil {
		if err := m.Instruction.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("instruction")
			}
			return err
		}
	}

	return nil
}

func (m *TflAPIPresentationEntitiesJourneyPlannerLeg) validateMode(formats strfmt.Registry) error {

	if swag.IsZero(m.Mode) { // not required
		return nil
	}

	if m.Mode != nil {
		if err := m.Mode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mode")
			}
			return err
		}
	}

	return nil
}

func (m *TflAPIPresentationEntitiesJourneyPlannerLeg) validateObstacles(formats strfmt.Registry) error {

	if swag.IsZero(m.Obstacles) { // not required
		return nil
	}

	for i := 0; i < len(m.Obstacles); i++ {
		if swag.IsZero(m.Obstacles[i]) { // not required
			continue
		}

		if m.Obstacles[i] != nil {
			if err := m.Obstacles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("obstacles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TflAPIPresentationEntitiesJourneyPlannerLeg) validatePath(formats strfmt.Registry) error {

	if swag.IsZero(m.Path) { // not required
		return nil
	}

	if m.Path != nil {
		if err := m.Path.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("path")
			}
			return err
		}
	}

	return nil
}

func (m *TflAPIPresentationEntitiesJourneyPlannerLeg) validatePlannedWorks(formats strfmt.Registry) error {

	if swag.IsZero(m.PlannedWorks) { // not required
		return nil
	}

	for i := 0; i < len(m.PlannedWorks); i++ {
		if swag.IsZero(m.PlannedWorks[i]) { // not required
			continue
		}

		if m.PlannedWorks[i] != nil {
			if err := m.PlannedWorks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("plannedWorks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TflAPIPresentationEntitiesJourneyPlannerLeg) validateRouteOptions(formats strfmt.Registry) error {

	if swag.IsZero(m.RouteOptions) { // not required
		return nil
	}

	for i := 0; i < len(m.RouteOptions); i++ {
		if swag.IsZero(m.RouteOptions[i]) { // not required
			continue
		}

		if m.RouteOptions[i] != nil {
			if err := m.RouteOptions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("routeOptions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesJourneyPlannerLeg) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesJourneyPlannerLeg) UnmarshalBinary(b []byte) error {
	var res TflAPIPresentationEntitiesJourneyPlannerLeg
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
