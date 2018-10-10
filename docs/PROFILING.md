
# Profiling

## Create Profiles

~~~
./urlinfo -urlFile=./_TestData/urls.txt -cpuprofile=./cpu.prof -memprofile=./mem.prof
~~~

## Interpret Profiling

> CPU

~~~
go tool pprof urlinfo ./cpu.prof
~~~

> Memory

~~~
go tool pprof urlinfo ./mem_Done.prof
~~~

## Convert

> CPU, PDF

~~~
go tool pprof --pdf ./cpu.prof > cpu.pdf
~~~

> Memory, PDF

~~~
go tool pprof --pdf ./mem.prof > mem_Done.pdf
~~~

## References

- <https://flaviocopes.com/golang-profiling/>
- <https://blog.golang.org/profiling-go-programs>
