package main

import "fmt"
import "time"
import "math"
import "sync"

func source(nums []int) chan float64{
  out := make(chan float64, len(nums))

  for num := range nums {
    out <- float64(num)
  }
  close(out)
  return out
}

func sqrt(in chan float64) chan float64 {
  out := make(chan float64)

  go func() {
    for i := range in {
      time.Sleep(200 * time.Millisecond)
      out <- math.Sqrt(i)
    }
    close(out)
  }()
  return out
}

func mutiTen(in chan float64) chan float64 {
  out := make(chan float64)

  go func() {
    for i := range in {
      time.Sleep(100 * time.Millisecond)
      out <- i * 10
    }
    close(out)
  }()
  return out
}

func merge(cs ...chan float64) chan float64 {
  var wg sync.WaitGroup
  out := make(chan float64)

  output := func(c <-chan float64) {
      for n := range c {
          out <- n
      }
      wg.Done()
  }
  wg.Add(len(cs))
  for _, c := range cs {
      go output(c)
  }

  go func() {
      wg.Wait()
      close(out)
  }()
  return out
}

func main() {
  var nums []int
  for i := 1; i<100; i++ {
    nums = append(nums, i)
  }
  s := source(nums)

  s1 := sqrt(s)
  s2 := sqrt(s)
  result := mutiTen( merge(s1, s2) )

  for i := range result {
    fmt.Println(i)
  }
}
