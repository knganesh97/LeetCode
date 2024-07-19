package problem1530

func countPairs(root *TreeNode, distance int) int {
	var count int
	var dfs func(node *TreeNode) []int

	dfs = func(node *TreeNode) []int {

		if node == nil {
			return nil
		}

		left := dfs(node.Left)
		right := dfs(node.Right)

		// fmt.Println(node, left, right)

		for i := range left {
			for j := range right {
				if left[i]+right[j] <= distance {
					count++
				}
			}
			left[i]++
		}

		for i := range right {
			right[i]++
			left = append(left, right[i])
		}

		if left == nil {
			left = append(left, 1)
		}

		return left
	}

	dfs(root)
	return count
}
