# Algorytmy i złożoność obliczeniowa

## Wymagania:
- Python
- Matplotlib
- numpy
- go 1.22
- make

## Uruchomienie:
```bash
make
```

### Build
```bash
make build
```

or

```bash
go build .\cmd\main.go
```
### Run
```bash
make run
```

or

```bash
chcp 65001 & .\main.exe -c 100 1000 2000 4000 8000 16000 32000 64000 > data.csv
```
main.exe -c [count of tests] [sizes of test arrays] > data.csv

### Plot
```bash
make plot
```

or

```bash
python .\scripts\plottime.py .\data.csv plots
```

