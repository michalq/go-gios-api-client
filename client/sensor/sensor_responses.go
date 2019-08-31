// Code generated by go-swagger; DO NOT EDIT.

package sensor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "gios-api-client/models"
)

// SensorReader is a Reader for the Sensor structure.
type SensorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SensorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSensorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSensorOK creates a SensorOK with default headers values
func NewSensorOK() *SensorOK {
	return &SensorOK{}
}

/*SensorOK handles this case with default header values.

Ok
*/
type SensorOK struct {
	Payload *models.SensorResponse200
}

func (o *SensorOK) Error() string {
	return fmt.Sprintf("[GET /data/getData/{sensorId}][%d] sensorOK  %+v", 200, o.Payload)
}

func (o *SensorOK) GetPayload() *models.SensorResponse200 {
	return o.Payload
}

func (o *SensorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SensorResponse200)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
