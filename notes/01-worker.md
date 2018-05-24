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
* infinite for loop to consume the final channel
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
