package main

import (
	"fmt"
	"maps"
)

func main() {
	// var mapVariableName map[keyType]valueType
	// mapVariableName = make(map[keyType]valueType)

	// using Map Literal
	// mapVariableName := map[keyType]valueType{
	// 	key1: value1,
	// 	key2: value2,
	// }

	myMap := make(map[string]int)
	fmt.Println(myMap)

	myMap["key1"] = 9
	myMap["code"] = 18

	fmt.Println(myMap)
	fmt.Println(myMap["key1"])
	fmt.Println(myMap["key"])

	myMap["code"] = 21
	fmt.Println(myMap)

	delete(myMap, "key1")
	fmt.Println(myMap)

	myMap["key1"] = 10
	myMap["key2"] = 11
	myMap["key3"] = 12
	fmt.Println(myMap)

	//clear(myMap)
	//fmt.Println(myMap)

	value, ok := myMap["key1"]
	fmt.Println("value", value)
	fmt.Println("Is a value associated with key1", ok)

	myMap2 := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println("myMap2:", myMap2)

	myMap3 := map[string]int{"a": 1, "b": 2, "c": 5}
	fmt.Println("myMap2:", myMap3)

	if maps.Equal(myMap2, myMap3) {
		fmt.Println("myMap2 and myMap3 are equal")
	}

	for _, v := range myMap {
		fmt.Println("value:", v)
	}

	myMap4 := make(map[string]string)
	myMap4["name"] = "Juan"
	//if myMap4 == nil {
	//	fmt.Println("myMap4 is nil")
	//} else {
	//	fmt.Println("myMap4 is not nil")
	//}

	myMap7 := make(map[string]string)
	myMap7["name"] = "Juana"

	myMap5 := make(map[string]map[string]string)
	myMap5["student"] = myMap4
	myMap5["teacher"] = myMap7

	fmt.Println("myMap5", myMap5)

	myMap6 := make(map[int]int)
	for i := range 2 {
		for j := range 2 {
			myMap6[i] = j
		}
	}
	fmt.Println("myMap6", myMap6)

	for i, v := range myMap5 {
		fmt.Println("Index:", i)
		fmt.Println("Value:", v)
	}

}
