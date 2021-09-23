/*

*/

type Pair struct {
	iStart, iEnd int
}

func partitionLabels(s string) []int {
	iStrLen := len(s)
	stCharSlice := []Pair{}
	stResult := []int{}

	var i byte
	for i = 'a'; i <= 'z'; i++ {
		bFlag := false
		iTmpStart, iTmpEnd := 0, 0
		for k := 0; k < iStrLen; k++ {
			if i == s[k] {
				bFlag = true
				iTmpStart = k
				break
			}
		}

		for k := iStrLen - 1; k >= 0; k-- {
			if i == s[k] {
				bFlag = true
				iTmpEnd = k
				break
			}
		}

		if bFlag {
			stCharSlice = append(stCharSlice, Pair{iTmpStart, iTmpEnd})
		}
	}

	sort.Slice(stCharSlice, func(i, j int) bool {
		return stCharSlice[i].iStart < stCharSlice[j].iStart
	})

	iSliceLen := len(stCharSlice)
	f := stCharSlice[0]
	for i := 1; i < iSliceLen; i++{
		val := stCharSlice[i]
		if val.iStart < f.iEnd {
			if val.iEnd > f.iEnd {
				f.iEnd = val.iEnd
			} else {
				continue
			}
		} else {
			stResult = append(stResult, f.iEnd-f.iStart+1)
			f = val
		}
	}

	stResult = append(stResult, f.iEnd-f.iStart+1)

	return stResult
}
