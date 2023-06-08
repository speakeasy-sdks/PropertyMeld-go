# Estimates

### Available Operations

* [EstimatesList](#estimateslist)
* [EstimatesRetrieve](#estimatesretrieve)

## EstimatesList

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
    res, err := s.Estimates.EstimatesList(ctx, operations.EstimatesListRequest{
        XPmOrg: 978619,
        EstimateStatus: []EstimatesListEstimateStatus{
            operations.EstimatesListEstimateStatusEstimateSubmitted,
            operations.EstimatesListEstimateStatusEstimateSubmitted,
        },
        Limit: meldapi.Int64(461479),
        Offset: meldapi.Int64(520478),
        Ordering: meldapi.String("porro"),
    }, operations.EstimatesListSecurity{
        PMOAuth2Authentication: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedEstimateSerializerListList != nil {
        // handle response
    }
}
```

## EstimatesRetrieve

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
    res, err := s.Estimates.EstimatesRetrieve(ctx, operations.EstimatesRetrieveRequest{
        XPmOrg: 678880,
        ID: "1ba928fc-8167-442c-b739-205929396fea",
    }, operations.EstimatesRetrieveSecurity{
        PMOAuth2Authentication: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.EstimateSerializerDetail != nil {
        // handle response
    }
}
```
