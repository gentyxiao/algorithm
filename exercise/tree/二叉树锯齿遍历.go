package tree

// 二叉树节点定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉树锯齿层遍历
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	ans := [][]int{}
	// 广度优先算法实现，需要借助队列保存各个节点的值
	queue := []*TreeNode{root}
	for level := 0; len(queue) > 0; level++ {
		q := queue
		queue = nil
		vals := []int{}
		for _, node := range q {
			vals = append(vals, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		// 判断奇偶来确定正反序
		if level%2 == 1 {
			reverseArray(vals)
		}
		ans = append(ans, vals)
	}
	return ans
}

func reverseArray(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

// ----------------------------
// 方法二，深度搜素实现
func zigzagDfs(node *TreeNode, level int, result *[][]int) {
	if node == nil {
		return
	}

	if level == len(*result) {
		*result = append(*result, []int{})
	}

	if level%2 == 0 {
		// 正序插入
		(*result)[level] = append((*result)[level], node.Val)
	} else {
		// 反序插入
		(*result)[level] = append([]int{node.Val}, (*result)[level]...)
	}

	zigzagDfs(node.Left, level+1, result)
	zigzagDfs(node.Right, level+1, result)
}

func zigzagLevelOrder2(root *TreeNode) [][]int {
	var result [][]int
	zigzagDfs(root, 0, &result)
	return result
}
