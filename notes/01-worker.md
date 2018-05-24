# architecture
![image](/images/worker.png)
# code
[src/worker.go](/src/worker.go)
# notes
* Able to pass function as parameters
* Use channel as queue, workers consume a channel and put result into another channel
```
func worker(f func(float64) float64, jobs chan float64, result chan float64) {
  for job := range jobs {
    result <- f(job)
  }
}
```
* Infinite for loop to consume the final channel
```
  for {
    select {
    case result := <-results:
      fmt.Println("Receive result: ", result)
    default:
      continue
    }
  }
```
* Start more workers for more expensive jobs(simulate with time.Sleep())
```
func sqrt(x float64) float64{
  time.Sleep(2 * time.Second)
  return math.Sqrt(x)
}

func mutiTen(x float64) float64{
  time.Sleep(time.Second)
  return x * 10
}

for i := 0; i < 10; i++ {
  go worker(sqrt, jobs, mid)
}
for i := 0; i < 5; i++ {
  go worker(mutiTen, mid, results)
}
```
