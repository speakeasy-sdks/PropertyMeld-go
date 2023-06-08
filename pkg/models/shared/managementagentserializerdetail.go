// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

type ManagementAgentSerializerDetail struct {
	AgentPreferences        int64     `json:"agent_preferences"`
	Contact                 *int64    `json:"contact,omitempty"`
	Created                 time.Time `json:"created"`
	DefaultInvoiceFilter    *int64    `json:"default_invoice_filter,omitempty"`
	DefaultMeldFilter       *int64    `json:"default_meld_filter,omitempty"`
	FirstName               *string   `json:"first_name,omitempty"`
	ID                      int64     `json:"id"`
	IsActive                *bool     `json:"is_active,omitempty"`
	LastName                *string   `json:"last_name,omitempty"`
	Management              int64     `json:"management"`
	NewNotificationSettings *int64    `json:"new_notification_settings,omitempty"`
	PropertyGroups          []int64   `json:"property_groups"`
	Role                    int64     `json:"role"`
	Title                   string    `json:"title"`
	Updated                 time.Time `json:"updated"`
}
