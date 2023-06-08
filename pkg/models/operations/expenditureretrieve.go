// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"MeldAPI/pkg/models/shared"
	"net/http"
)

type ExpenditureRetrieveSecurity struct {
	PMOAuth2Authentication string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
}

type ExpenditureRetrieveRequest struct {
	// The management ID (MID), found in the first number of your URL when logged in:  https://app.propertymeld.com/{MID}/m/123
	XPmOrg int64  `header:"style=simple,explode=false,name=X-Pm-Org"`
	ID     string `pathParam:"style=simple,explode=false,name=id"`
}

type ExpenditureRetrieveResponse struct {
	ContentType                              string
	MeldExpendituresListViewSerializerDetail *shared.MeldExpendituresListViewSerializerDetail
	StatusCode                               int
	RawResponse                              *http.Response
}
