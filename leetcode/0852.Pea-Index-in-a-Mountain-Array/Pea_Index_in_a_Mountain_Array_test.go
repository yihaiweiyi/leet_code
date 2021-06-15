package _852_Pea_Index_in_a_Mountain_Array

import (
	"fmt"
	"testing"
)

func Test_peakIndexInMountainArray(t *testing.T) {
	arr := []int{0, 1, 3, 4}
	index := peakIndexInMountainArray2(arr)
	fmt.Println(index)

	arr = []int{0,10,5,2}
	index = peakIndexInMountainArray2(arr)
	fmt.Println(index)

	arr = []int{24,69,100,99,79,78,67,36,26,19}
	index = peakIndexInMountainArray2(arr)
	fmt.Println(index)
}
