package xfile

import (
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"os"
	"strings"
)

var cmdTemplate = `

	s.Group("/$name$", func(g *ghttp.RouterGroup) {
		g.Middleware(middleware.Auth)
		g.GET("/list", handler.$Name$.List)
		g.GET("/getById", handler.$Name$.GetById)
		g.Middleware(middleware.LockAction)
		g.POST("/add", handler.$Name$.Add)
		g.PUT("/update", handler.$Name$.Put)
		g.DELETE("/del", handler.$Name$.Del)
	})
}`

func genCmd(c TemplateConfig) {
	delLastLine(c)
	file, err := gfile.OpenFile(c.RootPath+c.PathCmd, os.O_WRONLY|os.O_APPEND, 0600)
	template := cmdTemplate
	template = strings.ReplaceAll(template, "$Name$", c.EntityName)
	template = strings.ReplaceAll(template, "$name$", gstr.CaseCamelLower(c.EntityName))
	if _, err = file.WriteString(template); err != nil {
		panic(err)
	}
	if err = file.Close(); err != nil {
		panic(err)
	}
}

func delLastLine(c TemplateConfig) {
	file, err := gfile.OpenFile(c.RootPath+c.PathCmd, os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	stat, err := file.Stat()
	if err = file.Truncate(stat.Size() - 2); err != nil {
		panic(err)
	}
	if err = file.Close(); err != nil {
		panic(err)
	}
}
