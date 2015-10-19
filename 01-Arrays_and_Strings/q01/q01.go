// Implement an algorithm to determine if a string has all unique characters. What
//  if you cannot use additional data structures?

package q01

import (
	"ccint/ccintcommon"
)

// O(n)
func AllUniqueCharsAssumeASCII(s string) bool {
	letters := make([]byte, 'z'-'a'+1)

	// Accessing the string by index yeilds a byte
	for i := 0; i < len(s); i++ {
		if letters[s[i]%'a'] == 1 {
			return false
		}
		letters[s[i]%'a'] = 1
	}

	return true
}

// O(n) + hash map cost
func AllUniqueCharsHashMap(s string) bool {
	charsInStr := make(map[rune]int, len(s))

	for _, c := range s {
		if charsInStr[c] == 1 {
			return false
		}
		charsInStr[c] = 1
	}

	return true
}

// O(n^2)
func AllUniqueCharsNoDataStructsOne(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				return false
			}
		}
	}

	return true
}

// O(nlog(n) + n)
func AllUniqueCharsNoDataStructsTwo(s string) bool {
	s = ccintcommon.SortString(s)
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			return false
		}
	}

	return true
}
