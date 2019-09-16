sudo apt  install protobuf-compiler
go get -u github.com/golang/protobuf/{protoc-gen-go,proto}

protoc --go_out=. test.proto
