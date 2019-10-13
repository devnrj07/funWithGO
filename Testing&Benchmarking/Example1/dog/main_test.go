package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	dogYears := Years(7)
	if dogYears != 49 {
		t.Error("Expected : 49 but got : ", dogYears)
	}

}
func TestYearsTwo(t *testing.T) {
	dogYears := YearsTwo(7)
	if dogYears != 49 {
		t.Error("Expected : ", 49, " but got : ", dogYears)
	}

}

func ExampleYears() {
	dogYears := Years(10)
	fmt.Println(dogYears)
	//Output :
	// 70
}

func ExampleYearsTwo() {
	dogYears := YearsTwo(10)
	fmt.Println(dogYears)
	//Output :
	// 70
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(10)
	}
}
