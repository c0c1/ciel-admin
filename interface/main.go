package main

import (
	_ "interface/boot"
	_ "interface/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
