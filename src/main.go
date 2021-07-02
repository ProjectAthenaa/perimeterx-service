package main

import (
	"google.golang.org/grpc"
	"log"
	"main/services"
	px "main/services/protos"
	"net"
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
