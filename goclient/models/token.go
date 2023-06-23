// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Token token
//
// swagger:model Token
type Token struct {

	// account
	Account string `json:"account,omitempty"`

	// company
	Company string `json:"company,omitempty"`

	// details url
	DetailsURL string `json:"details_url,omitempty"`

	// exp
	Exp int64 `json:"exp,omitempty"`

	// iat
	Iat int64 `json:"iat,omitempty"`

	// included users
	IncludedUsers int64 `json:"included_users,omitempty"`

	// iss
	Iss string `json:"iss,omitempty"`

	// jti
	Jti string `json:"jti,omitempty"`

	// lexp
	Lexp int64 `json:"lexp,omitempty"`

	// lic exp warn days
	LicExpWarnDays int64 `json:"lic_exp_warn_days,omitempty"`

	// lid
	Lid string `json:"lid,omitempty"`

	// limit by
	LimitBy string `json:"limit_by,omitempty"`

	// max concurrent user sessions
	MaxConcurrentUserSessions int64 `json:"max_concurrent_user_sessions,omitempty"`

	// nbf
	Nbf int64 `json:"nbf,omitempty"`

	// prod
	Prod []string `json:"prod"`

	// slug
	Slug string `json:"slug,omitempty"`

	// status
	Status TokenStatus `json:"status,omitempty"`

	// sub
	Sub string `json:"sub,omitempty"`

	// tok exp warn days
	TokExpWarnDays int64 `json:"tok_exp_warn_days,omitempty"`

	// trial
	Trial bool `json:"trial,omitempty"`

	// trial exp
	TrialExp int64 `json:"trial_exp,omitempty"`

	// update days
	UpdateDays int64 `json:"update_days,omitempty"`

	// usage billing
	UsageBilling bool `json:"usage_billing,omitempty"`
}

// Validate validates this token
func (m *Token) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Token) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("status")
		}
		return err
	}

	return nil
}

// ContextValidate validate this token based on the context it is used
func (m *Token) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Token) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("status")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Token) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Token) UnmarshalBinary(b []byte) error {
	var res Token
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
