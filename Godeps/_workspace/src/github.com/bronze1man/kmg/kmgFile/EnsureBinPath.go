package kmgFile

import (
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/bronze1man/kmg/kmgCmd"
	"github.com/bronze1man/kmg/kmgTime"
)

func MustEnsureBinPath(finalPath string) {
	basePath := filepath.Base(finalPath)
	path, err := exec.LookPath(basePath)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		panic(err)
	}
	if path != finalPath {
		backPathDir := "/var/backup/bin/" + basePath + time.Now().Format(kmgTime.FormatFileName)
		MustMkdirAll(backPathDir)
		kmgCmd.MustRun("mv " + path + " " + backPathDir)
	}
}
