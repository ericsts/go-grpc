#para compilar


```
go get google.golang.org/protobuf/cmd/protoc-gen-go google.golang.org/grpc/cmd/protoc-gen-go-grpc

$ protoc --proto_path=proto proto/*.proto --go_out=pb
$ protoc --proto_path=proto proto/*.proto --go_out=pb --go-grpc_out=pb
```

