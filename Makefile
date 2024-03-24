SHELL=cmd.exe
.PHONY: build run plot prof prod test bench

all: prod run plot

build:
	go build .\cmd\main.go

run:
	chcp 65001 & .\main.exe -a -h -i -q -s -c 100 1000 2000 4000 8000 16000 32000 64000 > data.csv

plot:
	python .\scripts\plottime.py .\data.csv plots

prof:
	go tool pprof -http 127.0.0.1:8080 cpu_profile.prof

prod:
	go build -ldflags "-s -w" .\cmd\main.go

test:
	go test .\algo\sort\...

bench:
	go test -bench=. .\algo\sort\...
