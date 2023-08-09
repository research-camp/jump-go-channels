# Jump Golang Channels

![](https://img.shields.io/badge/language-go-32a89e)
![](https://img.shields.io/badge/topic-scheduling-DD5511)
![](https://img.shields.io/badge/version-v0.3-AA5533)

Implementing scheduling feature for Golang channels in order to jump over low priority tasks in them.
The reason for this feature is to send important data sooner.
For scheduling buffered and unbuffered channels
we used [pyramid](https://github.com/amirhnajafiz/pyramid) which is a heap
data structure in Golang.

- [Interface](#interface)
  - [create](#create-channel)
  - [schedule](#scheduling)
- [Example](#example)
- [Benchmark](#benchmark)

## Interface

Just like Golang channels, we created an interface in order to simulate the channel behaviour.
The interface goes like this:

```go
type Channel interface {
    Send(interface{})
    Recv() (interface{}, bool)
    Close()
    Next() bool
}
```

### create channel

In order to create channel, you can set size for it. If the size is zero then it would create an
unbuffered channel. Otherwise it would create a buffered channel.

```go
// buffered channel
ch := internal.NewChannel(2, false)

// unbuffered channel
uch := internal.NewChannel(0, false)
```

### scheduling

To make a scheduled channel you need to create a channel with ```true``` input in second argument.

```go
// buffered channel
ch := internal.NewChannel(10, true)
```

Make sure that your input data follows the following interface to have a priority method:

```go
type (
	// Schedulable object has a priority method which
	// return the priority of that object for scheduling.
	Schedulable interface {
		Priority() int
	}
)
```

This method is used to sort channel data based on the priority.

## Example

As you can see, I created a buffered channel with capacity of 2. Then we send
data with order of low priority to high. But in consuming, it will get the high
priority first and low priority later.

```go
// buffered channel
ch := internal.NewChannel(2, true)
wg := sync.WaitGroup{}

wg.Add(2)

ch.Send(Data{
    Value: "low value",
    p:     1,
})
ch.Send(Data{
    Value: "high value",
    p:     2,
})

// create a go routine
go func() {
    for ch.Next() {
        msg, _ := ch.Recv()

        log.Println(msg.(Data).Value)

        wg.Done()
    }
}()

wg.Wait()
ch.Close()
```

### output

```shell
2023/07/16 09:47:02 high value
2023/07/16 09:47:02 low value
```

## Benchmark

In order to execute benchmarks, use ```make bench``` command.

```shell
BenchmarkChannel-8       8279013               121.0 ns/op
BenchmarkChannel-8      10276252               115.6 ns/op
BenchmarkChannel-8      10379642               115.7 ns/op
BenchmarkChannel-8      10325314               116.1 ns/op
BenchmarkChannel-8      10378026               115.6 ns/op
BenchmarkChannel-8      10111676               120.9 ns/op
BenchmarkChannel-8      10028256               118.9 ns/op
BenchmarkCJump-8             100          13387163 ns/op
BenchmarkCJump-8             100          13141886 ns/op
BenchmarkCJump-8             100          13122004 ns/op
BenchmarkCJump-8             100          13126526 ns/op
BenchmarkCJump-8             100          13286897 ns/op
BenchmarkCJump-8             100          13156385 ns/op
BenchmarkCJump-8             100          12833818 ns/op
PASS
ok      github.com/research-camp/go-channels-scheduling/test    6.967s
```
