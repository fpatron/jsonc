# Better Json

[![Go CI](https://github.com/FrancisPatron/betterjson/actions/workflows/go.yml/badge.svg)](https://github.com/fpatron/betterjson/actions/workflows/go.yml)

`Better Json` is a Go library that enhances JSON parsing by supporting comments within JSON files.

## Instalation
To install `betterjson`, use `go get`:   
```sh
 go get github.com/fpatron/betterjson
```

## Features

- **Comment Support**: Easily parse JSON files with single-line and multi-line comments.
- **Simple API**: Uses a familiar API similar to the standard `encoding/json` package.

## Usage

Here's a basic example of how to use `betterjson`:

```go
package main

import (
	"fmt"
	"log"
	"os"
	"github.com/fpatron/betterjson"
)

func main() {
	data, err := os.ReadFile("example.json")
	if err != nil {
		log.Fatalf("Failed to read example.json: %v", err)
	}

	var example ExampleData
	err = betterjson.Unmarshal(data, &example)
	if err != nil {
		log.Fatalf("Failed to parse example.json: %v", err)
	}

	fmt.Printf("Parsed data: %+v\n", example)
}

```
for a complete example checkout this [example code](https://github.com/fpatron/betterjson/blob/main/example/main.go)

## v2 — encoding/json/v2 support

A separate `betterjson/v2` module is available for projects that want to use Go's experimental [`encoding/json/v2`](https://pkg.go.dev/encoding/json/v2) backend. It exposes the same behaviour with the richer v2 `Unmarshal(data, v, ...json.Options)` signature.

> **Note:** `encoding/json/v2` requires Go 1.26 and the `GOEXPERIMENT=jsonv2` build flag.

See the [v2 README](v2/README.md) for installation instructions, a full API comparison, and usage examples.
