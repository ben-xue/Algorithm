package main

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
)

/*
 	对于k,有 n * n >= k
	求n的最小值
*/

func kthSmallest1(matrix [][]int, k int) int {
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

type Node struct { 
	x,y int
	val int
}

type IntHeapWithIndex []Node

func (h IntHeapWithIndex) Len() int           { return len(h) }
func (h IntHeapWithIndex) Less(i, j int) bool { return h[i].val < h[j].val }
func (h IntHeapWithIndex) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeapWithIndex) Push(x interface{}) {
	*h = append(*h, x.(Node))
}

func (h *IntHeapWithIndex) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kthSmallest(matrix [][]int, k int) int {
	iN := len(matrix)
	stHeap := IntHeapWithIndex{}
	heap.Init(&stHeap)

	stFlagmatrix := make([][]bool,iN)
	for i := 0; i < iN; i++ {
		stFlagmatrix[i] = make([]bool,iN)
		for n := 0; n < iN; n++ {
			stFlagmatrix[i][n] = false
		}
	}

	iIndxK := 0
	heap.Push(&stHeap,Node{0,0,matrix[0][0]})
	stFlagmatrix[0][0] = true

	stVal := Node{}
	for stHeap.Len() > 0 {
		stVal = heap.Pop(&stHeap).(Node)
		iIndxK++
		if iIndxK >= k{
			return stVal.val
		}

		if stVal.x + 1 < iN && stFlagmatrix[stVal.x + 1][stVal.y] == false{
			heap.Push(&stHeap,Node{stVal.x+1,stVal.y,matrix[stVal.x + 1][stVal.y]})
			stFlagmatrix[stVal.x + 1][stVal.y] = true
		}

		if stVal.y + 1 < iN && stFlagmatrix[stVal.x][stVal.y + 1] == false{
			heap.Push(&stHeap,Node{stVal.x,stVal.y+1,matrix[stVal.x][stVal.y+1]})
			stFlagmatrix[stVal.x][stVal.y + 1] = true
		}

	}

	return 0
}

func main() {
	//matrix := [][]int{{1,5,9},{10,11,13},{12,13,15}}
	matrix := [][]int{{1,3,5},{6,7,12},{11,14,14}}
	//matrix := [][]int{{1,2},{1,3}}
	iRet := kthSmallest(matrix,3)
	fmt.Println(iRet)
}
