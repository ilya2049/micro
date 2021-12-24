# Rest api generator
go install github.com/go-swagger/go-swagger/cmd/swagger@v0.28.0

# Protoc
PROTO_OS_VERSION="linux-x86_64"
PROTO_VERSION="3.19.1"
PROTO_ZIP="protoc-${PROTO_VERSION}-${PROTO_OS_VERSION}.zip"
PB_REL="https://github.com/protocolbuffers/protobuf/releases"
curl -L ${PB_REL}/download/v${PROTO_VERSION}/${PROTO_ZIP} -o /tmp/${PROTO_ZIP}
unzip /tmp/${PROTO_ZIP} -d ${HOME}/.local

# Protoc grpc plugins
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1

# GRPC UI
go install github.com/fullstorydev/grpcui/cmd/grpcui@v1.2.0

# Unit test browser goconvey
go install github.com/smartystreets/goconvey@v1.7.2