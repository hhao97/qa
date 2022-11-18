package main

import (
	_ "qa/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"qa/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
