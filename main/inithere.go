package main

import "fmt"

func init() {
	fmt.Println("Я должна быть вызвана первоочередно!")
}

func haha() {
	fmt.Println("HAHAHA")
}
