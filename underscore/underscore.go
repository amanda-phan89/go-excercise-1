package underscore

import "reflect"

// Underscore Object
type Underscore struct{}

// Map Produces a new slice of values by mapping each value in list through a transformation function
func (u Underscore) Map(arr interface{}, fn interface{}) interface{} {
	valueFunc := reflect.ValueOf(fn)
	valueArr := reflect.ValueOf(arr)

	typeOfResult := reflect.SliceOf(valueFunc.Type().Out(0))
	arrResultMap := reflect.MakeSlice(typeOfResult, valueArr.Len(), valueArr.Len())

	for i := 0; i < valueArr.Len(); i++ {
		funcResult := valueFunc.Call([]reflect.Value{valueArr.Index(i)})[0]
		arrResultMap.Index(i).Set(funcResult)
	}
	return arrResultMap.Interface()
}

// Filter Looks through each value in the list, returning a slice of all the values that pass a truth test
func (u Underscore) Filter(arr interface{}, fn interface{}) interface{} {
	valueFunc := reflect.ValueOf(fn)
	valueArr := reflect.ValueOf(arr)

	typeOfFuncResult := reflect.SliceOf(valueArr.Index(0).Type())
	arrResultMap := reflect.MakeSlice(typeOfFuncResult, 0, valueArr.Len())

	for i := 0; i < valueArr.Len(); i++ {
		funcResult := valueFunc.Call([]reflect.Value{valueArr.Index(i)})[0]
		if funcResult.Interface().(bool) == true {
			arrResultMap = reflect.Append(arrResultMap, valueArr.Index(i))
		}
	}
	return arrResultMap.Interface()
}

// Find Looks through each value in the list, returning the first one that passes a truth test
func (u Underscore) Find(arr interface{}, fn interface{}) interface{} {
	valueFunc := reflect.ValueOf(fn)
	valueArr := reflect.ValueOf(arr)

	result := reflect.Zero(valueArr.Index(0).Type())

	for i := 0; i < valueArr.Len(); i++ {
		funcResult := valueFunc.Call([]reflect.Value{valueArr.Index(i)})[0]
		if funcResult.Interface().(bool) == true {
			result = valueArr.Index(i)
			break
		}
	}
	return result.Interface()
}
