package generator

import (
	"errors"
	"fmt"
	"github.com/ProjectAthenaa/perimeterx-service/responsedeob"
	"github.com/ProjectAthenaa/perimeterx-service/siteconstants"
	"github.com/ProjectAthenaa/pxutils"
	jsoniter "github.com/json-iterator/go"
	"math/rand"
	"strconv"
	"strings"
)

var json = jsoniter.ConfigFastest


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
	Cs		string `json:"cs"`
	Ci		string `json:"ci"`
}

func GenPX2(sitedata *siteconstants.SiteData, UUID string) ([]byte, error){
	px2obj := GeneratePX2(UUID, sitedata)
	rp, err := json.Marshal(px2obj)
	if err != nil{
		return nil, err
	}

	return json.Marshal(&PayloadOut{
		Payload: pxutils.EncodePayload(string(rp), 50),
		AppID:   sitedata.AppId,
		Tag:     sitedata.Tag,
		Uuid:    UUID,
		Ft:      sitedata.Ft,
		Seq:     "0",
		En:      "NTA",
		Pc:      pxutils.CreatePC(string(rp), fmt.Sprintf(`%s:%s:%s`, UUID, sitedata.Tag, sitedata.Ft)),
		Rsc:     "1",
	})
}


func GenPX3(sitedata *siteconstants.SiteData, UUID string, resobj responsedeob.ResponseJSON) ([]byte, error){
	raw, err := json.Marshal(GeneratePX3(UUID, resobj, sitedata))
	rp := string(raw)
	sts, err := strconv.Atoi(resobj.STS)
	if err != nil{
		return nil, errors.New("could not read sts value")
	}
	rp = strings.Replace(rp, "calckey", pxutils.SimpleTextEncode(resobj.CLS0, (sts%10)+2), 1)
	rp = strings.Replace(rp, "calcval", pxutils.SimpleTextEncode(resobj.CLS0, (sts%10)+1), 1)
	if err != nil{
		return nil, err
	}
	return json.Marshal(&PayloadOut{
		Payload: pxutils.EncodePayload(rp, 50),
		AppID:   sitedata.AppId,
		Tag:     sitedata.Tag,
		Uuid:    UUID,
		Ft:      sitedata.Ft,
		Seq:     "1",
		En:      "NTA",
		Pc:      pxutils.CreatePC(rp, fmt.Sprintf(`%s:%s:%s`, UUID, sitedata.Tag, sitedata.Ft)),
		Sid:     resobj.SID,
		Vid:     resobj.VID,
		Cts:     resobj.CTS,
		Rsc:     "2",
		Cs:		resobj.CS,
	})
}

func GenEvents(sitedata *siteconstants.SiteData, UUID string, resobj responsedeob.ResponseJSON, rsc int) ([]byte, error){
	var events []interface{}
	switch rand.Intn(4){
	case 0:
		for _, event := range GenerateMouseOverEvent(resobj.SID, resobj.VID, UUID, sitedata, rsc){
			events = append(events, event)
		}
	case 1:
		for _, event := range GenerateMouseMoveEvent(UUID, sitedata, rsc){
			events = append(events, event)
		}
	case 2:
		for _, event := range GenerateRequestEvent(sitedata){
			events = append(events, event)
		}
	case 3:
		for _, event := range GenerateUserAgentEvent(resobj.SID, resobj.VID, UUID, sitedata){
			events = append(events, event)
		}
	}

	rp, err := json.Marshal(events)
	if err != nil{
		return nil, err
	}

	return json.Marshal(&PayloadOut{
		Payload: pxutils.EncodePayload(string(rp), 50),
		AppID:   sitedata.AppId,
		Tag:     sitedata.Tag,
		Uuid:    UUID,
		Ft:      sitedata.Ft,
		Seq:     "3",
		En:      "NTA",
		Pc:      pxutils.CreatePC(string(rp), fmt.Sprintf(`%s:%s:%s`, UUID, sitedata.Tag, sitedata.Ft)),
		Sid:     resobj.SID,
		Vid:     resobj.VID,
		Cts:     resobj.CTS,
		Rsc:     "3",
		Cs:		resobj.CS,

	})
}

func GenHoldCaptcha(sitedata *siteconstants.SiteData, UUID string, resobj responsedeob.ResponseJSON, token string) ([]byte, error){
	duration, _ := strconv.Atoi(resobj.CI)
	rp, err := json.Marshal(GenerateHCAP(UUID, resobj.VID, token, resobj.CP, resobj.CITOKEN, duration, sitedata))
	if err != nil{
		return nil, err
	}
	return json.Marshal(&PayloadOut{
		Payload: pxutils.EncodePayload(string(rp), 50),
		AppID:   sitedata.AppId,
		Tag:     sitedata.Tag,
		Uuid:    UUID,
		Ft:      sitedata.Ft,
		Seq:     "3",
		En:      "NTA",
		Pc:      pxutils.CreatePC(string(rp), fmt.Sprintf(`%s:%s:%s`, UUID, sitedata.Tag, sitedata.Ft)),
		Sid:     resobj.SID,
		Vid:     resobj.VID,
		Cts:     resobj.CTS,
		Rsc:     "3",
		Ci: resobj.CI,
	})
}

//todo implement
//func GenReCaptcha(sitedata *siteconstants.SiteData, UUID string, resobj responsedeob.ResponseJSON) ([]byte, error){
//	var payload PayloadOut
//
//	GenReCaptcha(sitedata, UUID, resobj)
//	return json.Marshal(&payload)
//}