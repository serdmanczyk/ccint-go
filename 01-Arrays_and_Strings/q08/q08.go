// Assume you have a method isSubstring which checks if one word is a substring of
//  another. Given two strings, s1 and s2, write code to check if s2 is a rotation
//  of si using only one call to isSubstring (e.g.,'waterbottle'is a rota- tion of'erbottlewat').

package q08

import (
	"strings"
)

func IsRotation(s1, s2 string) bool {
	return strings.Contains(s1+s1, s2)
}
