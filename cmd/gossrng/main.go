package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/e-kucheriavyi/gossrng/configs"
	"github.com/e-kucheriavyi/gossrng/pkg/articles"
	"github.com/e-kucheriavyi/gossrng/pkg/page"
	"github.com/e-kucheriavyi/gossrng/pkg/sitemap"
	"github.com/e-kucheriavyi/gossrng/pkg/static"
)

func main() {
	fmt.Println("Serving on", configs.Host)

	sitemap.ServeSitemap()
	static.ServeStatic()
	page.ServePages()
	articles.ServeArticlesList()

	err := http.ListenAndServe(configs.Host, nil)

	if err != nil {
		log.Fatal(err.Error())
	}
}
