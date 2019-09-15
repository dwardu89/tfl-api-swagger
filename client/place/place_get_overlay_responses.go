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

// PlaceGetOverlayReader is a Reader for the PlaceGetOverlay structure.
type PlaceGetOverlayReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PlaceGetOverlayReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPlaceGetOverlayOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPlaceGetOverlayOK creates a PlaceGetOverlayOK with default headers values
func NewPlaceGetOverlayOK() *PlaceGetOverlayOK {
	return &PlaceGetOverlayOK{}
}

/*PlaceGetOverlayOK handles this case with default header values.

OK
*/
type PlaceGetOverlayOK struct {
	Payload models.SystemObject
}

func (o *PlaceGetOverlayOK) Error() string {
	return fmt.Sprintf("[GET /Place/{type}/overlay/{z}/{Lat}/{Lon}/{width}/{height}][%d] placeGetOverlayOK  %+v", 200, o.Payload)
}

func (o *PlaceGetOverlayOK) GetPayload() models.SystemObject {
	return o.Payload
}

func (o *PlaceGetOverlayOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}