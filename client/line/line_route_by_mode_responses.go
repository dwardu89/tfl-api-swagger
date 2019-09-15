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

// LineRouteByModeReader is a Reader for the LineRouteByMode structure.
type LineRouteByModeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LineRouteByModeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLineRouteByModeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewLineRouteByModeOK creates a LineRouteByModeOK with default headers values
func NewLineRouteByModeOK() *LineRouteByModeOK {
	return &LineRouteByModeOK{}
}

/*LineRouteByModeOK handles this case with default header values.

OK
*/
type LineRouteByModeOK struct {
	Payload []*models.TflAPIPresentationEntitiesLine
}

func (o *LineRouteByModeOK) Error() string {
	return fmt.Sprintf("[GET /Line/Mode/{modes}/Route][%d] lineRouteByModeOK  %+v", 200, o.Payload)
}

func (o *LineRouteByModeOK) GetPayload() []*models.TflAPIPresentationEntitiesLine {
	return o.Payload
}

func (o *LineRouteByModeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}