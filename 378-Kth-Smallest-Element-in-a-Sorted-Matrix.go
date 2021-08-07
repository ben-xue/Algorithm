package main

import (
	"fmt"
	"math"
	"sort"
)

/*
 	对于k,有 n * n >= k
	求n的最小值
*/

func kthSmallest(matrix [][]int, k int) int {
	iSize := len(matrix)

	iMinSize := 0
	sqrtK := int(math.Sqrt(float64(k)))
	if sqrtK * sqrtK == k{
		iMinSize = sqrtK
	}else{
		iMinSize = sqrtK + 1
	}

	if iMinSize < iSize{
		iSize = iMinSize
	}

	iMaxVal := matrix[iSize - 1][iSize - 1]

	stMap := make(map[int]int)
	for i := 0; i < iSize; i++ {
		for n := 0; n < iSize; n++ {
			_,_ok := stMap[matrix[i][n]]
			if _ok{
				stMap[matrix[i][n]]++
			}else{
				stMap[matrix[i][n]] = 1
			}
		}
	}

	keys := []int{}
	for key,_ := range stMap {
		keys = append(keys, key )
	}
	sort.Ints(keys)

	iTmpRank := 0
	for _, key := range keys {
		iTmpRank += stMap[key]
		if iTmpRank >= k{
			return key
		}
	}

	return iMaxVal
}

func main() {
	//matrix := [][]int{{1,5,9},{10,11,13},{12,13,15}}
	matrix := [][]int{{1,3,5},{6,7,12},{11,14,14}}

	iRet := kthSmallest(matrix,3)
	fmt.Println(iRet)
}
