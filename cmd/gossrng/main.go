package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/e-kucheriavyi/gossrng/configs"
	"github.com/e-kucheriavyi/gossrng/pkg/export"
	"github.com/e-kucheriavyi/gossrng/pkg/pages"
	"github.com/e-kucheriavyi/gossrng/pkg/sitemap"
	"github.com/e-kucheriavyi/gossrng/pkg/static"
)

func main() {
	args := os.Args

	if len(args) != 2 {
		fmt.Println("Check out README.md for instructions")
		return
	}

	mode := args[1]

	switch mode {
	case "serve":
		serve()
		return
	case "export":
		exportSite()
		return
	}

	fmt.Println("Invalid mode:", mode)
}

func serve() {
	fmt.Println("Serving on", configs.Host)

	sitemap.ServeSitemap()
	static.ServeStatic()
	pages.ServePages()

	err := http.ListenAndServe(configs.Host, nil)

	if err != nil {
		log.Fatal(err.Error())
	}
}

func exportSite() {
	fmt.Println("exporting")

	err := export.Export()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("done")
}
