SHELL=cmd.exe
.PHONY: build run plot prof prod test bench manual auto

manual: build
	.\cmd.exe

auto: prod run plot

build:
	go build .\cmd

run:
	chcp 65001 & .\cmd.exe -a -h -i -q -s -c 100 1000 2000 4000 8000 16000 32000 64000 > data.csv

plot:
	python .\scripts\plottime.py .\data.csv plots

prof:
	go tool pprof -http 127.0.0.1:8080 cpu_profile.prof

prod:
	go build -ldflags "-s -w" .\cmd

test:
	go test .\algo\sort\...

bench:
	go test -bench=. .\algo\sort\...
