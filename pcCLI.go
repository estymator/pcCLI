package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	valid "github.com/estymator/pcCLI/validate"
	web "github.com/estymator/pcCLI/web"
)

const CLI_VERSION string = "v1.0.0"

func main() {
	command, err := valid.CheckCommand()

	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}

	switch command {
	case "help":
		flag.Usage()
	case "version":
		fmt.Printf("version: %s", CLI_VERSION)
	default:
		web.ServeFile(command)
	}
}
