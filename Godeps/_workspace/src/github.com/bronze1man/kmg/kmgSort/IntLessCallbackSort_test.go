package kmgSort
import (
	"testing"
	"github.com/bronze1man/kmg/kmgTest"
)

func TestIntLessCallbackSort(t *testing.T) {
	a := []int{2, 1, 3, 0}
	IntLessCallbackSort(a, func(i int, j int) bool {
		return a[i] < a[j]
	})
	kmgTest.Equal(a, []int{0, 1, 2, 3})
}
