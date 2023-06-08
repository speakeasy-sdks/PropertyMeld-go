// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"MeldAPI/pkg/models/shared"
	"net/http"
)

type OwnerCreateSecurity struct {
	PMOAuth2Authentication string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
}

type OwnerCreateResponse struct {
	ContentType string
	Owner       *shared.Owner
	StatusCode  int
	RawResponse *http.Response
}
