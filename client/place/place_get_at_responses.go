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

// PlaceGetAtReader is a Reader for the PlaceGetAt structure.
type PlaceGetAtReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PlaceGetAtReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPlaceGetAtOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPlaceGetAtOK creates a PlaceGetAtOK with default headers values
func NewPlaceGetAtOK() *PlaceGetAtOK {
	return &PlaceGetAtOK{}
}

/*PlaceGetAtOK handles this case with default header values.

OK
*/
type PlaceGetAtOK struct {
	Payload models.SystemObject
}

func (o *PlaceGetAtOK) Error() string {
	return fmt.Sprintf("[GET /Place/{type}/At/{Lat}/{Lon}][%d] placeGetAtOK  %+v", 200, o.Payload)
}

func (o *PlaceGetAtOK) GetPayload() models.SystemObject {
	return o.Payload
}

func (o *PlaceGetAtOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
