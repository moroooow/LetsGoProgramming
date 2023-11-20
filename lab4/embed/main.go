package main

import (
	"embed"
)

//go:embed testString.txt
var fileString string

//go:embed testByte.txt
var fileByte []byte

//go:embed folderText/*.txt
var folderText embed.FS

func main() {
	println(fileString)
	println(string(fileByte))
	content1, _ := folderText.ReadFile("folderText/folderFile1.txt")
	println(string(content1))
	content2, _ := folderText.ReadFile("folderText/folderFile2.txt")
	println(string(content2))
}
