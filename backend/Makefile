run-accountsrv:
	go run src/account-service/cmd/main.go

run-productsrv:
	go run src/product-service/cmd/main.go

gen-protoc-account:
	protoc --go_out=. --go-grpc_out=. .\pkg\proto\account.proto

gen-protoc-user:
	protoc --go_out=. --go-grpc_out=. .\pkg\proto\user.proto

gen-protoc-product:
	protoc --go_out=. --go-grpc_out=. .\pkg\proto\product.proto

mock-gen-controller:
	mockgen -source="./src/product-service/repository/product-repository.go" -destination="./src/product-service/mock/product-repository_mock.go" -package=service

