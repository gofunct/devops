protobuf/hello.pb.go: protobuf/hello.proto
		protoc -I=protobuf $< --go_out=plugins=grpc:protobuf
		
