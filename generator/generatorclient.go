package generator

import (
	"github.com/ProjectAthenaa/perimeterx/responsedeob"
	"github.com/ProjectAthenaa/perimeterx/siteconstants"
	"encoding/json"
	"math/rand"
	"strconv"
)


type PayloadOut struct{
	Payload string `json:"payload"`
	AppID	string `json:"appId"`
	Tag		string `json:"tag"`
	Uuid	string `json:"uuid"`
	Ft		string `json:"ft"`
	Seq		string `json:"seq"`
	En		string `json:"en"`
	Pc		string `json:"pc"`
	Sid		string `json:"sid,omitempty"`
	Vid		string `json:"vid,omitempty"`
	Cts		string `json:"cts,omitempty"`
	Rsc		string `json:"rsc"`
}

func GenInit(sitedata *siteconstants.SiteData, UUID string) ([]byte, error){
	var payload PayloadOut
	GeneratePX2(UUID, sitedata)

	var tempresobj responsedeob.ResponseJSON
	GeneratePX3(UUID, tempresobj, sitedata)
	return json.Marshal(payload)
}

func GenEvents(sitedata *siteconstants.SiteData, UUID string, resobj responsedeob.ResponseJSON, rsc int) ([]byte, error){
	var payload PayloadOut

	switch rand.Intn(4){
	case 0:
		GenerateMouseOverEvent(resobj.SID, resobj.VID, UUID, sitedata, rsc)
	case 1:
		GenerateMouseMoveEvent(UUID, sitedata, rsc)
	case 2:
		GenerateRequestEvent(sitedata)
	case 3:
		GenerateUserAgentEvent(resobj.SID, resobj.VID, UUID, sitedata)
	}

	return json.Marshal(payload)
}

func GenHoldCaptcha(sitedata *siteconstants.SiteData, UUID string, resobj responsedeob.ResponseJSON, token string) ([]byte, error){
	var payload PayloadOut
	duration, _ := strconv.Atoi(resobj.CI)
	GenerateHCAP(UUID, resobj.VID, token, resobj.CP, resobj.CITOKEN, duration, sitedata)
	return json.Marshal(payload)
}

func GenReCaptcha(sitedata *siteconstants.SiteData, UUID string, resobj responsedeob.ResponseJSON) ([]byte, error){
	var payload PayloadOut

	GenReCaptcha(sitedata, UUID, resobj)
	return json.Marshal(payload)
}