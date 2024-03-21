SHELL=cmd.exe
.PHONY: build run plot

all: build run plot

build:
	go build .\cmd\main.go

run:
	chcp 65001 & .\main.exe -a -h -i -q -s -c 100 1000 2000 4000 8000 16000 32000 64000 > data.csv

plot:
	python .\scripts\plottime.py .\data.csv plots
