/**
6. ZigZag Conversion
**/
package zigzag_conversion

func convert(s string, numRows int) string {
	//make slice, length is numRows
	storage := make([][]rune, numRows)
	runes := []rune(s)
	len := len(s)
	if numRows == 1 || len == 0 {
		return s
	}

	for i := 0; i < len; {
		row := 0
		for row < numRows && i < len {
			storage[row] = append(storage[row], runes[i])
			i++
			row++
		}
		for row-2 > 0 && i < len {
			storage[row-2] = append(storage[row-2], runes[i])
			i++
			row--
		}
	}

	rst := ""
	for i := 0; i < numRows; i++ {
		rst = rst + string(storage[i][:])
	}

	return rst

}
