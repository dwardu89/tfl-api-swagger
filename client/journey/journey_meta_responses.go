// Code generated by go-swagger; DO NOT EDIT.

package journey

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/dwardu89/tfl-api-swagger/models"
)

// JourneyMetaReader is a Reader for the JourneyMeta structure.
type JourneyMetaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JourneyMetaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewJourneyMetaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewJourneyMetaOK creates a JourneyMetaOK with default headers values
func NewJourneyMetaOK() *JourneyMetaOK {
	return &JourneyMetaOK{}
}

/*JourneyMetaOK handles this case with default header values.

OK
*/
type JourneyMetaOK struct {
	Payload []*models.TflAPIPresentationEntitiesMode
}

func (o *JourneyMetaOK) Error() string {
	return fmt.Sprintf("[GET /Journey/Meta/Modes][%d] journeyMetaOK  %+v", 200, o.Payload)
}

func (o *JourneyMetaOK) GetPayload() []*models.TflAPIPresentationEntitiesMode {
	return o.Payload
}

func (o *JourneyMetaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
