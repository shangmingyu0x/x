package sort

// 912. 排序数组 https://leetcode.cn/problems/sort-an-array/S

// HeapSort .
// 堆是一个完全二叉树，通常采用数组存储
func HeapSort(nums []int) {
	n := len(nums)

	// 建堆
	//「整体」从前往后，「每个节点」从上往下
	// 时间复杂度是 O(n)
	// 叶子节点不需要堆化，从倒数第二层节点开始
	for i := n/2 - 1; i >= 0; i-- {
		heapify(nums, i, n-1)
	}

	// 排序
	// 建堆结束之后，数组中的数据已经是按照大顶堆的特性来组织的。
	// 把下标为 n 的元素放到堆顶，然后再通过堆化的方法，将剩下的 n−1 个元素重新构建成堆。
	// 一直重复这个过程，直到最后堆中只剩下标为 1 的一个元素，排序工作就完成了。
	for i := n - 1; i > 0; i-- {
		swap(nums, 0, i)
		heapify(nums, 0, i-1)
	}
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

// heapify 堆化区间[s, e]的数组，s表示起始索引，e表示结束索引
// down 自上而下，通常用于「删除」操作
func heapify(nums []int, s, e int) {
	for {
		i := s
		l, r := 2*s+1, 2*s+2
		if l > e || l < 0 {
			break
		}

		if l <= e && nums[l] > nums[i] {
			i = l
		}
		if r <= e && nums[r] > nums[i] {
			i = r
		}

		if i <= s {
			break
		}
		swap(nums, s, i)
		s = i
	}
}

// up 自下而上，通常用于「插入」操作
func up(nums []int, i int) {
	for {
		j := (i - 1) / 2 // parent
		if i == j || nums[j] <= nums[i] {
			break
		}
		swap(nums, i, j)
		i = j
	}
}
