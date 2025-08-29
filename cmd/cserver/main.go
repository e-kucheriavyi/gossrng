package main

import (
	"fmt"
	"gossrng/configs"
	"gossrng/pkg/page"
	"gossrng/pkg/sitemap"
	"gossrng/pkg/static"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Serving on", configs.Host)

	sitemap.ServeSitemap()
	static.ServeStatic()

	page.ServePages()

	err := http.ListenAndServe(configs.Host, nil)

	if err != nil {
		log.Fatal(err.Error())
	}
}
