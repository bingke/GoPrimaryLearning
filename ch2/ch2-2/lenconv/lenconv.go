// Package lenconv performs Meter and Centimeter conversions.
package lenconv

import "fmt"

type Meter float64
type Centimeter float64

func (m Meter) String() string      { return fmt.Sprintf("%g M", m) }
func (c Centimeter) String() string { return fmt.Sprintf("%g CM", c) }
