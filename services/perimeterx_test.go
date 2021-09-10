package services_test

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"github.com/ProjectAthenaa/perimeterx-service/helpers"
	"github.com/ProjectAthenaa/perimeterx-service/services"
	px "github.com/ProjectAthenaa/perimeterx-service/services/protos"
	"net"
	"testing"
)

const bufSize = 1024 * 1024

var (
	ctx       = context.Background()
	lis       *bufconn.Listener
	userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"
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

func TestServer_GetCookie(t *testing.T) {
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	helpers.AssertEqual(t, nil, err)
	defer conn.Close()

	client := px.NewPerimeterXClient(conn)

	resp, err := client.GetCookie(ctx, &px.CookieRequest{
		Site:      px.SITE_WALMART,
		UserAgent: userAgent,
	})
	helpers.AssertEqual(t, nil, err)
	helpers.AssertEqual(t, 2, len(resp.Cookie))

}
