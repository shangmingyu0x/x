package main

import "fmt"

func main() {
	v1 := Value{Name: "煎鱼", Gender: "男"}
	v2 := Value{Name: "煎鱼", Gender: "男"}

	if v1 == v2 {
		fmt.Println("11111")
	} else {
		fmt.Println("2222")
	}
}

type Value struct {
	Name   string
	Gender string
}
