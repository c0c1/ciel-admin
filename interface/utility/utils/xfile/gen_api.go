package xfile

import (
	"fmt"
	"github.com/gogf/gf/v2/os/gfile"
	"os"
	"strings"
)

var apiTemplate = fmt.Sprintf(`
type $Name$Req struct {
    Page     int %vd:"1"%v
    PageSize int %vd:"10"%v
    entity.$Name$
}
`, "`", "`", "`", "`")

func genApi(c TemplateConfig) {
	file, err := gfile.OpenFile(c.RootPath+c.PathApi, os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	apiTemplate = strings.ReplaceAll(apiTemplate, "$Name$", c.EntityName)
	if _, err = file.WriteString(apiTemplate); err != nil {
		panic(err)
	}
	if err = file.Close(); err != nil {
		panic(err)
	}
}
