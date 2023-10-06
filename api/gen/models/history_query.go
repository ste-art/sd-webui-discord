// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HistoryQuery history query
//
// swagger:model HistoryQuery
type HistoryQuery struct {

	// command
	Command string `json:"command,omitempty"`
}

// Validate validates this history query
func (m *HistoryQuery) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this history query based on context it is used
func (m *HistoryQuery) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HistoryQuery) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HistoryQuery) UnmarshalBinary(b []byte) error {
	var res HistoryQuery
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
