package main

import (
	"fmt"
	"github.com/gymnasy55/polcraz-go-lessons/002_package/nested/hello"
	"github.com/gymnasy55/polcraz-go-lessons/002_package/nested/say"
)

func main() {
	fmt.Println("Start program...")
	fmt.Println(hello.CallFromHello() + say.CallFromSay())
}
