// Implement a function void reverse(char* str) in C or C++ which reverses a null-
//  terminated string.

package q02

func Reverse(s string) string {
	arr := []rune(s)
	lens := len(arr)

	for i := 0; i < lens/2; i++ {
		arr[i], arr[lens-i-1] = arr[lens-i-1], arr[i]
	}

	return string(arr)
}
