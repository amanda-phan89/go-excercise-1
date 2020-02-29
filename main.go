package main

import (
	"fmt"

	"github.com/amanda-phan89/go-excercise-1/underscore"
)

func main() {
	mapping()
	filtering()
	finding()
	rejecting()
	every()
}

func mapping() {
	arr := []int{1, 2, 3, 4, 8, 9, 10}
	newArr := underscore.Map(arr, func(val int) int {
		return val * 2
	})
	fmt.Printf("Type: %T - Value: %v\n", newArr, newArr)

	arrStr := []string{"a", "b", "c"}
	newArrStr := underscore.Map(arrStr, func(val string) string {
		return val + " test"
	})
	fmt.Printf("Type: %T - Value: %v\n", newArrStr, newArrStr)
}

func filtering() {
	arr := []int{1, 2, 3, 4, 8, 9, 10}
	resFilter := underscore.Filter(arr, func(el int) bool {
		// return even number
		return el%2 == 0
	})
	fmt.Printf("Type: %T - Value: %v\n", resFilter, resFilter)
}

func finding() {
	arr := []int{1, 3, 3, 4, 8, 9, 10}
	resFinding := underscore.Find(arr, func(el int) bool {
		return el%2 == 0
	})
	fmt.Printf("Type: %T - Value: %v\n", resFinding, resFinding)
}

func rejecting() {
	arr := []int{1, 2, 3, 4, 8, 9, 10}
	resReject := underscore.Reject(arr, func(el int) bool {
		// reject even number
		return el%2 == 0
	})
	fmt.Printf("Type: %T - Value: %v\n", resReject, resReject)
}

func every() {
	arr := []int{1, 2, 3, 4, 8, 9, 10}
	result := underscore.Every(arr, func(el int) bool {
		return el%2 == 0
	})
	if result == true {
		fmt.Printf("All elements are even numbers")
	} else {
		fmt.Printf("Some elements are odd numbers")
	}
}
