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
