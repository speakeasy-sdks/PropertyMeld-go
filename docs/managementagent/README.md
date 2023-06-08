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
