// Code generated by go-swagger; DO NOT EDIT.

package occupancy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/dwardu89/tfl-api-swagger/models"
)

// OccupancyGetBikePointsOccupanciesReader is a Reader for the OccupancyGetBikePointsOccupancies structure.
type OccupancyGetBikePointsOccupanciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OccupancyGetBikePointsOccupanciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOccupancyGetBikePointsOccupanciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewOccupancyGetBikePointsOccupanciesOK creates a OccupancyGetBikePointsOccupanciesOK with default headers values
func NewOccupancyGetBikePointsOccupanciesOK() *OccupancyGetBikePointsOccupanciesOK {
	return &OccupancyGetBikePointsOccupanciesOK{}
}

/*OccupancyGetBikePointsOccupanciesOK handles this case with default header values.

OK
*/
type OccupancyGetBikePointsOccupanciesOK struct {
	Payload []*models.TflAPIPresentationEntitiesBikePointOccupancy
}

func (o *OccupancyGetBikePointsOccupanciesOK) Error() string {
	return fmt.Sprintf("[GET /Occupancy/BikePoints/{ids}][%d] occupancyGetBikePointsOccupanciesOK  %+v", 200, o.Payload)
}

func (o *OccupancyGetBikePointsOccupanciesOK) GetPayload() []*models.TflAPIPresentationEntitiesBikePointOccupancy {
	return o.Payload
}

func (o *OccupancyGetBikePointsOccupanciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}