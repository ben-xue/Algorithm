func IsPalindromeTimeout(num uint64) bool {
	bitNums := []int{}
	tmp := num
	for tmp != 0 {
		bitNums = append(bitNums, int(tmp%10))
		tmp = tmp / 10
	}

	iStart := 0
	iEnd := len(bitNums) - 1

	bRes := true
	for iStart <= iEnd {
		if bitNums[iStart] != bitNums[iEnd] {
			bRes = false
			break
		}

		iStart++
		iEnd--
	}
	return bRes
}

func largestPalindromeTimeout(n int) int {
	iMaxRes := uint64(0)
	iMax := int(math.Pow(10, float64(n)))
	iMin := int(math.Pow(10, float64(n-1)))
	for i := iMax - 1; i >= iMin; i-- {
		for k := iMax - 1; k >= iMin; k-- {
			var tmp uint64 = uint64(i * k)
			if tmp > iMaxRes {
				if IsPalindromeTimeout(tmp) {
					iMaxRes = tmp
				}
			}
		}
	}
	return int(iMaxRes % 1337)
}

//————————————————————————————————————————————————————————
func ReverseStr(str string) string {
	runes := []rune(str)

	for from, to := 0, len(str)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}

	return string(runes)
}

func largestPalindrome(n int) int {
	if n == 1{
		return 9
	}

	iMax := int64(math.Pow(10, float64(n)))
	iMin := iMax / 10

	iMultiMax := int64((iMax - 1) * (iMax - 1))
	for i := iMax - 1; i >= iMin; i-- {
		//构造回文数
		head := strconv.FormatInt(i,10)
		tail := ReverseStr(head)
		tmpNum, _ := strconv.ParseInt(head+tail, 10, 64)

		if iMultiMax < tmpNum {
			continue
		}
		for k := int64(iMax - 1); k*k >= tmpNum; k-- {
			if tmpNum%k == 0 {
				return int(tmpNum % 1337)
			}
		}
	}
	return 0
}
