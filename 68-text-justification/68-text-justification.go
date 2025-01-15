package textjustification

import "slices"

func fullJustify(words []string, maxWidth int) []string {
	var out []string = make([]string, 0, len(words))
	var line []byte = make([]byte, 0, maxWidth)
	var spaceByte []byte = slices.Repeat([]byte{' '}, maxWidth)
	var spaceNum, last, next, curr int = 0, 0, 0, 0

	for next < len(words)-1 {
		next++
		curr += len(words[next-1])

		if curr+len(words[next])+(next-last) > maxWidth {
			spaceNum = (maxWidth - curr) / max(1, next-last-1)

			for i := last; i < next-1; i++ {
				line = append(line, []byte(words[i])...)
				line = append(line, spaceByte[:spaceNum]...)
				if i-last < (maxWidth-curr)%(next-last-1) {
					line = append(line, ' ')
				}
			}
			line = append(line, []byte(words[next-1])...)
			if next-last == 1 {
				line = append(line, spaceByte[:spaceNum]...)
			}
			out = append(out, string(line))
			line = line[:0]
			last = next
			curr = 0
		}
	}

	for i := last; i < next; i++ {
		line = append(line, []byte(words[i])...)
		line = append(line, ' ')
	}
	line = append(line, []byte(words[next])...)
	line = append(line, spaceByte[:maxWidth-len(line)]...)
	out = append(out, string(line))

	return out
}
