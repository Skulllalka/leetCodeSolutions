package solution

// import (
// 	"fmt"
// )

func Average(salary []int) float64 {
	tempMin:=1000001
	tempMax:=0
	retSum:=0.0

	for _, number:=range salary{
		retSum+=float64(number)
		tempMin= func(a,b  int) int{
			if a>b{
				return b
			}else{
				return a
			}
		}(tempMin, number)
		tempMax= func(a, b int) int{
			if a>b{
				return a
			}else{
				return b
			}
		}(tempMax, number)
				
	}
	retSum=(retSum-float64(tempMax)-float64(tempMin))/float64((len(salary)-2))

	return retSum
}