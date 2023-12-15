package main

//https://leetcode.com/problems/contains-duplicate

func containsDuplicate(nums []int) bool {
	tmp := make(map[int]int)

	for _, v := range nums {
		tmp[v]++

		if tmp[v] > 1 {
			return true
		}
	}
	return false
}

func main() {
	var arr = []int{2, 3, 5, 6, 11, 13}
	print(containsDuplicate(arr))
}
