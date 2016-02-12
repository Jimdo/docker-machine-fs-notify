package kmgCmd

import (
	"github.com/bronze1man/kmg/kmgTest"
	"testing"
)

func TestExist(t *testing.T) {
	kmgTest.Ok(!Exist("absadfklsjdfa"))
	kmgTest.Ok(Exist("ls"))
	kmgTest.Ok(Exist("top"))
}

// 研究当前信号的处理的模式
func TestSingle(ot *testing.T) {

}
