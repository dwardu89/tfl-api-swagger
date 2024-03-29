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

// StopPointMetaCategoriesReader is a Reader for the StopPointMetaCategories structure.
type StopPointMetaCategoriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StopPointMetaCategoriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStopPointMetaCategoriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStopPointMetaCategoriesOK creates a StopPointMetaCategoriesOK with default headers values
func NewStopPointMetaCategoriesOK() *StopPointMetaCategoriesOK {
	return &StopPointMetaCategoriesOK{}
}

/*StopPointMetaCategoriesOK handles this case with default header values.

OK
*/
type StopPointMetaCategoriesOK struct {
	Payload []*models.TflAPIPresentationEntitiesStopPointCategory
}

func (o *StopPointMetaCategoriesOK) Error() string {
	return fmt.Sprintf("[GET /StopPoint/Meta/Categories][%d] stopPointMetaCategoriesOK  %+v", 200, o.Payload)
}

func (o *StopPointMetaCategoriesOK) GetPayload() []*models.TflAPIPresentationEntitiesStopPointCategory {
	return o.Payload
}

func (o *StopPointMetaCategoriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
