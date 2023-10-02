package main

import (
	"fmt"
	"testing"
)

func TestAbs(t *testing.T) {
	if Abs(-1) != 1 {
		t.Error("Error")
	}
	if Abs(1) != 1 {
		t.Error("Error")
	}
}

func BenchmarkAbs(b *testing.B) {
	//data := generateTestData() // Generate test data
	//b.ResetTimer()             // Reset the benchmark timer (optional)

	for i := 0; i < b.N; i++ {
		// Code you want to benchmark
		result := Abs(-7)
		_ = result // Prevent compiler optimization of the result
	}
}

func TestFactorial(t *testing.T) {
	test := []struct {
		input  int
		output int
	}{
		{0, 1},
		{1, 1},
        {2, 2},
        {3, 6},
        {4, 24},
        {5, 120},
	}

	for _, tt := range test {
		t.Run(fmt.Sprintf("Input %d", tt.input), func(t *testing.T) {
			result := Factorial(tt.input)
			if result != tt.output {
				t.Errorf("Expected %d but got %d",tt.output,result)
			}
		})
	}
}

func BenchmarkFactorial(b *testing.B){
	for i:=0;i<=b.N;i++{
		result:=Factorial(10)
		_=result
	} 
}