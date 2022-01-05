package xfile

type TemplateConfig struct {
	RootPath      string // 项目绝对路径地址 eg :/home/howl/learn/ciel-admin
	FileName      string // 文件名称 u_login_log
	EntityName    string // 实体名称 LoginLog
	PathService   string // 服务层文件夹 eg: /interface/apiv1/user_api.go
	PathApi       string // 请求实体文件夹  eg:/home/howl/learn/ciel-admin/interface/apiv1
	PathCmd       string // Cmd文件夹 /interface/internal/cmd/admin.go
	PathDataFetch string // js axios 请求文件(没有请先创建) /admin/data/user.js
	PathIndex     string // js 页面 /admin/pages/user
	group         string // axios 分组名称 user
	P1            string // 主要字段1 field
	P1Name        string // 主要字段1 名称
	P2            string // 主要字段2 field
	P2Name        string // 主要字段2 名称
	P3            string // 主要字段3 field
	P3Name        string // 主要字段3 名称
	NameZh        string // 页面标题
}

func Gen(c TemplateConfig) {
	genService(c)
	genApi(c)
	genHandler(c)
	genCmd(c)
	genDataFetch(c)
	genIndex(c)
}
