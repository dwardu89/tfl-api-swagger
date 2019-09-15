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

// LineRouteReader is a Reader for the LineRoute structure.
type LineRouteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LineRouteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLineRouteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewLineRouteOK creates a LineRouteOK with default headers values
func NewLineRouteOK() *LineRouteOK {
	return &LineRouteOK{}
}

/*LineRouteOK handles this case with default header values.

OK
*/
type LineRouteOK struct {
	Payload []*models.TflAPIPresentationEntitiesLine
}

func (o *LineRouteOK) Error() string {
	return fmt.Sprintf("[GET /Line/Route][%d] lineRouteOK  %+v", 200, o.Payload)
}

func (o *LineRouteOK) GetPayload() []*models.TflAPIPresentationEntitiesLine {
	return o.Payload
}

func (o *LineRouteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
