package content

import (
	"fmt"
	"os"

	"github.com/e-kucheriavyi/gossrng/configs"
)

func ReadArticle(urlPath string) ([]byte, error) {
	var filePath string

	isDir := urlPath[len(urlPath)-1] == '/'

	if isDir {
		filePath = configs.ContentDirectory + urlPath + "index.md"
	} else {
		filePath = configs.ContentDirectory + urlPath + ".md"
	}

	fmt.Println(filePath)

	f, err := os.ReadFile(filePath)

	if err != nil {
		if os.IsNotExist(err) {
			if !isDir {
				return ReadArticle(urlPath + "/")
			}
			return []byte{}, err
		}
		fmt.Println(err.Error())
		return []byte{}, err
	}

	return f, nil
}
