package constants

import (
	"os"
	"path"
)

var Pwd, _ = os.Getwd()
var DumpsDir = path.Join(Pwd, "dumps")
