// Code generated by go-swagger; DO NOT EDIT.

package travel_time

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/dwardu89/tfl-api-swagger/models"
)

// TravelTimeGetCompareOverlayReader is a Reader for the TravelTimeGetCompareOverlay structure.
type TravelTimeGetCompareOverlayReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TravelTimeGetCompareOverlayReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTravelTimeGetCompareOverlayOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTravelTimeGetCompareOverlayOK creates a TravelTimeGetCompareOverlayOK with default headers values
func NewTravelTimeGetCompareOverlayOK() *TravelTimeGetCompareOverlayOK {
	return &TravelTimeGetCompareOverlayOK{}
}

/*TravelTimeGetCompareOverlayOK handles this case with default header values.

OK
*/
type TravelTimeGetCompareOverlayOK struct {
	Payload models.SystemObject
}

func (o *TravelTimeGetCompareOverlayOK) Error() string {
	return fmt.Sprintf("[GET /TravelTimes/compareOverlay/{z}/mapcenter/{mapCenterLat}/{mapCenterLon}/pinlocation/{pinLat}/{pinLon}/dimensions/{width}/{height}][%d] travelTimeGetCompareOverlayOK  %+v", 200, o.Payload)
}

func (o *TravelTimeGetCompareOverlayOK) GetPayload() models.SystemObject {
	return o.Payload
}

func (o *TravelTimeGetCompareOverlayOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
