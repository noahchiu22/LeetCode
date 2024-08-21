package main

import (
	"fmt"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	median := 0.0

	fmt.Println("nums1:", len(nums1), "nums2:", len(nums2))

	if len(nums2) > 0 {
		i := 0
		j := 0
		for i < len(nums1) {
			for nums2[j] < nums1[i] {
				nums1 = append(nums1[:i+1], nums1[i:]...)
				nums1[i] = nums2[j]
				j++
				if j == len(nums2) {
					break
				}
			}
			if j == len(nums2) {
				break
			}
			i++
		}
		if j < len(nums2) {
			nums1 = append(nums1, nums2[j:]...)
		}
	}

	fmt.Println("merged array =", nums1)

	if len(nums1)%2 == 0 {
		median = float64(nums1[len(nums1)/2-1]+nums1[len(nums1)/2]) / 2
	} else {
		median = float64(nums1[len(nums1)/2])
	}

	return median
}
