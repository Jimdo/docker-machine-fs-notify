package kmgFile

import (
	"os"
	"path/filepath"

	"github.com/bronze1man/kmg/errors"
)

var NotFoundError = errors.New("Not found")

//向上搜索某个目录下面的某个文件
// 比如搜索 .git 目录 搜索 .kmg.yml 文件
// 返回 .git 所在的上级目录.
func SearchFileInParentDir(startDirPath string, fileName string) (file string, err error) {
	startDirPath, err = filepath.Abs(startDirPath)
	if err != nil {
		return
	}
	p := startDirPath
	var kmgFilePath string
	for {
		kmgFilePath = filepath.Join(p, fileName)
		_, err = os.Stat(kmgFilePath)
		if err == nil {
			//found it
			return p, nil
		}
		if !os.IsNotExist(err) {
			return
		}
		thisP := filepath.Dir(p)
		if p == thisP {
			//到底
			return "", NotFoundError
		}
		p = thisP
	}
}
