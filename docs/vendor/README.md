# Vendor

### Available Operations

* [VendorDestroy](#vendordestroy)
* [VendorList](#vendorlist)
* [VendorRetrieve](#vendorretrieve)

## VendorDestroy

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
    res, err := s.Vendor.VendorDestroy(ctx, operations.VendorDestroyRequest{
        ID: "e189dbb3-0fcb-433e-a055-b197cd44e2f5",
    }, operations.VendorDestroySecurity{
        PMOAuth2Authentication: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## VendorList

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
    res, err := s.Vendor.VendorList(ctx, operations.VendorListRequest{
        XPmOrg: 164532,
        Limit: meldapi.Int64(813880),
        Offset: meldapi.Int64(512905),
        Ordering: meldapi.String("odit"),
    }, operations.VendorListSecurity{
        PMOAuth2Authentication: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedVendorList != nil {
        // handle response
    }
}
```

## VendorRetrieve

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
    res, err := s.Vendor.VendorRetrieve(ctx, operations.VendorRetrieveRequest{
        ID: "d3513bb6-f48b-4656-bcdb-35ff2e4b2753",
    }, operations.VendorRetrieveSecurity{
        PMOAuth2Authentication: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Vendor != nil {
        // handle response
    }
}
```
