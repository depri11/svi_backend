init: proto
	go mod tidy

proto:
	@echo Generating article microservice proto
	cd pkg/proto && protoc --go_out=. --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=. article.proto

api:
	go run ./api_gateway/cmd/main.go

article-service:
	cd article_service & go run ./article_service/cmd/main.go