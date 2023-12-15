/*
https://leetcode.com/problems/two-sum/

Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.
*/
package main

import "fmt"


func twoSum(nums []int, target int) []int {
    mymap := make(map[int]int)
    //Функция make создает обнуленный массив и возвращает срез, который ссылается на этот массив.
    for i := 0; i < len(nums); i++ {
        j, ok := mymap[target-nums[i]]
        if ok {
            result := []int{j, i}
            return result
        }
        mymap[nums[i]] = i
    }
    result := []int{-1, -1}
    return result
}

func main() {
	res:= twoSum([]int{2,7,11,15},9)
	fmt.Println(res)
}
