package cmd

import "fmt"

/*
type slice struct {
    array unsafe.Pointer
    len   int
    cap   int
}
*/

func main() {
	s1 := make([]int, 3, 4)
	s2 := s1[1:3]

	fmt.Println(s1) // 0 0 0
	fmt.Println(s2) // 0 0

	s1[1] = 1
	fmt.Println("Changed")
	fmt.Println(s1) // 0 1 0
	fmt.Println(s2) // 1 0

	s2 = append(s2, 2)
	fmt.Println("Changed")
	fmt.Println(s1) // 0 1 0
	fmt.Println(s2) // 1 0 2

	s1 = append(s1, 3)
	fmt.Println("Changed")
	fmt.Println(s1) // 0 1 0 3
	fmt.Println(s2) // 1 0 3

	s2 = append(s2, 4)
	fmt.Println("Changed")
	fmt.Println(s1) // 0 1 0 3
	fmt.Println(s2) // 1 0 3 4

	s1[2] = 5
	fmt.Println("Changed")
	fmt.Println(s1) // 0 1 5 3
	fmt.Println(s2) // 1 0 3 4
}
