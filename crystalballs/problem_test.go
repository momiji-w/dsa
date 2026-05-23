package crystalballs

import "testing"

func TestCrystalBall(t *testing.T) {
	testArr := make([]bool, 100)
	breakPoint := 79
	for i := range len(testArr) - breakPoint {
		testArr[breakPoint+i] = true
	}

	if v := TwoCrystalBalls(testArr); v != breakPoint {
		t.Fatalf("Expecting %d, got %d", breakPoint, v)
	}
}
