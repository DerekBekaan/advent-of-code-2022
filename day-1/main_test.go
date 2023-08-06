package main

import (
	"fmt"
	"testing"
)

func TestMaxCaloires(t *testing.T) {
	calories := `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`
	expected := 45000

	actual := maxCalories(calories)

	if actual != expected {
		t.Fatalf(fmt.Sprintf("Expected %d got %d", expected, actual))
	}
}
