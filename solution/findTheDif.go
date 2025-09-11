package solution

import (
// "strconv"
// "strings"
)

func FindTheDifference(s string, t string) byte {
	var retValue byte
	flag := false

	for i := 0; i < len(s); i++ {
		if s[i] != t[i] {
			flag = true
			retValue = t[i]
		}
	}
	if flag {
		return retValue
	}
	return t[len(t)-1]
}
