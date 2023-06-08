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
