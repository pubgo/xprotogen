package main

import (
	"log"

	"github.com/dave/jennifer/jen"
	"github.com/pubgo/xerror"
	"github.com/pubgo/xprotogen"
)

const (
	rpcPkgPath = "net/rpc"
)

func main() {
	defer xerror.Resp(func(err xerror.XErr) {
		log.Println(err.Println())
	})

	hello := xprotogen.New("hello")
	hello.Parameter(func(key, value string) {
		log.Println("params:", key, "=", value)
	})

	xerror.Panic(hello.Service(func(ss *xprotogen.Service) {
		j := ss.J
		srv := ss.Name

		j.ImportAlias(rpcPkgPath, "rpc")

		j.Comment("// test")
		j.Comment("/* ssss */")
		j.Type().Id(srv).InterfaceFunc(func(group *jen.Group) {
			for _, m := range ss.GetMethod() {
				mthName := xprotogen.CamelCase(m.GetName())

				mthOpt, err := xprotogen.ExtractAPIOptions(m)
				xerror.Panic(err)

				method, path := xprotogen.ExtractHttpMethod(mthOpt)
				log.Println(srv, mthName, method, path)

				group.Id(mthName).Params(
					jen.Id("in *"+ss.TypeName(m.GetInputType())),
					jen.Id("out *"+ss.TypeName(m.GetOutputType())),
				).Error()
			}
		})

		// method
		j.Func().Id("Register"+srv).Params(
			jen.Id("srv *").Qual(rpcPkgPath, "Server"),
			jen.Id("x "+srv),
		).Error().BlockFunc(func(group *jen.Group) {
			group.Return().Nil()
		})
	}))

	xerror.Panic(hello.Save())
}
