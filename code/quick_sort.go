package code

func quickSort(nums *[]int, start, end int) {
	if start >= end {
		return
	}

	left, right := start, end
	// 1. pivot，nums[start], nums[end]: get value not index
	pivot := (*nums)[(start+end)/2]
	// 2. left <= right not left < right
	for left <= right {
		// 3. nums[left] < pivot  not <=
		for left <= right && (*nums)[left] < pivot {
			left++
		}

		for left <= right && (*nums)[right] > pivot {
			right--
		}

		if left <= right {
			(*nums)[left], (*nums)[right] = (*nums)[right], (*nums)[left]
			// 4. don't forget left++ right--
			left++
			right--
		}
	}

	quickSort(nums, start, right)
	quickSort(nums, left, end)
}
