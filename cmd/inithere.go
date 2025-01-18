package cmd

import "fmt"

func init() {
	fmt.Println("Я должна быть вызвана первоочередно!")
}

func Haha() {
	fmt.Println("HAHAHA")
}
