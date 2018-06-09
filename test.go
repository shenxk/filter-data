package main

import (
	"fmt"
	. "math"
	"time"
)

func main() {
	data()
}
func direct() {

}

var cmd float64

func data() {
	fmt.Println("hello")
	for i := 0.0; ; i++ {
		time.Sleep(time.Millisecond * 1)
		cmd = Sin(i)
	}
}
