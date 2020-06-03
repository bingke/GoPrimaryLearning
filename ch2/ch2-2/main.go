// Cf converts its numeric argument to Centimeter and Meter.
package main

import (
    "fmt"
    "os"
    "strconv"

    "github.com/bingke/GoPractice/src/lenconv"
)

func main() {
    for _, arg := range os.Args[1:] {
        t, err := strconv.ParseFloat(arg, 64)
        if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			fmt.Printf("wrong")
            os.Exit(1)
        }
        m := lenconv.Meter(t)
        cm := lenconv.Centimeter(t)
        fmt.Printf("%s = %s, %s = %s\n",
            m, lenconv.MToCM(m), cm, lenconv.CMToM(cm))
    }
}