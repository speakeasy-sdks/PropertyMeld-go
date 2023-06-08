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
