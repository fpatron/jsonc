# betterjson/v2

`betterjson/v2` is the experimental v2 variant of [betterjson](https://github.com/FrancisPatron/betterjson), backed by Go's [`encoding/json/v2`](https://pkg.go.dev/encoding/json/v2) package instead of `encoding/json`.

## Requirements

`encoding/json/v2` is experimental in Go 1.26 and must be enabled with the `GOEXPERIMENT=jsonv2` environment variable at build time.

## Installation

```sh
GOEXPERIMENT=jsonv2 go get github.com/fpatron/betterjson/v2
```

## Differences from v1

| | `betterjson` (v1) | `betterjson/v2` |
|---|---|---|
| JSON backend | `encoding/json` | `encoding/json/v2` |
| `Unmarshal` signature | `([]byte, interface{})` | `([]byte, any, ...json.Options)` |
| Duplicate keys | allowed | rejected by default |
| Case-sensitive field matching | no (loose) | yes (strict) by default |
| Invalid UTF-8 | replaced silently | rejected by default |
| Build requirement | none | `GOEXPERIMENT=jsonv2` |

## Usage

```go
package main

import (
	"fmt"
	"log"
	"os"

	betterjson "github.com/fpatron/betterjson/v2"
)

func main() {
	data, err := os.ReadFile("example.json")
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	var example ExampleData
	if err := betterjson.Unmarshal(data, &example); err != nil {
		log.Fatalf("Failed to parse: %v", err)
	}

	fmt.Printf("Parsed data: %+v\n", example)
}
```

### Passing `encoding/json/v2` options

The variadic `opts ...json.Options` parameter is passed directly to `encoding/json/v2`, so any v2 option works:

```go
import (
	betterjson "github.com/fpatron/betterjson/v2"
	"encoding/json/v2"
)

// Reject unknown fields
err := betterjson.Unmarshal(data, &cfg, json.RejectUnknownMembers(true))

// Allow duplicate keys (opt back into v1-like behaviour)
err := betterjson.Unmarshal(data, &cfg, jsontext.AllowDuplicateNames(true))
```

## Building and testing

All build and test commands must be run with `GOEXPERIMENT=jsonv2`:

```sh
GOEXPERIMENT=jsonv2 go build ./...
GOEXPERIMENT=jsonv2 go test ./...
```
