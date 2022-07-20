package sort

// 912. 排序数组 https://leetcode.cn/problems/sort-an-array/S

func quickSort(nums []int, start, end int) {
	if start >= end {
		return
	}

	pivot := nums[(start+end)/2]
	left, right := start, end
	for left <= right {
		for left <= right && nums[left] < pivot {
			left++
		}

		for left <= right && nums[right] > pivot {
			right--
		}

		if left <= right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}

	quickSort(nums, start, right)
	quickSort(nums, left, end)
}
