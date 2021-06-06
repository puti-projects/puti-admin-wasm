# puti-admin-wasm
Puti 后台管理，前端部分。基于 Go + Wasm + go-app 实现。

## Wasm Side
```bash
GOARCH=wasm GOOS=js go build -o web/app.wasm cmd/wasm/main.go
```

## Server Side
Run dev：
```
go run cmd/server/test-server.go
```

Build and run：
```
go build cmd/server/test-server.go 

./test-server 
```

打开：http://localhost:8000/
