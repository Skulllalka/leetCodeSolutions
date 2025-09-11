package solution

import "reflect"

func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	value1, value2 := make(map[rune]int, 0), make(map[rune]int, 0)

	for i := 0; i < len(s); i++ {
		value1[rune(s[i])]++
		value2[rune(t[i])]++
	}
	if reflect.DeepEqual(value1, value2) {
		return true
	}
	return false
}
