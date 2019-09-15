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

// StopPointGetByTypeReader is a Reader for the StopPointGetByType structure.
type StopPointGetByTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StopPointGetByTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStopPointGetByTypeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStopPointGetByTypeOK creates a StopPointGetByTypeOK with default headers values
func NewStopPointGetByTypeOK() *StopPointGetByTypeOK {
	return &StopPointGetByTypeOK{}
}

/*StopPointGetByTypeOK handles this case with default header values.

OK
*/
type StopPointGetByTypeOK struct {
	Payload []*models.TflAPIPresentationEntitiesStopPoint
}

func (o *StopPointGetByTypeOK) Error() string {
	return fmt.Sprintf("[GET /StopPoint/Type/{types}][%d] stopPointGetByTypeOK  %+v", 200, o.Payload)
}

func (o *StopPointGetByTypeOK) GetPayload() []*models.TflAPIPresentationEntitiesStopPoint {
	return o.Payload
}

func (o *StopPointGetByTypeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}