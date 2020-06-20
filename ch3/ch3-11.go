//comma function to support floating flaot number processing and an optional sign processing.
package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	s := "1234567890.56789"
	out := comma(s)
	fmt.Printf("%s\n", out)
}

func comma(s string) string {
	var buffer bytes.Buffer

	var symble byte
	if s[0]=='-' || s[0]=='+' {
		symble = s[0]
		s = s[1:]
	}
	
    buffer.WriteByte(symble)

	arr := strings.Split(s, ".")
	s = arr[0]
	l := len(s)

	for i := 0; i < len(s); i++ {

		buffer.WriteString(string(s[i]))
		
        if (i+1)%3 == l%3 && i != l-1 {
            buffer.WriteString(",")
		}
	}	

	if len(arr) > 1 { 
        buffer.WriteString(".")
        buffer.WriteString(arr[1])
    }

	return buffer.String()

}