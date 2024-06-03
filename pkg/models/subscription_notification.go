// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SubscriptionNotification subscription notification
// swagger:model subscriptionNotification
type SubscriptionNotification struct {

	// event type
	EventType EventType `json:"eventType,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// version
	Version int64 `json:"version,omitempty"`

	// x apps
	XApps AllDeployedXapps `json:"xApps,omitempty"`
}

// Validate validates this subscription notification
func (m *SubscriptionNotification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEventType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateXApps(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SubscriptionNotification) validateEventType(formats strfmt.Registry) error {

	if swag.IsZero(m.EventType) { // not required
		return nil
	}

	if err := m.EventType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("eventType")
		}
		return err
	}

	return nil
}

func (m *SubscriptionNotification) validateXApps(formats strfmt.Registry) error {

	if swag.IsZero(m.XApps) { // not required
		return nil
	}

	if err := m.XApps.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("xApps")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SubscriptionNotification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubscriptionNotification) UnmarshalBinary(b []byte) error {
	var res SubscriptionNotification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
