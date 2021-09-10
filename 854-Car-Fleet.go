/*
设
	Sx:x车的位置、Vx：x车的速度,St:target的位置
证明：
任意三个车 a、b、c,且Sa < Sb < Sc.若b可以追上c且Va > Vc，那么a是否追得上c与b无关。
a本来可以追的上c，结果让b挡了导致没追上。有没有这种情况？

答：没有。因为b本来就可以追上c,a被b挡了，Va变为Vb，仍然可以追上。
所以，b可以追上c的情况下，a可以追上c 等价于 (St - Sa)/Va <= (St - Sc)/Vc。

*/

func carFleet(target int, postion []int, speed []int) int {
	iPosNum := len(postion)
	s := make([][2]float64,iPosNum)	//一个slice，元素类型为 [2]float64

	for i := 0; i < iPosNum; i++ {
		s[i][0] = float64(speed[i])
		s[i][1] = float64(target - postion[i])
	}

	// i,j是数组下标
	sort.Slice(s, func(i, j int) bool {
		return s[i][1] < s[j][1]
	})

	iResult := 1
	i:= 0
	j := i+1
	for j < iPosNum{
		if s[j][1] / s[j][0] <= s[i][1] / s[i][0]{
			j++
		}else{
			i = j
			j = i + 1
			iResult++
		}
	}
	return iResult
}
