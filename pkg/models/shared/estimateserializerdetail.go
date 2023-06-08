// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

type EstimateSerializerDetailOutput struct {
	ApprovalRequested    *bool               `json:"approval_requested,omitempty"`
	DateAccepted         time.Time           `json:"date_accepted"`
	DaysToComplete       *int64              `json:"days_to_complete,omitempty"`
	DaysUntilStart       *int64              `json:"days_until_start,omitempty"`
	Deadline             *time.Time          `json:"deadline,omitempty"`
	EstimateMeld         int64               `json:"estimate_meld"`
	EstimateStatus       *EstimateStatusEnum `json:"estimate_status,omitempty"`
	GetVendorEmail       string              `json:"get_vendor_email"`
	GetVendorInviteEmail string              `json:"get_vendor_invite_email"`
	GetVendorName        string              `json:"get_vendor_name"`
	ID                   int64               `json:"id"`
	LineItems            []int64             `json:"line_items"`
	Notes                *string             `json:"notes,omitempty"`
	OriginalMeld         int64               `json:"original_meld"`
	PmFee                *string             `json:"pm_fee,omitempty"`
	Servicer             *int64              `json:"servicer,omitempty"`
	Total                string              `json:"total"`
	Vendor               int64               `json:"vendor"`
}
