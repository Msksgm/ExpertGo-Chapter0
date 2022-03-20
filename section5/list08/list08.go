package main

import "fmt"

type T struct {
	Number int
}

func main() {
	s := []T{{1}, {2}, {3}, {4}, {5}}
	s2 := []*T{}
	for _, v := range s {
		s2 = append(s2, &v)
	}
	for _, v := range s2 {
		fmt.Printf("%+v\n", v)
	}
}
