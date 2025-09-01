package pages

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/e-kucheriavyi/gossrng/configs"
)

func ServePages() {
	root := configs.ContentDirectory

	http.HandleFunc("/articles", func(w http.ResponseWriter, r *http.Request) {
		pages, err := PreparePagesList()

		if err != nil {
			w.WriteHeader(500)
			return
		}

		links := []string{}

		for _, v := range FilterUtilityPages(pages) {
			links = append(links, fmt.Sprintf("<a href='%s'>%s</a>", v.Route, v.Meta["title"]))
		}

		f := Page{
			Content: []byte(strings.Join(links, "<br>")),
			Meta: NewMetaMap(map[string]string{
				"title": "Список статей",
			}),
		}

		tmp, err := ReadTemplateFile()

		if err != nil {
			w.WriteHeader(500)
			return
		}

		result := FormatTemplate(tmp, f)

		w.WriteHeader(200)
		w.Header().Add("Content-Type", "text/html")
		w.Write([]byte(result))
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		paths, err := ScanAllFilepaths(root)

		if err != nil {
			w.WriteHeader(500)
			return
		}

		var filepath string

		for _, v := range paths {
			if FormatFilepathToRoute(root, v) == r.URL.Path {
				filepath = v
				break
			}
		}

		if filepath == "" {
			w.WriteHeader(404)
			return
		}

		page, err := ReadPageFile(filepath)

		if err != nil {
			w.WriteHeader(500)
			return
		}

		tmp, err := ReadTemplateFile()

		if err != nil {
			w.WriteHeader(500)
			return
		}

		result := FormatTemplate(tmp, page)

		w.WriteHeader(200)
		w.Header().Add("Content-Type", "text/html")
		w.Write([]byte(result))
	})
}
