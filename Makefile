protoCompile:
	protoc --go_out=./src/services/protos --go_opt=paths=source_relative --go-grpc_out=./src/services/protos --go-grpc_opt=paths=source_relative ./PerimeterX.proto
