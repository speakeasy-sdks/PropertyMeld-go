# ManagerFile

### Available Operations

* [ManagerFileList](#managerfilelist)

## ManagerFileList

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
    res, err := s.ManagerFile.ManagerFileList(ctx, operations.ManagerFileListRequest{
        XPmOrg: 968962,
        Limit: meldapi.Int64(652103),
        Offset: meldapi.Int64(320997),
        Ordering: meldapi.String("eum"),
    }, operations.ManagerFileListSecurity{
        PMOAuth2Authentication: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedMeldFileList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ManagerFileListRequest](../../models/operations/managerfilelistrequest.md)   | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.ManagerFileListSecurity](../../models/operations/managerfilelistsecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


### Response

**[*operations.ManagerFileListResponse](../../models/operations/managerfilelistresponse.md), error**

