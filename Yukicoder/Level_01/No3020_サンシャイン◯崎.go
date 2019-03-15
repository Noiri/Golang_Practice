package main

import (
	"bufio"
	"os"
	"strings"
	"fmt"
)

var sc = bufio.NewScanner(os.Stdin)

func read() string {
	sc.Scan()
	return sc.Text()
}

func main(){
	//scanner.Scan()は、バッファサイズを超えると読み込みに失敗するので今回は使用しない（大きいケースでエラー）
	//s := read()

	var s string
	fmt.Scan(&s)
	var YEAH = []string{"Y", "E", "A", "H", "!"}
	for i := 0; i < len(YEAH); i++ {
		var cnt int = strings.Count(s, YEAH[i])

		if i != (len(YEAH) - 1) {
			fmt.Print(cnt, " ")
		} else {
			fmt.Println(cnt)
		}
	}
}
