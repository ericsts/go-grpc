package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/ericsts/go-grpc/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could Not Log do grpServer", err)
	}
	defer connection.Close()

	client := pb.NewUserServiceClient(connection)
	// AddUser(client)
	// AddUserVerbose(client)
	// AddUsers(client)
	AddUsersStreamBoth(client)
}

func AddUser(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Jony",
		Email: "jony@jony.com",
	}

	res, err := client.AddUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Erro call client RPC request %v", err)
	}
	fmt.Println(res)
}

func AddUserVerbose(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Jony",
		Email: "jony@jony.com",
	}

	resStream, err := client.AddUserVerbose(context.Background(), req)
	if err != nil {
		log.Fatalf("Erro call client RPC request %v", err)
	}

	for {
		stream, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Erro could not receive %v", err)
		}
		fmt.Println("Status: ", stream.Status, " - ", stream.GetUser())
	}
}

func AddUsers(client pb.UserServiceClient) {

	reqs := []*pb.User{
		&pb.User{
			Id:    "w1",
			Name:  "Nome",
			Email: "Emil",
		},
		&pb.User{
			Id:    "w2",
			Name:  "Nome2",
			Email: "Emil",
		},
		&pb.User{
			Id:    "w3",
			Name:  "Nome3",
			Email: "Emil",
		},
		&pb.User{
			Id:    "w4",
			Name:  "Nome4",
			Email: "Emil",
		},
		&pb.User{
			Id:    "w5",
			Name:  "Nome5",
			Email: "Emil",
		},
	}
	stream, err := client.AddUsers(context.Background())
	if err != nil {
		log.Fatalf("Error creating request", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 2)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receiving response", err)
	}

	fmt.Println(res)
}

func AddUsersStreamBoth(client pb.UserServiceClient) {

	stream, err := client.AddUserStreamBoth(context.Background())
	if err != nil {
		log.Fatalf("Error creating request Stream", err)
	}
	reqs := []*pb.User{
		&pb.User{
			Id:    "w1",
			Name:  "Nome",
			Email: "Emil",
		},
		&pb.User{
			Id:    "w2",
			Name:  "Nome2",
			Email: "Emil",
		},
		&pb.User{
			Id:    "w3",
			Name:  "Nome3",
			Email: "Emil",
		},
		&pb.User{
			Id:    "w4",
			Name:  "Nome4",
			Email: "Emil",
		},
		&pb.User{
			Id:    "w5",
			Name:  "Nome5",
			Email: "Emil",
		},
	}

	wait := make(chan int)

	go func() {
		for _, req := range reqs {
			fmt.Println("Sending user: ", req.Name)
			stream.Send(req)
			time.Sleep(time.Second * 2)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Erro receiving data", err)
				break
			}
			fmt.Printf("Receiving User %v with statu %v\n", res.GetUser().GetName(), res.GetStatus())
		}
		close(wait)
	}()

	<-wait
}
