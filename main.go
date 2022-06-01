package main

// #cgo CFLAGS: -I${SRCDIR}/libs -Wall -g
// #cgo LDFLAGS: -L${SRCDIR}/libs/bin -lrastrigin
// #include <rastrigin.h>
import "C"

import (
  "fmt"
  "time"
  "math"
)

func rastrigin(n int) float64 {
  const A int = 10
  var res, x, inc float64

  inc = 10.24 / float64(n)
  x = -5.12

  for i := 0; i < n; i++ {
    res += (x * x) - float64(A) * math.Cos(2.0 * math.Pi * x)
    x += inc
  }
  res = float64(A * n) + res

  return res
}

func main() {
  n := (1 << 24)

  fmt.Printf("GO...\n")
  go_timeStart := time.Now()
  go_res := rastrigin(n)
  go_timeElapsed := time.Since(go_timeStart)
  fmt.Printf("GO result: %v, time: %s\n", go_res, go_timeElapsed)

  fmt.Printf("CGO...\n")
  c_timeStart := time.Now()
  c_res := C.rastrigin(C.int(n))
  c_timeElapsed := time.Since(c_timeStart)
  fmt.Printf("C result: %v, time: %s\n", c_res, c_timeElapsed)
}
