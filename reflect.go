package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := 42
	v := reflect.ValueOf(x)
	t := v.Type()

	fmt.Println("Value", v)
	fmt.Println("Type", t)
	fmt.Println("Kind", t.Kind())
	fmt.Println("Is Int", t.Kind() == reflect.Int)
	fmt.Println("Is string", t.Kind() == reflect.String)
	fmt.Println("Is zero", v.IsZero())

	y := 10
	v1 := reflect.ValueOf(&y).Elem()
	v2 := reflect.ValueOf(&y)
	fmt.Println("v2 type", v2.Type())
	fmt.Println("Original value", v1.Int())

	v1.SetInt(18)
	fmt.Println("Modified value", v1.Int())

	var itf interface{} = "Hello"
	v3 := reflect.ValueOf(itf)

	fmt.Println("V3 Type", v3.Type())
	if v3.Kind() == reflect.String {
		fmt.Println("V3 is string", v3.String())
	}
}
