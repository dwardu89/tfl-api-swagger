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

// LineMetaModesReader is a Reader for the LineMetaModes structure.
type LineMetaModesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LineMetaModesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLineMetaModesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewLineMetaModesOK creates a LineMetaModesOK with default headers values
func NewLineMetaModesOK() *LineMetaModesOK {
	return &LineMetaModesOK{}
}

/*LineMetaModesOK handles this case with default header values.

OK
*/
type LineMetaModesOK struct {
	Payload []*models.TflAPIPresentationEntitiesMode
}

func (o *LineMetaModesOK) Error() string {
	return fmt.Sprintf("[GET /Line/Meta/Modes][%d] lineMetaModesOK  %+v", 200, o.Payload)
}

func (o *LineMetaModesOK) GetPayload() []*models.TflAPIPresentationEntitiesMode {
	return o.Payload
}

func (o *LineMetaModesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
