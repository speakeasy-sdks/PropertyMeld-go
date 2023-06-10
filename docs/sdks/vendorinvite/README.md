# VendorInvite

### Available Operations

* [VendorInviteCreate](#vendorinvitecreate)

## VendorInviteCreate

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
    res, err := s.VendorInvite.VendorInviteCreate(ctx, operations.VendorInviteCreateRequest{
        VendorInvite: shared.VendorInvite{
            CompanyName: meldapi.String("fugiat"),
            CompanyPhone: meldapi.String("unde"),
            Email: "Jeromy62@hotmail.com",
            FirstName: meldapi.String("Owen"),
            LastName: meldapi.String("Brown"),
            Line1: meldapi.String("dignissimos"),
            Postcode: meldapi.String("83139"),
        },
        XPmOrg: 482584,
    }, operations.VendorInviteCreateSecurity{
        PMOAuth2Authentication: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.VendorInvite != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.VendorInviteCreateRequest](../../models/operations/vendorinvitecreaterequest.md)   | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `security`                                                                                     | [operations.VendorInviteCreateSecurity](../../models/operations/vendorinvitecreatesecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |


### Response

**[*operations.VendorInviteCreateResponse](../../models/operations/vendorinvitecreateresponse.md), error**

