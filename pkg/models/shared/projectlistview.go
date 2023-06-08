// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

type ProjectListView struct {
	Coordinators []ProjectListViewSerializerManagementAgent `json:"coordinators"`
	Created      time.Time                                  `json:"created"`
	Description  *string                                    `json:"description,omitempty"`
	DueDate      time.Time                                  `json:"due_date"`
	ID           int64                                      `json:"id"`
	Melds        []int64                                    `json:"melds"`
	Name         string                                     `json:"name"`
	ProjectType  *ProjectTypeEnum                           `json:"project_type,omitempty"`
	StartDate    time.Time                                  `json:"start_date"`
	Unit         *ProjectListViewSerializerUnit             `json:"unit,omitempty"`
	Updated      time.Time                                  `json:"updated"`
}
