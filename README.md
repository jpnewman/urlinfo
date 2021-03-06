
# URLInfo

Go Program to get page information from URLs.

It can create a standard output (stdout) or Markdown report and also logs to a separate log file.

~~~
go get -u github.com/jpnewman/urlinfo
~~~

## Install Dependencies

- <https://github.com/golang/dep>

~~~
dep ensure
~~~

## Build

~~~
go build
~~~

> Print optimization decisions

~~~
go build -o ./urlinfo -gcflags -m *.go
~~~

> Print more detailed optimization decisions

~~~
go build -o ./urlinfo -gcflags '-m -m' *.go 2> op_decisions.log
~~~

## Test

[TESTING.md](docs/TESTING.md)

## Code Coverage

[COVERAGE.md](docs/COVERAGE.md)

## Benchmarking

[BENCHMARKING.md](docs/BENCHMARKING.md)

## Run

~~~
./urlinfo -urlFile=./_TestData/urls.txt
~~~

> Debug logging

~~~
LOG_LEVEL=Debug ./urlinfo -urlFile=./_TestData/urls.txt
~~~

> Use only HTTP HEAD Method

~~~
./urlinfo -urlFile=./_TestData/urls.txt -getHeadOny
~~~

**N.B.** Modern dynamic websites will not return HTTP header ContentLength.

> Generate Markdown Report

~~~
./urlinfo -urlFile=./_TestData/urls.txt -reportFormat=Markdown -reportFile=urlinfo.md
~~~

> Dry-Run

In Dry-Run mode no HTTP request are made, But they are simulated by sleeping for ```-httpTimeout```.

~~~
./urlinfo -urlFile=./_TestData/urls.txt -httpTimeout=1 -dryrun
~~~

## Profiling

[PROFILING.md](docs/PROFILING.md)

## Tracing

[TRACING.md](docs/TRACING.md)

## Debugging

[DEBUGGING.md](docs/DEBUGGING.md)

## Documentation

[DOCUMENTATION.md](docs/DOCUMENTATION.md)

## Editing

Visual Studio Code (VSCode), on Mac OS X, was used as an IDE.  
This project contains VSCode build and debug configurations.
