# Go Methodology

`github.com/bagaking/aplug/methodology` is the public package in this
repository. It provides an in-memory, thread-safe container for methodology
records that can be listed, queried, updated, and cloned across API boundaries.

## Scope

The package includes:

- `Methodology`: the structured record type.
- `Methodologies`: a mutex-protected container keyed by methodology id.
- Default loading from the embedded `default.toml` file.
- Query helpers for all records, one or more keys, and scenario matches.
- Mutation helpers for set, delete, and callback-style updates.

It does not provide persistence, external plugin discovery, cross-process
synchronization, or an agent runtime.

## Install

```sh
go get github.com/bagaking/aplug/methodology
```

## Example

```go
package main

import (
	"fmt"

	"github.com/bagaking/aplug/methodology"
)

func main() {
	methods := methodology.NewContainer()

	methods.Set("planning", &methodology.Methodology{
		ID:       "planning",
		Usage:    "Use when a task needs a concrete next action.",
		Scenario: []string{"planning", "execution"},
		Steps:    []string{"Define the goal", "Choose the next action"},
	})

	found := methods.TryGet("planning")
	if found != nil {
		fmt.Println(found.Usage)
	}
}
```

## Validation

From the repository root:

```sh
go mod verify
go test ./...
```

Use `go test -race ./...` before changing container locking or clone semantics.
