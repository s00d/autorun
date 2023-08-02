# autorun-go

[![GoDoc](https://godoc.org/github.com/s00d/autorun-go?status.svg)](https://godoc.org/github.com/s00d/autorun-go)

A Go library to run a command after login.

## Usage

```go
package main

import (
	"log"
	"github.com/s00d/autorun-go"
)

func main() {
	autorun := NewAutoStart()

	if autorun.IsEnabled() {
		log.Println("App is already enabled, removing it...")

		if err := autorun.Disable(); err != nil {
			log.Fatal(err)
		}
	} else {
		log.Println("Enabling app...")

		if err := autorun.Enable(); err != nil {
			log.Fatal(err)
		}
	}

	log.Println("Done!")
}
```

## License

MIT