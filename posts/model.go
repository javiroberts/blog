package posts

import (
	"backend/markdown"
	"fmt"
	yaml2 "gopkg.in/yaml.v3"
	"html/template"
	"os"
	"path/filepath"
	"regexp"
)

type Post struct {
	Title      string   `yaml:"title"`
	Date       string   `yaml:"date"`
	Author     string   `yaml:"author"`
	Tags       []string `yaml:"tags"`
	Categories []string `yaml:"categories"`
	Draft      bool     `yaml:"draft"`
	Summary    string   `yaml:"summary"`
	Slug       string
	Article    template.HTML
}

func (p *Post) Load(slug string) {
	content, err := os.ReadFile(fmt.Sprintf("public/markdown/posts/%s", slug))
	if err != nil {
		panic(err)
	}

	r := regexp.MustCompile(`---((?s).*)---`)
	yaml := r.FindStringSubmatch(string(content))
	article := r.ReplaceAllString(string(content), "")

	err = yaml2.Unmarshal([]byte(yaml[0]), &p)

	ext := filepath.Ext(slug)
	p.Slug = slug[0 : len(slug)-len(ext)]

	p.Article = markdown.LoadMDArticle(article)
}
