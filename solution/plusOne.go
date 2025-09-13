package solution



func PlusOne(digits []int) []int {
	for i:=len(digits)-1;i>=0;i--{
		if digits[i]<9{
			digits[i]++
			return digits
		}
		digits[i]=0
	}
	retValue:=make([]int, len(digits)+1)
	retValue[0]=1
	return retValue
}