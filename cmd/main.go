package main

import (
	"gitea.flying-lama.com/apomedica/product-management/internal/app"
	"log"
	"os"
)

func main() {
	application := app.NewApp()

	if len(os.Args) == 1 {
		application.Usage()
		os.Exit(0)
	}

	function := os.Args[1]
	params := os.Args[2:]

	switch function {
	case "serve":
		err := application.Serve(params)
		if err != nil {
			log.Fatal(err)
		}
	case "migrate":
		err := application.Migrate(params)
		if err != nil {
			log.Fatal(err)
		}
	default:
		application.Usage()
	}

}
