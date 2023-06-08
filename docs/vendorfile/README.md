# VendorFile

### Available Operations

* [VendorFileList](#vendorfilelist)

## VendorFileList

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
    res, err := s.VendorFile.VendorFileList(ctx, operations.VendorFileListRequest{
        XPmOrg: 480061,
        Limit: meldapi.Int64(664965),
        Offset: meldapi.Int64(522176),
        Ordering: meldapi.String("eligendi"),
    }, operations.VendorFileListSecurity{
        PMOAuth2Authentication: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedVendorMeldFileList != nil {
        // handle response
    }
}
```
