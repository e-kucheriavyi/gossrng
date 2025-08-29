package main

import (
	"cserver/configs"
	"cserver/pkg/page"
	"cserver/pkg/sitemap"
	"cserver/pkg/static"
	"fmt"
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
