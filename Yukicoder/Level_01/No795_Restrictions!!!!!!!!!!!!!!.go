package main

import (
	"fmt"
	"strconv"
)

func getINT() int {
	var s string
	fmt.Scan(&s)
	n, _ := strconv.Atoi(s)
	return n
}

func main(){
	_ = getINT()
	_ = getINT()
	fmt.Println("Yes");
}