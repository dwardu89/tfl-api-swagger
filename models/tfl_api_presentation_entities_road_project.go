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

// TflAPIPresentationEntitiesRoadProject tfl Api presentation entities road project
// swagger:model Tfl.Api.Presentation.Entities.RoadProject
type TflAPIPresentationEntitiesRoadProject struct {

	// boroughs benefited
	BoroughsBenefited []string `json:"boroughsBenefited"`

	// construction end date
	// Format: date-time
	ConstructionEndDate strfmt.DateTime `json:"constructionEndDate,omitempty"`

	// construction start date
	// Format: date-time
	ConstructionStartDate strfmt.DateTime `json:"constructionStartDate,omitempty"`

	// consultation end date
	// Format: date-time
	ConsultationEndDate strfmt.DateTime `json:"consultationEndDate,omitempty"`

	// consultation page Url
	ConsultationPageURL string `json:"consultationPageUrl,omitempty"`

	// consultation start date
	// Format: date-time
	ConsultationStartDate strfmt.DateTime `json:"consultationStartDate,omitempty"`

	// contact email
	ContactEmail string `json:"contactEmail,omitempty"`

	// contact name
	ContactName string `json:"contactName,omitempty"`

	// cycle superhighway Id
	CycleSuperhighwayID string `json:"cycleSuperhighwayId,omitempty"`

	// external page Url
	ExternalPageURL string `json:"externalPageUrl,omitempty"`

	// phase
	// Enum: [Unscoped Concept ConsultationEnded Consultation Construction Complete]
	Phase string `json:"phase,omitempty"`

	// project description
	ProjectDescription string `json:"projectDescription,omitempty"`

	// project Id
	ProjectID string `json:"projectId,omitempty"`

	// project name
	ProjectName string `json:"projectName,omitempty"`

	// project page Url
	ProjectPageURL string `json:"projectPageUrl,omitempty"`

	// project summary page Url
	ProjectSummaryPageURL string `json:"projectSummaryPageUrl,omitempty"`

	// scheme name
	SchemeName string `json:"schemeName,omitempty"`
}

// Validate validates this tfl Api presentation entities road project
func (m *TflAPIPresentationEntitiesRoadProject) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConstructionEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConstructionStartDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConsultationEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConsultationStartDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhase(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TflAPIPresentationEntitiesRoadProject) validateConstructionEndDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ConstructionEndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("constructionEndDate", "body", "date-time", m.ConstructionEndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TflAPIPresentationEntitiesRoadProject) validateConstructionStartDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ConstructionStartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("constructionStartDate", "body", "date-time", m.ConstructionStartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TflAPIPresentationEntitiesRoadProject) validateConsultationEndDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ConsultationEndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("consultationEndDate", "body", "date-time", m.ConsultationEndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TflAPIPresentationEntitiesRoadProject) validateConsultationStartDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ConsultationStartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("consultationStartDate", "body", "date-time", m.ConsultationStartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

var tflApiPresentationEntitiesRoadProjectTypePhasePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Unscoped","Concept","ConsultationEnded","Consultation","Construction","Complete"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tflApiPresentationEntitiesRoadProjectTypePhasePropEnum = append(tflApiPresentationEntitiesRoadProjectTypePhasePropEnum, v)
	}
}

const (

	// TflAPIPresentationEntitiesRoadProjectPhaseUnscoped captures enum value "Unscoped"
	TflAPIPresentationEntitiesRoadProjectPhaseUnscoped string = "Unscoped"

	// TflAPIPresentationEntitiesRoadProjectPhaseConcept captures enum value "Concept"
	TflAPIPresentationEntitiesRoadProjectPhaseConcept string = "Concept"

	// TflAPIPresentationEntitiesRoadProjectPhaseConsultationEnded captures enum value "ConsultationEnded"
	TflAPIPresentationEntitiesRoadProjectPhaseConsultationEnded string = "ConsultationEnded"

	// TflAPIPresentationEntitiesRoadProjectPhaseConsultation captures enum value "Consultation"
	TflAPIPresentationEntitiesRoadProjectPhaseConsultation string = "Consultation"

	// TflAPIPresentationEntitiesRoadProjectPhaseConstruction captures enum value "Construction"
	TflAPIPresentationEntitiesRoadProjectPhaseConstruction string = "Construction"

	// TflAPIPresentationEntitiesRoadProjectPhaseComplete captures enum value "Complete"
	TflAPIPresentationEntitiesRoadProjectPhaseComplete string = "Complete"
)

// prop value enum
func (m *TflAPIPresentationEntitiesRoadProject) validatePhaseEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, tflApiPresentationEntitiesRoadProjectTypePhasePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TflAPIPresentationEntitiesRoadProject) validatePhase(formats strfmt.Registry) error {

	if swag.IsZero(m.Phase) { // not required
		return nil
	}

	// value enum
	if err := m.validatePhaseEnum("phase", "body", m.Phase); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesRoadProject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TflAPIPresentationEntitiesRoadProject) UnmarshalBinary(b []byte) error {
	var res TflAPIPresentationEntitiesRoadProject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
