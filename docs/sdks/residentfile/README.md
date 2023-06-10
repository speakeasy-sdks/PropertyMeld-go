# ResidentFile

### Available Operations

* [ResidentFileList](#residentfilelist)

## ResidentFileList

### Example Usage

```go
package main

import(
	"context"
	"log"
	"MeldAPI"
	"MeldAPI/pkg/models/operations"
)

func main() {
    s := meldapi.New()

    ctx := context.Background()
    res, err := s.ResidentFile.ResidentFileList(ctx, operations.ResidentFileListRequest{
        XPmOrg: 487676,
        Limit: meldapi.Int64(144691),
        Offset: meldapi.Int64(545),
        Ordering: meldapi.String("magni"),
    }, operations.ResidentFileListSecurity{
        PMOAuth2Authentication: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedTenantMeldFileList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ResidentFileListRequest](../../models/operations/residentfilelistrequest.md)   | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `security`                                                                                 | [operations.ResidentFileListSecurity](../../models/operations/residentfilelistsecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |


### Response

**[*operations.ResidentFileListResponse](../../models/operations/residentfilelistresponse.md), error**

