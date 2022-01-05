package xfile

type TemplateConfig struct {
	RootPath      string
	FileName      string // u_login_log
	EntityName    string // LoginLog
	PathService   string // /home/howl/learn/ciel-admin/interface/internal/service
	PathApi       string // /home/howl/learn/ciel-admin/interface/apiv1
	PathCmd       string // /home/howl/learn/ciel-admin/interface/internal/cmd/admin.go
	PathDataFetch string // /home/howl/learn/ciel-admin/admin/data
	PathIndex     string // /home/howl/learn/ciel-admin/admin/pages
	group         string //  user
	P1            string
	P1Name        string
	P2            string
	P2Name        string
	P3            string
	P3Name        string
	NameZh        string
}

func Gen(c TemplateConfig) {
	genService(c)
	genApi(c)
	genHandler(c)
	genCmd(c)
	genDataFetch(c)
	genIndex(c)
}
