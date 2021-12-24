# Generate GRPC hasher api. 
protoc \
  --go-grpc_out=internal/common/hasherproto \
  --go_out=internal/common/hasherproto \
  api/hasher.proto