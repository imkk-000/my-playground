package main

import (
	"embed"
)

func main() {
	//go:embed hello.txt
	var s string
	print(s)
	println()

	//go:embed hello.txt
	var b []byte
	print(string(b))
	println()

	//go:embed hello.txt
	var f embed.FS
	data, _ := f.ReadFile("hello.txt")
	print(string(data))
	println()
}
