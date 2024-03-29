// Code generated by go-swagger; DO NOT EDIT.

package bike_point

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/dwardu89/tfl-api-swagger/models"
)

// BikePointSearchReader is a Reader for the BikePointSearch structure.
type BikePointSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BikePointSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBikePointSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewBikePointSearchOK creates a BikePointSearchOK with default headers values
func NewBikePointSearchOK() *BikePointSearchOK {
	return &BikePointSearchOK{}
}

/*BikePointSearchOK handles this case with default header values.

OK
*/
type BikePointSearchOK struct {
	Payload []*models.TflAPIPresentationEntitiesPlace
}

func (o *BikePointSearchOK) Error() string {
	return fmt.Sprintf("[GET /BikePoint/Search][%d] bikePointSearchOK  %+v", 200, o.Payload)
}

func (o *BikePointSearchOK) GetPayload() []*models.TflAPIPresentationEntitiesPlace {
	return o.Payload
}

func (o *BikePointSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
