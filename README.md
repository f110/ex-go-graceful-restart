# goでgraceful restartなhttpサーバーのサンプル

golangでgraceful restartができるアプリケーションサーバーのサンプル


## 使い方

### 1. サーバーを起動する

```
$ go build server.go
$ gom exec start_server --port=8000 --pid-file=./server.pid -- ./server
```

### 2. クライアントを起動する

```
$ go run client.go
```

### 3. graceful restartしてみる

```
$ kill -HUP $(< server.pid)
```
