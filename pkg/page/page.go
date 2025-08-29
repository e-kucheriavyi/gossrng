package page

import (
	"cserver/configs"
	"cserver/pkg/mdparcer"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type PageFile struct {
	Content []byte
	Meta    map[string]string
}

type PageResponse struct {
	Content []byte
	Code    int
}

func ServePages() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		page, err := PreparePage(r.URL.Path)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		w.WriteHeader(page.Code)
		w.Header().Add("Content-Type", "text/html")
		w.Write(page.Content)
	})
}

func PreparePage(path string) (PageResponse, error) {
	f, err := ReadPageFile(path)

	if err != nil {
		if os.IsNotExist(err) {
			return PreparePage("/404")
		}
		return PreparePage("/500")
	}

	html := mdparcer.MdToHTML(f.Content)

	template, err := ReadTemplate()

	if err != nil {
		fmt.Println(err.Error())
		return PreparePage("/500")
	}

	text := template

	for key, value := range f.Meta {
		text = strings.Replace(text, "%"+key+"%", value, 1)
	}

	text = strings.Replace(text, "%CONTENT%", string(html), 1)

	code := 200

	switch path {
	case "/404":
		code = 404
	case "/500":
		code = 500
	case "/403":
		code = 403
	}

	return PageResponse{
		Content: []byte(text),
		Code:    code,
	}, nil
}

func ReadPageFile(urlPath string) (PageFile, error) {
	var filePath string

	isDir := urlPath[len(urlPath)-1] == '/'

	if isDir {
		filePath = configs.ContentDirectory + urlPath + "index.md"
	} else {
		filePath = configs.ContentDirectory + urlPath + ".md"
	}

	f, err := os.ReadFile(filePath)

	if err != nil {
		if os.IsNotExist(err) {
			if !isDir {
				return ReadPageFile(urlPath + "/")
			}
			return PageFile{}, err
		}
		return PageFile{}, err
	}

	return ParsePageInfo(f)
}

func ParsePageInfo(f []byte) (PageFile, error) {
	meta := map[string]string{
		"title":       configs.FallbackTitle,
		"description": configs.FallbackDescription,
		"keywords":    configs.FallbackKeywords,
		"date":        "",
	}

	text := string(f)

	var contentStr string

	closed := false

	for i, v := range strings.Split(text, "\n") {
		if i == 0 && v != "---" {
			return PageFile{
				Content: f,
				Meta:    meta,
			}, nil
		}

		if i == 0 {
			continue
		}

		if i != 0 && v == "---" {
			closed = true
			continue
		}

		if closed {
			contentStr = contentStr + "\n" + v
			continue
		}

		s := strings.Split(v, ":")

		if len(s) < 2 {
			return PageFile{}, errors.New("meta string should have key and value")
		}

		if len(s) == 2 {
			meta[s[0]] = strings.Trim(s[1], " ")
			continue
		}

		meta[s[0]] = strings.Trim(strings.Join(s[1:], ":"), " ")
	}

	return PageFile{
		Content: []byte(contentStr),
		Meta:    meta,
	}, nil
}

func ReadTemplate() (string, error) {
	f, err := os.ReadFile(configs.TemplatesDirectory + "/content.html")

	if err != nil {
		return "", err
	}

	return string(f), nil
}
