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

// AddDataSourceOKBody add data source o k body
//
// swagger:model addDataSourceOKBody
type AddDataSourceOKBody struct {

	// datasource
	// Required: true
	Datasource *DataSource `json:"datasource"`

	// ID Identifier of the new data source.
	// Example: 65
	// Required: true
	ID *int64 `json:"id"`

	// Message Message of the deleted dashboard.
	// Example: Data source added
	// Required: true
	Message *string `json:"message"`

	// Name of the new data source.
	// Example: My Data source
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this add data source o k body
func (m *AddDataSourceOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDatasource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AddDataSourceOKBody) validateDatasource(formats strfmt.Registry) error {

	if err := validate.Required("datasource", "body", m.Datasource); err != nil {
		return err
	}

	if m.Datasource != nil {
		if err := m.Datasource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("datasource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("datasource")
			}
			return err
		}
	}

	return nil
}

func (m *AddDataSourceOKBody) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *AddDataSourceOKBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

func (m *AddDataSourceOKBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this add data source o k body based on the context it is used
func (m *AddDataSourceOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDatasource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AddDataSourceOKBody) contextValidateDatasource(ctx context.Context, formats strfmt.Registry) error {

	if m.Datasource != nil {

		if err := m.Datasource.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("datasource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("datasource")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AddDataSourceOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AddDataSourceOKBody) UnmarshalBinary(b []byte) error {
	var res AddDataSourceOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
