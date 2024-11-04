package reverseinteger

import "math"

func reverse(x int) int {
	var value int64 = int64(x)
	var out int64 = 0
	for value != 0 {
		out = out*10 + value%10
		if math.MaxInt32 < out || math.MinInt32 > out {
			return 0
		} else {
			value = value / 10
		}
	}
	return int(out)
}
