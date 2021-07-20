# gRpc
middleware gRpc escrito em golang para a cadeira programação concorrente

Modúlos para instalar: 

go get google.golang.org/grpc
go get google.golang.org/protobuf
google.golang.org/grpc/cmd/protoc-gen-go-grpc

para rodar o proto: protoc --go_out=. --go-grpc_out=. proto/*.proto

O serviço factorial define um serviço, e o rpc define a fucionalidade do serviço.
Nesser serviço factorialServiço espera como entrada(factorialRequest) um numero e devolve(factorialResponse) o fatorial dele.
