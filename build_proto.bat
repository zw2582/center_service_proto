echo "构建proto"

protoc --plugin=protoc-gen-go=%GOPATH%\bin\protoc-gen-go.exe --plugin=protoc-gen-micro=%GOPATH%\bin\protoc-gen-micro.exe --micro_out=. --go_out=. proto_app.proto

protoc --plugin=protoc-gen-go=%GOPATH%\bin\protoc-gen-go.exe --plugin=protoc-gen-micro=%GOPATH%\bin\protoc-gen-micro.exe --micro_out=. --go_out=. proto_coin.proto

protoc --plugin=protoc-gen-go=%GOPATH%\bin\protoc-gen-go.exe --plugin=protoc-gen-micro=%GOPATH%\bin\protoc-gen-micro.exe --micro_out=. --go_out=. proto_device.proto

protoc --plugin=protoc-gen-go=%GOPATH%\bin\protoc-gen-go.exe --plugin=protoc-gen-micro=%GOPATH%\bin\protoc-gen-micro.exe --micro_out=. --go_out=. proto_push.proto

protoc --plugin=protoc-gen-go=%GOPATH%\bin\protoc-gen-go.exe --plugin=protoc-gen-micro=%GOPATH%\bin\protoc-gen-micro.exe --micro_out=. --go_out=. proto_user.proto

echo "构建proto完成"