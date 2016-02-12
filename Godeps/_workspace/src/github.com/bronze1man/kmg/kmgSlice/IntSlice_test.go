package kmgSlice

import (
	"testing"

	"github.com/bronze1man/kmg/kmgTest"
)

func TestIntSliceRemoveAt(ot *testing.T) {
	s := []int{1, 2, 3}
	IntSliceRemoveAt(&s, 1)
	kmgTest.Equal(s, []int{1, 3})
}
