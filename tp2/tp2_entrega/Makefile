.PHONY: all clean

all: algogram

algogram: main.go pkg_algogram/*.go
	go build -o algogram main.go

clean:
	rm -f algogram

