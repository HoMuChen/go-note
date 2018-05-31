# architecture
![image](/images/worker.compose.png)
# code
[src/worker.compose.go](/src/worker.compose.go)
# notes
* Refactoring from worker.go
* Every worker function return a new chan
```go
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
```
