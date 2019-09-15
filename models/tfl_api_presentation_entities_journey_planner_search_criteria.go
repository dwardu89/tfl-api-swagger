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

// TflAPIPresentationEntitiesJourneyPlannerSearchCriteria tfl Api presentation entities journey planner search criteria
// swagger:model Tfl.Api.Presentation.Entities.JourneyPlanner.SearchCriteria
type TflAPIPresentationEntitiesJourneyPlannerSearchCriteria struct {

	// date time
	// Format: date-time
	DateTime strfmt.DateTime `json:"dateTime,omitempty"`

	// date time type
	// Enum: [Arriving Departing]
	DateTimeType string `json:"dateTimeType,omitempty"`

	// time adjustments
	TimeAdjustments *TflAPIPresentationEntitiesJourneyPlannerTimeAdjustments `json:"timeAdjustments,omitempty"`
}

// Validate validates this tfl Api presentation entities journey planner search criteria
func (m *TflAPIPresentationEntitiesJourneyPlannerSearchCriteria) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateTimeType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeAdjustments(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TflAPIPresentationEntitiesJourneyPlannerSearchCriteria) validateDateTime(formats strfmt.Registry) error {

	if swag.IsZero(m.DateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("dateTime", "body", "date-time", m.DateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

var tflApiPresentationEntitiesJourneyPlannerSearchCriteriaTypeDateTimeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Arriving","Departing"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tflApiPresentationEntitiesJourneyPlannerSearchCriteriaTypeDateTimeTypePropEnum = append(tflApiPresentationEntitiesJourneyPlannerSearchCriteriaTypeDateTimeTypePropEnum, v)
	}
}

const (

	// TflAPIPresentationEntitiesJourneyPlannerSearchCriteriaDateTimeTypeArriving captures enum value "Arriving"
	TflAPIPresentationEntitiesJourneyPlannerSearchCriteriaDateTimeTypeArriving string = "Arriving"

	// TflAPIPresentationEntitiesJourneyPlannerSearchCriteriaDateTimeTypeDeparting captures enum value "Departing"
	TflAPIPresentationEntitiesJourneyPlannerSearchCriteriaDateTimeTypeDeparting string = "Departing"
)

// prop value enum
func (m *TflAPIPresentationEntitiesJourneyPlannerSearchCriteria) validateDateTimeTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, tflApiPresentationEntitiesJourneyPlannerSearchCriteriaTypeDateTimeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TflAPIPresentationEntitiesJourneyPlannerSearchCriteria) validateDateTimeType(formats strfmt.Registry) error {

	if swag.IsZero(m.DateTimeType) { // not required
		return nil
	}

	// value enum
	if err := m.validateDateTimeTypeEnum("dateTimeType", "body", m.DateTimeType); err != nil {
		return err
	}

	return nil
}

func (m *TflAPIPresentationEntitiesJourneyPlannerSearchCriteria) validateTimeAdjustments(formats strfmt.Registry) error {

	if swag.IsZero(m.TimeAdjustments) { // not required
		return nil
	}

	if m.TimeAdjustments != nil {
		if err := m.TimeAdjustments.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("timeAdjustments")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesJourneyPlannerSearchCriteria) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesJourneyPlannerSearchCriteria) UnmarshalBinary(b []byte) error {
	var res TflAPIPresentationEntitiesJourneyPlannerSearchCriteria
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}