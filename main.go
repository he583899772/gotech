package main

import (
	_ "gotech/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"gotech/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
