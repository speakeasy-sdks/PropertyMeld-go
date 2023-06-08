// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

type PMAPIMeldfile struct {
	Cloned           *bool     `json:"cloned,omitempty"`
	Created          time.Time `json:"created"`
	Filename         string    `json:"filename"`
	HiddenFromOwner  *bool     `json:"hidden_from_owner,omitempty"`
	HiddenFromTenant *bool     `json:"hidden_from_tenant,omitempty"`
	HiddenFromVendor *bool     `json:"hidden_from_vendor,omitempty"`
	ID               int64     `json:"id"`
	Meld             int64     `json:"meld"`
	Updated          time.Time `json:"updated"`
	Uploader         *int64    `json:"uploader,omitempty"`
	URL              string    `json:"url"`
}
