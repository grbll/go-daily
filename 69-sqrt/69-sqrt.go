package sqrt

func mySqrt(x int) int {
	if x == 1 {
		return 1
	}
	var bot, piv, top int = 0, x / 2, x
	for bot+1 < top {
		if x/piv >= piv {
			bot = piv
		} else {
			top = piv
		}
		piv = (bot + top) / 2
	}
	return bot
}
