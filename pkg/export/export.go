package export

import (
	"fmt"
	"os"
	"strings"

	"github.com/e-kucheriavyi/gossrng/configs"
	"github.com/e-kucheriavyi/gossrng/pkg/pages"
)

const distPath = "./dist"

func Export() error {
	err := os.Mkdir(distPath, 0755)

	if err != nil {
		if os.IsExist(err) {
			os.RemoveAll(distPath)
			return Export()
		}
		return err
	}

	err = exportStatic()

	if err != nil {
		return err
	}

	err = exportPages()

	if err != nil {
		return err
	}

	return nil
}

func exportStatic() error {
	fmt.Println("exporting static files...")

	rootFs := os.DirFS(configs.RootStaticDirectory)
	err := os.CopyFS(distPath+"/", rootFs)

	if err != nil {
		return err
	}

	staticFs := os.DirFS(configs.StaticDirectory)
	err = os.CopyFS(distPath+"/static", staticFs)

	if err != nil {
		return err
	}

	assetsFs := os.DirFS(configs.AssetsDirectory)
	err = os.CopyFS(distPath+"/assets", assetsFs)

	if err != nil {
		return err
	}

	err = exportPagesList()

	if err != nil {
		return err
	}

	return nil
}

func exportPages() error {
	root := configs.ContentDirectory

	fmt.Println("exporting pages...")

	paths, err := pages.ScanAllFilepaths(root)

	if err != nil {
		return err
	}

	tmp, err := pages.ReadTemplateFile()

	if err != nil {
		return err
	}

	for _, v := range paths {
		page, err := pages.ReadPageFile(v)

		if err != nil {
			return err
		}

		formatted := pages.FormatTemplate(tmp, page)

		newPath := strings.Replace(page.Filepath, configs.ContentDirectory, distPath, 1)
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

func exportPagesList() error {
	fmt.Println("exporting pages list...")

	result, err := pages.FormatPageList()

	if err != nil {
		return err
	}

	os.Mkdir(distPath+"/articles", 0755)

	return os.WriteFile(distPath+"/articles/index.html", []byte(result), 0666)
}
