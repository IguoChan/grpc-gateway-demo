go mod tidy

# 下载proto文件引用proto文件的包
go get -u github.com/googleapis/googleapis
go get -u github.com/protocolbuffers/protobuf

# 安装对应的go工具
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.31.0

go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0

go install github.com/envoyproxy/protoc-gen-validate@v1.0.2

go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.18.1 \
            github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.18.1
