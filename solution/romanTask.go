package solution


func RomainToInt (s string) int{
	plusMap:= map[string]int{
		"I":1,
		"V":5,
		"X":10,
		"L":50,
		"C":100,
		"D":500,
		"M":1000,
	}
	
	retValue:=0

	for i:=0;i<len(s)-1;i++{
		tempCur:=plusMap[string(s[i])]
		tempNext:=plusMap[string(s[i+1])]
		if tempCur<tempNext{
			retValue-=tempCur
		}else{
			retValue+=tempCur
		}
	}

	retValue+=plusMap[string(s[len(s)-1])]

	return retValue
}