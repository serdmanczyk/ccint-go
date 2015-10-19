// Given two strings, write a method to decide if one is a
// permutation of the other.

package q03

// O(nlog(n) + klog(k) + min(n, k))
func IsPermutation(one, two string) bool {
	if len(one) != len(two) {
		return false
	}

	one = SortString(one)
	two = SortString(two)

	for i := 0; i < len(one); i++ {
		if one[i] != two[i] {
			return false
		}
	}

	return true
}
