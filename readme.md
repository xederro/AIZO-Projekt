# Algorytmy i złożoność obliczeniowa

## Requirements:
- Python
- Matplotlib
- numpy
- go 1.22
- make

## How to run program:
```bash
make
```

or for manual tests

```bash
make manual
```

or for automated tests

```bash
make auto
```

### Build
```bash
make build
```

or

```bash
make prod
```

or

```bash
go build .\cmd
```
### Run
```bash
make run
```

or

```bash
.\cmd.exe -p -a -h -i -q -s -c 100 1000 2000 4000 8000 16000 32000 64000
```
cmd.exe [params] -c [count of tests] [sizes of test arrays]

-p - profile mode

-a - asynchronous mode

-h - test heap sort

-i - test insertion sort

-q - test quick sort

-s - test selection sort

> If you want to plot the data on windows, 
> you need to run the program with `chcp 65001 & <run command>`
> or manually convert csv to utf8

### Plot
```bash
make plot
```

or

```bash
python .\scripts\plottime.py .\data.csv plots
```

### Tests
```bash
make test
```

or

```bash
go test .\algo\sort\...
```

### Benchmarks
```bash
make bench
```

or

```bash
go test -bench=. .\algo\sort\...
```