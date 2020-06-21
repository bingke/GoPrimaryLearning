//Compare whether the characters contained in the two strings are the same.

package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "asdsg"
	s2 := "asdg"
	b := comString(s1, s2)
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