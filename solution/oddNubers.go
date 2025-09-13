package solution

func СountOdds(low int, high int) int {
	count, start:=0, 0
	
	if low%2==0 {
		start=low+1
	}else{
		start=low
	}
	for start<=high{
		count++
		start+=2
	}
	return count
}