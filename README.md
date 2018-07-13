# gofizzbuzz
A simple fizz buzz served hot from REST

`http://localhost:8020/fizzbuzz/100`

### Prerequisites

Will need golang 1.9.3 or above

### Setup

Do a git clone from this repo (preferable to your $GOPATH\src) or set GOPATH explicitly

```
git clone https://github.com/ashwathps/gofizzbuzz.git

cd $GOPATH/src/fizzbuzz-server/src

```
Then run

` got get -v -t .` to get and resolve all package deps

### Building & Running

```
cd src
go build server.go fizzbuzz.go

./server.go

```

### Tests
The code uses golang's default testing framework

Run ` go test -run TestFizzBuzz15`
