package xfile

import (
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"os"
	"strings"
)

var dataFetchTemplate = `
//$name$
export const $name$Fetcher = ({url, params}) => axios.get(url, {params: params, headers: {'token': token()}}).then(res => res.data)
export const add$Name$ = (data) => axios.post(Server+"/$name$/add", data, {headers: {'token': token()}}).then(res => res.data)
export const del$Name$ = (id) => axios.delete(Server+"/$name$/del?id="+id, {headers: {'token': token()}}).then(res => res.data)
export const update$Name$ = (data) => axios.put(Server+"/$name$/update", data, {headers: {'token': token()}}).then(res => res.data)
export const get$Name$ById = (id) => axios.get(Server+"/$name$/getById?id="+id, {headers: {'token': token()}}).then(res => res.data)
`

func genDataFetch(c TemplateConfig) {
	file, err := gfile.OpenFile(c.RootPath+c.PathDataFetch, os.O_WRONLY|os.O_APPEND, 0600)
	template := dataFetchTemplate
	template = strings.ReplaceAll(template, "$Name$", c.EntityName)
	template = strings.ReplaceAll(template, "$name$", gstr.CaseCamelLower(c.EntityName))
	if _, err = file.WriteString(template); err != nil {
		panic(err)
	}
	if err = file.Close(); err != nil {
		panic(err)
	}
}
