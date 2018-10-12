
# Tracing

## Trace

> Report package

~~~
cd report
~~~

~~~
go test . -v -trace trace.out
~~~

> Cannot use -trace flag with multiple packages

## View, web

> Report package

~~~
go tool trace trace.out
~~~

## Generate a pprof-like profile

### Network Blocking Profile

~~~
go tool trace -pprof=net trace.out > net.pprof
~~~

### Synchronization Blocking Profile

~~~
go tool trace -pprof=sync trace.out > sync.pprof
~~~

### SysCall Blocking Profile

~~~
go tool trace -pprof=syscall trace.out > syscall.pprof
~~~

## Scheduler Latency Profile

~~~
go tool trace -pprof=sched trace.out > sched.pprof
~~~

## Generate flamegraph

> syscall profile

~~~
go tool pprof -http=":8081" ./urlinfo syscall.pprof
~~~

> sched profile

~~~
go tool pprof -http=":8081" ./urlinfo sched.pprof
~~~
