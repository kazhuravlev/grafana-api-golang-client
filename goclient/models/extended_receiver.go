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

// ExtendedReceiver extended receiver
//
// swagger:model ExtendedReceiver
type ExtendedReceiver struct {

	// email configs
	EmailConfigs *EmailConfig `json:"email_configs,omitempty"`

	// grafana managed receiver
	GrafanaManagedReceiver *PostableGrafanaReceiver `json:"grafana_managed_receiver,omitempty"`

	// opsgenie configs
	OpsgenieConfigs *OpsGenieConfig `json:"opsgenie_configs,omitempty"`

	// pagerduty configs
	PagerdutyConfigs *PagerdutyConfig `json:"pagerduty_configs,omitempty"`

	// pushover configs
	PushoverConfigs *PushoverConfig `json:"pushover_configs,omitempty"`

	// slack configs
	SlackConfigs *SlackConfig `json:"slack_configs,omitempty"`

	// victorops configs
	VictoropsConfigs *VictorOpsConfig `json:"victorops_configs,omitempty"`

	// webhook configs
	WebhookConfigs *WebhookConfig `json:"webhook_configs,omitempty"`

	// wechat configs
	WechatConfigs *WechatConfig `json:"wechat_configs,omitempty"`
}

// Validate validates this extended receiver
func (m *ExtendedReceiver) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmailConfigs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGrafanaManagedReceiver(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOpsgenieConfigs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePagerdutyConfigs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePushoverConfigs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSlackConfigs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVictoropsConfigs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWebhookConfigs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWechatConfigs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExtendedReceiver) validateEmailConfigs(formats strfmt.Registry) error {
	if swag.IsZero(m.EmailConfigs) { // not required
		return nil
	}

	if m.EmailConfigs != nil {
		if err := m.EmailConfigs.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("email_configs")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("email_configs")
			}
			return err
		}
	}

	return nil
}

func (m *ExtendedReceiver) validateGrafanaManagedReceiver(formats strfmt.Registry) error {
	if swag.IsZero(m.GrafanaManagedReceiver) { // not required
		return nil
	}

	if m.GrafanaManagedReceiver != nil {
		if err := m.GrafanaManagedReceiver.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("grafana_managed_receiver")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("grafana_managed_receiver")
			}
			return err
		}
	}

	return nil
}

func (m *ExtendedReceiver) validateOpsgenieConfigs(formats strfmt.Registry) error {
	if swag.IsZero(m.OpsgenieConfigs) { // not required
		return nil
	}

	if m.OpsgenieConfigs != nil {
		if err := m.OpsgenieConfigs.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("opsgenie_configs")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("opsgenie_configs")
			}
			return err
		}
	}

	return nil
}

func (m *ExtendedReceiver) validatePagerdutyConfigs(formats strfmt.Registry) error {
	if swag.IsZero(m.PagerdutyConfigs) { // not required
		return nil
	}

	if m.PagerdutyConfigs != nil {
		if err := m.PagerdutyConfigs.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pagerduty_configs")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pagerduty_configs")
			}
			return err
		}
	}

	return nil
}

func (m *ExtendedReceiver) validatePushoverConfigs(formats strfmt.Registry) error {
	if swag.IsZero(m.PushoverConfigs) { // not required
		return nil
	}

	if m.PushoverConfigs != nil {
		if err := m.PushoverConfigs.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pushover_configs")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pushover_configs")
			}
			return err
		}
	}

	return nil
}

func (m *ExtendedReceiver) validateSlackConfigs(formats strfmt.Registry) error {
	if swag.IsZero(m.SlackConfigs) { // not required
		return nil
	}

	if m.SlackConfigs != nil {
		if err := m.SlackConfigs.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("slack_configs")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("slack_configs")
			}
			return err
		}
	}

	return nil
}

func (m *ExtendedReceiver) validateVictoropsConfigs(formats strfmt.Registry) error {
	if swag.IsZero(m.VictoropsConfigs) { // not required
		return nil
	}

	if m.VictoropsConfigs != nil {
		if err := m.VictoropsConfigs.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("victorops_configs")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("victorops_configs")
			}
			return err
		}
	}

	return nil
}

func (m *ExtendedReceiver) validateWebhookConfigs(formats strfmt.Registry) error {
	if swag.IsZero(m.WebhookConfigs) { // not required
		return nil
	}

	if m.WebhookConfigs != nil {
		if err := m.WebhookConfigs.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("webhook_configs")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("webhook_configs")
			}
			return err
		}
	}

	return nil
}

func (m *ExtendedReceiver) validateWechatConfigs(formats strfmt.Registry) error {
	if swag.IsZero(m.WechatConfigs) { // not required
		return nil
	}

	if m.WechatConfigs != nil {
		if err := m.WechatConfigs.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("wechat_configs")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("wechat_configs")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this extended receiver based on the context it is used
func (m *ExtendedReceiver) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEmailConfigs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGrafanaManagedReceiver(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOpsgenieConfigs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePagerdutyConfigs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePushoverConfigs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSlackConfigs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVictoropsConfigs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWebhookConfigs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWechatConfigs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExtendedReceiver) contextValidateEmailConfigs(ctx context.Context, formats strfmt.Registry) error {

	if m.EmailConfigs != nil {

		if swag.IsZero(m.EmailConfigs) { // not required
			return nil
		}

		if err := m.EmailConfigs.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("email_configs")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("email_configs")
			}
			return err
		}
	}

	return nil
}

func (m *ExtendedReceiver) contextValidateGrafanaManagedReceiver(ctx context.Context, formats strfmt.Registry) error {

	if m.GrafanaManagedReceiver != nil {

		if swag.IsZero(m.GrafanaManagedReceiver) { // not required
			return nil
		}

		if err := m.GrafanaManagedReceiver.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("grafana_managed_receiver")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("grafana_managed_receiver")
			}
			return err
		}
	}

	return nil
}

func (m *ExtendedReceiver) contextValidateOpsgenieConfigs(ctx context.Context, formats strfmt.Registry) error {

	if m.OpsgenieConfigs != nil {

		if swag.IsZero(m.OpsgenieConfigs) { // not required
			return nil
		}

		if err := m.OpsgenieConfigs.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("opsgenie_configs")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("opsgenie_configs")
			}
			return err
		}
	}

	return nil
}

func (m *ExtendedReceiver) contextValidatePagerdutyConfigs(ctx context.Context, formats strfmt.Registry) error {

	if m.PagerdutyConfigs != nil {

		if swag.IsZero(m.PagerdutyConfigs) { // not required
			return nil
		}

		if err := m.PagerdutyConfigs.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pagerduty_configs")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pagerduty_configs")
			}
			return err
		}
	}

	return nil
}

func (m *ExtendedReceiver) contextValidatePushoverConfigs(ctx context.Context, formats strfmt.Registry) error {

	if m.PushoverConfigs != nil {

		if swag.IsZero(m.PushoverConfigs) { // not required
			return nil
		}

		if err := m.PushoverConfigs.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pushover_configs")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pushover_configs")
			}
			return err
		}
	}

	return nil
}

func (m *ExtendedReceiver) contextValidateSlackConfigs(ctx context.Context, formats strfmt.Registry) error {

	if m.SlackConfigs != nil {

		if swag.IsZero(m.SlackConfigs) { // not required
			return nil
		}

		if err := m.SlackConfigs.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("slack_configs")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("slack_configs")
			}
			return err
		}
	}

	return nil
}

func (m *ExtendedReceiver) contextValidateVictoropsConfigs(ctx context.Context, formats strfmt.Registry) error {

	if m.VictoropsConfigs != nil {

		if swag.IsZero(m.VictoropsConfigs) { // not required
			return nil
		}

		if err := m.VictoropsConfigs.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("victorops_configs")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("victorops_configs")
			}
			return err
		}
	}

	return nil
}

func (m *ExtendedReceiver) contextValidateWebhookConfigs(ctx context.Context, formats strfmt.Registry) error {

	if m.WebhookConfigs != nil {

		if swag.IsZero(m.WebhookConfigs) { // not required
			return nil
		}

		if err := m.WebhookConfigs.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("webhook_configs")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("webhook_configs")
			}
			return err
		}
	}

	return nil
}

func (m *ExtendedReceiver) contextValidateWechatConfigs(ctx context.Context, formats strfmt.Registry) error {

	if m.WechatConfigs != nil {

		if swag.IsZero(m.WechatConfigs) { // not required
			return nil
		}

		if err := m.WechatConfigs.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("wechat_configs")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("wechat_configs")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExtendedReceiver) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExtendedReceiver) UnmarshalBinary(b []byte) error {
	var res ExtendedReceiver
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
