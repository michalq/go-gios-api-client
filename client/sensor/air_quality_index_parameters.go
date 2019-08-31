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

// NewAirQualityIndexParams creates a new AirQualityIndexParams object
// with the default values initialized.
func NewAirQualityIndexParams() *AirQualityIndexParams {
	var ()
	return &AirQualityIndexParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAirQualityIndexParamsWithTimeout creates a new AirQualityIndexParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAirQualityIndexParamsWithTimeout(timeout time.Duration) *AirQualityIndexParams {
	var ()
	return &AirQualityIndexParams{

		timeout: timeout,
	}
}

// NewAirQualityIndexParamsWithContext creates a new AirQualityIndexParams object
// with the default values initialized, and the ability to set a context for a request
func NewAirQualityIndexParamsWithContext(ctx context.Context) *AirQualityIndexParams {
	var ()
	return &AirQualityIndexParams{

		Context: ctx,
	}
}

// NewAirQualityIndexParamsWithHTTPClient creates a new AirQualityIndexParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAirQualityIndexParamsWithHTTPClient(client *http.Client) *AirQualityIndexParams {
	var ()
	return &AirQualityIndexParams{
		HTTPClient: client,
	}
}

/*AirQualityIndexParams contains all the parameters to send to the API endpoint
for the air quality index operation typically these are written to a http.Request
*/
type AirQualityIndexParams struct {

	/*StationID
	  ID of pet to return

	*/
	StationID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the air quality index params
func (o *AirQualityIndexParams) WithTimeout(timeout time.Duration) *AirQualityIndexParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the air quality index params
func (o *AirQualityIndexParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the air quality index params
func (o *AirQualityIndexParams) WithContext(ctx context.Context) *AirQualityIndexParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the air quality index params
func (o *AirQualityIndexParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the air quality index params
func (o *AirQualityIndexParams) WithHTTPClient(client *http.Client) *AirQualityIndexParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the air quality index params
func (o *AirQualityIndexParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithStationID adds the stationID to the air quality index params
func (o *AirQualityIndexParams) WithStationID(stationID int64) *AirQualityIndexParams {
	o.SetStationID(stationID)
	return o
}

// SetStationID adds the stationId to the air quality index params
func (o *AirQualityIndexParams) SetStationID(stationID int64) {
	o.StationID = stationID
}

// WriteToRequest writes these params to a swagger request
func (o *AirQualityIndexParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param stationId
	if err := r.SetPathParam("stationId", swag.FormatInt64(o.StationID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
