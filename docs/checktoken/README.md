# CheckToken

### Available Operations

* [CheckTokenRetrieve](#checktokenretrieve) - Used to check the validity of an oauth2 token

## CheckTokenRetrieve

Used to check the validity of an oauth2 token

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
    res, err := s.CheckToken.CheckTokenRetrieve(ctx, operations.CheckTokenRetrieveSecurity{
        PMOAuth2Authentication: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CheckTokenRetrieve200ApplicationJSONObject != nil {
        // handle response
    }
}
```
