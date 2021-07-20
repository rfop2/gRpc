package main

import (
	"context"
	"fmt"
	"project-go/pb"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	fmt.Println("Oi do cliente ...")

	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	start := time.Now()

	client := pb.NewFactorialServiceClient(cc)
	request := &pb.FactorialRequest{Numero: "20"}

	resp, _ := client.Factorial(context.Background(), request)
	end := time.Now()

	fmt.Printf("Resultado recebido => [%v]", resp.Resultado)
	
	tempo := end.Sub(start)
	fmt.Println("tempo: ", tempo.Seconds())
}