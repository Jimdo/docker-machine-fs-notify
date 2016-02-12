package kmgFile

import "github.com/bronze1man/kmg/kmgRand"

// 返回一个新的临时文件目录,保证父级目录存在,保证文件不存在.
func NewTmpFilePath() string {
	return "/tmp/kmg_" + kmgRand.MustCryptoRandToReadableAlphaNum(8)
}
