package tree

func mirrirDfs(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	tmp := root.Left
	root.Left = mirrirDfs(root.Right)
	root.Right = mirrirDfs(tmp)
	return root
}
func mirrorBfs(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}

		tmp := node.Left
		node.Left = node.Right
		node.Right = tmp
	}
	return root
}

func mirrorTree(root *TreeNode) *TreeNode {
	// return mirrirDfs(root)
	return mirrorBfs(root)
}
