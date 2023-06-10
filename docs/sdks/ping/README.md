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

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.PingRetrieveResponse](../../models/operations/pingretrieveresponse.md), error**

