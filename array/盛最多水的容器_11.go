package main

//题解，头尾指针，跳过短板状态,o(n)
func maxArea(height []int) int {
	left, right, value := 0, len(height)-1, 0
	for left < right {
		//left为短板，移动left，因为：
		//(1)当left为短板时，底L=right-left,高H=height[left],此时容器面积 S0=L*H
		//(2)若移动长板，则容器高h <= H，此时底 l = right - left < L
		//(3)则面积 S=l*h < S0=L*H 恒成立，无需计算长板移动后的面积
		if height[left] < height[right]{
			if area := (right - left)*height[left]; area > value { value = area}
			left++
		//否则移动right
		} else {
			if area := (right - left)*height[right]; area > value { value = area}
			right--
		}
	}
	return value
}

//暴力法，o(n^2)
func maxArea_(height []int) int {
	value := 0
	for i, v := range height{
		for j:=i+1;j<len(height);j++{
			min := v
			if height[j] < v {min=height[j]}
			area := (j-i)*min
			if area > value{value = area}
		}
	}
	return value
}