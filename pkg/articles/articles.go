package articles

import (
	"fmt"
	"maps"
	"net/http"
	"os"
	"slices"
	"strings"

	"github.com/e-kucheriavyi/gossrng/configs"
	"github.com/e-kucheriavyi/gossrng/pkg/page"
)

func ServeArticlesList() {
	http.HandleFunc("/articles", func(w http.ResponseWriter, r *http.Request) {
		articles, err := ScanArticles("/")

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		tmp, _ := page.ReadTemplate()

		links := []string{}

		for k, v := range articles {
			links = append(links, fmt.Sprintf("<a href='%s'>%s</a>", k, v.Meta["title"]))
		}

		f := page.PageFile{
			Content: []byte(strings.Join(links, "<br>")),
			Meta: page.NewMetaMap(map[string]string{
				"title": "Список статей",
			}),
		}

		t := page.FormatTemplate(tmp, f)

		w.Header().Add("Content-Type", "text/html")
		w.Write([]byte(t))
	})
}

func ScanArticles(root string) (map[string]page.PageFile, error) {
	entries, err := os.ReadDir(configs.ContentDirectory + root)

	if err != nil {
		fmt.Println(err.Error())
		return map[string]page.PageFile{}, err
	}

	links := map[string]page.PageFile{}

	for _, v := range entries {
		s := strings.Split(v.Name(), ".")
		if s[len(s)-1] == "md" {
			if IsSkipped(v.Name()) {
				continue
			}
			if s[0] == "index" {
				links[root], _ = page.ReadPageFile(root)
				continue
			}
			links[root+s[0]], _ = page.ReadPageFile(root + s[0])
			continue
		}

		if v.IsDir() {
			l, _ := ScanArticles(root + v.Name() + "/")

			maps.Copy(links, l)
			continue
		}
	}

	return links, nil
}

func IsSkipped(name string) bool {
	skiplist := []string{"404.md", "403.md", "500.md"}

	return slices.Contains(skiplist, name)
}
