# aplug

`aplug` provides Go methodology utilities. The current public API is the
`methodology` package, which keeps a collection of structured
methodologies in memory and supports loading built-in defaults from an embedded
TOML file.

## Packages

- `methodology`: defines the `Methodology` model and a thread-safe
  `Methodologies` container.
- Default methodology loading: `DefaultContainer` loads the embedded
  `methodology/default.toml` data once, and `NewContainer` starts from a cloned
  snapshot of those defaults.
- Query helpers: list all methodologies, get one or more by key, and find
  methodologies by scenario.
- Update helpers: set, delete, or update a methodology while storing cloned
  values so callers do not mutate container state accidentally.

## API Notes / Behavior Boundaries

- Default methodologies are embedded from `methodology/default.toml` and loaded
  once into the process-wide `DefaultContainer`.
- `NewContainer` starts from a locked, cloned snapshot of the current default
  container. Later changes to that new container do not mutate the defaults.
- `Methodologies` protects map access with an internal `sync.RWMutex`, and the
  exported methods clone values crossing the API boundary. `Set` stores a copy;
  `List`, `TryGet`, `MGet`, and `GetByScene` return copies; `UpdatedBy` passes a
  copy to the callback and stores a copy of the returned value.
- `Set` upserts by key and can initialize a zero-value container. It does not
  reconcile the map key with `Methodology.ID`.
- `Delete` is a no-op for missing keys, `TryGet` returns `nil` for a missing
  key, `MGet` skips missing keys, and `UpdatedBy` returns an error when the key
  does not exist.
- The package is an in-memory container. It does not provide persistence,
  plugin discovery/loading, or synchronization beyond the in-process mutex.

## Install

```sh
go get github.com/bagaking/aplug/methodology
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/bagaking/aplug/methodology"
)

func main() {
	methods := methodology.NewContainer()

	methods.Set("example", &methodology.Methodology{
		ID:       "example",
		Usage:    "Describe when this methodology should be used.",
		Scenario: []string{"planning"},
		Steps:    []string{"Define the goal", "Choose the next action"},
	})

	found := methods.TryGet("example")
	fmt.Println(found.Usage)
}
```

## Local Validation

```sh
go test ./...
```

## License

MIT. See [LICENSE](LICENSE).
