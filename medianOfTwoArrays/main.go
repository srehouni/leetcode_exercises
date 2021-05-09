package main

import "fmt"

func main() {
	array1 := []int{1, 2}
	array2 := []int{3, 4}

	fmt.Println(findMedianSortedArrays(array1, array2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	mergedSlice := []int{}

	for i := 0; len(nums1)+len(nums2) > 0; i++ {
		if 0 < len(nums1) && 0 < len(nums2) {
			if nums1[0] > nums2[0] {
				mergedSlice = append(mergedSlice, nums2[0])
				nums2 = nums2[1:]
			} else {
				mergedSlice = append(mergedSlice, nums1[0])
				nums1 = nums1[1:]
			}
		} else if len(nums1) > 0 {
			mergedSlice = append(mergedSlice, nums1[0])
			nums1 = nums1[1:]
		} else {
			mergedSlice = append(mergedSlice, nums2[0])
			nums2 = nums2[1:]
		}
	}

	if len(mergedSlice)%2 == 0 {
		return float64(mergedSlice[len(mergedSlice)/2]+mergedSlice[len(mergedSlice)/2-1]) / 2
	} else {
		return float64(mergedSlice[len(mergedSlice)/2])
	}
}
