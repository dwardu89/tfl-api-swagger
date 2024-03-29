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

// OccupancyGetAllChargeConnectorStatusReader is a Reader for the OccupancyGetAllChargeConnectorStatus structure.
type OccupancyGetAllChargeConnectorStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OccupancyGetAllChargeConnectorStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOccupancyGetAllChargeConnectorStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewOccupancyGetAllChargeConnectorStatusOK creates a OccupancyGetAllChargeConnectorStatusOK with default headers values
func NewOccupancyGetAllChargeConnectorStatusOK() *OccupancyGetAllChargeConnectorStatusOK {
	return &OccupancyGetAllChargeConnectorStatusOK{}
}

/*OccupancyGetAllChargeConnectorStatusOK handles this case with default header values.

OK
*/
type OccupancyGetAllChargeConnectorStatusOK struct {
	Payload []*models.TflAPIPresentationEntitiesChargeConnectorOccupancy
}

func (o *OccupancyGetAllChargeConnectorStatusOK) Error() string {
	return fmt.Sprintf("[GET /Occupancy/ChargeConnector][%d] occupancyGetAllChargeConnectorStatusOK  %+v", 200, o.Payload)
}

func (o *OccupancyGetAllChargeConnectorStatusOK) GetPayload() []*models.TflAPIPresentationEntitiesChargeConnectorOccupancy {
	return o.Payload
}

func (o *OccupancyGetAllChargeConnectorStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
