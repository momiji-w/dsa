package factorial

import "testing"

func TestFactorial(t *testing.T) {
	if res := Factorial(0); res != 1 {
		t.Fatalf("Expecting F(0) == 1, got %d", res)
	}
	if res := Factorial(5); res != 120 {
		t.Fatalf("Expecting F(5) == 120, got %d", res)
	}
}
