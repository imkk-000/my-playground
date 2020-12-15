package main

import "fmt"

func main() {
	sliceLiteral := []int{}
	sliceMake := make([]int, 0)

	fmt.Printf("%+v | %v | %T | addr(%p) | len(%d) | cap(%d) | nil(%v)\n",
		sliceLiteral, sliceLiteral, sliceLiteral, sliceLiteral, len(sliceLiteral), cap(sliceLiteral), sliceLiteral == nil)
	fmt.Printf("%+v | %v | %T | addr(%p) | len(%d) | cap(%d) | nil(%v)\n",
		sliceMake, sliceMake, sliceMake, sliceMake, len(sliceMake), cap(sliceMake), sliceMake == nil)
}
