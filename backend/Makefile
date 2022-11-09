run-accountsrv:
	go run internal/account-service/cmd/main.go

run-productsrv:
	go run internal/product-service/cmd/main.go

gen-protoc-account:
	protoc --go_out=. --go-grpc_out=. .\pkg\proto\account.proto

gen-protoc-user:
	protoc --go_out=. --go-grpc_out=. .\pkg\proto\user.proto

gen-protoc-product:
	protoc --go_out=. --go-grpc_out=. .\pkg\proto\product.proto

gen-protoc-auction:
	protoc --go_out=. --go-grpc_out=. .\pkg\proto\auction.proto

mock-gen-controller:
	mockgen -source="./internal/product-service/repository/product-repository.go" -destination="./internal/product-service/mock/product-repository_mock.go" -package=service

