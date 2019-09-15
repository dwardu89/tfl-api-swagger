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

// StopPointReachableFromReader is a Reader for the StopPointReachableFrom structure.
type StopPointReachableFromReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StopPointReachableFromReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStopPointReachableFromOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStopPointReachableFromOK creates a StopPointReachableFromOK with default headers values
func NewStopPointReachableFromOK() *StopPointReachableFromOK {
	return &StopPointReachableFromOK{}
}

/*StopPointReachableFromOK handles this case with default header values.

OK
*/
type StopPointReachableFromOK struct {
	Payload []*models.TflAPIPresentationEntitiesStopPoint
}

func (o *StopPointReachableFromOK) Error() string {
	return fmt.Sprintf("[GET /StopPoint/{id}/CanReachOnLine/{lineId}][%d] stopPointReachableFromOK  %+v", 200, o.Payload)
}

func (o *StopPointReachableFromOK) GetPayload() []*models.TflAPIPresentationEntitiesStopPoint {
	return o.Payload
}

func (o *StopPointReachableFromOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}