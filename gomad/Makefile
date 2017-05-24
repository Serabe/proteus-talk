compile:
	grpc_tools_node_protoc \
		--js_out=import_style=commonjs,binary,Mgithub.com/gogo/protobuf/gogoproto/gogo.proto,'':../ \
		--grpc_out=../ \
		--plugin=protoc-gen-grpc=`which grpc_tools_node_protoc_plugin` \
		--proto_path=$(GOPATH)/src/:. generated.proto
