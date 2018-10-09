
# Coverage

## Run

~~~
go test  ./... -cover -coverprofile=cover.out
~~~

## Convert coverage report to HTML

~~~
go tool cover -html=cover.out
~~~
