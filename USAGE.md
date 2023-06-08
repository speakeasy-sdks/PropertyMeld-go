<!-- Start SDK Example Usage -->
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
    res, err := s.Building.BuildingList(ctx, operations.BuildingListRequest{
        XPmOrg: 548814,
        Limit: meldapi.Int64(592845),
        Offset: meldapi.Int64(715190),
        Ordering: meldapi.String("quibusdam"),
    }, operations.BuildingListSecurity{
        PMOAuth2Authentication: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedBuildingSerializerListList != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->