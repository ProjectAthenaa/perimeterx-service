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
	//log.Info("args received")
	rand.Seed(time.Now().UnixNano())
	site := siteconstants.SiteDataHolder[payload.Site]

	switch payload.Type {
	case px.PXType_PX2:
		bytePayload, err := generator.GenPX2(&site, payload.Uuid)
		return &px.ConstructPayloadResponse{
			Cookie:  "",
			Payload: bytePayload,
		}, err
	case px.PXType_PX3, px.PXType_PX34:
		cookie, resObj := responsedeob.SplitResponse(payload.ResponseObject)
		bytePayload, err := generator.GenPX3(&site, payload.Uuid, resObj)
		return &px.ConstructPayloadResponse{
			Cookie:  cookie,
			Payload: bytePayload,
		}, err
	case px.PXType_RE, px.PXType_BRE, px.PXType_MME, px.PXType_MOE, px.PXType_UAE, px.PXType_EVENT:
		cookie, resObj := responsedeob.SplitResponse(payload.ResponseObject)
		bytePayload, err := generator.GenEvents(&site, payload.Uuid, resObj, int(payload.RSC))
		return &px.ConstructPayloadResponse{
			Cookie:  cookie,
			Payload: bytePayload,
		}, err
	case px.PXType_HCAPHIGH, px.PXType_HCAPLOW:
		cookie, resobj := responsedeob.SplitResponse(payload.ResponseObject)
		bytePayload, err := generator.GenHoldCaptcha(&site, payload.Uuid, resobj, resobj.HTMLCITOKEN)
		return &px.ConstructPayloadResponse{
			Cookie:  cookie,
			Payload: bytePayload,
		}, err
	//todo implement
	//case px.PXType_RECAP:
	//	cookie, resobj := responsedeob.SplitResponse(payload.ResponseObject)
	//	bytePayload, err := generator.GenReCaptcha(&site, payload.Uuid, resobj)
	//	return &px.ConstructPayloadResponse{
	//		Cookie:  cookie,
	//		Payload: bytePayload,
	//	}, err
	}

	return nil, errors.New("could not create payload")
}

func (s Server) GetCookie(ctx context.Context, req *px.GetCookieRequest) (*px.Cookie, error){
	cookie, _ := responsedeob.SplitResponse(req.PXResponse)
	return &px.Cookie{Name: "_px3", Value: cookie}, nil
}

func (s Server) GetPXde(ctx context.Context, req *px.GetCookieRequest) (*px.Cookie, error){
	_, resobj := responsedeob.SplitResponse(req.PXResponse)
	return &px.Cookie{Name:"_pxde", Value: resobj.EN}, nil
}