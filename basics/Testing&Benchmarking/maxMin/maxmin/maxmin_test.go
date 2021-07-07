package maxmin

import (
	"fmt"
	"testing"
)

func BenchMaxMin(b *testing.B) {
	values := []int{10, -10, 20, 22, 39, -9, 8}

	for i := 0; i < b.N; i++ {
		MaxMin(values...)
	}

}
func ExampleMaxmin() {
	values := []int{10, -10, 20, 22, 39, -9, 8}
	fmt.Println(MaxMin(values...))
	//Output :
	// 39 -10
}
func TestMaxMin(t *testing.T) {
	values := []int{10, -10, 20, 22, 39, -9, 8}
	max, min := MaxMin(values...)

	if max != 39 && min != -9 {
		t.Error("expected : ", 39, -9, "got : ", max, min)
	}

}
