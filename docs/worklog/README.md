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
