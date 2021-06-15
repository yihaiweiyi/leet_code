package _852_Pea_Index_in_a_Mountain_Array

// 解法一
func peakIndexInMountainArray(arr []int) int {
	if len(arr) < 3 {
		return 0
	}

	length := len(arr)
	for i:= 1;i<length;i++ {
		if i == length - 1 {
			return i
		}
		if arr[i-1] < arr[i] && arr[i] > arr[i+1] {
			return i
		}
	}
	return 0
}

// 解法二
func peakIndexInMountainArray2(arr []int) int {
	if len(arr) < 3 {
		return 0
	}

	left, right := 0, len(arr) - 1
	for left < right {
		middle := left + (right - left) >> 1
		if arr[middle] <= arr[middle+1] {
			left = middle +1
		} else {
			right = middle
		}
	}
	return left
}