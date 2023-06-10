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

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.PropertyGroupListRequest](../../models/operations/propertygrouplistrequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.PropertyGroupListSecurity](../../models/operations/propertygrouplistsecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


### Response

**[*operations.PropertyGroupListResponse](../../models/operations/propertygrouplistresponse.md), error**


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

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.PropertyGroupRetrieveRequest](../../models/operations/propertygroupretrieverequest.md)   | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `security`                                                                                           | [operations.PropertyGroupRetrieveSecurity](../../models/operations/propertygroupretrievesecurity.md) | :heavy_check_mark:                                                                                   | The security requirements to use for the request.                                                    |


### Response

**[*operations.PropertyGroupRetrieveResponse](../../models/operations/propertygroupretrieveresponse.md), error**

