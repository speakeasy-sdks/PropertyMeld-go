// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type StatusD0bEnum string

const (
	StatusD0bEnumDraft    StatusD0bEnum = "DRAFT"
	StatusD0bEnumInReview StatusD0bEnum = "IN_REVIEW"
	StatusD0bEnumHold     StatusD0bEnum = "HOLD"
	StatusD0bEnumApproved StatusD0bEnum = "APPROVED"
	StatusD0bEnumBilled   StatusD0bEnum = "BILLED"
)

func (e StatusD0bEnum) ToPointer() *StatusD0bEnum {
	return &e
}

func (e *StatusD0bEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "DRAFT":
		fallthrough
	case "IN_REVIEW":
		fallthrough
	case "HOLD":
		fallthrough
	case "APPROVED":
		fallthrough
	case "BILLED":
		*e = StatusD0bEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for StatusD0bEnum: %v", v)
	}
}
