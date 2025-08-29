package static

import (
	"net/http"

	"github.com/e-kucheriavyi/gossrng/configs"
)

func ServeStatic() {
	fs := http.FileServer(http.Dir(configs.StaticDirectory))

	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fsAssets := http.FileServer(http.Dir(configs.AssetsDirectory))

	http.Handle("/assets/", http.StripPrefix("/assets/", fsAssets))
}
