package main

import (
	"fmt"
	"reflect"
)

func main() {
	//Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.
	fmt.Println(GetTypeBySwitchAssertion(42)) //v1
	fmt.Println(GetTypeByTypeAssertion("Word")) //v2
	fmt.Println(GetTypeByRefelction(true)) //v3
	fmt.Println(GetTypeBySwitchAssertion(make(chan string)))
	fmt.Println(GetTypeByTypeAssertion(3.14)) //undefined type
}

func GetTypeBySwitchAssertion(i interface{}) (string) {
	switch i.(type) {
	case int:
		return fmt.Sprintf("Type = %T, val = %v", i, i)
	case string:
		return fmt.Sprintf("Type = %T, val = %v", i, i)
	case bool:
		return fmt.Sprintf("Type = %T, val = %v", i, i)
	case chan int, chan bool, chan string:
		return fmt.Sprintf("Type = %T, val = %v", i, i)
	default:
		return "Undedined type"
	}
}

func GetTypeByTypeAssertion(i interface{}) (string) {
	if _, ok := i.(int); ok {
		return "int"
	}
	if _, ok := i.(string); ok {
		return "string"
	}
	if _, ok := i.(bool); ok {
		return "bool"
	}
	if _, ok := i.(chan int); ok {
		return "chan int"
	}
	if _, ok := i.(chan string); ok {
		return "chan string"
	}
	if _, ok := i.(chan bool); ok {
		return "chan bool"
	}

	return "Undefined type"
}

func GetTypeByRefelction(i interface{}) (string) { //by reflect. But there will be processed any type without additional (maybe useless) checks
	return reflect.TypeOf(i).String()
}
