package main

import (
	"main/services"
	px "main/services/protos"
)

func main() {
	services.LocalGetCookie(px.SITE_WALMART, "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Safari/537.36")

	//listener, err := net.Listen("tcp", ":3000")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//server := grpc.NewServer()
	//
	//px.RegisterPerimeterXServer(server, services.Server{})
	//
	//if err = server.Serve(listener); err != nil{
	//	log.Fatal(err)
	//}

}
