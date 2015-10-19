// Write a method to replace all spaces in a string with'%%20'. You may assume that
//  the string has sufficient space at the end of the string to hold the additional
//  characters, and that you are given the 'true' length of the string. (Note: if
//  imple- menting in Java, please use a character array so that you can perform this
//  opera- tion in place.) EXAMPLE Input: 'Mr John Smith Output: 'Mr%%20Dohn%%20Smith'

package q04

func EncodeSpaces(s []rune, inlen int) {
	ini := inlen - 1
	outi := len(s) - 1

	for outi > -1 && ini > -1 {
		if s[ini] == ' ' {
			copy(s[outi-2:outi+1], []rune("%20"))
			outi -= 3
		} else {
			s[outi] = s[ini]
			outi--
		}

		ini--
	}
}
