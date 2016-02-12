package kmgFile

import (
	"testing"

	. "github.com/bronze1man/kmg/kmgTest"
)

func TestKmgFile(ot *testing.T) {
	err := WriteFile(".kmgFileTest", []byte(""))
	Equal(err, nil)
	MustDeleteFile(".kmgFileTest")
	MustDeleteFile(".kmgFileTest")
	MustDeleteFileOrDirectory(".kmgFileTest")
}
