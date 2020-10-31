package main

import (
	"github.com/pubgo/xerror"
	"github.com/pubgo/xprotogen/gen"
	"log"
)

func main() {
	defer xerror.RespDebug()

	m := gen.New("micro")
	m.Parameter(func(key, value string) {
		log.Println("params:", key, "=", value)
	})

	m.Header(Header)
	m.Service(Service)
	xerror.Panic(m.Save())
}
