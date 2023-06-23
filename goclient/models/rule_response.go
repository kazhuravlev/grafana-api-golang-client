// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RuleResponse rule response
//
// swagger:model RuleResponse
type RuleResponse struct {

	// data
	Data *RuleDiscovery `json:"data,omitempty"`

	// error
	Error string `json:"error,omitempty"`

	// error type
	ErrorType ErrorType `json:"errorType,omitempty"`

	// status
	// Required: true
	Status *string `json:"status"`
}

// Validate validates this rule response
func (m *RuleResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateErrorType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RuleResponse) validateData(formats strfmt.Registry) error {
	if swag.IsZero(m.Data) { // not required
		return nil
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

func (m *RuleResponse) validateErrorType(formats strfmt.Registry) error {
	if swag.IsZero(m.ErrorType) { // not required
		return nil
	}

	if err := m.ErrorType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("errorType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("errorType")
		}
		return err
	}

	return nil
}

func (m *RuleResponse) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this rule response based on the context it is used
func (m *RuleResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateErrorType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RuleResponse) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if m.Data != nil {

		if swag.IsZero(m.Data) { // not required
			return nil
		}

		if err := m.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

func (m *RuleResponse) contextValidateErrorType(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.ErrorType) { // not required
		return nil
	}

	if err := m.ErrorType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("errorType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("errorType")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RuleResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RuleResponse) UnmarshalBinary(b []byte) error {
	var res RuleResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
