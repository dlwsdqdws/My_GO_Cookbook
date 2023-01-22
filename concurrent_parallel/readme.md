- [Concurrent and Parallel](#concurrent-and-parallel)
  - [Goroutine](#goroutine)
  - [Channel](#channel)
  - [Synchronization](#synchronization)
    - [Mutex](#mutex)
    - [WaitGroup](#waitgroup)


# Concurrent and Parallel

## Goroutine

1. Differences between Goroutine and Threads

- Goroutine : works on user mode, lightweight thread, KB level in stack.
- Thread : works on kernel mode, MB level in stack.

2. Use `go` start a new goroutine.

```go
go func(){}
```

Function can be anonymous. When a goroutine needs to be blocked, the system will move other goroutines waiting to run on this thread to other threads that can run programs, so other goroutines will not be stuck.

## Channel

1. Definition

Go advocates sharing memory through communication instead of communication through sharing memory. Sharing memory through communication needs a **channel**. A channel is a data structure used to transfer data. It can be used between two goroutines to synchronize operation and communication by passing a value of a specified type.

2. Usage

We can create a channel by using

```go
make(chan mem_type ,[buffer_size])
```

- If buffer is NOT used, the channel is a synchronous channel. The sender will block until a receiver has received a value from the channel. The receiver blocks until there is a value to receive.
- If buffer is used, the channel is a producer-consumer model. The sender will block until the sent value is sent into the buffer. If the buffer is full, the sender will block until a receiver receives a value. The receiver blocks until there is a value to receive.

The operator `<-` is used to specify the direction of the channel to achieve sending or receiving. In particular, if no direction is specified, it is a bidirectional channel.

```go
// Send data to channel ch
ch <- data

// Receive data from channel ch and assign it to d
d := <- ch
```

The channel should be closed after being used, otherwise it is easy to cause deadlock.

```go
defer close(chan)
```

## Synchronization

Communication through sharing memory can lead to multiple goroutines accessing the same data at the same time, and thus we need a **Lock**.

### Mutex

Each thread tries to lock the data before accessing it. The operation(R/W) can only be performed after successful locking, and then unlocked after the operation is completed. In other words, only one goroutine can access the data at a time when a mutex is used.

```go
var lock sync.Mutex // Declaration
lock.Lock() // add lock
//code...
lock.Unlock() // unlock
```

### WaitGroup

WaitGroup can help set 'sleep_time' accurately. WatiGroup can wait until all goroutines are executed and block the execution of the main thread until all goroutines are executed. WaitGroup has three methods:

| method |                                      meaning                                       |
| :----: | :--------------------------------------------------------------------------------: |
| `Add`  |                Add the number of waiting goroutines to the counter                 |
| `Done` | Decrement the value of the counter, should be executed at the end of the goroutine |
| `Wait` |                     block until all WaitGroup counts become 0                      |

Note that the number set by `Add()` must be consistent with the number of waiting goroutines otherwise a deadlock will happen. For more examples please refer to https://www.cnblogs.com/sunshineliulu/p/14779158.html.