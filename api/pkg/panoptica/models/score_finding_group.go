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
	"github.com/go-openapi/validate"
)

// ScoreFindingGroup ScoreFindingGroup
//
// swagger:model ScoreFindingGroup
type ScoreFindingGroup struct {

	// Count
	// Required: true
	Count *int64 `json:"count"`

	// Findings
	// Required: true
	Findings []*ScoreFinding `json:"findings"`
}

// Validate validates this score finding group
func (m *ScoreFindingGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFindings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScoreFindingGroup) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("count", "body", m.Count); err != nil {
		return err
	}

	return nil
}

func (m *ScoreFindingGroup) validateFindings(formats strfmt.Registry) error {

	if err := validate.Required("findings", "body", m.Findings); err != nil {
		return err
	}

	for i := 0; i < len(m.Findings); i++ {
		if swag.IsZero(m.Findings[i]) { // not required
			continue
		}

		if m.Findings[i] != nil {
			if err := m.Findings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("findings" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("findings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this score finding group based on the context it is used
func (m *ScoreFindingGroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFindings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScoreFindingGroup) contextValidateFindings(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Findings); i++ {

		if m.Findings[i] != nil {
			if err := m.Findings[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("findings" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("findings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ScoreFindingGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScoreFindingGroup) UnmarshalBinary(b []byte) error {
	var res ScoreFindingGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}