package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main(){
	var s string
	fmt.Scan(&s)

	r := regexp.MustCompile(`c[^w]*w[^w]*w`)

	var ans int = -1

	for {
		index := strings.Index(s, "c")
		if index == -1 {
			break
		}

		s = s[index:]

		var res string = r.FindString(s)

		if res != "" && (ans < 0 || len(res) < ans) {
			ans = len(res)
		}

		s = s[1:]
	}
	fmt.Println(ans)
}