
# Benchmarking

~~~
go test ./... -bench=.
~~~

## Comparing Benchmarks

### Install Tools

~~~
go get golang.org/x/tools/cmd/benchcmp
go get github.com/ajstarks/svgo/benchviz
~~~

### Generate OLD Benchmark

~~~
go test ./... -bench=. -run="^$" 2>&1 | tee benchmarks_OLD.log
~~~

> Don't run any tests ```-run="^$"```.

### Make Change

Make changes to source code.

### Generate NEW Benchmark

~~~
go test ./... -bench=. -run="^$" 2>&1 | tee benchmarks_NEW.log
~~~

> Don't run any tests ```-run="^$"```.

### Compare Benchmark Differences

~~~
benchcmp benchmarks_OLD.log benchmarks_NEW.log
~~~

### Visualize Benchmark Differences

~~~
benchcmp benchmarks_OLD.log benchmarks_NEW.log | benchviz > benchmarks.svg; open benchmarks.svg
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
