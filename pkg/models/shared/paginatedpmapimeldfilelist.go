// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type PaginatedPMAPIMeldfileList struct {
	Count    *int64          `json:"count,omitempty"`
	Next     *string         `json:"next,omitempty"`
	Previous *string         `json:"previous,omitempty"`
	Results  []PMAPIMeldfile `json:"results,omitempty"`
}
