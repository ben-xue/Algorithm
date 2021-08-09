
func specialArray(nums []int) int {
	iLeft ,iRight := 0,0
	for _,v := range nums {
		if v > iRight{
			iRight = v
		}
	}

	iLeftCnt,iRightCnt := 0,0
	for _,v := range nums {
		if v >= iLeft{
			iLeftCnt++
		}
		if v >= iRight{
			iRight++
		}
	}
	if iLeftCnt == iLeft{
		return iLeft
	}
	if iRightCnt == iRight{
		return iRight
	}

	for iLeft < iRight {
		iMid := (iLeft + iRight) / 2

		iCnt := 0
		for _,v := range nums {
			if v >= iMid{
				iCnt++
			}
		}

		if iCnt == iMid{
			return iMid
		}else if iCnt < iMid {
			if iRight <= iMid{
				return -1
			}
			iRight = iMid
		}else{
			if iLeft >= iMid{
				return -1
			}
			iLeft = iMid
		}
	}

	return -1
}
