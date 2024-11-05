package containerwithmostwater

func maxArea(height []int) int {

	area := 0
	currentMax := 0
	i := 0
	j := len(height) - 1

	for i < j {

		if height[i] <= height[j] {
			area = (j - i) * height[i]
			i++
		} else {
			area = (j - i) * height[j]
			j--
		}

		if area > currentMax {
			currentMax = area
		}

	}
	return currentMax
}
