package tree

func preorderDfs(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	*result = append(*result, root.Val)
	preorderDfs(root.Left, result)
	preorderDfs(root.Right, result)
}

func preorderTraversal(root *TreeNode) []int {
	ans := []int{}
	preorderDfs(root, &ans)
	return ans
}
