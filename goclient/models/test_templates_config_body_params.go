// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TestTemplatesConfigBodyParams test templates config body params
//
// swagger:model TestTemplatesConfigBodyParams
type TestTemplatesConfigBodyParams struct {

	// Alerts to use as data when testing the template.
	Alerts []*PostableAlert `json:"alerts"`

	// Name of the template file.
	Name string `json:"name,omitempty"`

	// Template string to test.
	Template string `json:"template,omitempty"`
}

// Validate validates this test templates config body params
func (m *TestTemplatesConfigBodyParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlerts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestTemplatesConfigBodyParams) validateAlerts(formats strfmt.Registry) error {
	if swag.IsZero(m.Alerts) { // not required
		return nil
	}

	for i := 0; i < len(m.Alerts); i++ {
		if swag.IsZero(m.Alerts[i]) { // not required
			continue
		}

		if m.Alerts[i] != nil {
			if err := m.Alerts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("alerts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("alerts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this test templates config body params based on the context it is used
func (m *TestTemplatesConfigBodyParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAlerts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestTemplatesConfigBodyParams) contextValidateAlerts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Alerts); i++ {

		if m.Alerts[i] != nil {

			if swag.IsZero(m.Alerts[i]) { // not required
				return nil
			}

			if err := m.Alerts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("alerts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("alerts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TestTemplatesConfigBodyParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestTemplatesConfigBodyParams) UnmarshalBinary(b []byte) error {
	var res TestTemplatesConfigBodyParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
