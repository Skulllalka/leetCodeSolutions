package solution

func IsMonotonic(nums []int) bool {
	
	dlina:=len(nums)
	if dlina==1{
		return true
	}	
	
	isDec:=true
	isInc:=true

	for i:=0;i<dlina-1;i++{
		if !isDec && !isInc{
			return false
		}
		
		if nums[i+1]>nums[i]{
			isInc=false
		}
		if nums[i+1]<nums[i]{
			isDec=false
		}

	}

	return isDec || isInc
}