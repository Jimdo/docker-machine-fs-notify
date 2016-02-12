package kmgFile

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

/*
func AllDirectory(root string)(out []string,err error){
    err=filepath.Walk(root,func(path string, info os.FileInfo, err error) error {

    })
}
*/

type StatAndFullPath struct {
	Fi       os.FileInfo
	FullPath string
}

// 获取这个路径的所有文件的状态和完整路径
//   如果输入是一个文件,则返回这个文件的完整路径
//   如果输入是一个目录,则返回这个目录和下面所有目录和文件的信息和完整路径
//   目前暂不明确symlink的文件会如何处理
func GetAllFileAndDirectoryStat(root string) (out []StatAndFullPath, err error) {
	root, err = Realpath(root)
	if err != nil {
		return nil, err
	}
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		out = append(out, StatAndFullPath{
			FullPath: path,
			Fi:       info,
		})
		return nil
	})
	return
}

//返回这个目录下面所有的文件,返回格式为完整文件名
func GetAllFiles(root string) (out []string, err error) {
	root, err = Realpath(root)
	if err != nil {
		return nil, err
	}
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			out = append(out, path)
		}
		return nil
	})
	return
}

// 获得所有文件,不包括目录
func MustGetAllFiles(root string) (out []string) {
	out, err := GetAllFiles(root)
	if err != nil {
		panic(err)
	}
	return out
}

// 只返回这个目录额一层文件.
// 获得所有文件,不包括目录
// 返回绝对路径
func MustGetAllFileOneLevel(path string) (fileList []string) {
	fiList, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, fi := range fiList {
		if !fi.IsDir() {
			fileList = append(fileList, filepath.Join(path, fi.Name()))
		}
	}
	return fileList
}

// 获得所有文件,不包括目录
func MustGetAllFileFromPathList(pathlist []string) (outList []string) {
	for _, root := range pathlist {
		out, err := GetAllFiles(root)
		if err != nil {
			panic(err)
		}
		outList = append(outList, out...)
	}
	return outList
}

// 获得所有目录,不包括文件
func MustGetAllDir(root string) (out []string) {
	root, err := Realpath(root)
	if err != nil {
		panic(err)
	}
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			out = append(out, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return
}

// 返回一个目录下面的所有文件的名字
// 文件名是相对于这个目录的
// 只返回第一层,没有更多层
// 只返回文件,不返回目录
func ReadDirFileOneLevel(path string) (fileList []string, err error) {
	fiList, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	for _, fi := range fiList {
		if !fi.IsDir() {
			fileList = append(fileList, fi.Name())
		}
	}
	return fileList, nil
}

func MustReadDirFileOneLevel(path string) (fileList []string) {
	fileList, err := ReadDirFileOneLevel(path)
	if err != nil {
		panic(err)
	}
	return fileList
}

/*
// 回调只会给文件名,回调返回true表示需要这个文件,返回false表示不需要这个文件
func MustGetAllFilesWithCallback(root string,cb func(absPath string)bool) (out []string) {
	root, err := Realpath(root)
	if err != nil {
		panic(err)
	}
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir(){
			return nil
		}
		if cb(path) {
			out = append(out,path)
		}
		return nil
	})
	if err!=nil{
		panic(err)
	}
	return
}
*/
