package main
import (
	"fmt"
)
type ListNode struct {
	Val int
	*Next *ListNode
}

func main()  {
	nums := []int{1,2,3}
	target := 3
	i := twoSum(nums,target)
	fmt.Println(i)
}
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	
	return l1
}

func twoSum(nums []int, target int) []int {
	hashTable := make(map[int]int)
	for i, i2 := range nums {
		if ii,ok := hashTable[target - i2];ok {
			return []int{ii,i}
		}
		hashTable[i2] = i
	}
	return nil
}
func twoSum1(nums []int, target int) []int {
	c := make([]int,2)
	for inx, val := range nums {
		c[0] = inx
		for inx2, val2 := range nums {
			if v := val+val2; inx!=inx2 && v == target {
				c[1] = inx2
				return c
			}
		}
	}
	return nil
}