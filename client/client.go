package main

import (
	"context"
	"fmt"
	"project-go/pb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("Oi do cliente ...")

	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	client := pb.NewFactorialServiceClient(cc)
	request := &pb.FactorialRequest{Numero: "5"}

	resp, _ := client.Factorial(context.Background(), request)
	fmt.Printf("Resultado recebido => [%v]", resp.Resultado)
}