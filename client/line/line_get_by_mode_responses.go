// Code generated by go-swagger; DO NOT EDIT.

package line

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/dwardu89/tfl-api-swagger/models"
)

// LineGetByModeReader is a Reader for the LineGetByMode structure.
type LineGetByModeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LineGetByModeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLineGetByModeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewLineGetByModeOK creates a LineGetByModeOK with default headers values
func NewLineGetByModeOK() *LineGetByModeOK {
	return &LineGetByModeOK{}
}

/*LineGetByModeOK handles this case with default header values.

OK
*/
type LineGetByModeOK struct {
	Payload []*models.TflAPIPresentationEntitiesLine
}

func (o *LineGetByModeOK) Error() string {
	return fmt.Sprintf("[GET /Line/Mode/{modes}][%d] lineGetByModeOK  %+v", 200, o.Payload)
}

func (o *LineGetByModeOK) GetPayload() []*models.TflAPIPresentationEntitiesLine {
	return o.Payload
}

func (o *LineGetByModeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}