package main

import (
	"fmt"

	filtr "github.com/intervinn/filter"
)

func main() {
	str := "holy shit @ss bananas"
	nodes := filtr.Parse(str)

	for _, v := range nodes {
		fmt.Println(str[v.Pos:v.End])
	}
}