package helloworld

import rpc "net/rpc"

// test
/* ssss */
type TestApi interface {
	Version(in *TestReq, out *TestApiOutput) error
	VersionTest(in *TestReq, out *TestApiOutput) error
}

func RegisterTestApi(srv *rpc.Server, x TestApi) error {
	return nil
}

// test
/* ssss */
type TestApiV2 interface {
	Version(in *TestReq, out *TestApiOutput) error
	VersionTest(in *TestReq, out *TestApiOutput) error
}

func RegisterTestApiV2(srv *rpc.Server, x TestApiV2) error {
	return nil
}
