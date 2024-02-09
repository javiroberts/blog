package markdown

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type MDFile struct {
	FrontMatter string
	Slug        string
	Body        string
}

func (f *MDFile) Load(path string) {
	content, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	r := regexp.MustCompile(`---((?s).*)---`)
	f.FrontMatter = r.FindStringSubmatch(string(content))[0]
	f.Body = r.ReplaceAllString(string(content), "")
	f.Slug = path[strings.LastIndex(path, "/")+1:]

	ext := filepath.Ext(f.Slug)
	f.Slug = f.Slug[0 : len(f.Slug)-len(ext)]
}
