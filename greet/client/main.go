package main

import (
	"log"
	"time"

	pb "golang-grpc/greet/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {

	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error ", err.Error())
	}
	defer conn.Close()
	client := pb.NewGreetServiceClient(conn)

	// doGreet(client)
	// doGreetManytimes(client)
	// doGreetEveryone(client)
	// doLongGreet(client)

	doGreetWithDeadline(client, 1*time.Second)

	// router := gin.Default()

	// router.GET("/test", func(ctx *gin.Context) {

	// 	client := pb.NewGreetServiceClient(conn)
	// 	fmt.Println(client)

	// 	ctx.JSON(200, DoLongGreet(client))

	// })
	// router.Run(":5000")

}
