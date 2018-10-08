
# Profiling

## Create Profiles

~~~
./urlinfo -urlFile=./_TestData/urls.txt -cpuprofile=./_OUTPUT/cpu.prof -memprofile=./_OUTPUT/mem.prof
~~~

## Interpret

> CPU

~~~
go tool pprof urlinfo ./_OUTPUT/cpu.prof
~~~

> Memory

~~~
go tool pprof urlinfo ./_OUTPUT/mem.prof
~~~

## Convert

> CPU, PDF

~~~
go tool pprof --pdf ./_OUTPUT/cpu.prof > cpu_prof.pdf
~~~

> Memory, PDF

~~~
go tool pprof --pdf ./_OUTPUT/mem.prof > mem_prof.pdf
~~~

## References

- <https://flaviocopes.com/golang-profiling/>
- <https://blog.golang.org/profiling-go-programs>
