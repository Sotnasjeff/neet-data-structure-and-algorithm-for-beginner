package main

import (
	"fmt"
	"strconv"
)

func main() {
	//array1 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	//array2 := []int{1, 1, 2}
	//array3 := []int{0, 1, 2, 2, 3, 0, 4, 2}

	//fmt.Println(getConcatenation(array1))

	//fmt.Println(removeElement(array1, 3)) //8, {0,0,1,1,1,2,2,4}
	//fmt.Println(removeElement(array2, 1)) //1, {2,1,1}
	//fmt.Println(removeElement(array3, 2)) //5, {0,1,3,0,4}

	ops := []string{"5", "2", "C", "D", "+"}
	fmt.Println(calPoints(ops))

	s := "{[]}"

	fmt.Println(isValid(s))
}

func isValid(s string) bool {
	return false
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
