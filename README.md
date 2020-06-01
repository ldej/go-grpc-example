# A Go gRPC application
Source: https://www.youtube.com/watch?v=BdzYdN_Zd9Q

- https://developers.google.com/protocol-buffers/docs/downloads
- https://github.com/protocolbuffers/protobuf/releases/

```shell script
$ sudo cp ~/Downloads/protoc-3.12.2-linux-x86_64/bin/protoc /usr/local/bin/
$ sudo chmod +x /usr/local/bin/protoc
$ sudo cp -r ~/Downloads/protoc-3.12.2-linux-x86_64/include/google /usr/local/include/
```

```shell script
$ go get -u github.com/golang/protobuf/protoc-gen-go
$ protoc --go_out=plugins=grpc:chat chat.proto
```
