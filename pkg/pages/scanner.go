package pages

import (
	"os"
	"slices"
	"strings"

	"github.com/e-kucheriavyi/gossrng/configs"
)

func IsSkipped(name string) bool {
	skiplist := []string{"404.md", "403.md", "500.md"}

	return slices.Contains(skiplist, name)
}

func PreparePagesList() ([]Page, error) {
	paths, err := ScanAllFilepaths(configs.ContentDirectory)

	if err != nil {
		return []Page{}, err
	}

	pages := make([]Page, 0, len(paths))

	for _, path := range paths {
		p, err := ReadPageFile(path)

		if err != nil {
			return []Page{}, err
		}

		pages = append(pages, p)
	}

	return pages, nil
}

func FilterUtilityPages(pages []Page) []Page {
	filtered := make([]Page, 0, len(pages))
	for _, v := range pages {
		s := strings.Split(v.Filepath, "/")

		if IsSkipped(s[len(s)-1]) {
			continue
		}
		filtered = append(filtered, v)
	}

	return filtered
}

func ReadPageFile(path string) (Page, error) {
	f, err := os.ReadFile(path)

	if err != nil {
		return Page{}, err
	}

	page, err := ParsePageInfo(f)

	if err != nil {
		return page, err
	}

	page.Filepath = path
	page.Route = FormatFilepathToRoute(configs.ContentDirectory, path)

	return page, nil
}

func ScanAllFilepaths(root string) ([]string, error) {
	paths := []string{}

	entries, err := os.ReadDir(root)

	if err != nil {
		return []string{}, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			if entry.Name()[0] == '.' {
				continue
			}

			var p, err = ScanAllFilepaths(root + "/" + entry.Name())

			if err != nil {
				return []string{}, err
			}

			paths = append(paths, p...)
			continue
		}

		s := strings.Split(entry.Name(), ".")

		if s[len(s)-1] != "md" {
			continue
		}

		paths = append(paths, root+"/"+entry.Name())
	}

	return paths, nil
}
