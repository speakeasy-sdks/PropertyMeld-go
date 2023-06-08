// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"MeldAPI/pkg/models/shared"
	"net/http"
)

type ManagerFileListSecurity struct {
	PMOAuth2Authentication string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
}

type ManagerFileListRequest struct {
	// The management ID (MID), found in the first number of your URL when logged in:  https://app.propertymeld.com/{MID}/m/123
	XPmOrg int64 `header:"style=simple,explode=false,name=X-Pm-Org"`
	// Number of results to return per page.
	Limit *int64 `queryParam:"style=form,explode=true,name=limit"`
	// The initial index from which to return the results.
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
	// Which field to use when ordering the results.
	Ordering *string `queryParam:"style=form,explode=true,name=ordering"`
}

type ManagerFileListResponse struct {
	ContentType           string
	PaginatedMeldFileList *shared.PaginatedMeldFileList
	StatusCode            int
	RawResponse           *http.Response
}
