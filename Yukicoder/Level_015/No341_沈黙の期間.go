package main

import (
	"fmt"
	"regexp"
)

func main(){
	var s string
	fmt.Scan(&s);

	r := regexp.MustCompile(`…*…`);
	matchs := r.FindAllString(s, -1);

	ans := 0;
	for i:=0; i<len(matchs); i++ {
		size := len(matchs[i])/3;
		if ans < size {
			ans = size;
		}
	}
	fmt.Println(ans);
}
