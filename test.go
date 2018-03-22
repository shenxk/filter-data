package main

import (
	. "math"
	"time"
)

func main() {

}
func direct() {

}

var cmd float64

func data() {
	for i := 0.0; ; i++ {
		time.Sleep(time.Millisecond * 1)
		cmd = Sin(i)
	}
}
