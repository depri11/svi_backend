proto:
	@echo Generating article microservice proto
	cd pkg/proto && protoc --go_out=. --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=. article.proto