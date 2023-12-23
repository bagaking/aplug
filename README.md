# aplug

`aplug` provides Go plugin and methodology utilities. The current public API is
focused on the `methodology` package, which keeps a collection of structured
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

## Install

```sh
go get github.com/bagaking/aplug
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
