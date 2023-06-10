# Unit

### Available Operations

* [UnitCreate](#unitcreate)
* [UnitDestroy](#unitdestroy)
* [UnitList](#unitlist)
* [UnitPartialUpdate](#unitpartialupdate)
* [UnitRetrieve](#unitretrieve)
* [UnitUpdate](#unitupdate)

## UnitCreate

### Example Usage

```go
package main

import(
	"context"
	"log"
	"MeldAPI"
	"MeldAPI/pkg/models/shared"
	"MeldAPI/pkg/models/operations"
)

func main() {
    s := meldapi.New()

    ctx := context.Background()
    res, err := s.Unit.UnitCreate(ctx, shared.UnitInput{
        ApprovalCurrencyLimit: meldapi.String("quae"),
        CurrentResidents: []int64{
            208253,
            348357,
        },
        MaintenanceNotes: meldapi.String("itaque"),
        PropertyID: 88248,
        Unit: meldapi.String("ipsum"),
        UnitAddress: &shared.Address{
            City: "Stiedemanncester",
            Country: meldapi.String("Sierra Leone"),
            CountyProvince: "quia",
            Line1: "quia",
            Line2: meldapi.String("nostrum"),
            Line3: meldapi.String("omnis"),
            Postcode: "16786-5804",
        },
    }, operations.UnitCreateSecurity{
        PMOAuth2Authentication: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Unit != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [shared.UnitInput](../../models/shared/unitinput.md)                           | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `security`                                                                     | [operations.UnitCreateSecurity](../../models/operations/unitcreatesecurity.md) | :heavy_check_mark:                                                             | The security requirements to use for the request.                              |


### Response

**[*operations.UnitCreateResponse](../../models/operations/unitcreateresponse.md), error**


## UnitDestroy

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
    res, err := s.Unit.UnitDestroy(ctx, operations.UnitDestroyRequest{
        ID: "0e1084cb-0672-4d1a-9879-eeb9665b85ef",
    }, operations.UnitDestroySecurity{
        PMOAuth2Authentication: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.UnitDestroyRequest](../../models/operations/unitdestroyrequest.md)   | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `security`                                                                       | [operations.UnitDestroySecurity](../../models/operations/unitdestroysecurity.md) | :heavy_check_mark:                                                               | The security requirements to use for the request.                                |


### Response

**[*operations.UnitDestroyResponse](../../models/operations/unitdestroyresponse.md), error**


## UnitList

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
    res, err := s.Unit.UnitList(ctx, operations.UnitListRequest{
        Limit: meldapi.Int64(737279),
        Offset: meldapi.Int64(872303),
        Ordering: meldapi.String("alias"),
    }, operations.UnitListSecurity{
        PMOAuth2Authentication: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedUnitList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.UnitListRequest](../../models/operations/unitlistrequest.md)   | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `security`                                                                 | [operations.UnitListSecurity](../../models/operations/unitlistsecurity.md) | :heavy_check_mark:                                                         | The security requirements to use for the request.                          |


### Response

**[*operations.UnitListResponse](../../models/operations/unitlistresponse.md), error**


## UnitPartialUpdate

### Example Usage

```go
package main

import(
	"context"
	"log"
	"MeldAPI"
	"MeldAPI/pkg/models/operations"
	"MeldAPI/pkg/models/shared"
)

func main() {
    s := meldapi.New()

    ctx := context.Background()
    res, err := s.Unit.UnitPartialUpdate(ctx, operations.UnitPartialUpdateRequest{
        PatchedUnitInput: &shared.PatchedUnitInput{
            ApprovalCurrencyLimit: meldapi.String("quia"),
            CurrentResidents: []int64{
                684126,
                919508,
                34070,
            },
            MaintenanceNotes: meldapi.String("expedita"),
            PropertyID: meldapi.Int64(885208),
            Unit: meldapi.String("eos"),
            UnitAddress: &shared.Address{
                City: "Lawton",
                Country: meldapi.String("Lithuania"),
                CountyProvince: "odit",
                Line1: "explicabo",
                Line2: meldapi.String("corporis"),
                Line3: meldapi.String("error"),
                Postcode: "29626-3164",
            },
        },
        ID: "f92443da-7ce5-42b8-95c5-37c6454efb0b",
    }, operations.UnitPartialUpdateSecurity{
        PMOAuth2Authentication: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Unit != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.UnitPartialUpdateRequest](../../models/operations/unitpartialupdaterequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.UnitPartialUpdateSecurity](../../models/operations/unitpartialupdatesecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


### Response

**[*operations.UnitPartialUpdateResponse](../../models/operations/unitpartialupdateresponse.md), error**


## UnitRetrieve

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
    res, err := s.Unit.UnitRetrieve(ctx, operations.UnitRetrieveRequest{
        ID: "34896c3c-a5ac-4fbe-afd5-707577929177",
    }, operations.UnitRetrieveSecurity{
        PMOAuth2Authentication: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Unit != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.UnitRetrieveRequest](../../models/operations/unitretrieverequest.md)   | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `security`                                                                         | [operations.UnitRetrieveSecurity](../../models/operations/unitretrievesecurity.md) | :heavy_check_mark:                                                                 | The security requirements to use for the request.                                  |


### Response

**[*operations.UnitRetrieveResponse](../../models/operations/unitretrieveresponse.md), error**


## UnitUpdate

### Example Usage

```go
package main

import(
	"context"
	"log"
	"MeldAPI"
	"MeldAPI/pkg/models/operations"
	"MeldAPI/pkg/models/shared"
)

func main() {
    s := meldapi.New()

    ctx := context.Background()
    res, err := s.Unit.UnitUpdate(ctx, operations.UnitUpdateRequest{
        UnitInput: shared.UnitInput{
            ApprovalCurrencyLimit: meldapi.String("pariatur"),
            CurrentResidents: []int64{
                627735,
                763165,
                401428,
                311486,
            },
            MaintenanceNotes: meldapi.String("commodi"),
            PropertyID: 888616,
            Unit: meldapi.String("placeat"),
            UnitAddress: &shared.Address{
                City: "Heidenreichport",
                Country: meldapi.String("Cote d'Ivoire"),
                CountyProvince: "modi",
                Line1: "ipsa",
                Line2: meldapi.String("sint"),
                Line3: meldapi.String("vero"),
                Postcode: "97193",
            },
        },
        ID: "a2b12eb0-7f11-46db-9954-5fc95fa88970",
    }, operations.UnitUpdateSecurity{
        PMOAuth2Authentication: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Unit != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.UnitUpdateRequest](../../models/operations/unitupdaterequest.md)   | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `security`                                                                     | [operations.UnitUpdateSecurity](../../models/operations/unitupdatesecurity.md) | :heavy_check_mark:                                                             | The security requirements to use for the request.                              |


### Response

**[*operations.UnitUpdateResponse](../../models/operations/unitupdateresponse.md), error**

