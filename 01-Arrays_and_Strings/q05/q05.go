// Implement a method to perform basic string compression using the counts of repeated
//  characters. For example, the string aabcccccaaa would become a2blc5a3. If the
//  'compressed' string would not become smaller than the orig- inal string, your
//  method should return the original string.

package q05

import (
	"strconv"
)

func Compress(ins string) string {
	s := []rune(ins)
	comp := make([]rune, len(s))
	cur, cnt, pos := s[0], 1, 1

	comp[0] = s[0]

	for i := 1; i < len(s); i++ {
		switch {
		case s[i] == cur:
			cnt++
		case s[i] != cur && cnt == 1:
			// append next character
			comp[pos] = s[i]
			cur = comp[pos]
			pos++
		case s[i] != cur && cnt > 1:
			// append number of character
			strcnt := []rune(strconv.Itoa(cnt))
			sclen := len(strcnt)
			copy(comp[pos:], strcnt)
			// add next character
			pos += sclen
			comp[pos] = s[i]
			cur = comp[pos]
			// update tracking vars
			pos++
			cnt = 1
		}
	}

	if cnt > 1 {
		// add count of last character
		strcnt := []rune(strconv.Itoa(cnt))
		copy(comp[pos:], strcnt)
		pos += len(strcnt)
	}

	if pos < len(s) {
		return string(comp[:pos])
	}

	return ins
}
