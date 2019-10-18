//Package maxmin is a custom package for finding maximum and minimum in a given slice of ints
package maxmin

//MaxMin takes slice of int as a parameter and returns max,min as output
func MaxMin(x ...int) (int, int) {

	var max, min int

	for value := range x {
		if value > max {
			max = value
		} else if value < min {
			min = value
		}
	}
	return max, min
}
