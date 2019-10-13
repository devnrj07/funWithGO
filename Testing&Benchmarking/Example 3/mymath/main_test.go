package mymath

import (
	"fmt"
	"testing"
)

//Benchmark is used measuring the performance of the written method
func BenchmarkCenteredAvg(b *testing.B) {

	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{1, 2, 4, 5, 55, 6443, 663})
	}
}

//ExampleCenteredAvg documents the working of CenteredAvg
func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{1, 5, 4, 5, 55, 66, 663}))
	//Output:
	//27
}
//TestCentredAvg uses Table test
func TestCenteredAvg(t *testing.T) {
	type dummyData struct {
		values []int
		answer float64
	}

	tests := []dummyData {
				dummyData{[]int{1,2,3,4,5},3.0},
				dummyData{[]int{9,8,7,6,5},7.0},
				dummyData{[]int{12,8,90,-8,7},30.0},
	}

	for _, v := range tests {

		actualValue := CenteredAvg(v.values)
		if actualValue != v.answer {
			t.Error("expected : ",v.answer,"got : ", actualValue)
		}
	}

	}

}
