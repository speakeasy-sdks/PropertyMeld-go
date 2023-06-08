// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

type MeldListViewProperty struct {
	City                *string                `json:"city,omitempty"`
	Country             *string                `json:"country,omitempty"`
	CountyProvince      *string                `json:"county_province,omitempty"`
	CreateBy            map[string]interface{} `json:"create_by,omitempty"`
	Created             time.Time              `json:"created"`
	ID                  int64                  `json:"id"`
	IsActive            *bool                  `json:"is_active,omitempty"`
	Latitude            *string                `json:"latitude,omitempty"`
	Line1               *string                `json:"line_1,omitempty"`
	Line2               *string                `json:"line_2,omitempty"`
	Line3               *string                `json:"line_3,omitempty"`
	Longitude           *string                `json:"longitude,omitempty"`
	MaintenanceNotes    *string                `json:"maintenance_notes,omitempty"`
	OtherAddressDetails *string                `json:"other_address_details,omitempty"`
	Postcode            string                 `json:"postcode"`
	PropertyName        *string                `json:"property_name,omitempty"`
	SupplementalData    map[string]interface{} `json:"supplemental_data,omitempty"`
	UpdateBy            map[string]interface{} `json:"update_by,omitempty"`
	Updated             time.Time              `json:"updated"`
}
