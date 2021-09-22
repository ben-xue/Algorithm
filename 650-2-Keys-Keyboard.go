
/*

思考流程
最开始当做了用 2^n 组合的贪心算法,因为没仔细审题以为是每次都是乘2的关系。
仔细审题后发现是除了copy-paste乘2,还可以只patse做加法。

又想到按奇数偶数考虑,奇数则需要自身次操作，偶数一直除2,每次除2操作次数加2,得到奇数再计算。
提交未通过后,观察输入,看到输入9输出的并不是9而是6,立马想到还可以分解成3 * 3.进而意识到并不是奇偶的
问题,而是质数合数的问题.联想到分解质因数，聚了几个例子发现规律。

*/


func minSteps(n int) int {
	iResult := 0
	PrimeList := []int{}

	iN := n
	for i := 2; i*i <= iN; {
		if iN%i == 0 {
			PrimeList = append(PrimeList, i)
			iN /= i
			i = 2
		}else{
			i++
		}
	}

	if iN != 1 {
		PrimeList = append(PrimeList, iN)
	}

	for _, val := range PrimeList {
		iResult += val
	}
	return iResult
}
