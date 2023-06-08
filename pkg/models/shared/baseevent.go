// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"time"
)

type BaseEventClassificationType string

const (
	BaseEventClassificationTypeClassificationEnum BaseEventClassificationType = "ClassificationEnum"
	BaseEventClassificationTypeBlankEnum          BaseEventClassificationType = "BlankEnum"
)

type BaseEventClassification struct {
	ClassificationEnum *ClassificationEnum
	BlankEnum          *BlankEnum

	Type BaseEventClassificationType
}

func CreateBaseEventClassificationClassificationEnum(classificationEnum ClassificationEnum) BaseEventClassification {
	typ := BaseEventClassificationTypeClassificationEnum

	return BaseEventClassification{
		ClassificationEnum: &classificationEnum,
		Type:               typ,
	}
}

func CreateBaseEventClassificationBlankEnum(blankEnum BlankEnum) BaseEventClassification {
	typ := BaseEventClassificationTypeBlankEnum

	return BaseEventClassification{
		BlankEnum: &blankEnum,
		Type:      typ,
	}
}

func (u *BaseEventClassification) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	classificationEnum := new(ClassificationEnum)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&classificationEnum); err == nil {
		u.ClassificationEnum = classificationEnum
		u.Type = BaseEventClassificationTypeClassificationEnum
		return nil
	}

	blankEnum := new(BlankEnum)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&blankEnum); err == nil {
		u.BlankEnum = blankEnum
		u.Type = BaseEventClassificationTypeBlankEnum
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u BaseEventClassification) MarshalJSON() ([]byte, error) {
	if u.ClassificationEnum != nil {
		return json.Marshal(u.ClassificationEnum)
	}

	if u.BlankEnum != nil {
		return json.Marshal(u.BlankEnum)
	}

	return nil, nil
}

type BaseEvent struct {
	Classification *BaseEventClassification `json:"classification,omitempty"`
	CreateBy       map[string]interface{}   `json:"create_by,omitempty"`
	Created        time.Time                `json:"created"`
	Description    *string                  `json:"description,omitempty"`
	Dtend          *time.Time               `json:"dtend,omitempty"`
	Dtstart        time.Time                `json:"dtstart"`
	DurDay         *int64                   `json:"dur_day,omitempty"`
	DurWeek        *int64                   `json:"dur_week,omitempty"`
	Duration       *int64                   `json:"duration,omitempty"`
	ID             int64                    `json:"id"`
	Latitude       *string                  `json:"latitude,omitempty"`
	Longitude      *string                  `json:"longitude,omitempty"`
	Organizer      *string                  `json:"organizer,omitempty"`
	Rrule          *int64                   `json:"rrule,omitempty"`
	UpdateBy       map[string]interface{}   `json:"update_by,omitempty"`
	Updated        time.Time                `json:"updated"`
}
