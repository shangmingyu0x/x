package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 求二叉树的右视图
 * @param xianxu int整型一维数组 先序遍历
 * @param zhongxu int整型一维数组 中序遍历
 * @return int整型一维数组
 */
func solve(xianxu []int, zhongxu []int) []int {
	// write code here
	var ans []int

	root := buildTree(xianxu, zhongxu)

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		end := queue[len(queue)-1]
		ans = append(ans, end.Val)

		var tmp []*TreeNode
		for i := 0; i < len(queue); i++ {
			node := queue[i]
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
		}
		queue = tmp
	}
	return ans
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(xianxu []int, zhongxu []int) *TreeNode {
	if len(zhongxu) == 0 {
		return nil
	}

	root := &TreeNode{Val: xianxu[0]}

	var i int
	for i < len(zhongxu) {
		if zhongxu[i] == xianxu[0] {
			break
		}
		i++
	}

	root.Left = buildTree(xianxu[1:1+i], zhongxu[:i])
	root.Right = buildTree(xianxu[1+i:], zhongxu[i+1:])

	return root
}
