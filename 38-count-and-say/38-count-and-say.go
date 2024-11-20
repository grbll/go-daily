package countandsay

import "bytes"

func countAndSay(n int) string {
	var outBuf *bytes.Buffer = bytes.NewBuffer([]byte{'1'})
	var cur, nex, cou byte
	for i := 0; i < n-1; i++ {
		cur, _ = outBuf.ReadByte()
		cou = '1'
		outBuf.WriteByte('x')
		for {
			nex, _ = outBuf.ReadByte()
			if nex == 'x' {
				outBuf.WriteByte(cou)
				outBuf.WriteByte(cur)
				break
			} else {
				if cur == nex {
					cou++
				} else {
					outBuf.WriteByte(cou)
					outBuf.WriteByte(cur)
					cur = nex
					cou = '1'
				}
			}
		}
	}
	return outBuf.String()
}
