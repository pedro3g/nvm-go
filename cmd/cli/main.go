package main

import (
	"os"

	"github.com/pedro3g/nvm-go/handlers"
)

func main() {
	mainArg := os.Args[1]

	if mainArg == "list" {
		handlers.List()
	}
}
