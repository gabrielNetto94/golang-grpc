package main

import (
	"log"

	"golang-grpc/configs"
	pb "golang-grpc/greet/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {

	opts := []grpc.DialOption{}

	if configs.GetEnv("SSL") == "TRUE" {
		certFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")

		if err != nil {
			log.Fatalf("Error while loading CA trust certificate: %v\n", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		creds := grpc.WithTransportCredentials(insecure.NewCredentials())
		opts = append(opts, creds)
	}

	conn, err := grpc.NewClient(addr, opts...)
	if err != nil {
		log.Fatal("Error ", err.Error())
	}
	defer conn.Close()
	client := pb.NewGreetServiceClient(conn)

	doGreet(client)
	// doGreetManytimes(client)
	// doGreetEveryone(client)
	// doLongGreet(client)

	// doGreetWithDeadline(client, 1*time.Second)

	// router := gin.Default()

	// router.GET("/test", func(ctx *gin.Context) {

	// 	client := pb.NewGreetServiceClient(conn)
	// 	fmt.Println(client)

	// 	ctx.JSON(200, DoLongGreet(client))

	// })
	// router.Run(":5000")

}
