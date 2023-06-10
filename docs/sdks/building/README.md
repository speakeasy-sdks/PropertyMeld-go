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

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.BuildingListRequest](../../models/operations/buildinglistrequest.md)   | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `security`                                                                         | [operations.BuildingListSecurity](../../models/operations/buildinglistsecurity.md) | :heavy_check_mark:                                                                 | The security requirements to use for the request.                                  |


### Response

**[*operations.BuildingListResponse](../../models/operations/buildinglistresponse.md), error**


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

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.BuildingRetrieveRequest](../../models/operations/buildingretrieverequest.md)   | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `security`                                                                                 | [operations.BuildingRetrieveSecurity](../../models/operations/buildingretrievesecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |


### Response

**[*operations.BuildingRetrieveResponse](../../models/operations/buildingretrieveresponse.md), error**

