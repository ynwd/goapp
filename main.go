package main

import (
	"embed"
	"fmt"
)

//go:embed tes.xml
var templateDir embed.FS

func GetName() string {
	b, _ := templateDir.ReadFile("tes.xml")
	return string(b)
}

func main() {
	fmt.Println(GetName())
}
