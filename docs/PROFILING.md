
# Profiling

## Create Profiles

~~~
./urlinfo -urlFile=./_TestData/urls.txt -cpuprofile=./_OUTPUT/cpu.prof -memprofile=./_OUTPUT/mem.prof
~~~

## Interpret Profiling

> CPU

~~~
go tool pprof urlinfo ./_OUTPUT/cpu.prof
~~~

> Memory

~~~
go tool pprof urlinfo ./_OUTPUT/mem_Done.prof
~~~

## Convert

> CPU, PDF

~~~
go tool pprof --pdf ./_OUTPUT/cpu.prof > cpu.pdf
~~~

> Memory, PDF

~~~
go tool pprof --pdf ./_OUTPUT/mem.prof > mem_Done.pdf
~~~

## References

- <https://flaviocopes.com/golang-profiling/>
- <https://blog.golang.org/profiling-go-programs>
