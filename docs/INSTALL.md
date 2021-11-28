# INSTALL.md

## Requirements

1. go

```
$ go version
go version go1.17.3 linux/amd64
```

2. docker and docker-compose

3. protoc and following plugins

```
sudo apt install -y protobuf-compiler
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
```

4. make

## Build

1. run `make build`

```
$ make build
```

## Run

1. run `docker-compose run`

```
$ docker-compose run
```