#para compilar


`
$ protoc --proto_path=proto proto/*.proto --go_out=pb
$ protoc --proto_path=proto proto/*.proto --go_out=pb --go-grpc_out=pb
`
