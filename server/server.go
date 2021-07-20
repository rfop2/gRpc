package main

import (
	"context"
	"fmt"
	"project-go/pb"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)

func factorial(x int) (result int) {
	if x == 0 {
		result = 1;
	  } else {
		result = x * factorial(x - 1);
	  }
	  return;
}

type server struct {
	pb.UnimplementedFactorialServiceServer
}

func (*server) Factorial(ctx context.Context, request *pb.FactorialRequest) (*pb.FactorialResponse, error) {
	numero := request.Numero

	//conversao str -> int
	numeroAux := 0
	msg := string(numero)
	fmt.Sscan(msg, &numeroAux)
		
	resultado := factorial(numeroAux)

	//conversao int -> str
	novamensagem := strconv.Itoa(resultado)

	response := &pb.FactorialResponse{
		Resultado: "O fatorial é " + novamensagem,
	}
	return response, nil
}

func (service *server) mustEmbedUnimplementedFactorialServiceServer() {}



func main() {
	address := "0.0.0.0:50051"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	fmt.Printf("Servidor está ouvindo na porta:  %v ...", address)

	s := grpc.NewServer()
	pb.RegisterFactorialServiceServer(s, &server{})

	s.Serve(lis)
}