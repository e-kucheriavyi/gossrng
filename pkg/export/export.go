package export

import (
	"fmt"
	"os"
	"strings"

	"github.com/e-kucheriavyi/gossrng/pkg/pages"
)

const distPath = "./dist"

func Export(root string) error {
	err := os.Mkdir(distPath, 0755)

	if err != nil {
		if os.IsExist(err) {
			os.RemoveAll(distPath)
			return Export(root)
		}
		return err
	}

	err = exportStatic(root)

	if err != nil {
		return err
	}

	err = exportPages(root)

	if err != nil {
		return err
	}

	return nil
}

func exportStatic(root string) error {
	fmt.Println("exporting static files...")

	publicFs := os.DirFS(root + "/public")
	err := os.CopyFS(distPath+"/", publicFs)

	if err != nil {
		fmt.Println("Error while exporting `/public` directory:", err.Error())
	}

	assetsFs := os.DirFS(root + "/assets")
	err = os.CopyFS(distPath+"/assets", assetsFs)

	if err != nil {
		return err
	}

	err = exportPagesList(root)

	if err != nil {
		return err
	}

	return nil
}

func exportPages(root string) error {
	fmt.Println("exporting pages...")

	paths, err := pages.ScanAllFilepaths(root)

	if err != nil {
		return err
	}

	tmp, err := pages.ReadTemplateFile(root)

	if err != nil {
		return err
	}

	for _, v := range paths {
		page, err := pages.ReadPageFile(root, v)

		if err != nil {
			return err
		}

		formatted := pages.FormatTemplate(tmp, page)

		newPath := strings.Replace(page.Filepath, root, distPath, 1)
		newPath = strings.Replace(newPath, ".md", ".html", 1)

		s := strings.Split(newPath, "/")

		dirPath := strings.Join(s[0:len(s)-1], "/")

		os.MkdirAll(dirPath, 0777)

		err = os.WriteFile(newPath, []byte(formatted), 0666)

		if err != nil {
			return err
		}
	}

	return nil
}

func exportPagesList(root string) error {
	fmt.Println("exporting pages list...")

	result, err := pages.FormatPageList(root)

	if err != nil {
		return err
	}

	os.Mkdir(distPath+"/articles", 0755)

	return os.WriteFile(distPath+"/articles/index.html", []byte(result), 0666)
}
