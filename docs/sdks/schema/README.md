# Schema

### Available Operations

* [SchemaRetrieve](#schemaretrieve) - OpenApi3 schema for this API. Format can be selected via content negotiation.

- YAML: application/vnd.oai.openapi
- JSON: application/vnd.oai.openapi+json

## SchemaRetrieve

OpenApi3 schema for this API. Format can be selected via content negotiation.

- YAML: application/vnd.oai.openapi
- JSON: application/vnd.oai.openapi+json

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
    res, err := s.Schema.SchemaRetrieve(ctx, operations.SchemaRetrieveRequest{
        XPmOrg: 425402,
        Format: operations.SchemaRetrieveFormatJSON.ToPointer(),
    }, operations.SchemaRetrieveSecurity{
        PMOAuth2Authentication: meldapi.String(""),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SchemaRetrieve200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.SchemaRetrieveRequest](../../models/operations/schemaretrieverequest.md)   | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.SchemaRetrieveSecurity](../../models/operations/schemaretrievesecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


### Response

**[*operations.SchemaRetrieveResponse](../../models/operations/schemaretrieveresponse.md), error**

