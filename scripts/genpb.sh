set -ex

googleapisDir=$(go list -m -f "{{.Dir}}" github.com/googleapis/googleapis)
gatewayDir=$(go list -m -f "{{.Dir}}" github.com/grpc-ecosystem/grpc-gateway/v2)
validateDir=$(go list -m -f "{{.Dir}}" github.com/envoyproxy/protoc-gen-validate)

echo $googleapisDir
echo $gatewayDir
echo $validateDir

#protoc  -I../proto -I${googleapisDir} -I${gatewayDir} -I${validateDir} --go_out=../api/demopb  --go_opt paths=source_relative demo.proto

protoc -I../proto \
  -I${gatewayDir} \
  -I${validateDir} \
  -I${googleapisDir} \
  --experimental_allow_proto3_optional \
  --go_out=../api/demopb --go_opt paths=source_relative\
  --go-grpc_out=../api/demopb --go-grpc_opt paths=source_relative \
  --grpc-gateway_out=../api/demopb --grpc-gateway_opt paths=source_relative --grpc-gateway_opt logtostderr=true \
  --openapiv2_out=../swagger --openapiv2_opt use_go_templates=true --openapiv2_opt logtostderr=true \
  --validate_out="lang=go:../api/demopb" --validate_opt paths=source_relative  --validate_opt logtostderr=true \
  ../proto/*.proto