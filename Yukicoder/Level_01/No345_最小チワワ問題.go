package main

import (
	"fmt"
	"math"
)

func min(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}

func main(){
	var s string
	fmt.Scan(&s)
	var ans int = 1e9
	for i := 0; i <len(s); i++ {
		if s[i] == 'c'{
			var cnt int = 0
			for j := i; j < len(s); j++ {
				if s[j] == 'w' {
					cnt++;
				}
				if cnt == 2 {
					ans = min(ans, j-i+1)
				}
			}
		}
	}
	if ans == int(1e9) {
		fmt.Println(-1)
	} else {
		fmt.Println(ans)
	}
}
