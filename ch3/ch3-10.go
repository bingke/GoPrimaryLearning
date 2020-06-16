// comma inserts commas in a non-negative decimal integer string.
package main

import (
	"bytes"
	"fmt"
)


func main(){
	var s = "1234567890"
	fmt.Println(comma(s))
}


func comma(s string) string {
	
	const width = 3

	if len(s) == 0 {
		return "there is something wrong in your string"
	}

	var buf bytes.Buffer  

	if len(s) < width {
		return s
	}
	
	iterations := len(s)/width
	overflow := len(s)%width

	if len(s) % width != 0 {
		subS := s[ 0 : overflow]
		fmt.Fprintf(&buf, "%s", subS)
	} else {
		subS := s[0*width + overflow : (0+1)*width + overflow]
		fmt.Fprintf(&buf, "%s", subS)
	}

	for i := 1; i < iterations; i++ {

		buf.WriteString(", ")
		subS := s[i*width + overflow : (i+1)*width + overflow]
		fmt.Fprintf(&buf, "%s", subS)
		
	}

	return buf.String()
}

