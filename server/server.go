package main

import (
	"context"
	"fmt"
	"project-go/pb"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
	"time"
)

var tempoTotal = 0.0

// tempo que dura uma solicitacao
func timeTrack(start time.Time, name string) {
    elapsed := time.Since(start)
    log.Println("%s demorou %s", name, elapsed)
	tempoTotal = tempoTotal + elapsed.Seconds();
	fmt.Println("", tempoTotal, " segundos")
}


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
	address := ":50051"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	fmt.Println("Servidor está ouvindo na porta:   ...", address)

	s := grpc.NewServer()
	pb.RegisterFactorialServiceServer(s, &server{})

	s.Serve(lis)
}