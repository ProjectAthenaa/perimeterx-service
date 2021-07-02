package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	px "main/services/protos"
	"main/services"
)

func main() {
	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()

	px.RegisterPerimeterXServer(server, services.Server{})

	if err = server.Serve(listener); err != nil{
		log.Fatal(err)
	}

}
