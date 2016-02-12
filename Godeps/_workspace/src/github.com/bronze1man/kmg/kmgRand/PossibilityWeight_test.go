package kmgRand
import (
	"testing"
	"github.com/bronze1man/kmg/kmgTest"
)

func TestPossibilityWeight(t *testing.T) {
	r := NewInt64SeedKmgRand(0)
	rander := NewPossibilityWeightRander([]float64{1e-20})
	for i := 0; i < 100; i++ {
		kmgTest.Equal(rander.ChoiceOne(r), 0)
	}
	rander = NewPossibilityWeightRander([]float64{1, 2, 3, 4})
	for i := 0; i < 100; i++ {
		ret := rander.ChoiceOne(r)
		kmgTest.Ok(ret >= 0)
		kmgTest.Ok(ret <= 3)
	}
	rander = NewPossibilityWeightRander([]float64{1, 0, 3, 0, 1})
	for i := 0; i < 100; i++ {
		ret := rander.ChoiceOne(r)
		kmgTest.Ok(ret >= 0)
		kmgTest.Ok(ret <= 4)
	}
}
