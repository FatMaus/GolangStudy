package algorithm

// MinTrianglePath1 自底向上动态规划
func MinTrianglePath1(triangle [][]int) int {
	var (
		length  int
		tempArr [][]int
	)
	if len(triangle) == 0 || len(triangle[0]) == 0 {
		return 0
	}
	length = len(triangle)
	tempArr = make([][]int, length)
	// 准备解集容器
	for i := 0; i < length; i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if tempArr[i] == nil {
				tempArr[i] = make([]int, len(triangle[i]))
			}
			tempArr[i][j] = triangle[i][j]
		}
	}
	// 进行迭代寻找最佳解
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			tempArr[i][j] = min(tempArr[i+1][j], tempArr[i+1][j+1]) + triangle[i][j]
		}
	}
	return tempArr[0][0]
}

// MinTrianglePath2 自顶向下动态规划
func MinTrianglePath2(triangle [][]int) int {
	if len(triangle) == 0 || len(triangle[0]) == 0 {
		return 0
	}
	var (
		length  int     = len(triangle)
		tempArr [][]int = make([][]int, length)
		ret     int
	)
	// 准备解集容器
	for i := 0; i < length; i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if tempArr[i] == nil {
				tempArr[i] = make([]int, len(triangle[i]))
			}
			tempArr[i][j] = triangle[i][j]
		}
	}
	// 迭代寻找最佳解
	for i := 1; i < length; i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if j-1 < 0 {
				tempArr[i][j] = tempArr[i-1][j] + triangle[i][j]
			} else if j >= len(tempArr[i-1]) {
				tempArr[i][j] = tempArr[i-1][j-1] + triangle[i][j]
			} else {
				tempArr[i][j] = min(tempArr[i-1][j], tempArr[i-1][j-1]) + triangle[i][j]
			}
		}
	}
	ret = tempArr[length-1][0]
	for i := 1; i < len(tempArr[length-1]); i++ {
		ret = min(ret, tempArr[length-1][i])
	}
	return ret
}

func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	for i := 1; i < len(grid); i++ {
		grid[i][0] = grid[i][0] + grid[i-1][0]
	}
	for j := 1; j < len(grid[0]); j++ {
		grid[0][j] = grid[0][j] + grid[0][j-1]
	}
	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			grid[i][j] = min(grid[i][j-1], grid[i-1][j]) + grid[i][j]
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}

func numOfPaths(m int, n int) int {
	var tempArr [][]int = make([][]int, m)
	// 初始化地图
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if tempArr[i] == nil {
				tempArr[i] = make([]int, n)
			}
			tempArr[i][j] = 1
		}
	}
	// 求解路径
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			tempArr[i][j] = tempArr[i-1][j] + tempArr[i][j-1]
		}
	}
	return tempArr[m-1][n-1]
}

func min(num1, num2 int) int {
	if num1 > num2 {
		return num2
	}
	return num1
}
