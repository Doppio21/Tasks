package main

import (
	"fmt"
	"reflect"
)

// Разработать программу, которая в рантайме способна определить
// тип переменной: int, string, bool, channel из переменной типа interface{}.

func v1(v interface{}) string {
	switch v.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan string:
		return "channel string"
	default:
		return "unknown"
	}
}

func v2(v interface{}) string {
	return reflect.TypeOf(v).String()
}

func main() {
	var types = []interface{}{string("str"), int(1), true, make(chan string)}
	for _, t := range types {
		fmt.Printf("v1: %s v2: %s\n", v1(t), v2(t))
	}
}
