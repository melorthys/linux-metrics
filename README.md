# linux-metrics
Simple Golang tool for monitoring linux cpu, ram and disk usage.

### Usage
```
go run main.go
```

### Build and run
```
go build
chmod +x main
./main
```

### Output
```
ram  = % 26.9
cpu  = % 1.7
disk = % 4
```

### How to use binary
```
wget https://github.com/melorthys/linux-metrics/releases/download/v0.0.1/linux-metrics-amd64
chmod +x linux-metrics-amd64
./linux-metrics-amd64
```