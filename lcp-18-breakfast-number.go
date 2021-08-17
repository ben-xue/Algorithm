package main

import (
	"fmt"
	"math"
	"sort"
)

//小于等于val的元素数量
func CntLessThanVal(val int, sortArray []int) int {
	iSortArrLen := len(sortArray)
	if val < sortArray[0] {
		return 0
	}

	iCnt := 0
	if val == sortArray[0] {
		for i := 0; i < iSortArrLen; i++ {
			if val == sortArray[i] {
				iCnt++
			}
		}
		return iCnt
	}

	if val >= sortArray[iSortArrLen-1] {
		return iSortArrLen
	}

	iLeft := 0
	iRight := iSortArrLen - 1
	for val > sortArray[iLeft] && val < sortArray[iRight] {
		iMid := (iLeft + iRight) / 2
		if sortArray[iMid] > val {
			iRight = iMid
			if sortArray[iRight-1] < val {
				return iRight
			}
		} else if sortArray[iMid] == val {
			iIndex := iMid + 1
			for sortArray[iIndex] == val {
				iIndex++
			}
			return iIndex
		} else {
			iLeft = iMid
			if sortArray[iLeft+1] > val {
				return iLeft + 1
			}
		}
	}

	return iCnt
}

func breakfastNumber(staple []int, drinks []int, x int) int {
	var s []int
	for _, val := range staple {
		if val < x {
			s = append(s, val)
		}
	}

	var d []int
	for _, val := range drinks {
		if val < x {
			d = append(d, val)
		}
	}

	sort.Ints(d)

	iCnt := 0
	for _, val := range s {
		iCnt += CntLessThanVal(x-val, d)
	}

	return iCnt % (int(math.Pow10(9)) + 7)
}

/*
//内存最少
func breakfastNumber(staple []int, drinks []int, x int) int {
    sort.Ints(staple)
	sort.Ints(drinks)

	lS := len(staple)
	lD := len(drinks)
	count := 0
	i := 0
	j := lD - 1
	for i < lS && j >= 0 {
		if staple[i]+drinks[j] <= x {
			count = count + j + 1
			i++
		} else {
			j--
		}
	}

	return count % (1e9 + 7)

}
*/

/*
//速度最快: 计数排序
func breakfastNumber(staple []int, drinks []int, x int) int {
	drinkNum := make([]int, x + 1)
	for _, v := range drinks {
		if v < x {
			drinkNum[v]++
		}
	}

	for i := 1; i < len(drinkNum); i++ {
		drinkNum[i] += drinkNum[i - 1]
	}

	var result int
	for _, v := range staple {
		if (x - v) > 0 {
			result += drinkNum[x - v]
		}
	}

	return result % (1000000007)
}
*/

func main() {
	//s := []int{10,20,5}
	//d := []int{5,5,2}
	d := []int{2, 1, 1}
	s := []int{8, 9, 5, 1}
	fmt.Println(breakfastNumber(s, d, 9))
}
