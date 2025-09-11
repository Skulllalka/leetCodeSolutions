package solution

import "strings"

func LengthOfLastWord(s string) int {
	s = strings.Trim(s, " ")
	retCount := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			break
		}

		retCount++
	}
	return retCount

}
