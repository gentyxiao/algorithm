package tree

// 递归 vals必须是指针
func levelDfs(root *TreeNode, level int, vals *[][]int) {
	if root == nil {
		return
	}
	if level == len(*vals) {
		*vals = append(*vals, []int{})
	}
	(*vals)[level] = append((*vals)[level], root.Val)
	levelDfs(root.Left, level+1, vals)
	levelDfs(root.Right, level+1, vals)
}

// 二叉树层序遍历，深度搜索优先实现
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	ans := [][]int{}
	levelDfs(root, 0, &ans)
	return ans
}
