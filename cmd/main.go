package main

import (
	"fmt"
	"strconv"

	linkedlist "github.com/Sotnasjeff/neet-datastructure-and-algorithm-for-beginner/linked_list"
)

func main() {
	//array1 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	//array2 := []int{1, 1, 2}
	//array3 := []int{0, 1, 2, 2, 3, 0, 4, 2}

	//fmt.Println(getConcatenation(array1))

	//fmt.Println(removeElement(array1, 3)) //8, {0,0,1,1,1,2,2,4}
	//fmt.Println(removeElement(array2, 1)) //1, {2,1,1}
	//fmt.Println(removeElement(array3, 2)) //5, {0,1,3,0,4}

	l := linkedlist.NewLinkedList()
	l.InsertHead(10)
	l.InsertTail(5)
	l.InsertHead(6)

	fmt.Println(l.GetValues())
	fmt.Println(l.Get(1))

	ops := []string{"5", "2", "C", "D", "+"}
	fmt.Println(calPoints(ops))
	fmt.Println(isAnagram("cacc", "ccaa"))

	s := "({}"

	arr := []int{17, 18, 5, 4, 6, 1}

	fmt.Println(isValid(s))
	fmt.Println(replaceElements(arr))
	fmt.Println(isSubsequence("b", "abc"))

	newNode := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	newNode2 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	fmt.Println(mergeTwoLists(&newNode, &newNode2))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	newLinkedList := &ListNode{}
	tail := newLinkedList
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}

		tail = tail.Next
	}

	if list1 == nil {
		tail.Next = list2
	}

	if list2 == nil {
		tail.Next = list1
	}

	return newLinkedList.Next
}

func reverseList(head *ListNode) *ListNode {
	curr := head
	var prev *ListNode
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}

func isSubsequence(s string, t string) bool {
	array := []rune{}
	for _, v := range s {
		array = append(array, v)
	}

	j := len(array) - 1

	for i := len(t) - 1; i > -1; i-- {
		if len(array) != 0 && array[j] == rune(t[i]) {
			array = array[:len(array)-1]
			j = len(array) - 1
		}
	}

	return len(array) == 0
}

func replaceElements(arr []int) []int {
	lastElement := -1

	for i := len(arr) - 1; i > -1; i-- {
		greaterValue := max(lastElement, arr[i])
		arr[i] = lastElement
		lastElement = greaterValue
	}

	return arr
}

func isAnagram(s string, t string) bool {
	hashMap := make(map[rune]int)

	if len(s) != len(t) {
		return false
	}

	for _, v := range s {
		hashMap[v]++
	}

	for _, v := range t {
		hashMap[v]--
		if hashMap[v] < 0 {
			return false
		}
	}

	return true
}

func isValid(s string) bool {
	hashMap := map[rune]rune{
		'{': '}',
		'[': ']',
		'(': ')',
	}

	if len(s) == 1 {
		return false
	}

	array := []rune{}
	for _, v := range s {
		if _, b := hashMap[v]; b {
			array = append(array, v)
		} else if len(array) == 0 || hashMap[array[len(array)-1]] != v {
			return false
		} else {
			array = array[:len(array)-1]
		}

	}

	return len(array) == 0

}

func calPoints(operations []string) int {
	var result []int
	var total int

	for _, v := range operations {
		if v != "+" && v != "D" && v != "C" {
			r, _ := strconv.Atoi(v)
			result = append(result, r)
		}

		if v == "+" {
			penultimateNumber := result[len(result)-1]
			newValue := (penultimateNumber + result[len(result)-2])
			result = append(result, newValue)
		}

		if v == "D" {
			doubleValue := (result[len(result)-1] * 2)
			result = append(result, doubleValue)
		}

		if v == "C" {
			result = result[:len(result)-1]
		}
	}

	for _, v := range result {
		total += v
	}

	return total
}

func getConcatenation(nums []int) []int {
	newArray := nums
	for _, v := range newArray {
		nums = append(nums, v)
	}

	return nums
}

func shuffleSavingMemory(nums []int, n int) []int {
	newArray := []int{}
	for i := 0; i < n; i++ {
		newArray = append(newArray, nums[i], nums[n+i])
	}

	return newArray
}

func shuffle(nums []int, n int) []int {
	newArray := []int{}
	var counter int
	values, antivalues := nums[:n], nums[n:]
	for _, v := range values {
		newArray = append(newArray, v, antivalues[counter])
		counter++
	}

	return newArray
}

func removeElement(nums []int, val int) int {
	i := 0
	for _, v := range nums {
		if v != val {
			nums[i] = v
			i++
		}
	}
	return i
}

func removeDuplicates(nums []int) int {
	aux := make(map[int]bool)
	var counter int
	for _, v := range nums {
		if aux[v] != true {
			nums[counter] = v
			counter++
		}
		aux[v] = true
	}
	return counter
}
