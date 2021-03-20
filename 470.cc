
/*
已有方法 rand7 可生成 1 到 7 范围内的均匀随机整数，试写一个方法 rand10 生成 1 到 10 范围内的均匀随机整数。

不要使用系统的 Math.random() 方法。
*/

class Solution {
public:
    int rand10() {
    	int col,row,index;
    	do
    	{
    		col = rand7();
    		row = rand7();
    		index = col + (row - 1) * 7;
    	}wihle(index > 40);

        return 1 + index % 10;
    }
};


//do-while、while —— 满足条件继续循环,与当型、直到型条件语句不一样
//1. 随机数转有序数列
//2. 贝叶斯公式
//3. 等比数列求和公式,sn = a1 * (q^n - 1) / (q - 1),(q != 1);  |q| < 1时,sn收敛于 a1 / (1 - q)

//没解出来的原因:贝叶斯公式想到了,1没想到

//拓展: 用randM()模拟randN(),且N > M*M呢？没关系,那就用M*M*M模拟,这样期望值更小。


