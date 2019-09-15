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

// StopPointGetReader is a Reader for the StopPointGet structure.
type StopPointGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StopPointGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStopPointGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStopPointGetOK creates a StopPointGetOK with default headers values
func NewStopPointGetOK() *StopPointGetOK {
	return &StopPointGetOK{}
}

/*StopPointGetOK handles this case with default header values.

OK
*/
type StopPointGetOK struct {
	Payload []*models.TflAPIPresentationEntitiesStopPoint
}

func (o *StopPointGetOK) Error() string {
	return fmt.Sprintf("[GET /StopPoint/{ids}][%d] stopPointGetOK  %+v", 200, o.Payload)
}

func (o *StopPointGetOK) GetPayload() []*models.TflAPIPresentationEntitiesStopPoint {
	return o.Payload
}

func (o *StopPointGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
