package system

import (
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

func StringToMarkDown(input []byte) []byte {
	return blackfriday.MarkdownBasic(input)
}

func MarkDownToHTML(markdown string) string {
	HTMLFlags := 0 |
		blackfriday.HTML_USE_XHTML |
		blackfriday.HTML_USE_SMARTYPANTS |
		blackfriday.HTML_SMARTYPANTS_FRACTIONS |
		blackfriday.HTML_SMARTYPANTS_DASHES |
		blackfriday.HTML_SMARTYPANTS_LATEX_DASHES

	Extensions := 0 |
		blackfriday.EXTENSION_NO_INTRA_EMPHASIS |
		blackfriday.EXTENSION_TABLES |
		blackfriday.EXTENSION_FENCED_CODE |
		blackfriday.EXTENSION_AUTOLINK |
		blackfriday.EXTENSION_STRIKETHROUGH |
		blackfriday.EXTENSION_SPACE_HEADERS |
		blackfriday.EXTENSION_HEADER_IDS |
		blackfriday.EXTENSION_BACKSLASH_LINE_BREAK |
		blackfriday.EXTENSION_DEFINITION_LISTS |
		blackfriday.EXTENSION_HARD_LINE_BREAK

	renderer := blackfriday.HtmlRenderer(HTMLFlags, "", "")
	bytes := blackfriday.MarkdownOptions([]byte(markdown), renderer, blackfriday.Options{
		Extensions: Extensions,
	})
	HTML := string(bytes)
	return bluemonday.UGCPolicy().Sanitize(HTML)
}