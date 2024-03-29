// Code generated by go-swagger; DO NOT EDIT.

package accident_stats

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/dwardu89/tfl-api-swagger/models"
)

// AccidentStatsGetReader is a Reader for the AccidentStatsGet structure.
type AccidentStatsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AccidentStatsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAccidentStatsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAccidentStatsGetOK creates a AccidentStatsGetOK with default headers values
func NewAccidentStatsGetOK() *AccidentStatsGetOK {
	return &AccidentStatsGetOK{}
}

/*AccidentStatsGetOK handles this case with default header values.

OK
*/
type AccidentStatsGetOK struct {
	Payload []*models.TflAPIPresentationEntitiesAccidentStatsAccidentDetail
}

func (o *AccidentStatsGetOK) Error() string {
	return fmt.Sprintf("[GET /AccidentStats/{year}][%d] accidentStatsGetOK  %+v", 200, o.Payload)
}

func (o *AccidentStatsGetOK) GetPayload() []*models.TflAPIPresentationEntitiesAccidentStatsAccidentDetail {
	return o.Payload
}

func (o *AccidentStatsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
