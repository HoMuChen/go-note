package main

import "fmt"
import "time"
import "math"

func sqrt(x float64) float64{
  time.Sleep(2 * time.Second)
  return math.Sqrt(x)
}

func mutiTen(x float64) float64{
  time.Sleep(time.Second)
  return x  * 10
}

func worker(f func(float64) float64, jobs chan float64, result chan float64) {
  for job := range jobs {
    //fmt.Println("receive job: ", job)
    result <- f(job)
  }
}

func main() {
  jobs := make(chan float64, 100)
  mid := make(chan float64, 100)
  results := make(chan float64, 100)

  for i := 0; i < 10; i++ {
    go worker(sqrt, jobs, mid)
  }
  for i := 0; i < 5; i++ {
    go worker(mutiTen, mid, results)
  }


  for j := 1; j <= 100; j++ {
    jobs <- float64(j)
  }
  close(jobs)

  for {
    select {
      case result := <-results:
        fmt.Println("Receive result: ", result)
      default:
        continue
    }
  }
}
