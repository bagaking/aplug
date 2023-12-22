# Go Methodology

This package contains the source code for a Go implementation of a simple methodology management system, designed to be used by AI agents.
Overview

This methodology management system is designed to provide a structured way of managing and retrieving methodologies which are strategies or approaches to solve problems or complete tasks. Each methodology is associated with a unique key and contains detailed information such as its usage, main idea, applicable scenarios, strategy, steps and examples.
AI agent can use this system to store and retrieve methodologies. For instance, the agent can fetch a list of all methodologies, retrieve detailed information about a specific methodology by its key, get methodologies applicable to a specific scenario and update or delete a methodology.

# Files

There are 5 Go files included in this project:
- `conf.go`: This file contains the code to load the default configuration from a TOML file.
- `container.go`: This file includes the code for the methodology container that keeps all the methodologies in-memory and provides methods for retrieving them.
- `model.go`: This file defines the data model of a methodology.
- `modify.go`: This file includes the methods to modify the methodologies, like adding, updating or deleting a methodology.
- `query.go`: This file contains the methods to retrieve the methodologies, either all or a specific one by its key or scenario.

# Usage

This system is meant to be used in a Go environment. You can import it into your Go project and use its interfaces to manage methodologies. Here is a simple usage example:

```go
import "github.com/yourusername/go-methodology/methodology"

func main() {
// Create a new methodology container
container := methodology.NewContainer()

    // Use the Set method to add a new methodology
    container.Set("key", &methodology.Methodology{
        ID:     "key",
        Usage:  "usage",
        Scenario: []string{"scenario1", "scenario2"},
    })

    // Use the TryGet method to retrieve a methodology by its key
    method := container.TryGet("key")
    fmt.Println(method)
}
```
# Using with AI Agent

An AI assistant needs to be able to solve a wide variety of problems and complete tasks in different scenarios, 
and therefore needs to know various methodologies. This system allows an AI assistant to store the methodologies 
it learns or originates in a structured way, and retrieve them when encountering a problem or task, enhancing its 
ability to handle different situations.

Whether it's a simple question answering, helping developers write code, or more complex tasks like project management, 
methodologies stored in this system that AI agent can rely on to provide high-quality assistance.