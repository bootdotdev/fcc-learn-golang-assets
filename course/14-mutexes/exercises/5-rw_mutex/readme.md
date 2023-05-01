# RW Mutex

The standard library also exposes a [sync.RWMutex](https://golang.org/pkg/sync/#RWMutex)

In addition to these methods:

* [Lock()](https://golang.org/pkg/sync/#Mutex.Lock)
* [Unlock()](https://golang.org/pkg/sync/#Mutex.Unlock)

The `sync.RWMutex` also has these methods:

* [RLock()](https://golang.org/pkg/sync/#RWMutex.RLock)
* [RUnlock()](https://golang.org/pkg/sync/#RWMutex.RUnlock)

The `sync.RWMutex` can help with performance if we have a read-intensive process. Many goroutines can safely read from the map at the same time (multiple `Rlock()` calls can happen simultaneously). However, only one goroutine can hold a `Lock()` and all `RLock()`'s will also be excluded.

## Assignment

Let's update our same code from the last assignment, but this time we can speed it up by allowing readers to read from the map concurrently.

Update the `val` method to only lock the mutex for *reading*. Notice that if you run the code with a write lock it will block forever.
