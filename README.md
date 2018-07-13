# gofizzbuzz
A simple fizz buzz served hot from REST


### Setup

Do a git clone from this repo (preferable to your $GOPATH\src)

```
git clone https://github.com/ashwathps/gofizzbuzz.git

cd fizzbuzz-server

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
