all: bindata.go example

bindata.go: static/index.html static/main.js
	go-bindata -pkg='bindata' -o bindata/bindata.go static/

clean:
	rm -f bindata/bindata.go
	rm -f example

example:
	go build github.com/aalpern/go-metrics-charts/cmd/example
