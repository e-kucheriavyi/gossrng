package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/e-kucheriavyi/gossrng/pkg/content"
	"github.com/e-kucheriavyi/gossrng/pkg/export"
	"github.com/e-kucheriavyi/gossrng/pkg/pages"
	"github.com/e-kucheriavyi/gossrng/pkg/sitemap"
	"github.com/e-kucheriavyi/gossrng/pkg/static"
)

func main() {
	mode := flag.String("m", "export", "gossrng mode. Can be `init`, `serve` or `export`(default)")
	root := flag.String("r", "", "content directory path")
	port := flag.String("p", ":3030", "Serving port for SSMG (SSR) mode")
	flag.Parse()

	if *root == "" {
		fmt.Println("You have to specify the content path with `-r` flag ")
		return
	}

	switch *mode {
	case "serve":
		serve(*root, *port)
		return
	case "export":
		exportSite(*root)
		return
	case "init":
		initRoot(*root)
	}

	fmt.Println("Invalid mode:", mode)
}

func serve(root, port string) {
	fmt.Println("Serving on", port)

	sitemap.ServeSitemap()
	static.ServeStatic(root)
	pages.ServePages(root)

	err := http.ListenAndServe(port, nil)

	if err != nil {
		log.Fatal(err.Error())
	}
}

func exportSite(root string) {
	t0 := time.Now()

	fmt.Println("exporting")

	err := export.Export(root)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("done in", time.Since(t0))
}

func initRoot(root string) {
	t0 := time.Now()

	fmt.Println("Init", root)

	content.InitializeContentTemplate(root)
	fmt.Println("Done in", time.Since(t0))
}
