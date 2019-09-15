// Code generated by go-swagger; DO NOT EDIT.

package stop_point

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/dwardu89/tfl-api-swagger/models"
)

// StopPointCrowdingReader is a Reader for the StopPointCrowding structure.
type StopPointCrowdingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StopPointCrowdingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStopPointCrowdingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStopPointCrowdingOK creates a StopPointCrowdingOK with default headers values
func NewStopPointCrowdingOK() *StopPointCrowdingOK {
	return &StopPointCrowdingOK{}
}

/*StopPointCrowdingOK handles this case with default header values.

OK
*/
type StopPointCrowdingOK struct {
	Payload []*models.TflAPIPresentationEntitiesStopPoint
}

func (o *StopPointCrowdingOK) Error() string {
	return fmt.Sprintf("[GET /StopPoint/{id}/Crowding/{line}][%d] stopPointCrowdingOK  %+v", 200, o.Payload)
}

func (o *StopPointCrowdingOK) GetPayload() []*models.TflAPIPresentationEntitiesStopPoint {
	return o.Payload
}

func (o *StopPointCrowdingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
