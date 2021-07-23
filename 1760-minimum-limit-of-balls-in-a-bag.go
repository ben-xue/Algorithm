func minimumSize(nums []int, maxOperations int) int {
	iSum := 0
	iMaxRes := 0
	iLen := len(nums)
	for i := 0; i < iLen; i++ {
		if nums[i] > iMaxRes{
			iMaxRes = nums[i]
		}
		iSum += nums[i]
	}

	iMinRes := 1

	//边界条件
	iMaxOpCnt := iSum - iLen
	if maxOperations >= iMaxOpCnt{
		return iMinRes
	}

	iTmpMaxRes := iMaxRes
	iTmpMinRes := iMinRes
	iRes := iMaxRes

	for true {
		iTmpRes := (iTmpMaxRes + iTmpMinRes) / 2
		if iTmpRes == iTmpMinRes{
			return iRes
		}

		iTmpOpCnt := 0
		for i := 0; i < iLen; i++ {
			if nums[i] > iTmpRes{
				iTmpOpCnt += (nums[i] - 1) / iTmpRes
				if iTmpOpCnt > maxOperations{
					break
				}
			}
		}

		if iTmpOpCnt <= maxOperations{
			//还可以更小
			iTmpMaxRes = iTmpRes
			if iTmpRes < iRes{
				iRes = iTmpRes
			}
		}else{
			//约定操作次数不够,加大最优值
			iTmpMinRes = iTmpRes
		}
	}

	return iRes
}
