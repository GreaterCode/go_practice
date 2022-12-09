package main

import "fmt"

func main() {
	s := "hello, world!"
	s1 := s[:5]
	s2 := s[7:]
	s3 := s[1:5]
	fmt.Println(s, s1, s2, s3)

}
