package kmgFile

import (
	"path/filepath"
)

//使用workingPath把path更新为全路径
func FullPathOnPath(workingPath string, path string) string {
	if filepath.IsAbs(path) {
		return path
	}
	return filepath.Join(workingPath, path)
}
