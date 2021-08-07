package main

import (
	"fmt"
)

var iConstArrayLen  = 4
var stConstArray = [][]int{{4,0},{3,1},{2,2},{1,3}}
var DEFAULTVALUE = 2.0
var fMatrix [][]float64 = nil

func Calc(x int,y int) float64{
	if x == 0 || y == 0{
		return fMatrix[x][y]
	}

	if fMatrix[x][y] != DEFAULTVALUE {
		return fMatrix[x][y]
	}

	fResult := 0.0
	iTmpx ,iTmpy := 0,0
	for i := 0; i < iConstArrayLen; i++ {
		iTmpx = x-stConstArray[i][0]
		iTmpy = y-stConstArray[i][1]
		if iTmpx < 0{
			iTmpx = 0
		}
		if iTmpy < 0{
			iTmpy = 0
		}
		fResult += Calc(iTmpx,iTmpy) * 0.25
	}

	fMatrix[x][y] = fResult
	return fResult
}

func soupServings(iSerialn int) float64 {
	if iSerialn >= 12500{
		return 1.0
	}
	n := (iSerialn + 25 - 1) / 25
	fMatrix = make([][]float64,n+1)
	for i := range fMatrix {
		fMatrix[i] = make([]float64,n+1)
	}

	for i := 0; i <= n; i++ {
		for k := 0; k <= n; k++ {
			fMatrix[i][k] = DEFAULTVALUE
		}
	}

	fMatrix[0][0] = 0.5
	for i := 1; i <= n; i++ {
		fMatrix[0][i] = 1
		fMatrix[i][0] = 0
	}
	return Calc(n,n)
}

func main()  {
	fmt.Println(soupServings(7000))
}
