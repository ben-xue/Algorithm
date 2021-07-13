func subsets(nums []int) (ans [][]int) {
	stNumsMap := make(map[int]int)
	for index, num := range nums {
		stNumsMap[index] = num
	}

	numsLen := float64(len(nums))
	sliceCap := int(math.Pow(2,numsLen))

	stBitsMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		stBitsMap[i] = int(math.Pow(2,float64(i)))
	}

	ans = [][]int{}
	for i := 0; i < sliceCap; i++ {
		innerAnd := []int{}
		for iBit := 0; iBit < len(nums); iBit++ {
			if stBitsMap[iBit] & i != 0{
				innerAnd = append(innerAnd, stNumsMap[iBit])
			}
		}
		ans = append(ans, innerAnd)
	}

	return  ans
}
