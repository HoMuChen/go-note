# architecture
![image](/images/worker.compose.png)
# code
[src/worker.compose.go](/src/worker.compose.go)
# notes
* Refactoring from worker.go
* Every worker function return a new chan, removing side effects: putting result into a outside chan variable
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
* Now functions are composable
```go
output := multiTen(multiTen(multiTen(sqrt(input))))
```
* If you want to distribute works to many workers to parallelize CPU and I/O use, you can read from the same chan util the chan is closed
* Merge channels into one for next processing stage
```go
c1 := sqrt(input)
c2 := sqrt(input)
c3 := sqrt(input)

output := merge(c1, c2, c3)
```
* Merge function
```go
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
 ```
