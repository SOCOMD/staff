cd ../
protoc -I staff/ staff/staff.proto --go_out=plugins=grpc:staff
protoc \
	--plugin=protoc-gen-ts=./staff/website/node_modules/.bin/protoc-gen-ts \
	--js_out=import_style=commonjs,binary:staff \
	--ts_out=service=true:staff \
	-I ./staff \
	staff/staff.proto 

cd ./staff

mv staff_pb_service.ts ./website/src/rpc/staff_pb_service.ts
mv staff_pb.d.ts ./website/src/rpc/staff_pb.d.ts
mv staff_pb.js ./website/src/rpc/staff_pb.js