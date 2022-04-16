// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Product product
//
// swagger:model Product
type Product struct {

	// category
	Category *CategoryWithoutRequired `json:"category,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// price
	// Required: true
	Price *float64 `json:"price"`

	// product info
	ProductInfo *ProductInfo `json:"productInfo,omitempty"`

	// stock
	// Required: true
	Stock *int64 `json:"stock"`
}

// Validate validates this product
func (m *Product) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCategory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProductInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStock(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Product) validateCategory(formats strfmt.Registry) error {
	if swag.IsZero(m.Category) { // not required
		return nil
	}

	if m.Category != nil {
		if err := m.Category.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("category")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("category")
			}
			return err
		}
	}

	return nil
}

func (m *Product) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Product) validatePrice(formats strfmt.Registry) error {

	if err := validate.Required("price", "body", m.Price); err != nil {
		return err
	}

	return nil
}

func (m *Product) validateProductInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.ProductInfo) { // not required
		return nil
	}

	if m.ProductInfo != nil {
		if err := m.ProductInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("productInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("productInfo")
			}
			return err
		}
	}

	return nil
}

func (m *Product) validateStock(formats strfmt.Registry) error {

	if err := validate.Required("stock", "body", m.Stock); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this product based on the context it is used
func (m *Product) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCategory(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProductInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Product) contextValidateCategory(ctx context.Context, formats strfmt.Registry) error {

	if m.Category != nil {
		if err := m.Category.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("category")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("category")
			}
			return err
		}
	}

	return nil
}

func (m *Product) contextValidateProductInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.ProductInfo != nil {
		if err := m.ProductInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("productInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("productInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Product) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Product) UnmarshalBinary(b []byte) error {
	var res Product
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
