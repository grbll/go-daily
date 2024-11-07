package integertoroman

var values = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
var strings = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

func intToRoman(num int) string {
	out := ""
	var index uint8 = uint8(0)
	for num > 0 {
		if values[index] <= num {
			num -= values[index]
			out += strings[index]
		} else {
			index++
		}
	}
	return out
}
