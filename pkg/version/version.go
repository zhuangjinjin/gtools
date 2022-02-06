package version

import (
	"encoding/json"
	"fmt"
	"runtime"
)

type VersionInfo struct {
	GoVersion string `json:"goVersion"`
	Compiler string `json:"compiler"`
	Platform string `json:"platform"`
}

func Get() *VersionInfo {
	return &VersionInfo{
		GoVersion: runtime.Version(),
		Compiler: runtime.Compiler,
		Platform:   fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}

func (self *VersionInfo) ToString() string {
	b, _ := json.Marshal(self)
	return string(b)
}