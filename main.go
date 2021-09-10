package main

import (
	"github.com/ProjectAthenaa/perimeterx-service/services"
	protos "github.com/ProjectAthenaa/sonic-core/sonic/antibots/perimeterx"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()

	protos.RegisterPerimeterXServer(server, services.Server{})

	if err = server.Serve(listener); err != nil{
		log.Fatal(err)
	}
}
