package main

import (
	"google.golang.org/grpc"
	"log"
	"github.com/ProjectAthenaa/perimeterx-service/services"
	px "github.com/ProjectAthenaa/perimeterx-service/services/protos"
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
