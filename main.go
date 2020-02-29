package main

import (
	"fmt"

	utils "github.com/amanda-phan89/go-excercise-1/underscore"
)

func main() {
	u := utils.Underscore{}
	mapping(u)
	filtering(u)
	finding(u)
	rejecting(u)
	every(u)
}

func mapping(u utils.Underscore) {
	arr := []int{1, 2, 3, 4, 8, 9, 10}
	newArr := u.Map(arr, func(val int) int {
		return val * 2
	}).([]int)
	fmt.Printf("Type: %T - Value: %v\n", newArr, newArr)

	arrStr := []string{"a", "b", "c"}
	newArrStr := u.Map(arrStr, func(val string) string {
		return val + " test"
	}).([]string)
	fmt.Printf("Type: %T - Value: %v\n", newArrStr, newArrStr)
}

func filtering(u utils.Underscore) {
	arr := []int{1, 2, 3, 4, 8, 9, 10}
	resFilter := u.Filter(arr, func(el int) bool {
		// return even number
		return el%2 == 0
	})
	fmt.Printf("Type: %T - Value: %v\n", resFilter, resFilter)
}

func finding(u utils.Underscore) {
	arr := []int{1, 3, 3, 4, 8, 9, 10}
	resFinding := u.Find(arr, func(el int) bool {
		return el%2 == 0
	})
	fmt.Printf("Type: %T - Value: %v\n", resFinding, resFinding)
}

func rejecting(u utils.Underscore) {
	arr := []int{1, 2, 3, 4, 8, 9, 10}
	resReject := u.Reject(arr, func(el int) bool {
		// reject even number
		return el%2 == 0
	})
	fmt.Printf("Type: %T - Value: %v\n", resReject, resReject)
}

func every(u utils.Underscore) {
	arr := []int{1, 2, 3, 4, 8, 9, 10}
	result := u.Every(arr, func(el int) bool {
		return el%2 == 0
	})
	if result == true {
		fmt.Printf("All elements are even numbers")
	} else {
		fmt.Printf("Some elements are odd numbers")
	}
}
