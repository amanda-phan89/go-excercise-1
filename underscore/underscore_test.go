package underscore

import (
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	// Given
	arr := []int{1, 2, 3, 4, 8, 9, 10}

	// When
	actual := Map(arr, func(val int) int {
		return val * 2
	})

	// Then
	expected := []int{2, 4, 6, 8, 16, 18, 20}
	if reflect.DeepEqual(expected, actual) == false {
		t.Error("Failed")
	}
}

func TestFilter(t *testing.T) {
	// Given
	arr := []int{1, 2, 3, 4, 8, 9, 10}

	// When
	actual := Filter(arr, func(el int) bool {
		// return odd number
		return el%2 != 0
	})

	// Then
	expected := []int{1, 3, 9}
	if reflect.DeepEqual(expected, actual) == false {
		t.Error("Failed")
	}
}

func TestFind(t *testing.T) {
	// Given
	arr := []int{1, 2, 3, 4, 8, 9, 10}

	// When
	actual := Find(arr, func(el int) bool {
		// return odd number
		return el%2 != 0
	}).(int)

	// Then
	if actual != 1 {
		t.Error("Failed")
	}
}

func TestReject(t *testing.T) {
	// Given
	arr := []int{1, 2, 3, 4, 8, 9, 10}

	// When
	actual := Reject(arr, func(el int) bool {
		// reject odd number
		return el%2 != 0
	})

	// Then
	expected := []int{2, 4, 8, 10}
	if reflect.DeepEqual(expected, actual) == false {
		t.Error("Failed")
	}
}

func TestEvery(t *testing.T) {
	// Given
	arr := []int{1, 3, 5, 7, 9}

	// When
	actual := Every(arr, func(el int) bool {
		// check odd number
		return el%2 != 0
	})

	// Then
	if actual == false {
		t.Error("Failed")
	}
}
