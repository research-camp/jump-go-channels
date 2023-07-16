# GO Buffered Channels Scheduling

![](https://img.shields.io/badge/language-go-1102CA)
![](https://img.shields.io/badge/topic-scheduling-DD5511)
![](https://img.shields.io/badge/version-v0.1-AA5533)

Implementing scheduling feature for Golang buffered channels.
The reason for this feature is to send important data sooner.
Since unbuffered channels don’t have buffer and they work by ```ticket/token```
topology we cannot set scheduling algorithms on them. But for scheduling buffered channels
we used [pyramid](https://github.com/amirhnajafiz/pyramid) which is a heap
data structure in Golang.

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

```shell
2023/07/16 09:47:02 high value
2023/07/16 09:47:02 low value
```
