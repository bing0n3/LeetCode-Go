/**
 * 11. Container with most water
 */

package container_with_most_water

func maxArea(height []int) int {
	maxContainer := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			tmp := 0
			if height[i] > height[j] {
				tmp = height[j] * (j - i)
			} else {
				tmp = height[i] * (j - i)
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
