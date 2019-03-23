package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	fmt.Scan(&s)
	n, _ := strconv.Atoi(s)
	for i:=0; i<n-1; i++ {
		fmt.Print(3, " ")
	}
	fmt.Println(1);
}

