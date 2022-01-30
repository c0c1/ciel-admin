package main

import (
	_ "interface/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"
	"interface/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
