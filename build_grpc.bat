cd ..
protoc -I staff\ staff\staff.proto --go_out=plugins=grpc:staff
protoc --plugin=protoc-gen-ts=.\staff\website\node_modules\.bin\protoc-gen-ts.cmd --js_out=import_style=commonjs,binary:staff --ts_out=service=true:staff -I .\staff .\staff\staff.proto