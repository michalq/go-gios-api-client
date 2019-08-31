// Code generated by go-swagger; DO NOT EDIT.

package sensor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewSensorParams creates a new SensorParams object
// with the default values initialized.
func NewSensorParams() *SensorParams {
	var ()
	return &SensorParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSensorParamsWithTimeout creates a new SensorParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSensorParamsWithTimeout(timeout time.Duration) *SensorParams {
	var ()
	return &SensorParams{

		timeout: timeout,
	}
}

// NewSensorParamsWithContext creates a new SensorParams object
// with the default values initialized, and the ability to set a context for a request
func NewSensorParamsWithContext(ctx context.Context) *SensorParams {
	var ()
	return &SensorParams{

		Context: ctx,
	}
}

// NewSensorParamsWithHTTPClient creates a new SensorParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSensorParamsWithHTTPClient(client *http.Client) *SensorParams {
	var ()
	return &SensorParams{
		HTTPClient: client,
	}
}

/*SensorParams contains all the parameters to send to the API endpoint
for the sensor operation typically these are written to a http.Request
*/
type SensorParams struct {

	/*SensorID
	  Sensor identifier

	*/
	SensorID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the sensor params
func (o *SensorParams) WithTimeout(timeout time.Duration) *SensorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the sensor params
func (o *SensorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the sensor params
func (o *SensorParams) WithContext(ctx context.Context) *SensorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the sensor params
func (o *SensorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the sensor params
func (o *SensorParams) WithHTTPClient(client *http.Client) *SensorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the sensor params
func (o *SensorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSensorID adds the sensorID to the sensor params
func (o *SensorParams) WithSensorID(sensorID int64) *SensorParams {
	o.SetSensorID(sensorID)
	return o
}

// SetSensorID adds the sensorId to the sensor params
func (o *SensorParams) SetSensorID(sensorID int64) {
	o.SensorID = sensorID
}

// WriteToRequest writes these params to a swagger request
func (o *SensorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param sensorId
	if err := r.SetPathParam("sensorId", swag.FormatInt64(o.SensorID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}