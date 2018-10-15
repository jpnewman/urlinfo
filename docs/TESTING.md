
# Testing

## Run all tests

~~~
go test ./... -v
~~~

## ANSI Color Output

> Mac OS X / Bash

~~~
go test ./... -v | sed ''/PASS/s//$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$(printf "\033[31mFAIL\033[0m")/''
~~~

### RichGo, ANSI Color Output

#### Install Rich-Go

~~~
go get -u github.com/kyoh86/richgo
~~~

#### Run

~~~
richgo test ./... -v
~~~
