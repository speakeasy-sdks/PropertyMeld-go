# Property

### Available Operations

* [PropertyCreate](#propertycreate)
* [PropertyDestroy](#propertydestroy)
* [PropertyList](#propertylist)
* [PropertyPartialUpdate](#propertypartialupdate)
* [PropertyRetrieve](#propertyretrieve)
* [PropertyUpdate](#propertyupdate)

## PropertyCreate

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
    res, err := s.Property.PropertyCreate(ctx, shared.PropertyInput{
        City: "Destineecester",
        Country: meldapi.String("Svalbard & Jan Mayen Islands"),
        CountyProvince: "soluta",
        Line1: "accusantium",
        Line2: meldapi.String("aliquam"),
        Line3: meldapi.String("sapiente"),
        MaintenanceNotes: meldapi.String("dicta"),
        OtherAddressDetails: meldapi.String("ullam"),
        Owners: []int64{
            356707,
            391774,
        },
        Postcode: "51845",
        PropertyGroups: []int64{
            680270,
            99615,
            609178,
            945302,
        },
        PropertyName: meldapi.String("quasi"),
        Units: []int64{
            92027,
            454162,
            55965,
            326701,
        },
    }, operations.PropertyCreateSecurity{
        PMOAuth2Authentication: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Property != nil {
        // handle response
    }
}
```

## PropertyDestroy

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
    res, err := s.Property.PropertyDestroy(ctx, operations.PropertyDestroyRequest{
        ID: "1339d080-86a1-4840-b94c-26071f93f5f0",
    }, operations.PropertyDestroySecurity{
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

## PropertyList

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
    res, err := s.Property.PropertyList(ctx, operations.PropertyListRequest{
        Limit: meldapi.Int64(409054),
        Offset: meldapi.Int64(310067),
        Ordering: meldapi.String("consequuntur"),
    }, operations.PropertyListSecurity{
        PMOAuth2Authentication: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedPropertyList != nil {
        // handle response
    }
}
```

## PropertyPartialUpdate

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
    res, err := s.Property.PropertyPartialUpdate(ctx, operations.PropertyPartialUpdateRequest{
        PatchedPropertyInput: &shared.PatchedPropertyInput{
            City: meldapi.String("Oklahoma City"),
            Country: meldapi.String("Singapore"),
            CountyProvince: meldapi.String("dignissimos"),
            Line1: meldapi.String("officia"),
            Line2: meldapi.String("asperiores"),
            Line3: meldapi.String("nemo"),
            MaintenanceNotes: meldapi.String("quae"),
            OtherAddressDetails: meldapi.String("quaerat"),
            Owners: []int64{
                801836,
                288398,
                70447,
                241418,
            },
            Postcode: meldapi.String("63266-9584"),
            PropertyGroups: []int64{
                554688,
                427834,
            },
            PropertyName: meldapi.String("labore"),
            Units: []int64{
                706575,
                738227,
                414857,
                447144,
            },
        },
        ID: "5fd5e60b-375e-4d4f-afbe-e41f33317fe3",
    }, operations.PropertyPartialUpdateSecurity{
        PMOAuth2Authentication: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Property != nil {
        // handle response
    }
}
```

## PropertyRetrieve

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
    res, err := s.Property.PropertyRetrieve(ctx, operations.PropertyRetrieveRequest{
        ID: "5b60eb1e-a426-4555-ba3c-28744ed53b88",
    }, operations.PropertyRetrieveSecurity{
        PMOAuth2Authentication: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Property != nil {
        // handle response
    }
}
```

## PropertyUpdate

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
    res, err := s.Property.PropertyUpdate(ctx, operations.PropertyUpdateRequest{
        PropertyInput: shared.PropertyInput{
            City: "Compton",
            Country: meldapi.String("New Caledonia"),
            CountyProvince: "corrupti",
            Line1: "pariatur",
            Line2: meldapi.String("totam"),
            Line3: meldapi.String("hic"),
            MaintenanceNotes: meldapi.String("exercitationem"),
            OtherAddressDetails: meldapi.String("nobis"),
            Owners: []int64{
                699575,
            },
            Postcode: "91974",
            PropertyGroups: []int64{
                70869,
                611749,
                292794,
            },
            PropertyName: meldapi.String("laborum"),
            Units: []int64{
                447516,
            },
        },
        ID: "6b26916f-e1f0-48f4-a94e-3698f447f603",
    }, operations.PropertyUpdateSecurity{
        PMOAuth2Authentication: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Property != nil {
        // handle response
    }
}
```
