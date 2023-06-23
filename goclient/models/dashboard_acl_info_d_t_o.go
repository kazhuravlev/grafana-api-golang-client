// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DashboardACLInfoDTO dashboard ACL info d t o
//
// swagger:model DashboardACLInfoDTO
type DashboardACLInfoDTO struct {

	// created
	// Format: date-time
	Created strfmt.DateTime `json:"created,omitempty"`

	// dashboard Id
	DashboardID int64 `json:"dashboardId,omitempty"`

	// folder Id
	FolderID int64 `json:"folderId,omitempty"`

	// inherited
	Inherited bool `json:"inherited,omitempty"`

	// is folder
	IsFolder bool `json:"isFolder,omitempty"`

	// permission
	Permission PermissionType `json:"permission,omitempty"`

	// permission name
	PermissionName string `json:"permissionName,omitempty"`

	// role
	// Enum: [Viewer Editor Admin]
	Role string `json:"role,omitempty"`

	// slug
	Slug string `json:"slug,omitempty"`

	// team
	Team string `json:"team,omitempty"`

	// team avatar Url
	TeamAvatarURL string `json:"teamAvatarUrl,omitempty"`

	// team email
	TeamEmail string `json:"teamEmail,omitempty"`

	// team Id
	TeamID int64 `json:"teamId,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// uid
	UID string `json:"uid,omitempty"`

	// updated
	// Format: date-time
	Updated strfmt.DateTime `json:"updated,omitempty"`

	// url
	URL string `json:"url,omitempty"`

	// user avatar Url
	UserAvatarURL string `json:"userAvatarUrl,omitempty"`

	// user email
	UserEmail string `json:"userEmail,omitempty"`

	// user Id
	UserID int64 `json:"userId,omitempty"`

	// user login
	UserLogin string `json:"userLogin,omitempty"`
}

// Validate validates this dashboard ACL info d t o
func (m *DashboardACLInfoDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePermission(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdated(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DashboardACLInfoDTO) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DashboardACLInfoDTO) validatePermission(formats strfmt.Registry) error {
	if swag.IsZero(m.Permission) { // not required
		return nil
	}

	if err := m.Permission.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("permission")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("permission")
		}
		return err
	}

	return nil
}

var dashboardAclInfoDTOTypeRolePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Viewer","Editor","Admin"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		dashboardAclInfoDTOTypeRolePropEnum = append(dashboardAclInfoDTOTypeRolePropEnum, v)
	}
}

const (

	// DashboardACLInfoDTORoleViewer captures enum value "Viewer"
	DashboardACLInfoDTORoleViewer string = "Viewer"

	// DashboardACLInfoDTORoleEditor captures enum value "Editor"
	DashboardACLInfoDTORoleEditor string = "Editor"

	// DashboardACLInfoDTORoleAdmin captures enum value "Admin"
	DashboardACLInfoDTORoleAdmin string = "Admin"
)

// prop value enum
func (m *DashboardACLInfoDTO) validateRoleEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, dashboardAclInfoDTOTypeRolePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DashboardACLInfoDTO) validateRole(formats strfmt.Registry) error {
	if swag.IsZero(m.Role) { // not required
		return nil
	}

	// value enum
	if err := m.validateRoleEnum("role", "body", m.Role); err != nil {
		return err
	}

	return nil
}

func (m *DashboardACLInfoDTO) validateUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.Updated) { // not required
		return nil
	}

	if err := validate.FormatOf("updated", "body", "date-time", m.Updated.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this dashboard ACL info d t o based on the context it is used
func (m *DashboardACLInfoDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePermission(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DashboardACLInfoDTO) contextValidatePermission(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Permission) { // not required
		return nil
	}

	if err := m.Permission.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("permission")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("permission")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DashboardACLInfoDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DashboardACLInfoDTO) UnmarshalBinary(b []byte) error {
	var res DashboardACLInfoDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
