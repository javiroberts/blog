package pages

import (
	"backend/markdown"
	"html/template"
)

type Page struct {
	Slug    string
	Article template.HTML
}

func (p *Page) WithMarkdown(file markdown.MDFile) error {
	p.Slug = file.Slug
	p.Article = markdown.LoadMDArticle(file.Body)

	return nil
}
