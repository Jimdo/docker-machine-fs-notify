package kmgRand

import (
	"testing"

	"github.com/bronze1man/kmg/kmgTest"
)

func TestMustCryptoRand(ot *testing.T) {
	for _, f := range []func(length int) string{
		MustCryptoRandToHex,
		MustCryptoRandToReadableAlphaNum,
		MustCryptoRandToAlphaNum,
	} {
		ret := f(15)
		kmgTest.Equal(len(ret), 15)

		ret = f(1)
		kmgTest.Equal(len(ret), 1)

		ret = f(20)
		kmgTest.Equal(len(ret), 20)
	}
}
