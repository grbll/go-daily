package containerwithmostwater

func maxArea(height []int) int {

	currentMax := 0
	n := len(height)

	if n > 1 {

		heightPoints := make([]int, 1, n)
		heightPoints[0] = 0
		lastUpdate := 0
		value := 0

		index1 := 1
		for index1 < n {
			index2 := 0
			for index2 <= lastUpdate {
				value = min(height[index1], height[heightPoints[index2]]) * (index1 - heightPoints[index2])
				if value > currentMax {
					currentMax = value
				}
				index2++
			}
			if height[index1] > height[lastUpdate] {
				heightPoints = append(heightPoints, index1)
				lastUpdate++
			}
			index1++
		}
	}
	return currentMax
}
