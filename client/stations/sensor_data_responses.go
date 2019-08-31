// Code generated by go-swagger; DO NOT EDIT.

package stations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/michalq/go-gios-api-client/models"
)

// SensorDataReader is a Reader for the SensorData structure.
type SensorDataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SensorDataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSensorDataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSensorDataOK creates a SensorDataOK with default headers values
func NewSensorDataOK() *SensorDataOK {
	return &SensorDataOK{}
}

/*SensorDataOK handles this case with default header values.

Ok
*/
type SensorDataOK struct {
	Payload models.SensorsResponse200
}

func (o *SensorDataOK) Error() string {
	return fmt.Sprintf("[GET /station/sensors/{stationId}][%d] sensorDataOK  %+v", 200, o.Payload)
}

func (o *SensorDataOK) GetPayload() models.SensorsResponse200 {
	return o.Payload
}

func (o *SensorDataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}