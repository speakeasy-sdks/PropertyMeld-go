// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type OwnerDestroySecurity struct {
	PMOAuth2Authentication string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
}

type OwnerDestroyRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type OwnerDestroyResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
