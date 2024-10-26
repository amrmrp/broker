package app

import (
    "fmt"
    "async-entity-fetcher/pkg/example"
)

// Start initializes the application
func Start() {
    fmt.Println("Application is running...")
    result := example.SayHello("World")
    fmt.Println(result)
}
