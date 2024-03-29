// Code generated by go-swagger; DO NOT EDIT.

package place

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/dwardu89/tfl-api-swagger/models"
)

// PlaceMetaPlaceTypesReader is a Reader for the PlaceMetaPlaceTypes structure.
type PlaceMetaPlaceTypesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PlaceMetaPlaceTypesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPlaceMetaPlaceTypesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPlaceMetaPlaceTypesOK creates a PlaceMetaPlaceTypesOK with default headers values
func NewPlaceMetaPlaceTypesOK() *PlaceMetaPlaceTypesOK {
	return &PlaceMetaPlaceTypesOK{}
}

/*PlaceMetaPlaceTypesOK handles this case with default header values.

OK
*/
type PlaceMetaPlaceTypesOK struct {
	Payload []*models.TflAPIPresentationEntitiesPlaceCategory
}

func (o *PlaceMetaPlaceTypesOK) Error() string {
	return fmt.Sprintf("[GET /Place/Meta/PlaceTypes][%d] placeMetaPlaceTypesOK  %+v", 200, o.Payload)
}

func (o *PlaceMetaPlaceTypesOK) GetPayload() []*models.TflAPIPresentationEntitiesPlaceCategory {
	return o.Payload
}

func (o *PlaceMetaPlaceTypesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
