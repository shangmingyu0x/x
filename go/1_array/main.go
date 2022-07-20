package main

import "fmt"

func main() {
	s1 := [5]int{1, 2, 3, 4, 5}
	s2 := s1[1:4]
	// s3 := s2[2:5]
	fmt.Println(s1, s2)

	// fmt.Println(s1, s2, s3)

	s22 := s1[1:2]
	s33 := s22[2:3]
	fmt.Println(s22, s33)
}
