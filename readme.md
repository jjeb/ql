# Getting Started

# Building package with Go

```
 GOARCH=wasm GOOS=js go build -o main.wasm main.go
```

# Building package with tinyGo

```
 tinygo build  --no-debug  -o main.wasm -target wasm main.go
```
