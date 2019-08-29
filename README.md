# Fact 
this package implements factorial function in 3 different approaches
### SimpleFact
most simple implementation and only implements tail-call optimization for avoiding stack over flow
### Goroutine Fact
this implementation is a bit more complex, it will break the sequence into smaller sequences and spin up a goroutine for each one and all results goes into a channel which after all results arrived we will start adding them together
### Worker Fact
this implementation is most complex version and it will create a worker pool, and is optimized for large numbers