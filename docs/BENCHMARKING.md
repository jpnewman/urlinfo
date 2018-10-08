
# Benchmarking

~~~
go test ./... -bench=.
~~~

## CPU Profiling

> Report package

~~~
cd report
~~~

~~~
go test . -bench=. -cpuprofile cpu.out
~~~

> Cannot use ```-cpuprofile``` flag with multiple packages

## Interpret CPU Profiling

> Report package

~~~
go tool pprof report.test cpu.out
~~~

### Get Top CPU

~~~
(pprof) top10
~~~

## Memory Profiling

> Report package

~~~
cd report
~~~

~~~
go test . -bench=. -memprofile mem.out
~~~

## Interpret Memory Profiling

> Report package

~~~
go tool pprof report.test mem.out
~~~

### Get Top CPU

~~~
(pprof) top10
~~~
