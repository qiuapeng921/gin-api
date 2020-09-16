package templates

import (
	"gin-api/app/utility/system"
	"html/template"
)

//自定义函数
func FuncMap() template.FuncMap {
	return template.FuncMap{
		"formatDate":     formatDate,
		"markdownToHtml": markdownToHtml,
	}
}

// 时间戳转时间格式
func formatDate(timestamp int64) string {
	return system.FormatDate(timestamp)
}

// markdown转html
func markdownToHtml(markdown string) string {
	return system.MarkDownToHTML(markdown)
}