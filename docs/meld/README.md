# Meld

### Available Operations

* [MeldList](#meldlist)
* [MeldManagerFilesList](#meldmanagerfileslist)
* [MeldRetrieve](#meldretrieve)
* [MeldTenantFilesList](#meldtenantfileslist)
* [MeldVendorFilesList](#meldvendorfileslist)

## MeldList

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
    res, err := s.Meld.MeldList(ctx, operations.MeldListRequest{
        XPmOrg: 221262,
        AckApproval: operations.MeldListAckApprovalFalse.ToPointer(),
        AssignedToMe: meldapi.String("odit"),
        Categories: operations.MeldListCategoriesGeneral.ToPointer(),
        CommentsGte: meldapi.String("quasi"),
        CommentsLte: meldapi.String("iure"),
        CompletedGte: meldapi.String("doloribus"),
        CompletedInterval: meldapi.String("debitis"),
        CompletedLte: meldapi.String("eius"),
        CreatedGte: meldapi.String("maxime"),
        CreatedInterval: meldapi.String("deleniti"),
        CreatedLte: meldapi.String("facilis"),
        CreatedOffsetLte: meldapi.String("in"),
        DueDateChoice: operations.MeldListDueDateChoiceOverdue.ToPointer(),
        DueDateGte: meldapi.String("architecto"),
        DueDateLte: meldapi.String("repudiandae"),
        EverBeenAssigned: operations.MeldListEverBeenAssignedTrue.ToPointer(),
        Exp: operations.MeldListExpFalse.ToPointer(),
        ExpS: meldapi.String("nihil"),
        FilesGte: meldapi.String("repellat"),
        FilesLte: meldapi.String("quibusdam"),
        HasEstimates: operations.MeldListHasEstimatesFalse.ToPointer(),
        InvoiceS: meldapi.String("saepe"),
        Limit: meldapi.Int64(868126),
        Maint: meldapi.String("accusantium"),
        MaintType: operations.MeldListMaintTypeOne.ToPointer(),
        MarkedCompleteGte: meldapi.String("praesentium"),
        MarkedCompleteInterval: meldapi.String("natus"),
        MarkedCompleteLte: meldapi.String("magni"),
        Meldinvoice: operations.MeldListMeldinvoiceTrue.ToPointer(),
        Offset: meldapi.Int64(779051),
        OrderBy: operations.MeldListOrderByMinusOwnerApprovalRequestDate.ToPointer(),
        OwnerApproval: operations.MeldListOwnerApprovalOwnerApprovalRequestedEstimates.ToPointer(),
        Pg: meldapi.String("maxime"),
        Priority: operations.MeldListPriorityMedium.ToPointer(),
        Project: meldapi.String("excepturi"),
        ProjectAssigned: operations.MeldListProjectAssignedTrue.ToPointer(),
        Prop: meldapi.String("ea"),
        Rating: operations.MeldListRatingOne.ToPointer(),
        Recurring: operations.MeldListRecurringTrue.ToPointer(),
        RecurringMeld: meldapi.String("maiores"),
        Remindees: meldapi.String("quidem"),
        ReminderChoice: operations.MeldListReminderChoiceToday.ToPointer(),
        ReminderGte: meldapi.String("voluptate"),
        ReminderLte: meldapi.String("autem"),
        ScheduledGte: meldapi.String("nam"),
        ScheduledInterval: meldapi.String("eaque"),
        ScheduledLte: meldapi.String("pariatur"),
        Scheduling: meldapi.String("nemo"),
        Status: operations.MeldListStatusPendingAssignment.ToPointer(),
        Tags: meldapi.String("perferendis"),
        TagsEx: meldapi.String("fugiat"),
        Tpr: operations.MeldListTprFalse.ToPointer(),
        Unit: meldapi.String("aut"),
        UpdatedGte: meldapi.String("cumque"),
        UpdatedInterval: meldapi.String("corporis"),
        UpdatedLte: meldapi.String("hic"),
        UpdatedOffsetLte: meldapi.String("libero"),
        VendorScheduledGte: meldapi.String("nobis"),
        VendorScheduledInterval: meldapi.String("dolores"),
        VendorScheduledLte: meldapi.String("quis"),
        WlGte: meldapi.String("totam"),
        WlLte: meldapi.String("dignissimos"),
    }, operations.MeldListSecurity{
        PMOAuth2Authentication: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedMeldSerializerListList != nil {
        // handle response
    }
}
```

## MeldManagerFilesList

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
    res, err := s.Meld.MeldManagerFilesList(ctx, operations.MeldManagerFilesListRequest{
        XPmOrg: 54338,
        ID: "53202c73-d5fe-49b9-8c28-909b3fe49a8d",
        Limit: meldapi.Int64(589910),
        Offset: meldapi.Int64(750844),
        Ordering: meldapi.String("libero"),
    }, operations.MeldManagerFilesListSecurity{
        PMOAuth2Authentication: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedPMAPIMeldfileList != nil {
        // handle response
    }
}
```

## MeldRetrieve

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
    res, err := s.Meld.MeldRetrieve(ctx, operations.MeldRetrieveRequest{
        XPmOrg: 964490,
        ID: "48633323-f9b7-47f3-a410-0674ebf69280",
    }, operations.MeldRetrieveSecurity{
        PMOAuth2Authentication: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.MeldSerializerDetail != nil {
        // handle response
    }
}
```

## MeldTenantFilesList

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
    res, err := s.Meld.MeldTenantFilesList(ctx, operations.MeldTenantFilesListRequest{
        XPmOrg: 854614,
        ID: "1ba77a89-ebf7-437a-a420-3ce5e6a95d8a",
        Limit: meldapi.Int64(55),
        Offset: meldapi.Int64(872651),
        Ordering: meldapi.String("quaerat"),
    }, operations.MeldTenantFilesListSecurity{
        PMOAuth2Authentication: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedPmAPITenantMeldList != nil {
        // handle response
    }
}
```

## MeldVendorFilesList

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
    res, err := s.Meld.MeldVendorFilesList(ctx, operations.MeldVendorFilesListRequest{
        XPmOrg: 273542,
        ID: "6ce2af7a-73cf-43be-853f-870b326b5a73",
        Limit: meldapi.Int64(277628),
        Offset: meldapi.Int64(186458),
        Ordering: meldapi.String("cupiditate"),
    }, operations.MeldVendorFilesListSecurity{
        PMOAuth2Authentication: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedPMAPIVendorMeldFileList != nil {
        // handle response
    }
}
```
