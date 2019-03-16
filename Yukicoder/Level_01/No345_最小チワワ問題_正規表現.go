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
		//こうしないとwcwに捕まってans=0になってしまう
		if r.MatchString(s) == false {
			break
		}

		s = s[index:]
		res := r.FindString(s)

		if ans < 0 || len(res) < ans {
			ans = len(res)
		}

		s = s[1:]
	}
	fmt.Println(ans)
}
