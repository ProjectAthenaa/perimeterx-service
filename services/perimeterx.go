package services

import (
	"context"
	"errors"
	"github.com/ProjectAthenaa/perimeterx-service/generator"
	"github.com/ProjectAthenaa/perimeterx-service/responsedeob"
	"github.com/ProjectAthenaa/perimeterx-service/siteconstants"
	px "github.com/ProjectAthenaa/sonic-core/sonic/antibots/perimeterx"
	"math/rand"
	"time"
)

type Server struct {
	px.UnimplementedPerimeterXServer
}

func (s Server) ConstructPayload(ctx context.Context, payload *px.Payload) (*px.ConstructPayloadResponse, error) {
	rand.Seed(time.Now().UnixNano())
	SITE := siteconstants.SiteDataHolder[payload.Site]
	TYPE := payload.Type

	var UUID, TOKEN, COOKIE string
	var RSC int
	var resobj responsedeob.ResponseJSON

	if TYPE == px.PXType_PX3 || TYPE == px.PXType_PX34 {
		UUID = generator.GenerateUUID()
	} else if TYPE == px.PXType_HCAPHIGH || TYPE == px.PXType_RECAP || TYPE == px.PXType_HCAPLOW {
		COOKIE, resobj = responsedeob.SplitResponse(payload.ResponseObject)
		TOKEN = resobj.HTMLCITOKEN
	} else {
		COOKIE, resobj = responsedeob.SplitResponse(payload.ResponseObject)
		RSC = int(payload.RSC)
	}

	switch TYPE {
	case px.PXType_PX3, px.PXType_PX34:
		bytePayload, err := generator.GenInit(&SITE, UUID)
		return &px.ConstructPayloadResponse{
			Cookie:  COOKIE,
			Payload: bytePayload,
		}, err
	case px.PXType_RE, px.PXType_BRE, px.PXType_MME, px.PXType_MOE, px.PXType_UAE, px.PXType_EVENT:
		bytePayload, err := generator.GenEvents(&SITE, UUID, resobj, RSC)
		return &px.ConstructPayloadResponse{
			Cookie:  COOKIE,
			Payload: bytePayload,
		}, err
	case px.PXType_HCAPHIGH, px.PXType_HCAPLOW:
		bytePayload, err := generator.GenHoldCaptcha(&SITE, UUID, resobj, TOKEN)
		return &px.ConstructPayloadResponse{
			Cookie:  COOKIE,
			Payload: bytePayload,
		}, err
	case px.PXType_RECAP:
		bytePayload, err := generator.GenReCaptcha(&SITE, UUID, resobj)
		return &px.ConstructPayloadResponse{
			Cookie:  COOKIE,
			Payload: bytePayload,
		}, err
	}

	return nil, errors.New("could not create payload")
}

func (s Server) GetCookie(ctx context.Context, req *px.GetCookieRequest) (*px.Cookie, error){
	cookie, _ := responsedeob.SplitResponse(req.PXResponse)
	return &px.Cookie{Value: cookie}, nil
}