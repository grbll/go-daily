package medianoftwosortedarrays

func findMedianNumAndSortedArray(num int, nums []int) float64 {
	low := (len(nums) - 1) / 2
	high := (len(nums)) / 2
	if len(nums) > 1 {
		if num <= nums[low] {
			if len(nums)%2 == 0 {
				return float64(nums[low])
			} else {
				return float64(max(num, nums[low-1])+nums[low]) / 2
			}
		} else if num <= nums[high] { //here num>nums[low], thus nums[low] < nums[high], thus len(nums) even
			return float64(num)
		} else {
			if len(nums)%2 == 0 {
				return float64(nums[high])
			} else {
				return float64(nums[high]+min(nums[high+1], num)) / 2
			}
		}
	} else {
		return float64(num+nums[0]) / 2
	}

}

func findMedianSortedArraysPreordered(nums1 []int, nums2 []int) float64 {
	low1 := (len(nums1) - 1) / 2
	high1 := (len(nums1)) / 2
	low2 := (len(nums2) - 1) / 2
	high2 := (len(nums2)) / 2
	if len(nums1) != 1 {
		if nums1[high1] <= nums2[low2] {
			return findMedianSortedArraysPreordered(nums1[high1:], nums2[:len(nums2)-high1])
		} else if nums1[high1] <= nums2[high2] { //low2<high1 && high1<=high2
			if nums1[low1] <= nums2[low2] { //low1<=low2<high1<=high2
				return float64(nums2[low2]+nums1[high1]) / 2
			} else { //low2<low1<=high1<=high2
				return float64(nums1[low1]+nums1[high1]) / 2
			}
		} else { //high2<high1
			if nums1[low1] <= nums2[low2] { //low1<=low2<=high2<high1
				return float64(nums2[low2]+nums2[high2]) / 2
			} else if nums1[low1] <= nums2[high2] { //low2<low1<=high2<high1
				return float64(nums1[low1]+nums2[high2]) / 2
			} else { // low2<=high2<low1<=high1
				return findMedianSortedArraysPreordered(nums1[:len(nums1)-high1], nums2[high1:])
			}
		}
	} else {
		return findMedianNumAndSortedArray(nums1[0], nums2)
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 && len(nums2) == 0 {
		return 0.0
	} else if len(nums1) == 0 {
		return float64(nums2[(len(nums2)-1)/2]+nums2[len(nums2)/2]) / 2
	} else if len(nums2) == 0 {
		return float64(nums1[(len(nums1)-1)/2]+nums1[len(nums1)/2]) / 2
	} else {
		if len(nums2) < len(nums1) {
			return findMedianSortedArraysPreordered(nums2, nums1)
		} else {
			return findMedianSortedArraysPreordered(nums1, nums2)
		}
	}
}
