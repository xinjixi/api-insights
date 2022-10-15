// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// ScoreExitStatusEnum ScoreExitStatusEnum
//
// An enumeration.
//
// swagger:model ScoreExitStatusEnum
type ScoreExitStatusEnum string

func NewScoreExitStatusEnum(value ScoreExitStatusEnum) *ScoreExitStatusEnum {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ScoreExitStatusEnum.
func (m ScoreExitStatusEnum) Pointer() *ScoreExitStatusEnum {
	return &m
}

const (

	// ScoreExitStatusEnumSUCCESS captures enum value "SUCCESS"
	ScoreExitStatusEnumSUCCESS ScoreExitStatusEnum = "SUCCESS"

	// ScoreExitStatusEnumNETWORKFAILURE captures enum value "NETWORK_FAILURE"
	ScoreExitStatusEnumNETWORKFAILURE ScoreExitStatusEnum = "NETWORK_FAILURE"

	// ScoreExitStatusEnumTOOLONG captures enum value "TOO_LONG"
	ScoreExitStatusEnumTOOLONG ScoreExitStatusEnum = "TOO_LONG"

	// ScoreExitStatusEnumPARSEERROR captures enum value "PARSE_ERROR"
	ScoreExitStatusEnumPARSEERROR ScoreExitStatusEnum = "PARSE_ERROR"

	// ScoreExitStatusEnumGENERICFAILURE captures enum value "GENERIC_FAILURE"
	ScoreExitStatusEnumGENERICFAILURE ScoreExitStatusEnum = "GENERIC_FAILURE"
)

// for schema
var scoreExitStatusEnumEnum []interface{}

func init() {
	var res []ScoreExitStatusEnum
	if err := json.Unmarshal([]byte(`["SUCCESS","NETWORK_FAILURE","TOO_LONG","PARSE_ERROR","GENERIC_FAILURE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		scoreExitStatusEnumEnum = append(scoreExitStatusEnumEnum, v)
	}
}

func (m ScoreExitStatusEnum) validateScoreExitStatusEnumEnum(path, location string, value ScoreExitStatusEnum) error {
	if err := validate.EnumCase(path, location, value, scoreExitStatusEnumEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this score exit status enum
func (m ScoreExitStatusEnum) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateScoreExitStatusEnumEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this score exit status enum based on context it is used
func (m ScoreExitStatusEnum) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}