package main

import "fmt"

func minSubArrayLen(target int, nums []int) int {
	iLen := len(nums)
	iResult := iLen + 1
	iSum, iStart, iEnd := 0, 0, 0
	for {
		if iSum < target {
			if iEnd < iLen {
				iSum += nums[iEnd]
				iEnd++
			} else {
				break
			}
		} else {
			iTmpLen := iEnd - iStart
			if iTmpLen < iResult {
				iResult = iTmpLen
			}
			iSum -= nums[iStart]
			iStart++
		}
	}

	if iResult > iLen {
		return 0
	}
	return iResult
}

func minSubArrayLen1(target int, nums []int) int {
	iLen := len(nums)

	for i := 1; i <= iLen; i++ {
		iStart, iEnd := 0, 0
		iSum := 0
		for k := iStart; k < iStart+i; k++ {
			iSum += nums[k]
		}

		if iSum >= target {
			return i
		}

		iStart = 0
		iEnd = iStart + i
		for iEnd < iLen {
			iSum += nums[iEnd]
			iSum -= nums[iStart]

			if iSum >= target {
				return i
			}

			iStart++
			iEnd++
		}
	}

	return 0
}

func main() {
	//list := []int{1, 4, 4}
	list := []int{1, 4, 4}
	fmt.Println(minSubArrayLen(4, list))
}
