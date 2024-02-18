package tree

func inorderDfs(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	// 中序遍历，左、根、右
	preorderDfs(root.Left, result)
	*result = append(*result, root.Val)
	preorderDfs(root.Right, result)
}

func inorderTraversal(root *TreeNode) []int {
	ans := []int{}
	inorderDfs(root, &ans)
	return ans
}
