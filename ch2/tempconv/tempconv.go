// Package tempconv performs celcius to farenheit temparature compuatations
package tempconv

import "fmt"

type Celcius float64
type Farenheit float64

const (
    AbsoluteZeroC   Celcius = -273.15
    FreezingC       Celcius = 0
    BoilingC        Celcius = 100
)

//!-

func (c Celcius) String() string { return fmt.Sprintf("%g deg C", c) }

func (f Farenheit) String() string { return fmt.Sprintf("%g deg F", f) }
