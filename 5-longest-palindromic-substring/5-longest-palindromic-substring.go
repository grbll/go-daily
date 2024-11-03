// attempt on manachers algorithm
package longestpalindromicsubstring

func longestPalindrome(s string) string {
	if len(s) > 0 {

		largestRadius := make([]int, 2*len(s)+1)

		index := 0
		radius := 0
		maxRadiusIndex := 0

		for index < 2*len(s)+1 { // number of palindrome centers in s

			for index-(radius+1) >= 0 && index+(radius+1) <= 2*len(s) {
				if (index-radius)%2 == 1 {
					radius++
				} else {
					if s[(index-radius-1)/2] == s[(index+radius+1)/2] {
						radius++
					} else {
						break
					}
				}
			}

			largestRadius[index] = radius

			if radius > largestRadius[maxRadiusIndex] {
				maxRadiusIndex = index
			}

			mirrorPoint := index
			endMirrorPalindrome := mirrorPoint + radius

			radius = 0
			index++

			for index <= endMirrorPalindrome {
				if largestRadius[mirrorPoint-(index-mirrorPoint)] < endMirrorPalindrome-index {
					largestRadius[index] = largestRadius[mirrorPoint-(index-mirrorPoint)]
					index++
				} else if largestRadius[mirrorPoint-(index-mirrorPoint)] > endMirrorPalindrome-index {
					largestRadius[index] = endMirrorPalindrome - index
					index++
				} else { //largestRadius[mirrorPoint-(index-mirrorPoint)] == (endMirrorPalindrome-index)/2
					radius = endMirrorPalindrome - index
					break
				}
			}
		}
		return s[(maxRadiusIndex)/2-largestRadius[maxRadiusIndex]/2 : (maxRadiusIndex)/2+maxRadiusIndex%2+(largestRadius[maxRadiusIndex])/2]
	} else {
		return ""
	}
}
