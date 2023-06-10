# Expenditure

### Available Operations

* [ExpenditureList](#expenditurelist)
* [ExpenditureRetrieve](#expenditureretrieve)

## ExpenditureList

### Example Usage

```go
package main

import(
	"context"
	"log"
	"MeldAPI"
	"MeldAPI/pkg/models/operations"
	"MeldAPI/pkg/types"
)

func main() {
    s := meldapi.New()

    ctx := context.Background()
    res, err := s.Expenditure.ExpenditureList(ctx, operations.ExpenditureListRequest{
        XPmOrg: 449950,
        CreatedGte: types.MustTimeFromString("2022-05-22T05:33:50.280Z"),
        CreatedInterval: types.MustTimeFromString("2022-02-05T15:25:35.140Z"),
        CreatedLte: types.MustTimeFromString("2022-10-20T12:36:28.767Z"),
        Limit: meldapi.Int64(60225),
        Offset: meldapi.Int64(969810),
        Ordering: meldapi.String("est"),
        Status: []ExpenditureListStatus{
            operations.ExpenditureListStatusHold,
            operations.ExpenditureListStatusApproved,
            operations.ExpenditureListStatusBilled,
        },
    }, operations.ExpenditureListSecurity{
        PMOAuth2Authentication: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedMeldExpendituresListViewSerializerListList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ExpenditureListRequest](../../models/operations/expenditurelistrequest.md)   | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.ExpenditureListSecurity](../../models/operations/expenditurelistsecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


### Response

**[*operations.ExpenditureListResponse](../../models/operations/expenditurelistresponse.md), error**


## ExpenditureRetrieve

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
    res, err := s.Expenditure.ExpenditureRetrieve(ctx, operations.ExpenditureRetrieveRequest{
        XPmOrg: 358152,
        ID: "2c595590-7aff-41a3-a2fa-9467739251aa",
    }, operations.ExpenditureRetrieveSecurity{
        PMOAuth2Authentication: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.MeldExpendituresListViewSerializerDetail != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ExpenditureRetrieveRequest](../../models/operations/expenditureretrieverequest.md)   | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `security`                                                                                       | [operations.ExpenditureRetrieveSecurity](../../models/operations/expenditureretrievesecurity.md) | :heavy_check_mark:                                                                               | The security requirements to use for the request.                                                |


### Response

**[*operations.ExpenditureRetrieveResponse](../../models/operations/expenditureretrieveresponse.md), error**

