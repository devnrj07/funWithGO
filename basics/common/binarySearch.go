package main

import (
	"fmt"
)

func main() {
	x1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(binarySearch(x1, 15, 0, len(x1)-1))

	sorted := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(binarySearch(sorted, 3, 0, len(sorted)-1))

	sortedArray := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(binarySearch(sortedArray, 8, 0, len(sortedArray)-1))

	unSortedArray := []int{11, 2, 35, 43, 5, 16, 37, 8, 9} // does not work for unsorted array
	fmt.Println(binarySearch(unSortedArray, 9, 0, len(unSortedArray)-1))
}

func binarySearch(arr []int, target int, low int, high int) (bool, int) {

	if high > low {

		mid := low + (high-low)/2

		if target == arr[mid] {
			return true, mid
		}

		if target < arr[mid] {

			return binarySearch(arr, target, low, mid)
		}

		return binarySearch(arr, target, mid+1, high)
	}

	return false, -1
}
