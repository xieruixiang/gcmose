PROTO_PATH=./proto
API_PATH=./auth/api
mkdir -p $API_PATH
protoc -I=$PROTO_PATH --go_out=paths=source_relative:$API_PATH --go-grpc_out=paths=source_relative:$API_PATH wechat.proto
protoc -I=$PROTO_PATH --grpc-gateway_out=paths=source_relative,grpc_api_configuration=$PROTO_PATH/wechat.yaml:$API_PATH wechat.proto

function createProto(){
  API_PATH=$1
  PROTO_NAME=$2
  PROTO_PATH=$3
  mkdir -p $API_PATH
  protoc -I=$PROTO_PATH --go_out=paths=source_relative:$API_PATH --go-grpc_out=paths=source_relative:$API_PATH $PROTO_NAME.proto
  protoc -I=$PROTO_PATH --grpc-gateway_out=paths=source_relative,grpc_api_configuration=$PROTO_PATH/$PROTO_NAME.yaml:$API_PATH $PROTO_NAME.proto
}

createProto ./rental/api trip ./proto
createProto ./blob/api profile ./blob/proto

