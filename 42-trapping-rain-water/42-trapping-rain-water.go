package trappingrainwater

func trap(height []int) int {
	var out, curr, indexLow, indexHigh int

	indexLow, indexHigh = 0, len(height)-1

	out = 0
	if indexLow < indexHigh {
		curr = min(height[indexLow], height[indexHigh])
	}
	for indexLow < indexHigh-1 {
		if height[indexLow] <= height[indexHigh] {
			indexLow++
			curr = max(curr, min(height[indexLow], height[indexHigh]))
			out += max(0, curr-height[indexLow])
		} else {
			indexHigh--
			curr = max(curr, min(height[indexLow], height[indexHigh]))
			out += max(0, curr-height[indexHigh])
		}
	}
	return out
}
