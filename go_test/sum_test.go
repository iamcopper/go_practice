package sum

import "testing"

func TestSum1(t *testing.T) {
	if Sum(1, 2) != 3 {
		t.Error(`Sum(1, 2) != 3`)
	}
}

func TestSum2(t *testing.T) {
	if Sum(50, -50) != 0 {
		t.Error(`Sum(50, -50) != 0`)
	}
}
