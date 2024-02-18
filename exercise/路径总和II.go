package exercise

// 路径总和 II
// 给你二叉树的根节点 root 和一个整数目标和 targetSum
// 找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径
func HasPathSum2Dfs(root *TreeNode, targetSum, sum int, path []int, ans *[][]int) {
	if root == nil {
		return
	}
	path = append(path, root.Val)
	sum += root.Val
	if sum == targetSum && root.Left == nil && root.Right == nil {
		*ans = append(*ans, append([]int(nil), path...))
	}
	HasPathSum2Dfs(root.Left, targetSum, sum, path, ans)
	HasPathSum2Dfs(root.Right, targetSum, sum, path, ans)
}

// 路径总和2
func HasPathSum2(root *TreeNode, targetSum int) [][]int {
	ans := make([][]int, 0)
	path := make([]int, 0)
	HasPathSum2Dfs(root, targetSum, 0, path, &ans)
	return ans
}
