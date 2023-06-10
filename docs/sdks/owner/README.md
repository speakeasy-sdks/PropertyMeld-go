# Owner

### Available Operations

* [OwnerCreate](#ownercreate)
* [OwnerDestroy](#ownerdestroy)
* [OwnerList](#ownerlist)
* [OwnerPartialUpdate](#ownerpartialupdate)
* [OwnerRetrieve](#ownerretrieve)
* [OwnerUpdate](#ownerupdate)

## OwnerCreate

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
    res, err := s.Owner.OwnerCreate(ctx, shared.OwnerInput{
        Address: &shared.Address{
            City: "State College",
            Country: meldapi.String("Saint Helena"),
            CountyProvince: "dicta",
            Line1: "laborum",
            Line2: meldapi.String("totam"),
            Line3: meldapi.String("incidunt"),
            Postcode: "17734",
        },
        Contact: &shared.Contact{
            BusinessPhone: meldapi.String("molestias"),
            BusinessPhoneExt: meldapi.String("temporibus"),
            CellPhone: meldapi.String("qui"),
            CellPhoneExt: meldapi.String("neque"),
            Fax: meldapi.String("fugit"),
            HomePhone: meldapi.String("magni"),
            HomePhoneExt: meldapi.String("odio"),
            PrimaryEmail: meldapi.String("Fiona.Reichert76@gmail.com"),
            SecondaryEmail: meldapi.String("Nella.Bosco8@hotmail.com"),
            TertiaryEmail: meldapi.String("Katrine.Reynolds@hotmail.com"),
        },
        Email: "Corene24@hotmail.com",
        FirstName: "Marianne",
        HubAccessLevel: shared.HubAccessLevelEnumDirectOnly.ToPointer(),
        InvitedToHub: meldapi.Bool(false),
        LastName: "Berge",
        Properties: []int64{
            555649,
        },
    }, operations.OwnerCreateSecurity{
        PMOAuth2Authentication: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Owner != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [shared.OwnerInput](../../models/shared/ownerinput.md)                           | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `security`                                                                       | [operations.OwnerCreateSecurity](../../models/operations/ownercreatesecurity.md) | :heavy_check_mark:                                                               | The security requirements to use for the request.                                |


### Response

**[*operations.OwnerCreateResponse](../../models/operations/ownercreateresponse.md), error**


## OwnerDestroy

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
    res, err := s.Owner.OwnerDestroy(ctx, operations.OwnerDestroyRequest{
        ID: "e0adcf4b-9218-479f-8e95-3f73ef7fbc7a",
    }, operations.OwnerDestroySecurity{
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.OwnerDestroyRequest](../../models/operations/ownerdestroyrequest.md)   | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `security`                                                                         | [operations.OwnerDestroySecurity](../../models/operations/ownerdestroysecurity.md) | :heavy_check_mark:                                                                 | The security requirements to use for the request.                                  |


### Response

**[*operations.OwnerDestroyResponse](../../models/operations/ownerdestroyresponse.md), error**


## OwnerList

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
    res, err := s.Owner.OwnerList(ctx, operations.OwnerListRequest{
        Limit: meldapi.Int64(708548),
        Offset: meldapi.Int64(874288),
        Ordering: meldapi.String("ducimus"),
    }, operations.OwnerListSecurity{
        PMOAuth2Authentication: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedOwnerList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.OwnerListRequest](../../models/operations/ownerlistrequest.md)   | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `security`                                                                   | [operations.OwnerListSecurity](../../models/operations/ownerlistsecurity.md) | :heavy_check_mark:                                                           | The security requirements to use for the request.                            |


### Response

**[*operations.OwnerListResponse](../../models/operations/ownerlistresponse.md), error**


## OwnerPartialUpdate

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
    res, err := s.Owner.OwnerPartialUpdate(ctx, operations.OwnerPartialUpdateRequest{
        PatchedOwnerInput: &shared.PatchedOwnerInput{
            Address: &shared.Address{
                City: "Port Rosina",
                Country: meldapi.String("Comoros"),
                CountyProvince: "natus",
                Line1: "impedit",
                Line2: meldapi.String("aut"),
                Line3: meldapi.String("voluptatibus"),
                Postcode: "81799",
            },
            Contact: &shared.Contact{
                BusinessPhone: meldapi.String("iusto"),
                BusinessPhoneExt: meldapi.String("eligendi"),
                CellPhone: meldapi.String("ducimus"),
                CellPhoneExt: meldapi.String("alias"),
                Fax: meldapi.String("officia"),
                HomePhone: meldapi.String("tempora"),
                HomePhoneExt: meldapi.String("ipsam"),
                PrimaryEmail: meldapi.String("Brielle.Keebler18@yahoo.com"),
                SecondaryEmail: meldapi.String("Johnpaul98@yahoo.com"),
                TertiaryEmail: meldapi.String("Gustave_Stoltenberg@gmail.com"),
            },
            Email: meldapi.String("Victor.Schamberger77@yahoo.com"),
            FirstName: meldapi.String("Flossie"),
            HubAccessLevel: shared.HubAccessLevelEnumAllNotifications.ToPointer(),
            InvitedToHub: meldapi.Bool(false),
            LastName: meldapi.String("Jacobi"),
            Properties: []int64{
                301831,
            },
        },
        ID: "6c3e250f-b008-4c42-a141-aac366c8dd6b",
    }, operations.OwnerPartialUpdateSecurity{
        PMOAuth2Authentication: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Owner != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.OwnerPartialUpdateRequest](../../models/operations/ownerpartialupdaterequest.md)   | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `security`                                                                                     | [operations.OwnerPartialUpdateSecurity](../../models/operations/ownerpartialupdatesecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |


### Response

**[*operations.OwnerPartialUpdateResponse](../../models/operations/ownerpartialupdateresponse.md), error**


## OwnerRetrieve

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
    res, err := s.Owner.OwnerRetrieve(ctx, operations.OwnerRetrieveRequest{
        ID: "14429074-7477-48a7-bd46-6d28c10ab3cd",
    }, operations.OwnerRetrieveSecurity{
        PMOAuth2Authentication: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Owner != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.OwnerRetrieveRequest](../../models/operations/ownerretrieverequest.md)   | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `security`                                                                           | [operations.OwnerRetrieveSecurity](../../models/operations/ownerretrievesecurity.md) | :heavy_check_mark:                                                                   | The security requirements to use for the request.                                    |


### Response

**[*operations.OwnerRetrieveResponse](../../models/operations/ownerretrieveresponse.md), error**


## OwnerUpdate

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
    res, err := s.Owner.OwnerUpdate(ctx, operations.OwnerUpdateRequest{
        OwnerInput: shared.OwnerInput{
            Address: &shared.Address{
                City: "Parkerburgh",
                Country: meldapi.String("China"),
                CountyProvince: "voluptas",
                Line1: "ab",
                Line2: meldapi.String("cupiditate"),
                Line3: meldapi.String("consequatur"),
                Postcode: "83117",
            },
            Contact: &shared.Contact{
                BusinessPhone: meldapi.String("esse"),
                BusinessPhoneExt: meldapi.String("recusandae"),
                CellPhone: meldapi.String("aperiam"),
                CellPhoneExt: meldapi.String("distinctio"),
                Fax: meldapi.String("quod"),
                HomePhone: meldapi.String("dignissimos"),
                HomePhoneExt: meldapi.String("inventore"),
                PrimaryEmail: meldapi.String("Josiah48@yahoo.com"),
                SecondaryEmail: meldapi.String("Harvey64@yahoo.com"),
                TertiaryEmail: meldapi.String("Alfonzo_Sawayn@yahoo.com"),
            },
            Email: "Carol68@yahoo.com",
            FirstName: "Lysanne",
            HubAccessLevel: shared.HubAccessLevelEnumAllNotifications.ToPointer(),
            InvitedToHub: meldapi.Bool(false),
            LastName: "Lindgren",
            Properties: []int64{
                325685,
            },
        },
        ID: "62f222e9-817e-4e17-8be6-1e6b7b95bc0a",
    }, operations.OwnerUpdateSecurity{
        PMOAuth2Authentication: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Owner != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.OwnerUpdateRequest](../../models/operations/ownerupdaterequest.md)   | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `security`                                                                       | [operations.OwnerUpdateSecurity](../../models/operations/ownerupdatesecurity.md) | :heavy_check_mark:                                                               | The security requirements to use for the request.                                |


### Response

**[*operations.OwnerUpdateResponse](../../models/operations/ownerupdateresponse.md), error**

