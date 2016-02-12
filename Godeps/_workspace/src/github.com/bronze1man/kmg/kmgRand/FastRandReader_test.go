package kmgRand

import (
	"testing"
	"github.com/bronze1man/kmg/kmgTest"
)

func TestFastRandReader(ot *testing.T) {
	buf := make([]byte, 10*1024)
	n, err := FastRandReader.Read(buf)
	kmgTest.Equal(n, 10*1024)
	kmgTest.Equal(err, nil)
}

func BenchmarkFastRandReader(ot *testing.B) {
	ot.StopTimer()
	buf := make([]byte, ot.N)
	ot.SetBytes(int64(1))
	ot.StartTimer()
	n, err := FastRandReader.Read(buf)
	kmgTest.Equal(n, ot.N)
	kmgTest.Equal(err, nil)
}
