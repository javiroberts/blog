package markdown

import (
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"html/template"
)

//func LoadMarkdownPost(slug string) template.HTML {
//	content, err := os.ReadFile(fmt.Sprintf("public/markdown/%s.md", slug))
//	if err != nil {
//		panic(err)
//	}
//	return template.HTML(mdToHTML(content))
//}

func LoadMDArticle(content string) template.HTML {
	//content, err := os.ReadFile(fmt.Sprintf("public/markdown/%s.md", slug))
	//if err != nil {
	//	panic(err)
	//}
	return template.HTML(mdToHTML([]byte(content)))
}

func mdToHTML(md []byte) []byte {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock | parser.Titleblock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}
