package pages

import (
	"blog/markdown"
	"html/template"

	yaml2 "gopkg.in/yaml.v3"
)

type Page struct {
	Title   string `yaml:"title"`
	Slug    string
	Article template.HTML
}

func (p *Page) WithMarkdown(file markdown.MDFile) error {
	err := yaml2.Unmarshal([]byte(file.FrontMatter), &p)
	if err != nil {
		panic(err.Error())
	}
	p.Slug = file.Slug
	p.Article = markdown.LoadMDArticle(file.Body)

	return nil
}
