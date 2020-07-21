package algorithm

func searchMatrix(matrix [][]int, target int) bool {
	// 空矩阵直接判断
	if len(matrix) == 0 {
		return false
	} else if len(matrix[0]) == 0 {
		return false
	}
	var (
		rowleft     int = 0
		rowright    int = len(matrix) - 1
		row         int = (rowleft + rowright) / 2
		columnleft  int = 0
		columnright int = len(matrix[0]) - 1
		column      int = (columnleft + columnright) / 2
	)
	// 二分确定行
	if rowright >= 0 {
		for rowright >= rowleft {
			row = (rowleft + rowright) / 2
			switch {
			case matrix[row][columnright] < target:
				rowleft = row + 1
			case matrix[row][columnright] > target:
				rowright = row - 1
			default:
				return true
			}
		}
		// 参照二分法寻找插入位置，此处应返回left指针
	}
	// 消除指针异常
	if rowleft > len(matrix)-1 {
		rowleft = len(matrix) - 1
	}
	// 二分确定列
	if columnright >= 0 {
		for columnright >= columnleft {
			column = (columnleft + columnright) / 2
			switch {
			case matrix[rowleft][column] < target:
				columnleft = column + 1
			case matrix[rowleft][column] > target:
				columnright = column - 1
			default:
				return true
			}
		}
	}
	return false
}
