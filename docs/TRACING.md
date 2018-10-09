
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
