package floodFill

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if len(image) == 0 {
		return image
	}

	if image[sr][sc] == newColor {
		return image
	}

	dfs(&image, sr, sc, newColor, image[sr][sc])

	return image
}

func dfs(image *[][]int, sr int, sc int, newColor, nowColor int) {
	if (*image)[sr][sc] != nowColor {
		return
	}

	(*image)[sr][sc] = newColor

	if sr < len(*image)-1 {
		dfs(image, sr+1, sc, newColor, nowColor)
	}

	if sr > 0 {
		dfs(image, sr-1, sc, newColor, nowColor)
	}

	if sc < len((*image)[0])-1 {
		dfs(image, sr, sc+1, newColor, nowColor)
	}

	if sc > 0 {
		dfs(image, sr, sc-1, newColor, nowColor)
	}
}
