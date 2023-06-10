# Invoice

### Available Operations

* [InvoiceAttachmentRetrieve](#invoiceattachmentretrieve)
* [InvoiceList](#invoicelist)
* [InvoiceRetrieve](#invoiceretrieve)

## InvoiceAttachmentRetrieve

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
    res, err := s.Invoice.InvoiceAttachmentRetrieve(ctx, operations.InvoiceAttachmentRetrieveRequest{
        ID: "e6e13b99-d488-4e1e-91e4-50ad2abd4426",
    }, operations.InvoiceAttachmentRetrieveSecurity{
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

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.InvoiceAttachmentRetrieveRequest](../../models/operations/invoiceattachmentretrieverequest.md)   | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `security`                                                                                                   | [operations.InvoiceAttachmentRetrieveSecurity](../../models/operations/invoiceattachmentretrievesecurity.md) | :heavy_check_mark:                                                                                           | The security requirements to use for the request.                                                            |


### Response

**[*operations.InvoiceAttachmentRetrieveResponse](../../models/operations/invoiceattachmentretrieveresponse.md), error**


## InvoiceList

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
    res, err := s.Invoice.InvoiceList(ctx, operations.InvoiceListRequest{
        XPmOrg: 586513,
        CreatedGte: meldapi.String("quos"),
        CreatedInterval: meldapi.String("perferendis"),
        CreatedLte: meldapi.String("magni"),
        Declined: operations.InvoiceListDeclinedFalse.ToPointer(),
        Limit: meldapi.Int64(369808),
        Offset: meldapi.Int64(4695),
        Ordering: meldapi.String("fugit"),
        Paid: operations.InvoiceListPaidTrue.ToPointer(),
        PaymentRequestedGte: meldapi.String("excepturi"),
        PaymentRequestedInterval: meldapi.String("tempora"),
        PaymentRequestedLte: meldapi.String("facilis"),
        Status: operations.InvoiceListStatusApproved.ToPointer(),
        Vendor: meldapi.String("labore"),
        Vendors: meldapi.String("delectus"),
    }, operations.InvoiceListSecurity{
        PMOAuth2Authentication: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PaginatedMeldInvoiceSerializerListList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.InvoiceListRequest](../../models/operations/invoicelistrequest.md)   | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `security`                                                                       | [operations.InvoiceListSecurity](../../models/operations/invoicelistsecurity.md) | :heavy_check_mark:                                                               | The security requirements to use for the request.                                |


### Response

**[*operations.InvoiceListResponse](../../models/operations/invoicelistresponse.md), error**


## InvoiceRetrieve

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
    res, err := s.Invoice.InvoiceRetrieve(ctx, operations.InvoiceRetrieveRequest{
        XPmOrg: 433288,
        ID: "3c969e9a-3efa-477d-bb14-cd66ae395efb",
    }, operations.InvoiceRetrieveSecurity{
        PMOAuth2Authentication: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.MeldInvoiceSerializerDetail != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.InvoiceRetrieveRequest](../../models/operations/invoiceretrieverequest.md)   | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.InvoiceRetrieveSecurity](../../models/operations/invoiceretrievesecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


### Response

**[*operations.InvoiceRetrieveResponse](../../models/operations/invoiceretrieveresponse.md), error**

