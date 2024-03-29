// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SensorResponse200 sensor response200
// swagger:model SensorResponse200
type SensorResponse200 struct {

	// key
	// Required: true
	Key AcquisitionType `json:"key"`

	// values
	// Required: true
	Values []*SensorAcquisition `json:"values"`
}

// Validate validates this sensor response200
func (m *SensorResponse200) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValues(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SensorResponse200) validateKey(formats strfmt.Registry) error {

	if err := m.Key.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("key")
		}
		return err
	}

	return nil
}

func (m *SensorResponse200) validateValues(formats strfmt.Registry) error {

	if err := validate.Required("values", "body", m.Values); err != nil {
		return err
	}

	for i := 0; i < len(m.Values); i++ {
		if swag.IsZero(m.Values[i]) { // not required
			continue
		}

		if m.Values[i] != nil {
			if err := m.Values[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("values" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SensorResponse200) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SensorResponse200) UnmarshalBinary(b []byte) error {
	var res SensorResponse200
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
