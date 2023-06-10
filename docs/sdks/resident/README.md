# Resident

### Available Operations

* [ResidentCreateForm](#residentcreateform)
* [ResidentCreateJSON](#residentcreatejson)
* [ResidentCreateMultipart](#residentcreatemultipart)
* [ResidentDestroy](#residentdestroy)
* [ResidentList](#residentlist)
* [ResidentPartialUpdateForm](#residentpartialupdateform)
* [ResidentPartialUpdateJSON](#residentpartialupdatejson)
* [ResidentPartialUpdateMultipart](#residentpartialupdatemultipart)
* [ResidentRetrieve](#residentretrieve)
* [ResidentUpdateForm](#residentupdateform)
* [ResidentUpdateJSON](#residentupdatejson)
* [ResidentUpdateMultipart](#residentupdatemultipart)

## ResidentCreateForm

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
    res, err := s.Resident.ResidentCreateForm(ctx, shared.ResidentInput1{
        Address: meldapi.String("dolorum"),
        Contact: &shared.Contact{
            BusinessPhone: meldapi.String("nostrum"),
            BusinessPhoneExt: meldapi.String("officia"),
            CellPhone: meldapi.String("dolorum"),
            CellPhoneExt: meldapi.String("corrupti"),
            Fax: meldapi.String("accusamus"),
            HomePhone: meldapi.String("tempora"),
            HomePhoneExt: meldapi.String("atque"),
            PrimaryEmail: meldapi.String("Easter63@gmail.com"),
            SecondaryEmail: meldapi.String("Elinor.Adams@hotmail.com"),
            TertiaryEmail: meldapi.String("Justine.Lynch8@gmail.com"),
        },
        FirstName: "Jo",
        Invite: meldapi.Bool(false),
        LastName: "Jaskolski",
        MiddleName: meldapi.String("sed"),
        Notes: meldapi.String("sit"),
        NotificationSettings: &shared.NotificationSettingsInput{
            IncomingMeldFrequency: shared.IncomingMeldFrequencyEnumDaily.ToPointer(),
            SmsNotifications: meldapi.Bool(false),
            SuccessfulMeldFrequency: shared.SuccessfulMeldFrequencyEnumDaily.ToPointer(),
            Timezone: shared.TimezoneEnumPacificGambier.ToPointer(),
        },
        Units: []int64{
            8511,
            279068,
            968865,
        },
    }, operations.ResidentCreateFormSecurity{
        PMOAuth2Authentication: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Resident != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [shared.ResidentInput1](../../models/shared/residentinput1.md)                                 | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `security`                                                                                     | [operations.ResidentCreateFormSecurity](../../models/operations/residentcreateformsecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |


### Response

**[*operations.ResidentCreateFormResponse](../../models/operations/residentcreateformresponse.md), error**


## ResidentCreateJSON

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
    res, err := s.Resident.ResidentCreateJSON(ctx, shared.ResidentInput{
        Address: &shared.ResidentAddress{
            City: "Lake Bettie",
            Country: meldapi.String("Bhutan"),
            CountyProvince: "occaecati",
            Line1: "labore",
            Line2: meldapi.String("quidem"),
            Line3: meldapi.String("atque"),
            Postcode: "79302-6469",
        },
        Contact: &shared.Contact{
            BusinessPhone: meldapi.String("provident"),
            BusinessPhoneExt: meldapi.String("repellendus"),
            CellPhone: meldapi.String("delectus"),
            CellPhoneExt: meldapi.String("voluptates"),
            Fax: meldapi.String("perferendis"),
            HomePhone: meldapi.String("est"),
            HomePhoneExt: meldapi.String("quidem"),
            PrimaryEmail: meldapi.String("Raquel_Pfannerstill@yahoo.com"),
            SecondaryEmail: meldapi.String("Alessia.Schiller54@yahoo.com"),
            TertiaryEmail: meldapi.String("Yvette_Leannon@yahoo.com"),
        },
        FirstName: "Rae",
        Invite: meldapi.Bool(false),
        LastName: "Borer",
        MiddleName: meldapi.String("esse"),
        Notes: meldapi.String("amet"),
        NotificationSettings: &shared.NotificationSettingsInput{
            IncomingMeldFrequency: shared.IncomingMeldFrequencyEnumNever.ToPointer(),
            SmsNotifications: meldapi.Bool(false),
            SuccessfulMeldFrequency: shared.SuccessfulMeldFrequencyEnumDaily.ToPointer(),
            Timezone: shared.TimezoneEnumAsiaThimphu.ToPointer(),
        },
        Units: []int64{
            887265,
            886961,
            880107,
        },
    }, operations.ResidentCreateJSONSecurity{
        PMOAuth2Authentication: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Resident != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [shared.ResidentInput](../../models/shared/residentinput.md)                                   | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `security`                                                                                     | [operations.ResidentCreateJSONSecurity](../../models/operations/residentcreatejsonsecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |


### Response

**[*operations.ResidentCreateJSONResponse](../../models/operations/residentcreatejsonresponse.md), error**


## ResidentCreateMultipart

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
    res, err := s.Resident.ResidentCreateMultipart(ctx, shared.ResidentInput1{
        Address: meldapi.String("natus"),
        Contact: &shared.Contact{
            BusinessPhone: meldapi.String("minima"),
            BusinessPhoneExt: meldapi.String("aspernatur"),
            CellPhone: meldapi.String("ex"),
            CellPhoneExt: meldapi.String("maiores"),
            Fax: meldapi.String("corrupti"),
            HomePhone: meldapi.String("at"),
            HomePhoneExt: meldapi.String("error"),
            PrimaryEmail: meldapi.String("Genevieve_Walker@yahoo.com"),
            SecondaryEmail: meldapi.String("Theresia.Parisian96@gmail.com"),
            TertiaryEmail: meldapi.String("Tevin10@gmail.com"),
        },
        FirstName: "Chandler",
        Invite: meldapi.Bool(false),
        LastName: "Halvorson",
        MiddleName: meldapi.String("laboriosam"),
        Notes: meldapi.String("velit"),
        NotificationSettings: &shared.NotificationSettingsInput{
            IncomingMeldFrequency: shared.IncomingMeldFrequencyEnumNever.ToPointer(),
            SmsNotifications: meldapi.Bool(false),
            SuccessfulMeldFrequency: shared.SuccessfulMeldFrequencyEnumDaily.ToPointer(),
            Timezone: shared.TimezoneEnumAmericaPanama.ToPointer(),
        },
        Units: []int64{
            160467,
            580107,
            886305,
            597937,
        },
    }, operations.ResidentCreateMultipartSecurity{
        PMOAuth2Authentication: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Resident != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [shared.ResidentInput1](../../models/shared/residentinput1.md)                                           | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `security`                                                                                               | [operations.ResidentCreateMultipartSecurity](../../models/operations/residentcreatemultipartsecurity.md) | :heavy_check_mark:                                                                                       | The security requirements to use for the request.                                                        |


### Response

**[*operations.ResidentCreateMultipartResponse](../../models/operations/residentcreatemultipartresponse.md), error**


## ResidentDestroy

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
    res, err := s.Resident.ResidentDestroy(ctx, operations.ResidentDestroyRequest{
        ID: "73e922a5-7a15-4be3-a060-807e2b6e3ab8",
    }, operations.ResidentDestroySecurity{
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ResidentDestroyRequest](../../models/operations/residentdestroyrequest.md)   | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.ResidentDestroySecurity](../../models/operations/residentdestroysecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


### Response

**[*operations.ResidentDestroyResponse](../../models/operations/residentdestroyresponse.md), error**


## ResidentList

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
    res, err := s.Resident.ResidentList(ctx, operations.ResidentListRequest{
        Limit: meldapi.Int64(523006),
        Offset: meldapi.Int64(304446),
        Ordering: meldapi.String("ad"),
    }, operations.ResidentListSecurity{
        PMOAuth2Authentication: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedResidentList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.ResidentListRequest](../../models/operations/residentlistrequest.md)   | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `security`                                                                         | [operations.ResidentListSecurity](../../models/operations/residentlistsecurity.md) | :heavy_check_mark:                                                                 | The security requirements to use for the request.                                  |


### Response

**[*operations.ResidentListResponse](../../models/operations/residentlistresponse.md), error**


## ResidentPartialUpdateForm

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
    res, err := s.Resident.ResidentPartialUpdateForm(ctx, operations.ResidentPartialUpdateFormRequest{
        PatchedResidentInput1: &shared.PatchedResidentInput1{
            Address: meldapi.String("repellat"),
            Contact: &shared.Contact{
                BusinessPhone: meldapi.String("alias"),
                BusinessPhoneExt: meldapi.String("corporis"),
                CellPhone: meldapi.String("perspiciatis"),
                CellPhoneExt: meldapi.String("nihil"),
                Fax: meldapi.String("mollitia"),
                HomePhone: meldapi.String("voluptas"),
                HomePhoneExt: meldapi.String("alias"),
                PrimaryEmail: meldapi.String("Waldo.Daniel@hotmail.com"),
                SecondaryEmail: meldapi.String("Marilyne92@gmail.com"),
                TertiaryEmail: meldapi.String("Elsa.Kreiger@yahoo.com"),
            },
            FirstName: meldapi.String("Loren"),
            Invite: meldapi.Bool(false),
            LastName: meldapi.String("Ferry"),
            MiddleName: meldapi.String("debitis"),
            Notes: meldapi.String("laudantium"),
            NotificationSettings: &shared.NotificationSettingsInput{
                IncomingMeldFrequency: shared.IncomingMeldFrequencyEnumDaily.ToPointer(),
                SmsNotifications: meldapi.Bool(false),
                SuccessfulMeldFrequency: shared.SuccessfulMeldFrequencyEnumDaily.ToPointer(),
                Timezone: shared.TimezoneEnumPacificNiue.ToPointer(),
            },
            Units: []int64{
                592081,
                337477,
            },
        },
        ID: "6f9251a5-a9da-4660-bf57-bfaad4f9efc1",
    }, operations.ResidentPartialUpdateFormSecurity{
        PMOAuth2Authentication: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Resident != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.ResidentPartialUpdateFormRequest](../../models/operations/residentpartialupdateformrequest.md)   | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `security`                                                                                                   | [operations.ResidentPartialUpdateFormSecurity](../../models/operations/residentpartialupdateformsecurity.md) | :heavy_check_mark:                                                                                           | The security requirements to use for the request.                                                            |


### Response

**[*operations.ResidentPartialUpdateFormResponse](../../models/operations/residentpartialupdateformresponse.md), error**


## ResidentPartialUpdateJSON

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
    res, err := s.Resident.ResidentPartialUpdateJSON(ctx, operations.ResidentPartialUpdateJSONRequest{
        PatchedResidentInput: &shared.PatchedResidentInput{
            Address: &shared.PatchedResidentAddress{
                City: "Goldnerborough",
                Country: meldapi.String("Belgium"),
                CountyProvince: "fugit",
                Line1: "cumque",
                Line2: meldapi.String("quae"),
                Line3: meldapi.String("perferendis"),
                Postcode: "14258",
            },
            Contact: &shared.Contact{
                BusinessPhone: meldapi.String("impedit"),
                BusinessPhoneExt: meldapi.String("eos"),
                CellPhone: meldapi.String("sapiente"),
                CellPhoneExt: meldapi.String("eum"),
                Fax: meldapi.String("dicta"),
                HomePhone: meldapi.String("minima"),
                HomePhoneExt: meldapi.String("beatae"),
                PrimaryEmail: meldapi.String("Lambert_Weber84@yahoo.com"),
                SecondaryEmail: meldapi.String("Tremaine_Metz39@gmail.com"),
                TertiaryEmail: meldapi.String("Geo76@hotmail.com"),
            },
            FirstName: meldapi.String("Marlene"),
            Invite: meldapi.Bool(false),
            LastName: meldapi.String("Dickens"),
            MiddleName: meldapi.String("animi"),
            Notes: meldapi.String("necessitatibus"),
            NotificationSettings: &shared.NotificationSettingsInput{
                IncomingMeldFrequency: shared.IncomingMeldFrequencyEnumNever.ToPointer(),
                SmsNotifications: meldapi.Bool(false),
                SuccessfulMeldFrequency: shared.SuccessfulMeldFrequencyEnumImmediately.ToPointer(),
                Timezone: shared.TimezoneEnumAmericaAntigua.ToPointer(),
            },
            Units: []int64{
                497777,
            },
        },
        ID: "996312fd-e047-4717-b8ff-61d017476360",
    }, operations.ResidentPartialUpdateJSONSecurity{
        PMOAuth2Authentication: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Resident != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.ResidentPartialUpdateJSONRequest](../../models/operations/residentpartialupdatejsonrequest.md)   | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `security`                                                                                                   | [operations.ResidentPartialUpdateJSONSecurity](../../models/operations/residentpartialupdatejsonsecurity.md) | :heavy_check_mark:                                                                                           | The security requirements to use for the request.                                                            |


### Response

**[*operations.ResidentPartialUpdateJSONResponse](../../models/operations/residentpartialupdatejsonresponse.md), error**


## ResidentPartialUpdateMultipart

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
    res, err := s.Resident.ResidentPartialUpdateMultipart(ctx, operations.ResidentPartialUpdateMultipartRequest{
        PatchedResidentInput1: &shared.PatchedResidentInput1{
            Address: meldapi.String("laborum"),
            Contact: &shared.Contact{
                BusinessPhone: meldapi.String("sunt"),
                BusinessPhoneExt: meldapi.String("nostrum"),
                CellPhone: meldapi.String("fugiat"),
                CellPhoneExt: meldapi.String("expedita"),
                Fax: meldapi.String("aliquid"),
                HomePhone: meldapi.String("officia"),
                HomePhoneExt: meldapi.String("suscipit"),
                PrimaryEmail: meldapi.String("Alanis_Kemmer@yahoo.com"),
                SecondaryEmail: meldapi.String("Anthony_Muller65@yahoo.com"),
                TertiaryEmail: meldapi.String("Mollie.Hane@hotmail.com"),
            },
            FirstName: meldapi.String("Bernie"),
            Invite: meldapi.Bool(false),
            LastName: meldapi.String("Skiles"),
            MiddleName: meldapi.String("ex"),
            Notes: meldapi.String("quo"),
            NotificationSettings: &shared.NotificationSettingsInput{
                IncomingMeldFrequency: shared.IncomingMeldFrequencyEnumDaily.ToPointer(),
                SmsNotifications: meldapi.Bool(false),
                SuccessfulMeldFrequency: shared.SuccessfulMeldFrequencyEnumImmediately.ToPointer(),
                Timezone: shared.TimezoneEnumAmericaRegina.ToPointer(),
            },
            Units: []int64{
                29950,
                561577,
                737254,
            },
        },
        ID: "61891baa-0fe1-4ade-808e-6f8c5f350d8c",
    }, operations.ResidentPartialUpdateMultipartSecurity{
        PMOAuth2Authentication: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Resident != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.ResidentPartialUpdateMultipartRequest](../../models/operations/residentpartialupdatemultipartrequest.md)   | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `security`                                                                                                             | [operations.ResidentPartialUpdateMultipartSecurity](../../models/operations/residentpartialupdatemultipartsecurity.md) | :heavy_check_mark:                                                                                                     | The security requirements to use for the request.                                                                      |


### Response

**[*operations.ResidentPartialUpdateMultipartResponse](../../models/operations/residentpartialupdatemultipartresponse.md), error**


## ResidentRetrieve

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
    res, err := s.Resident.ResidentRetrieve(ctx, operations.ResidentRetrieveRequest{
        ID: "db5a3418-1430-4104-a181-3d5208ece7e2",
    }, operations.ResidentRetrieveSecurity{
        PMOAuth2Authentication: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Resident != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ResidentRetrieveRequest](../../models/operations/residentretrieverequest.md)   | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `security`                                                                                 | [operations.ResidentRetrieveSecurity](../../models/operations/residentretrievesecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |


### Response

**[*operations.ResidentRetrieveResponse](../../models/operations/residentretrieveresponse.md), error**


## ResidentUpdateForm

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
    res, err := s.Resident.ResidentUpdateForm(ctx, operations.ResidentUpdateFormRequest{
        ResidentInput1: shared.ResidentInput1{
            Address: meldapi.String("veniam"),
            Contact: &shared.Contact{
                BusinessPhone: meldapi.String("nesciunt"),
                BusinessPhoneExt: meldapi.String("expedita"),
                CellPhone: meldapi.String("eum"),
                CellPhoneExt: meldapi.String("vel"),
                Fax: meldapi.String("voluptatum"),
                HomePhone: meldapi.String("magnam"),
                HomePhoneExt: meldapi.String("exercitationem"),
                PrimaryEmail: meldapi.String("Oswald.Jones92@gmail.com"),
                SecondaryEmail: meldapi.String("Alessia.Heller39@gmail.com"),
                TertiaryEmail: meldapi.String("Shaina.Orn98@hotmail.com"),
            },
            FirstName: "Terrence",
            Invite: meldapi.Bool(false),
            LastName: "Rowe",
            MiddleName: meldapi.String("occaecati"),
            Notes: meldapi.String("nemo"),
            NotificationSettings: &shared.NotificationSettingsInput{
                IncomingMeldFrequency: shared.IncomingMeldFrequencyEnumDaily.ToPointer(),
                SmsNotifications: meldapi.Bool(false),
                SuccessfulMeldFrequency: shared.SuccessfulMeldFrequencyEnumDaily.ToPointer(),
                Timezone: shared.TimezoneEnumCanadaSaskatchewan.ToPointer(),
            },
            Units: []int64{
                254025,
                364912,
            },
        },
        ID: "84273a84-18d1-4623-89fb-0929921aefb9",
    }, operations.ResidentUpdateFormSecurity{
        PMOAuth2Authentication: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Resident != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ResidentUpdateFormRequest](../../models/operations/residentupdateformrequest.md)   | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `security`                                                                                     | [operations.ResidentUpdateFormSecurity](../../models/operations/residentupdateformsecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |


### Response

**[*operations.ResidentUpdateFormResponse](../../models/operations/residentupdateformresponse.md), error**


## ResidentUpdateJSON

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
    res, err := s.Resident.ResidentUpdateJSON(ctx, operations.ResidentUpdateJSONRequest{
        ResidentInput: shared.ResidentInput{
            Address: &shared.ResidentAddress{
                City: "Framingham",
                Country: meldapi.String("Liechtenstein"),
                CountyProvince: "maxime",
                Line1: "magnam",
                Line2: meldapi.String("temporibus"),
                Line3: meldapi.String("quos"),
                Postcode: "94592",
            },
            Contact: &shared.Contact{
                BusinessPhone: meldapi.String("nam"),
                BusinessPhoneExt: meldapi.String("vero"),
                CellPhone: meldapi.String("voluptatem"),
                CellPhoneExt: meldapi.String("ipsam"),
                Fax: meldapi.String("vel"),
                HomePhone: meldapi.String("alias"),
                HomePhoneExt: meldapi.String("quasi"),
                PrimaryEmail: meldapi.String("Wilford_Hamill@gmail.com"),
                SecondaryEmail: meldapi.String("Jaden_Hickle@yahoo.com"),
                TertiaryEmail: meldapi.String("Kiana_Thompson90@yahoo.com"),
            },
            FirstName: "Zachariah",
            Invite: meldapi.Bool(false),
            LastName: "Jakubowski",
            MiddleName: meldapi.String("voluptas"),
            Notes: meldapi.String("debitis"),
            NotificationSettings: &shared.NotificationSettingsInput{
                IncomingMeldFrequency: shared.IncomingMeldFrequencyEnumNever.ToPointer(),
                SmsNotifications: meldapi.Bool(false),
                SuccessfulMeldFrequency: shared.SuccessfulMeldFrequencyEnumImmediately.ToPointer(),
                Timezone: shared.TimezoneEnumEuropeSaratov.ToPointer(),
            },
            Units: []int64{
                675689,
                231070,
                244889,
            },
        },
        ID: "83c2beb4-7737-43c8-972f-64d1db1f2c43",
    }, operations.ResidentUpdateJSONSecurity{
        PMOAuth2Authentication: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Resident != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ResidentUpdateJSONRequest](../../models/operations/residentupdatejsonrequest.md)   | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `security`                                                                                     | [operations.ResidentUpdateJSONSecurity](../../models/operations/residentupdatejsonsecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |


### Response

**[*operations.ResidentUpdateJSONResponse](../../models/operations/residentupdatejsonresponse.md), error**


## ResidentUpdateMultipart

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
    res, err := s.Resident.ResidentUpdateMultipart(ctx, operations.ResidentUpdateMultipartRequest{
        ResidentInput1: shared.ResidentInput1{
            Address: meldapi.String("illo"),
            Contact: &shared.Contact{
                BusinessPhone: meldapi.String("accusantium"),
                BusinessPhoneExt: meldapi.String("vel"),
                CellPhone: meldapi.String("ea"),
                CellPhoneExt: meldapi.String("beatae"),
                Fax: meldapi.String("vero"),
                HomePhone: meldapi.String("excepturi"),
                HomePhoneExt: meldapi.String("eum"),
                PrimaryEmail: meldapi.String("Ed.Metz77@gmail.com"),
                SecondaryEmail: meldapi.String("Leonora39@hotmail.com"),
                TertiaryEmail: meldapi.String("Deangelo22@hotmail.com"),
            },
            FirstName: "Jayne",
            Invite: meldapi.Bool(false),
            LastName: "Bailey",
            MiddleName: meldapi.String("doloremque"),
            Notes: meldapi.String("consequatur"),
            NotificationSettings: &shared.NotificationSettingsInput{
                IncomingMeldFrequency: shared.IncomingMeldFrequencyEnumDaily.ToPointer(),
                SmsNotifications: meldapi.Bool(false),
                SuccessfulMeldFrequency: shared.SuccessfulMeldFrequencyEnumNever.ToPointer(),
                Timezone: shared.TimezoneEnumAsiaAshkhabad.ToPointer(),
            },
            Units: []int64{
                377406,
                705148,
                809365,
            },
        },
        ID: "9b8f759e-ac55-4a97-81d3-11352965bb8a",
    }, operations.ResidentUpdateMultipartSecurity{
        PMOAuth2Authentication: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Resident != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.ResidentUpdateMultipartRequest](../../models/operations/residentupdatemultipartrequest.md)   | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `security`                                                                                               | [operations.ResidentUpdateMultipartSecurity](../../models/operations/residentupdatemultipartsecurity.md) | :heavy_check_mark:                                                                                       | The security requirements to use for the request.                                                        |


### Response

**[*operations.ResidentUpdateMultipartResponse](../../models/operations/residentupdatemultipartresponse.md), error**

