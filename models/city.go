// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// City city
// swagger:model City
type City struct {

	// address street
	AddressStreet string `json:"addressStreet,omitempty"`

	// commune
	Commune *Commune `json:"commune,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this city
func (m *City) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommune(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *City) validateCommune(formats strfmt.Registry) error {

	if swag.IsZero(m.Commune) { // not required
		return nil
	}

	if m.Commune != nil {
		if err := m.Commune.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("commune")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *City) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *City) UnmarshalBinary(b []byte) error {
	var res City
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
