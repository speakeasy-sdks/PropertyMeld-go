# ManagementAgent

### Available Operations

* [ManagementAgentList](#managementagentlist)
* [ManagementAgentRetrieve](#managementagentretrieve)

## ManagementAgentList

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
    res, err := s.ManagementAgent.ManagementAgentList(ctx, operations.ManagementAgentListRequest{
        XPmOrg: 588465,
        Limit: meldapi.Int64(725255),
        Offset: meldapi.Int64(659669),
        Ordering: meldapi.String("blanditiis"),
    }, operations.ManagementAgentListSecurity{
        PMOAuth2Authentication: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedManagementAgentSerializerListList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ManagementAgentListRequest](../../models/operations/managementagentlistrequest.md)   | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `security`                                                                                       | [operations.ManagementAgentListSecurity](../../models/operations/managementagentlistsecurity.md) | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |


### Response

**[*operations.ManagementAgentListResponse](../../models/operations/managementagentlistresponse.md), error**


## ManagementAgentRetrieve

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
    res, err := s.ManagementAgent.ManagementAgentRetrieve(ctx, operations.ManagementAgentRetrieveRequest{
        XPmOrg: 533206,
        ID: "f3a66997-074b-4a44-a9b6-e2141959890a",
    }, operations.ManagementAgentRetrieveSecurity{
        PMOAuth2Authentication: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ManagementAgentSerializerDetail != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.ManagementAgentRetrieveRequest](../../models/operations/managementagentretrieverequest.md)   | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `security`                                                                                               | [operations.ManagementAgentRetrieveSecurity](../../models/operations/managementagentretrievesecurity.md) | :heavy_check_mark:                                                                                       | The security requirements to use for the request.                                                        |


### Response

**[*operations.ManagementAgentRetrieveResponse](../../models/operations/managementagentretrieveresponse.md), error**

