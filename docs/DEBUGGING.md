
# Debugging

## Install Delve

~~~
go get -u -v github.com/derekparker/delve/cmd/dlv
~~~

## Build with Debug symbols

~~~
GOCACHE=off go build -x -gcflags='all=-N -l' -tags nopkcs11 -ldflags='-linkmode internal' -o urlinfo *.go
~~~

## Dump Objects

> Mac OS X

~~~
objdump -section-headers ./urlinfo
~~~

## Disassemble

~~~
go tool objdump -S ./urlinfo > objdump.asm
~~~

## Run Delve (dlv)

### Build and Debug

~~~
dlv debug -- -urlFile=./_TestData/urls.txt
~~~

~~~
(dlv) break main.main
(dlv) continue
~~~
