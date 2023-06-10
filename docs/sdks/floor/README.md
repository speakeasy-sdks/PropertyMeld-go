# Floor

### Available Operations

* [FloorList](#floorlist)
* [FloorRetrieve](#floorretrieve)

## FloorList

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
    res, err := s.Floor.FloorList(ctx, operations.FloorListRequest{
        XPmOrg: 317202,
        Limit: meldapi.Int64(138183),
        Offset: meldapi.Int64(778346),
        Ordering: meldapi.String("sequi"),
    }, operations.FloorListSecurity{
        PMOAuth2Authentication: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedFloorSerializerListList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.FloorListRequest](../../models/operations/floorlistrequest.md)   | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `security`                                                                   | [operations.FloorListSecurity](../../models/operations/floorlistsecurity.md) | :heavy_check_mark:                                                           | The security requirements to use for the request.                            |


### Response

**[*operations.FloorListResponse](../../models/operations/floorlistresponse.md), error**


## FloorRetrieve

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
    res, err := s.Floor.FloorRetrieve(ctx, operations.FloorRetrieveRequest{
        XPmOrg: 949572,
        ID: "5ad019da-1ffe-478f-897b-0074f15471b5",
    }, operations.FloorRetrieveSecurity{
        PMOAuth2Authentication: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.FloorSerializerDetail != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.FloorRetrieveRequest](../../models/operations/floorretrieverequest.md)   | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `security`                                                                           | [operations.FloorRetrieveSecurity](../../models/operations/floorretrievesecurity.md) | :heavy_check_mark:                                                                   | The security requirements to use for the request.                                    |


### Response

**[*operations.FloorRetrieveResponse](../../models/operations/floorretrieveresponse.md), error**

