package main

import (
	"fmt"
	"github.com/Kambar-ZH/Golang_Onelab/utils"
)

func main() {
	fmt.Println(utils.Upper("Kambar"))
	fmt.Println(utils.Lower("Kambar"))

	fmt.Println(utils.Square(100))

	fmt.Println(utils.GetUUID())
}