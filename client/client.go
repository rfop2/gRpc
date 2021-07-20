package main

import (
	"context"
	"fmt"
	"log"
	"project-go/pb"
	"time"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Oi do cliente ...")

	tempoTotal := 0.0

	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("192.168.1.82:50051", opts)
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	requisicoes := 1000

	for i := 0; i <= requisicoes; i++ {
		start := time.Now()

		client := pb.NewFactorialServiceClient(cc)
		request := &pb.FactorialRequest{Numero: "20"}

		resp, _ := client.Factorial(context.Background(), request)
		end := time.Now()

		fmt.Printf("Resultado recebido => [%v]", resp.Resultado)

		tempo := end.Sub(start)
		fmt.Println("tempo em segundos: ", tempo.Seconds())

		tempoTotal = tempo.Seconds() + tempoTotal
	}

	fmt.Println(" media: ", tempoTotal/float64(requisicoes))

}
