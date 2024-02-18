package tree

func postorderDfs(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	// 后序遍历，左、右、根
	postorderDfs(root.Left, result)
	postorderDfs(root.Right, result)
	*result = append(*result, root.Val)

}

func postorderTraversal(root *TreeNode) []int {
	ans := []int{}
	postorderDfs(root, &ans)
	return ans
}
