package services_test

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"main/helpers"
	"main/services"
	px "main/services/protos"
	"net"
	"testing"
)

const bufSize = 1024 * 1024

var (
	ctx       = context.Background()
	lis       *bufconn.Listener
	userAgent = "TEST"
)

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	px.RegisterPerimeterXServer(s, services.Server{})
	go func() {
		log.Println("gRPC Test Server Started")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestPX2(t *testing.T) {
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	helpers.AssertEqual(t, nil, err)
	defer conn.Close()

	client := px.NewPerimeterXClient(conn)

	resp, err := client.PX2(ctx, &px.PX2Request{
		Site:      px.SITE_SOLEBOX,
		UserAgent: &userAgent,
	})
	helpers.AssertEqual(t, nil, err)
	helpers.AssertEqual(t, "pxur63h57z", resp.AppID)
	helpers.AssertEqual(t, "0", resp.SEQ)
	helpers.AssertEqual(t, "NTA", resp.EN)
	helpers.AssertUnequal(t, 0, len(resp.Value))
	helpers.AssertUnequal(t, 0, len(resp.PC))
	helpers.AssertUnequal(t, 0, len(resp.UUID))
}

//TODO Finish PX tests
//func TestPX3(t *testing.T){
//	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
//	helpers.AssertEqual(t, nil, err)
//	defer conn.Close()
//
//	client := px.NewPerimeterXClient(conn)
//
//	resp, err := client.PX3(ctx, &px.PXRequest{
//		Site:           px.SITE_ONYGO,
//		Type:           px.PXType_PX3,
//		VarObject:      "",
//		Version:        "",
//		Tag:            "",
//		UA:             "",
//		Count:          "",
//		UUID:           "",
//		SID:            nil,
//		VID:            "",
//		CS:             "",
//		PXHD:           nil,
//		RecaptchaToken: nil,
//	})
//}
