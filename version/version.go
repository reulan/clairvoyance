package version

import (
	"fmt"
	"runtime"
)

const Version = "0.2.5"

var GitCommit string // GitCommit returns the git commit that was compiled. This will be filled in by the compiler.
var BuildDate = ""
var GoVersion = runtime.Version()                               // GoVersion returns the version of the go runtime used to compile the binary
var OsArch = fmt.Sprintf("%s %s", runtime.GOOS, runtime.GOARCH) // OsArch returns the os and arch used to build the binary
