all: build

build:
	go build

linux: *.go
	GOOS=linux GOARCH=amd64 go build

clean:
	rm -f 200ok

dist: linux
	tar -cvzf linux-amd64-200ok.tar.gz ./200ok 
