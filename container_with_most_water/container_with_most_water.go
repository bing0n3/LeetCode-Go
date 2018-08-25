/**
 * 11. Container with most water
 */

package container_with_most_water

func maxArea(height []int) int {
	len := len(height)
	maxContainer := 0
	for i := 0; i < len-1; i++ {
		for j := i + 1; j < len; j++ {
			tmp := 0
			if height[i] > height[j] {
				tmp = height[j] * (j - i + 1)
			} else {
				tmp = height[i] * (j - i + 1)
			}
			if tmp > maxContainer {
				maxContainer = tmp
			}
		}
	}
	return maxContainer
}

func maxArea2(height []int) int {
	maxContainer := 0
	i := 0
	j := len(height) - 1
	tmp := 0
	for i < j {
		if height[i] > height[j] {
			tmp = height[j] * (j - i)
			j--
		} else {
			tmp = height[i] * (j - i)
			i++
		}
		if tmp > maxContainer {
			maxContainer = tmp
		}

	}
	return maxContainer
}
