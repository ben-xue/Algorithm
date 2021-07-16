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


/*多协程版本*/
func CalcSubSet(wg *sync.WaitGroup, ch chan [][]int, index int, iStart int, iEnd int, stNumsMap map[int]int, stBitsMap map[int]int) {
	iBitsCount := len(stBitsMap)
	ans := [][]int{}
	for i := iStart; i <= iEnd; i++ {
		innerAnd := []int{}
		for iBit := 0; iBit < iBitsCount; iBit++ {
			if stBitsMap[iBit]&i != 0 {
				innerAnd = append(innerAnd, stNumsMap[iBit])
			}
		}
		ans = append(ans, innerAnd)
	}

	//fmt.Println(index,":",ans)
	ch <- ans
	wg.Done()
}

func subsetsInGoroutine(nums []int) (ans [][]int) {
	stNumsMap := make(map[int]int)
	for index, num := range nums {
		stNumsMap[index] = num
	}

	numsLen := float64(len(nums))
	sliceCap := int(math.Pow(2, numsLen))

	stBitsMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		stBitsMap[i] = int(math.Pow(2, float64(i)))
	}

	iSeqCount := 1
	iSeqNums := sliceCap
	cores := runtime.NumCPU()
	if sliceCap/cores > 3 {
		iSeqCount = cores + 1
		iSeqNums = sliceCap / cores
	}

	fmt.Println(iSeqCount,":",iSeqNums)

	wg := sync.WaitGroup{}
	wg.Add(iSeqCount)

	ch := make(chan [][]int, iSeqCount)

	ans = [][]int{}
	for index, iBegin := 0, 0; index < iSeqCount; iBegin, index = iBegin+iSeqNums, index+1 {
		iEnd := iBegin + iSeqNums - 1
		if iEnd > sliceCap-1 {
			iEnd = sliceCap - 1
		}
		go CalcSubSet(&wg, ch, index, iBegin, iEnd, stNumsMap, stBitsMap)
	}

	wg.Wait()

	loopFlag := true
	for loopFlag {
		select {
		case stTmpAns := <-ch:
			{
				ans = append(ans, stTmpAns...)
			}
		default:
			{
				loopFlag = false
			}
		}
	}
	return ans
}