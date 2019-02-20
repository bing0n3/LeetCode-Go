package maxiamlRectangel

func maximalRectangle(matrix [][]byte) int {

	if len(matrix) == 0 || len(matrix[0]) == 0 || matrix == nil {
		return 0
	}

	rowCnt := len(matrix)
	colCnt := len(matrix[0])

	height := make([][]int, rowCnt)
	var maxArea int

	// O(n^2)
	for i := 0; i < rowCnt; i++ {
		height[i] = []int{}
		for j := 0; j < colCnt; j++ {
			if i == 0 {
				height[i] = append(height[i], int(matrix[i][j]-'0'))
			} else {
				//if height[i][j] equal to zero, the result should be zero
				height[i] = append(height[i], 0)
				if matrix[i][j] != '0' {
					height[i][j] = int(matrix[i][j]-'0') + height[i-1][j]
				}
			}
		}
	}

	monoStack := []int{}
	for i := 0; i < rowCnt; i++ {
		// initial stack
		monoStack = monoStack[:0]
		for j := 0; j <= colCnt; j++ {
			curHeight := -1
			if j != colCnt {
				curHeight = height[i][j]
			}

			// 寻找右界，当栈顶元素大于当前元素，说明到达了栈顶元素的右界
			for len(monoStack) != 0 && height[i][monoStack[len(monoStack)-1]] >= curHeight {
				// 弹出当前栈顶，因为已经打到右界
				//stack pop
				h := height[i][monoStack[len(monoStack)-1]]
				monoStack = monoStack[:len(monoStack)-1]

				// 当stack为空，说明当前的元素比左边所有元素小，宽度可以为它当前的索引，左界是0
				w := j
				if len(monoStack) != 0 {
					// 计算他可以到达的左界
					w = j - monoStack[len(monoStack)-1] - 1
				}
				if maxArea < w*h {
					maxArea = w * h
				}
			}
			monoStack = append(monoStack, j)
		}
	}

	return maxArea

}
