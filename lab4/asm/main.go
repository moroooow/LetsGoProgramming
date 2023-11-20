package asm

func main() {
	x := 1000
	y := 7
	if true {
		sub(x, y)
	}

}

func sub(a, b int) {
	println(a - b)
}
