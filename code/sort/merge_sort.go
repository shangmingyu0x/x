package sort

// 912. 排序数组 https://leetcode.cn/problems/sort-an-array/S

// mergeSort .
func mergeSort(nums []int) []int {

	n := len(nums)
	if n <= 1 {
		return nums
	}

	mid := n / 2
	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	var ans []int

	ls, rs := 0, 0
	le, re := len(left)-1, len(right)-1
	for ls <= le && rs <= re {
		if left[ls] < right[rs] {
			ans = append(ans, left[ls])
			ls++
		} else {
			ans = append(ans, right[rs])
			rs++
		}
	}

	for ls <= le {
		ans = append(ans, left[ls])
		ls++
	}

	for rs <= re {
		ans = append(ans, right[rs])
		rs++
	}

	// if ls <= le {
	// 	ans = append(ans, left[ls:]...)
	// }
	// if rs <= re {
	// 	ans = append(ans, right[rs:]...)
	// }

	return ans
}
