package main

import (
	"fmt"
	"bufio"
	"strconv"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func read() string {
	sc.Scan()
	return sc.Text()
}

func getInt() int {
	i, _ := strconv.Atoi(read())
	return i
}

func main(){
	sc.Split(bufio.ScanWords)
	a := getInt()
	b := getInt()
	s := read()
	fmt.Print(a + b, " ", s)
}