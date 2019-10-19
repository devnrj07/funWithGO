package main
import( 
	"fmt"
	"github.com/letsgo/basics/Testing&Benchmarking/maxMin/maxmin"
)

func main() {
	list := []int{10, 10, 12, -1, 29, 33, 45, 128, 0}
	max, min := maxmin.MaxMin(list...)
	fmt.Println("Max value found : ", max, "Min value found : ", min)
}
