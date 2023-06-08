# Ping

### Available Operations

* [PingRetrieve](#pingretrieve) - Used to double check that the api is up an running.

## PingRetrieve

Used to double check that the api is up an running.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"MeldAPI"
)

func main() {
    s := meldapi.New()

    ctx := context.Background()
    res, err := s.Ping.PingRetrieve(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.PingRetrieve200ApplicationJSONObject != nil {
        // handle response
    }
}
```
