package path

import (
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

const rootName = "api"
const dockerRootName = "app"

func GetProjectRootPath() string {
	_, b, _, _ := runtime.Caller(0)
	relativePath := path.Join(path.Dir(b))
	for !strings.HasSuffix(relativePath, rootName) &&
		!strings.HasSuffix(relativePath, dockerRootName) {
		relativePath = filepath.Dir(relativePath)
	}

	return relativePath
}
