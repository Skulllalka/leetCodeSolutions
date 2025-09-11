package solution

import "strings"

func StrStr(haystack string, needle string) int {
	if !strings.Contains(haystack, needle) {
		return -1
	}
	return strings.Index(haystack, needle)
}
