
# URLInfo

Go Program to get page information from URL.

It can create a standard output (stdout) or Markdown report and also logs to a separate log file.  

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
go build -gcflags -m *.go
~~~

## Test

[TESTING.md](docs/TESTING.md)

## Code Coverage

[COVERAGE.md](docs/COVERAGE.md)

## Benchmarking

[BENCHMARKING.MD](docs/BENCHMARKING.MD)

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

In Dry-Run mode no HTTP request are made, But they are simulated by sleeping for ```-httpTimeoutSeconds```.

~~~
./urlinfo -urlFile=./_TestData/urls.txt -httpTimeoutSeconds=1 -dryrun
~~~

## Profiling

[PROFILING.md](docs/PROFILING.md)

## Editing

Visual Studio Code (VSCode), on Mac OS X, was used as an IDE.  
This project contains VSCode build and debug configurations.
