package helpers

import (
	"fmt"
	"testing"
)

func TestFailSum(t *testing.T) {
	result := Sum(10, 10)

	if result != 40 {
		t.FailNow()
	}

	fmt.Println("TestFailSum eksekusi terhenti")
}

func TestSum(t *testing.T) {
	result := Sum(10, 10)

	if result != 20 {
		panic("Result harus 20 gans")
	}

	fmt.Println("TestSum terhenti")
}
