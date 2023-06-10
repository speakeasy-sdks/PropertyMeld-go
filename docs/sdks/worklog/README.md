# WorkLog

### Available Operations

* [WorkLogList](#workloglist)
* [WorkLogRetrieve](#worklogretrieve)

## WorkLogList

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
    res, err := s.WorkLog.WorkLogList(ctx, operations.WorkLogListRequest{
        XPmOrg: 491201,
        Limit: meldapi.Int64(729828),
        Offset: meldapi.Int64(72350),
        Ordering: meldapi.String("ab"),
    }, operations.WorkLogListSecurity{
        PMOAuth2Authentication: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedWorkEntrySerializerListList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.WorkLogListRequest](../../models/operations/workloglistrequest.md)   | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `security`                                                                       | [operations.WorkLogListSecurity](../../models/operations/workloglistsecurity.md) | :heavy_check_mark:                                                               | The security requirements to use for the request.                                |


### Response

**[*operations.WorkLogListResponse](../../models/operations/workloglistresponse.md), error**


## WorkLogRetrieve

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
    res, err := s.WorkLog.WorkLogRetrieve(ctx, operations.WorkLogRetrieveRequest{
        XPmOrg: 280114,
        ID: "eeb52ff7-85fc-4378-94d4-c98e0c2bb89e",
    }, operations.WorkLogRetrieveSecurity{
        PMOAuth2Authentication: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.WorkEntrySerializerDetail != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.WorkLogRetrieveRequest](../../models/operations/worklogretrieverequest.md)   | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.WorkLogRetrieveSecurity](../../models/operations/worklogretrievesecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


### Response

**[*operations.WorkLogRetrieveResponse](../../models/operations/worklogretrieveresponse.md), error**

