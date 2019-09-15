// Code generated by go-swagger; DO NOT EDIT.

package air_quality

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/dwardu89/tfl-api-swagger/models"
)

// AirQualityGetReader is a Reader for the AirQualityGet structure.
type AirQualityGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AirQualityGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAirQualityGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAirQualityGetOK creates a AirQualityGetOK with default headers values
func NewAirQualityGetOK() *AirQualityGetOK {
	return &AirQualityGetOK{}
}

/*AirQualityGetOK handles this case with default header values.

OK
*/
type AirQualityGetOK struct {
	Payload models.SystemObject
}

func (o *AirQualityGetOK) Error() string {
	return fmt.Sprintf("[GET /AirQuality][%d] airQualityGetOK  %+v", 200, o.Payload)
}

func (o *AirQualityGetOK) GetPayload() models.SystemObject {
	return o.Payload
}

func (o *AirQualityGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}