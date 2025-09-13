package solution

import "sort"

func CanMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)
	for i:=0; i<len(arr)-2;i++{
		if arr[i+1]-arr[i]!=arr[i+2]-arr[i+1]{
			return false
		}
	}
	return true
}