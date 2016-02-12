package kmgFile

import (
	"testing"

	"github.com/bronze1man/kmg/kmgTest"
)

func TestGetAllFileAndDirectoryStat(t *testing.T) {
	MustMkdirAll("testFile/d1/d2")
	MustWriteFile("testFile/d1/d2/f3", []byte("1"))
	out, err := GetAllFileAndDirectoryStat("testFile")
	kmgTest.Equal(err, nil)
	kmgTest.Equal(len(out), 4)

	out, err = GetAllFileAndDirectoryStat("testFile/d1/d2/f3")
	kmgTest.Equal(err, nil)
	kmgTest.Equal(len(out), 1)
}
