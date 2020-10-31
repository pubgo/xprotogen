package version

import (
	"runtime"
)

var BuildTime = ""
var Version = ""
var GoVersion = runtime.Version()
var GoPath = ""
var GoROOT = ""
var CommitID = ""
var Project = ""

func init() {
	if Version == "" {
		Version = "v0.0.1"
	}
}
