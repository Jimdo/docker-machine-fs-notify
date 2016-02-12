package kmgFile

import (
	"path/filepath"
	"strings"
)

// 返回文件名部分但是不带扩展名
// 如果没有扩展名就不变了.
// 有多个扩展名,就删除一个.
func PathBaseWithoutExt(path string) string {
	ext := filepath.Ext(path)
	return strings.TrimSuffix(filepath.Base(path), ext)
}

func PathTrimExt(path string) string {
	ext := filepath.Ext(path)
	return strings.TrimSuffix(path, ext)
}

type PathAndContentPair struct {
	Path    string
	Content []byte
}
