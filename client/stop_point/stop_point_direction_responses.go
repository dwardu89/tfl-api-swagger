// Code generated by go-swagger; DO NOT EDIT.

package stop_point

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// StopPointDirectionReader is a Reader for the StopPointDirection structure.
type StopPointDirectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StopPointDirectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStopPointDirectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStopPointDirectionOK creates a StopPointDirectionOK with default headers values
func NewStopPointDirectionOK() *StopPointDirectionOK {
	return &StopPointDirectionOK{}
}

/*StopPointDirectionOK handles this case with default header values.

OK
*/
type StopPointDirectionOK struct {
	Payload string
}

func (o *StopPointDirectionOK) Error() string {
	return fmt.Sprintf("[GET /StopPoint/{id}/DirectionTo/{toStopPointId}][%d] stopPointDirectionOK  %+v", 200, o.Payload)
}

func (o *StopPointDirectionOK) GetPayload() string {
	return o.Payload
}

func (o *StopPointDirectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
