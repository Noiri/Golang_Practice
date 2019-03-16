package main

import "fmt"

func main(){
	var s string
	fmt.Scan(&s)

	var ans = -1

	for i, c := range s {
		if string(c) != "c" {
			continue
		}
		w := 0
		for j, c := range s[i:] {
			if string(c) == "w" {
				w++
				if w == 2 {
					if ans < 0 || (len(s[i:i+j+1])) < ans {
						ans = (len(s[i:i+j+1]))
					}
					break
				}
			}
		}
	}
	fmt.Println(ans)
}
