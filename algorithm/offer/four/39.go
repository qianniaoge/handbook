package four

/**
数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。
你可以假设数组是非空的，并且给定的数组总是存在多数元素。
*/

func Hand39(nums []int) int {
	// 投票算法
	var itm int
	cnt := 0
	for _, num := range nums {
		if cnt == 0 {
			itm = num
		}
		if num == itm {
			cnt++
		} else {
			cnt--
		}
	}
	return itm
}
