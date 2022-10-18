package helpers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFailSum(t *testing.T) {
	result := Sum(10, 10)

	require.Equal(t, 40, result, "Result has to be 40")

	fmt.Println("TestFailSum Eksekusi Terhenti")
}

func TestSum(t *testing.T) {
	result := Sum(10, 10)

	require.Equal(t, 20, result, "Result has to be 20")

	fmt.Println("TestSum Eksekusi Terhenti")
}

func TestTableSum(t *testing.T) {

	tests := []struct {
		request  int
		expected int
		errMsg   string
	}{
		{
			request:  Sum(10, 10),
			expected: 20,
			errMsg:   "result has to be 20",
		},
		{
			request:  Sum(20, 20),
			expected: 40,
			errMsg:   "result has to be 40",
		},
		{
			request:  Sum(25, 25),
			expected: 50,
			errMsg:   "result has to be 50",
		},
		{
			request:  Sum(30, 30),
			expected: 60,
			errMsg:   "result has to be 60",
		},
		{
			request:  Sum(35, 35),
			expected: 70,
			errMsg:   "result has to be 70",
		},
		{
			request:  Sum(40, 40),
			expected: 80,
			errMsg:   "result has to be 80",
		},
		{
			request:  Sum(50, 50),
			expected: 100,
			errMsg:   "result has to be 100",
		},
		{
			request:  Sum(55, 55),
			expected: 110,
			errMsg:   "result has to be 110",
		},
		{
			request:  Sum(60, 60),
			expected: 120,
			errMsg:   "result has to be 120",
		},
		{
			request:  Sum(70, 70),
			expected: 140,
			errMsg:   "result has to be 140",
		},
		{
			request:  Sum(80, 80),
			expected: 160,
			errMsg:   "result has to be 160",
		},
		{
			request:  Sum(90, 95),
			expected: 185,
			errMsg:   "result has to be 180",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			require.Equal(t, test.expected, test.request, test.errMsg)
		})
	}
}

// Unit Test Failed
// func TestFailSum(t *testing.T) {
// 	result := Sum(10, 10)

// 	if result != 40 {
// 		t.FailNow()
// 	}

// 	fmt.Println("TestFailSum eksekusi terhenti")
// }

// // Error Method
// func TestFailSums(t *testing.T) {
// 	result := Sum(10, 10)

// 	if result != testing
// 		t.Error("Testing Error gans")
// 	}

// 	fmt.Println("Test Ekseskusi method terhenti")
// }

// func TestSum(t *testing.T) {
// 	result := Sum(10, 10)

// 	if result != 20 {
// 		t.FailNow()
// 	}

// 	fmt.Println("TestSum terhenti")
// }
