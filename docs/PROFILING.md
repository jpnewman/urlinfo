
# Profiling

## Timing

> Mac OS X

~~~
time ./urlinfo -urlFile=./_TestData/urls.txt
~~~

## Create Profiles

~~~
go build; ./urlinfo -urlFile=./_TestData/urls.txt -cpuprofile=./cpu.prof -memprofile=./mem.prof
~~~

> Output file ```./mem.prof``` will be changed to ```./mem_Done.prof```.

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
go tool pprof --pdf ./cpu.prof > cpu.pdf; open cpu.pdf
~~~

> Memory, PDF

~~~
go tool pprof --pdf ./mem_Done.prof > mem_Done.pdf; open mem_Done.pdf
~~~

## Web

~~~
go tool pprof -http=":8081" ./urlinfo mem_Done.prof
~~~

## References

- <https://flaviocopes.com/golang-profiling/>
- <https://blog.golang.org/profiling-go-programs>
