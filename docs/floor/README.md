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
