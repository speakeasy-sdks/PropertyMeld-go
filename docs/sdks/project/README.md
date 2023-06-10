# Project

### Available Operations

* [ProjectList](#projectlist)
* [ProjectRetrieve](#projectretrieve)

## ProjectList

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
    res, err := s.Project.ProjectList(ctx, operations.ProjectListRequest{
        XPmOrg: 731398,
        Limit: meldapi.Int64(240020),
        Offset: meldapi.Int64(766964),
        Ordering: meldapi.String("consequuntur"),
    }, operations.ProjectListSecurity{
        PMOAuth2Authentication: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedProjectListViewList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ProjectListRequest](../../models/operations/projectlistrequest.md)   | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `security`                                                                       | [operations.ProjectListSecurity](../../models/operations/projectlistsecurity.md) | :heavy_check_mark:                                                               | The security requirements to use for the request.                                |


### Response

**[*operations.ProjectListResponse](../../models/operations/projectlistresponse.md), error**


## ProjectRetrieve

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
    res, err := s.Project.ProjectRetrieve(ctx, operations.ProjectRetrieveRequest{
        XPmOrg: 9766,
        ID: "c4f3789f-d871-4f99-9d2e-fd121aa6f1e6",
    }, operations.ProjectRetrieveSecurity{
        PMOAuth2Authentication: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PmAPIProjectDetail != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ProjectRetrieveRequest](../../models/operations/projectretrieverequest.md)   | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.ProjectRetrieveSecurity](../../models/operations/projectretrievesecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


### Response

**[*operations.ProjectRetrieveResponse](../../models/operations/projectretrieveresponse.md), error**

