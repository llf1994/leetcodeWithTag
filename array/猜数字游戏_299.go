package main

import "fmt"

//题解，使用两个桶，时间复杂度为o(n)+o(10)=>o(n)，空间复杂度为o(10)=>o(1)
func getHint(secret string, guess string) string {
	bulls, cows, buckets1, buckets2, length := 0, 0, [10]int{}, [10]int{}, len(secret)
	//遍历数组，寻找公牛，并将非公牛的数分别放入两个桶中
	for i := 0; i < length; i++ {
		//如果是公牛
		if s, g := secret[i], guess[i]; s == g {
			bulls++
			//否则使用桶分别记录对应数字出现的次数
		} else {
			//通过下标s[i]取得的值为byte(ASCII字符)，这里将 byte -> ascii对应的值 -> int
			//s, _ := strconv.Atoi(string(s))
			//g, _ := strconv.Atoi(string(g))

			//也可直接 - '0'，'0'的ascii为48， s - '0' 后的值即为对应的数字(适用0~9的数字)
			buckets1[s - '0']++
			buckets2[g - '0']++
		}
	}
	//遍历长度为10的桶，两个桶均不为零时，其较小值为母牛
	for j := 0; j < 10; j++{
		if buckets1[j] > buckets2[j]{
			cows += buckets2[j]
		} else {
			cows += buckets1[j]
		}
	}
	//go不允许隐式转换
	return fmt.Sprintf("%dA%dB", bulls, cows)
}

//第一次通过答案，时间复杂度为O(n^2)，空间复杂度为o(n)
func getHint1(secret string, guess string) string {
	//A为公牛，B为母牛，tmp保存检测过的secret字符
	A, B, tmp := 0, 0, map[int]string{}
	for i := range guess {
		if guess[i] == secret[i] {
			A++
			//此时索引为i的secret字符使用过，加入map中
			tmp[i] = "deleted"
			continue
		}
		for j := range secret {
			//如果索引为j的secret字符未使用过，且该索引上的secret和guess不相等（相等时在之后的遍历中可找到），
			//寻找secret[j] == guess[i]的j加入map，找到时循环结束
			if _, ok := tmp[j]; ok == false && secret[j] != guess[j] && secret[j] == guess[i] {
				B++
				tmp[j] = "deleted"
				break
			}
		}
	}
	hint := fmt.Sprintf("%dA%dB", A, B)
	return hint
}

/* java，使用一个桶
https://leetcode.com/problems/bulls-and-cows/discuss/74621/One-pass-Java-solution
public String getHint(String secret, String guess) {
    int bulls = 0;
    int cows = 0;
    int[] numbers = new int[10];
    for (int i = 0; i<secret.length(); i++) {
        int s = Character.getNumericValue(secret.charAt(i));
        int g = Character.getNumericValue(guess.charAt(i));
        if (s == g) bulls++;
        else {
			//开始桶内值均为0，使用正负来标识secret和guess值；
			//numbers[s]为选中的实际值，当 numbers[s]<0 时，由于只有numbers[g] --，说明这个值之前猜到了，cows++
			//numbers[g]为猜测值，当 numbers[g] > 0 时，由于只有numbers[s] ++，说明这个值之前出现过，cows++
            if (numbers[s] < 0) cows++;
            if (numbers[g] > 0) cows++;
			//记录值，当实际出现时++，猜到时--
            numbers[s] ++;
            numbers[g] --;
        }
    }
    return bulls + "A" + cows + "B";
}
*/
