# Building

### Available Operations

* [BuildingList](#buildinglist)
* [BuildingRetrieve](#buildingretrieve)

## BuildingList

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
    res, err := s.Building.BuildingList(ctx, operations.BuildingListRequest{
        XPmOrg: 602763,
        Limit: meldapi.Int64(857946),
        Offset: meldapi.Int64(544883),
        Ordering: meldapi.String("illum"),
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

## BuildingRetrieve

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
    res, err := s.Building.BuildingRetrieve(ctx, operations.BuildingRetrieveRequest{
        XPmOrg: 423655,
        ID: "9a674e0f-467c-4c87-96ed-151a05dfc2dd",
    }, operations.BuildingRetrieveSecurity{
        PMOAuth2Authentication: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BuildingSerializerDetail != nil {
        // handle response
    }
}
```
