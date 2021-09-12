type CountEntry struct {
	iVal int
	iUnitIndex int
}

func MinNumToString(iMinNum int) string {
	stMinNumMap := map[int]string{
		1:  "One ",
		2:  "Two ",
		3:  "Three ",
		4:  "Four ",
		5:  "Five ",
		6:  "Six ",
		7:  "Seven ",
		8:  "Eight ",
		9:  "Nine ",
		10: "Ten ",
		11: "Eleven ",
		12: "Twelve ",
		13: "Thirteen ",
		14: "Fourteen ",
		15: "Fifteen ",
		16: "Sixteen ",
		17: "Seventeen ",
		18: "Eighteen ",
		19: "Nineteen "}

	stLessThanHundredDivMap := map[int]string{
		2: "Twenty ",
		3: "Thirty ",
		4: "Forty ",
		5: "Fifty ",
		6: "Sixty ",
		7: "Seventy ",
		8: "Eighty ",
		9: "Ninety "}

	strResult := ""

	iHundredUnitVal := 100
	iTmpMinNum := iMinNum
	if iTmpMinNum > 0{
		val,_ok := stMinNumMap[iTmpMinNum / iHundredUnitVal]
		if _ok{
			strResult += val + "Hundred "
		}
		iTmpMinNum %= iHundredUnitVal
	}

	iSpecialNum := 20

	iTenUnitVal := 10
	if iTmpMinNum >= iSpecialNum{
		val,_ok := stLessThanHundredDivMap[iTmpMinNum / iTenUnitVal]
		if _ok{
			strResult += val
		}
		iTmpMinNum %= iTenUnitVal
	}

	val,_ok := stMinNumMap[iTmpMinNum]
	if _ok{
		strResult += val
	}

	return strResult
}

func numberToWords(num int) string{
	if 0 == num{
		return "Zero"
	}

	stDivMap := map[int]string{
		0: "",
		1: "Thousand ",
		2: "Million ",
		3: "Billion ",
		4: "Trillion "}

	// 3个一组切分数字
	// 2 147 483 647
	stStack := make([]CountEntry,0)
	iNum := num
	iIndex := 0

	iThousandUnitVal := 1000
	for iNum != 0{
		stStack = append(stStack, CountEntry{iNum % iThousandUnitVal,iIndex})
		iNum /= iThousandUnitVal
		iIndex++
	}

	strResult := ""
	for i := range stStack {
		val := stStack[len(stStack) - i - 1]

		if val.iVal > 0{
			strResult += MinNumToString(val.iVal)
			mapVal ,_ok := stDivMap[val.iUnitIndex]
			if _ok{
				strResult += mapVal
			}
		}
	}

	return strings.TrimSpace(strResult)
}