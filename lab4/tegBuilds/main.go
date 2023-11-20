package main

//go:generate go build -tags=ultima,plus
import "fmt"

var features = []string{
	"TODO list",
	"Tech support",
}

func main() {
	for _, f := range features {
		fmt.Println("#", f)
	}
}
