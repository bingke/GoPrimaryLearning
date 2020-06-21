//Compare whether the characters contained in the two strings are the same.

package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "sjj"
	s2 := "jjs"
	b := isReverse(s1, s2)
	fmt.Println(b)

}

func comString(s1, s2 string) bool {
	
	for i:=0; i < len(s1); i++ {
		if !strings.Contains(s2, string(s1[i])) {
			return false
		}
	}

	if len(s1) == len(s2) {
		return true
	}

	return false
}

//Fix comString func to recognize character strings like "sjj", "ssj".
func isReverse(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	m := make(map[rune]int)
	n := make(map[rune]int)

	for _,v := range a {
		m[v]++
	}
	for _,v := range b {
		n[v]++
	}

	for i,v := range m {
		if n[i] != v {
			return false
		} 
	}

	return true


}