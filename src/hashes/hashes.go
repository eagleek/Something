package main

import (
	 "crypto/sha256"
	"fmt"
)


func main() {
	c1:= sha256.Sum256([]byte("X"))
	c2:= sha256.Sum256([]byte("X"))
	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Printf("%d",countHashes(c1,c2))
	
}
func countHashes(h1,h2 [32]byte)int{
	var counts int
	for i:= range h1{
		if h1[i]!=h2[i]{
		counts++}
		}
	return counts
	}

