package exercise

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// 路径总和是否存在
func HasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return targetSum == root.Val
	}
	return HasPathSum(root.Left, targetSum-root.Val) || HasPathSum(root.Right, targetSum-root.Val)
}
