package templates

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"path/filepath"
)

func InitTemplate(router *gin.Engine) {
	router.Static("/static", "./public/assets")
	router.StaticFile("/favicon.ico", "./public/favicon.ico")
	router.HTMLRender = loadTemplates("./templates")
}

//多模板（模板继承）
func loadTemplates(templatesDir string) multitemplate.Renderer {
	renderer := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/layouts/*.html")
	if err != nil {
		panic(err.Error())
	}

	pages, pagesError := filepath.Glob(templatesDir + "/pages/*.html")
	if pagesError != nil {
		panic(pagesError.Error())
	}

	for _, page := range pages {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, page)
		renderer.AddFromFilesFuncs(filepath.Base(page), FuncMap(), files...)
	}
	return renderer
}