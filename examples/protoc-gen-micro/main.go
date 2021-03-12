package main

import (
	"log"

	"github.com/pubgo/xerror"
	"github.com/pubgo/xprotogen/gen"
)

func main() {
	defer xerror.RespDebug()

	m := gen.New("micro")
	m.Parameter(func(key, value string) {
		log.Println("params:", key, "=", value)
	})

	xerror.Panic(m.GenWithTpl())
}
