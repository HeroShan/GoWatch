package leetcode
// DFS深度优先算法
// 判断边界
//         {
//             相应操作
//         }
//         尝试每一种可能
//         {
//                满足check条件
//                标记
//                继续下一步dfs(step+1)
//                恢复初始状态（回溯的时候要用到）
//         }
func NumIslands(grid [][]byte) int {
	count := 0
	for line := range grid {
		for column := range grid[line] {
			if grid[line][column] == '0' || grid[line][column] == '2' {
				continue
			}

			count++
			dfs(grid, line, column)
		}
	}
	return count
}

func dfs(grid [][]byte, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) {
		return
	}

	if grid[i][j] != '1' {
		return
	}

	grid[i][j] = '2'
	dfs(grid, i+1, j)
	dfs(grid, i-1, j)
	dfs(grid, i, j+1)
	dfs(grid, i, j-1)
}