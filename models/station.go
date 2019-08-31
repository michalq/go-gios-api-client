// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Station station
// swagger:model Station
type Station struct {

	// city
	City *City `json:"city,omitempty"`

	// gegr lat
	// Required: true
	GegrLat *string `json:"gegrLat"`

	// gegr lon
	// Required: true
	GegrLon *string `json:"gegrLon"`

	// id
	// Required: true
	ID *int64 `json:"id"`

	// station name
	// Required: true
	StationName *string `json:"stationName"`
}

// Validate validates this station
func (m *Station) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGegrLat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGegrLon(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStationName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Station) validateCity(formats strfmt.Registry) error {

	if swag.IsZero(m.City) { // not required
		return nil
	}

	if m.City != nil {
		if err := m.City.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("city")
			}
			return err
		}
	}

	return nil
}

func (m *Station) validateGegrLat(formats strfmt.Registry) error {

	if err := validate.Required("gegrLat", "body", m.GegrLat); err != nil {
		return err
	}

	return nil
}

func (m *Station) validateGegrLon(formats strfmt.Registry) error {

	if err := validate.Required("gegrLon", "body", m.GegrLon); err != nil {
		return err
	}

	return nil
}

func (m *Station) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Station) validateStationName(formats strfmt.Registry) error {

	if err := validate.Required("stationName", "body", m.StationName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Station) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Station) UnmarshalBinary(b []byte) error {
	var res Station
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
