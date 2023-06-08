# PropertyGroup

### Available Operations

* [PropertyGroupList](#propertygrouplist)
* [PropertyGroupRetrieve](#propertygroupretrieve)

## PropertyGroupList

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
    res, err := s.PropertyGroup.PropertyGroupList(ctx, operations.PropertyGroupListRequest{
        XPmOrg: 888044,
        Limit: meldapi.Int64(505866),
        Offset: meldapi.Int64(708609),
        Ordering: meldapi.String("quaerat"),
    }, operations.PropertyGroupListSecurity{
        PMOAuth2Authentication: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedPropertyGroupSerializerListList != nil {
        // handle response
    }
}
```

## PropertyGroupRetrieve

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
    res, err := s.PropertyGroup.PropertyGroupRetrieve(ctx, operations.PropertyGroupRetrieveRequest{
        XPmOrg: 277773,
        ID: "5e80ca55-efd2-40e4-97e1-858b6a89fbe3",
    }, operations.PropertyGroupRetrieveSecurity{
        PMOAuth2Authentication: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PropertyGroupSerializerDetail != nil {
        // handle response
    }
}
```
