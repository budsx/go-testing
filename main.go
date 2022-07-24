package main

import (
	"fmt"
	"runtime"
)

func main() {
	cek := runtime.GOOS
	fmt.Println(cek)
}
