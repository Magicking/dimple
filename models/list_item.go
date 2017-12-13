// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ListItem list item
// swagger:model list_item

type ListItem struct {

	// Some string
	Addr string `json:"addr,omitempty"`

	// amount sent
	Amount string `json:"amount,omitempty"`

	// txid
	Txid string `json:"txid,omitempty"`
}

/* polymorph list_item addr false */

/* polymorph list_item amount false */

/* polymorph list_item txid false */

// Validate validates this list item
func (m *ListItem) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ListItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListItem) UnmarshalBinary(b []byte) error {
	var res ListItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
