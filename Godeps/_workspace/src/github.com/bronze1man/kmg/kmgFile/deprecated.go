package kmgFile

import (
	"io/ioutil"
	"os"
)

//如果这个目录已经创建过了,不报错
// @deprecated
// use MustMkdir
func MustMkdirAll(dirname string) {
	err := os.MkdirAll(dirname, os.FileMode(0777))
	if err != nil {
		panic(err)
	}
}

// @deprecated
// use MustReadFile instead
func MustReadFileAll(path string) (content []byte) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return content
}

// @deprecated
// use ReadFile instead
func ReadFileAll(path string) (content []byte, err error) {
	return ioutil.ReadFile(path)
}

//delete file,ignore file not exist err
// @deprecated
// use MustDelete instead
func MustDeleteFile(path string) {
	err := os.Remove(path)
	if os.IsNotExist(err) {
		return
	}
	if err != nil {
		panic(err)
	}
}

// @deprecated
// use MustDelete instead
func MustDeleteFileOrDirectory(path string) {
	err := os.RemoveAll(path)
	if os.IsNotExist(err) {
		return
	}
	if err != nil {
		panic(err)
	}
	return
}
