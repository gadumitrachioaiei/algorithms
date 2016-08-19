Installing:
go get github.com/gadumitrachioaiei/algorythms/cmd/...

Example of calling and output:
go build -o main github.com/gadumitrachioaiei/algorythms/cmd/...
./main -limit 1000000 -n 10 -m 20
longest chain up to: 1000000 is for the number: 837799
shortest path between: 10 and: 20 is: [10 20]

Run the tests:
go test -v github.com/gadumitrachioaiei/algorythms/transformer
