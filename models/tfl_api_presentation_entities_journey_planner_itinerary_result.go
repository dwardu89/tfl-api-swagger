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

// TflAPIPresentationEntitiesJourneyPlannerItineraryResult A DTO representing a list of possible journeys.
// swagger:model Tfl.Api.Presentation.Entities.JourneyPlanner.ItineraryResult
type TflAPIPresentationEntitiesJourneyPlannerItineraryResult struct {

	// cycle hire docking station data
	CycleHireDockingStationData *TflAPIPresentationEntitiesJourneyPlannerJourneyPlannerCycleHireDockingStationData `json:"cycleHireDockingStationData,omitempty"`

	// journey vector
	JourneyVector *TflAPIPresentationEntitiesJourneyPlannerJourneyVector `json:"journeyVector,omitempty"`

	// journeys
	Journeys []*TflAPIPresentationEntitiesJourneyPlannerJourney `json:"journeys"`

	// lines
	Lines []*TflAPIPresentationEntitiesLine `json:"lines"`

	// recommended max age minutes
	RecommendedMaxAgeMinutes int32 `json:"recommendedMaxAgeMinutes,omitempty"`

	// search criteria
	SearchCriteria *TflAPIPresentationEntitiesJourneyPlannerSearchCriteria `json:"searchCriteria,omitempty"`

	// stop messages
	StopMessages []string `json:"stopMessages"`
}

// Validate validates this tfl Api presentation entities journey planner itinerary result
func (m *TflAPIPresentationEntitiesJourneyPlannerItineraryResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCycleHireDockingStationData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJourneyVector(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJourneys(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLines(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSearchCriteria(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TflAPIPresentationEntitiesJourneyPlannerItineraryResult) validateCycleHireDockingStationData(formats strfmt.Registry) error {

	if swag.IsZero(m.CycleHireDockingStationData) { // not required
		return nil
	}

	if m.CycleHireDockingStationData != nil {
		if err := m.CycleHireDockingStationData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cycleHireDockingStationData")
			}
			return err
		}
	}

	return nil
}

func (m *TflAPIPresentationEntitiesJourneyPlannerItineraryResult) validateJourneyVector(formats strfmt.Registry) error {

	if swag.IsZero(m.JourneyVector) { // not required
		return nil
	}

	if m.JourneyVector != nil {
		if err := m.JourneyVector.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("journeyVector")
			}
			return err
		}
	}

	return nil
}

func (m *TflAPIPresentationEntitiesJourneyPlannerItineraryResult) validateJourneys(formats strfmt.Registry) error {

	if swag.IsZero(m.Journeys) { // not required
		return nil
	}

	for i := 0; i < len(m.Journeys); i++ {
		if swag.IsZero(m.Journeys[i]) { // not required
			continue
		}

		if m.Journeys[i] != nil {
			if err := m.Journeys[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("journeys" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TflAPIPresentationEntitiesJourneyPlannerItineraryResult) validateLines(formats strfmt.Registry) error {

	if swag.IsZero(m.Lines) { // not required
		return nil
	}

	for i := 0; i < len(m.Lines); i++ {
		if swag.IsZero(m.Lines[i]) { // not required
			continue
		}

		if m.Lines[i] != nil {
			if err := m.Lines[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("lines" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TflAPIPresentationEntitiesJourneyPlannerItineraryResult) validateSearchCriteria(formats strfmt.Registry) error {

	if swag.IsZero(m.SearchCriteria) { // not required
		return nil
	}

	if m.SearchCriteria != nil {
		if err := m.SearchCriteria.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("searchCriteria")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesJourneyPlannerItineraryResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesJourneyPlannerItineraryResult) UnmarshalBinary(b []byte) error {
	var res TflAPIPresentationEntitiesJourneyPlannerItineraryResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
