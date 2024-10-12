package posts

import (
	"blog/markdown"
	"html/template"
	"time"

	yaml2 "gopkg.in/yaml.v3"
)

type Post struct {
	Title      string    `yaml:"title"`
	Date       time.Time `yaml:"date"`
	Author     string    `yaml:"author"`
	Tags       []string  `yaml:"tags"`
	Categories []string  `yaml:"categories"`
	Draft      bool      `yaml:"draft"`
	Summary    string    `yaml:"summary"`
	Slug       string
	Article    template.HTML
}

func (p *Post) WithMarkdown(file markdown.MDFile) {
	err := yaml2.Unmarshal([]byte(file.FrontMatter), &p)
	if err != nil {
		panic(err.Error())
	}
	p.Slug = file.Slug
	p.Article = markdown.LoadMDArticle(file.Body)
}
