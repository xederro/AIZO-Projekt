SHELL=cmd.exe
.PHONY: build run plot

build:
	go build .\cmd\main.go

run:
	chcp 65001 & .\main.exe > data.csv

plot:
	python .\scripts\plottime.py .\data.csv plots

all: build run plot