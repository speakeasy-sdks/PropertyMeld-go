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
