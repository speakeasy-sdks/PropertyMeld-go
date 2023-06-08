// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type PriorityEnum string

const (
	PriorityEnumLow    PriorityEnum = "LOW"
	PriorityEnumMedium PriorityEnum = "MEDIUM"
	PriorityEnumHigh   PriorityEnum = "HIGH"
)

func (e PriorityEnum) ToPointer() *PriorityEnum {
	return &e
}

func (e *PriorityEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "LOW":
		fallthrough
	case "MEDIUM":
		fallthrough
	case "HIGH":
		*e = PriorityEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PriorityEnum: %v", v)
	}
}
