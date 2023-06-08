// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type SchemaRetrieveSecurity struct {
	PMOAuth2Authentication *string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
}

// SchemaRetrieveFormat
type SchemaRetrieveFormat string

const (
	SchemaRetrieveFormatJSON SchemaRetrieveFormat = "json"
	SchemaRetrieveFormatYaml SchemaRetrieveFormat = "yaml"
)

func (e SchemaRetrieveFormat) ToPointer() *SchemaRetrieveFormat {
	return &e
}

func (e *SchemaRetrieveFormat) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "json":
		fallthrough
	case "yaml":
		*e = SchemaRetrieveFormat(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SchemaRetrieveFormat: %v", v)
	}
}

type SchemaRetrieveRequest struct {
	// The management ID (MID), found in the first number of your URL when logged in:  https://app.propertymeld.com/{MID}/m/123
	XPmOrg int64                 `header:"style=simple,explode=false,name=X-Pm-Org"`
	Format *SchemaRetrieveFormat `queryParam:"style=form,explode=true,name=format"`
}

type SchemaRetrieveResponse struct {
	Body                                                    []byte
	ContentType                                             string
	StatusCode                                              int
	RawResponse                                             *http.Response
	SchemaRetrieve200ApplicationJSONObject                  map[string]interface{}
	SchemaRetrieve200ApplicationVndOaiOpenapiPlusJSONObject map[string]interface{}
}
