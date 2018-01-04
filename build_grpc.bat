@echo off

cd ..
protoc -I staff\ staff\staff.proto --go_out=plugins=grpc:staff
protoc --plugin=protoc-gen-ts=.\staff\website\node_modules\.bin\protoc-gen-ts.cmd --js_out=import_style=commonjs,binary:staff --ts_out=service=true:staff -I .\staff .\staff\staff.proto

cd .\staff

move .\staff_pb_service.ts .\website\src\rpc
move .\staff_pb.d.ts .\website\src\rpc
move .\staff_pb.js .\website\src\rpc