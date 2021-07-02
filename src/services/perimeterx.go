package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	pxUtils "github.com/incizzle/perimeterx-utils-go"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"log"
	"main/helpers"
	"main/payloads"
	px "main/services/protos"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

type Server struct {
	px.UnimplementedPerimeterXServer
}

func (s Server) PX2(_ context.Context, request *px.PX2Request) (*px.Payload, error) {
	uid := uuid.NewV4().String()
	var encArr []interface{}

	site := helpers.ConvertToSite(request.Site)

	addReqObj := payloads.PX2{}
	addReqObj.Instantiate(site.Site, uid)
	encArr = append(encArr, addReqObj)

	rawPayload, err := json.Marshal(encArr)
	if err != nil {
		return nil, err
	}
	payload := &px.Payload{
		Value: pxUtils.EncodePayload(string(rawPayload), 50),
		AppID: site.Sitekey,
		UUID:  uid,
		SEQ:   "0",
		EN:    "NTA",
		PC:    pxUtils.CreatePC(string(rawPayload), fmt.Sprintf("%s:%s:%s", uid, request.Version, request.Tag)),
		Ua:    request.UserAgent,
		RSC:   int32(0),
	}

	return payload, nil
}

func (s Server) PX3(_ context.Context, request *px.PXRequest) (*px.Payload, error) {
	var encArr []interface{}

	site := helpers.ConvertToSite(request.Site)
	var sid string
	uid := request.UUID
	vid := strings.SplitN(request.VID, "|", 2)[0]
	cs := request.CS
	PXHD := request.PXHD
	count, err := strconv.Atoi(request.Count)
	if err != nil {
		return nil, err
	}

	if request.SID == nil {
		sid = uid
	} else {
		sid = *request.SID
	}

	switch request.Type {
	case px.PXType_PX3:
		if request.VarObject == "undefined" {
			return nil, noResponseToParse
		}

		addReqObj := payloads.PX3{}
		addReqObj.Instantiate(site.Site, uid, request.VarObject, request.UA)
		encArr = append(encArr, addReqObj)
	case px.PXType_PX4:
		if request.VarObject == "undefined" {
			return nil, noResponseToParse
		}
		px4add := payloads.PX4{}
		px4add.Instantiate(uid, site.Sitekey, request.UA, 2)
		encArr = append(encArr, px4add)
	case px.PXType_PX34:
		if request.VarObject == "undefined" {
			return nil, noResponseToParse
		}
		px3add := payloads.PX3{}
		px3add.Instantiate(site.Site, uid, request.VarObject, request.UA)
		px4add := payloads.PX4{}
		px4add.Instantiate(uid, site.Sitekey, request.UA, 2)
		encArr = append(encArr, px3add, px4add)
	case px.PXType_EVENT:
		encArr = payloads.MixedArrayGenerate(site.Site, uid, sid, vid, request.UA, 4+rand.Intn(10), 2+rand.Intn(2))
	case px.PXType_MOE:
		addReqObj := payloads.MouseOverEvent{}
		addReqObj.Instantiate(sid, uid, site.Sitekey, 2+rand.Intn(3))
		encArr = append(encArr, addReqObj)
	case px.PXType_MME:
		addReqObj := payloads.MouseMoveEvent{}
		addReqObj.Instantiate(sid, uid, 2+rand.Intn(3))
		encArr = append(encArr, addReqObj)
	case px.PXType_RE:
		for _, v := range payloads.RequestEventGenerate(site.Site, uid, sid, vid, 4+rand.Intn(10), 2+rand.Intn(2)) {
			encArr = append(encArr, v)
		}
	case px.PXType_UAE:
		for _, v := range payloads.UserAgentEventGenerate(site.Site, uid, sid, vid, request.UA, 4+rand.Intn(10), 2+rand.Intn(2)) {
			encArr = append(encArr, v)
		}
	case px.PXType_BRE:
		for _, v := range payloads.BrowserRequestEventGenerate(site.Site, uid, sid, vid, 4+rand.Intn(10), 2+rand.Intn(2)) {
			encArr = append(encArr, v)
		}
	case px.PXType_HCAPLOW:
		encArr = payloads.LOWSECHCAPGENERATE(uid, vid, site.Site, site.Sitekey)
	case px.PXType_HCAPHIGH:
		encArr = payloads.HIGHSECHCAPGENERATE(uid, vid, site.Site, site.Sitekey, request.UA)
	case px.PXType_RECAP:
		token := request.RecaptchaToken
		if token == nil {
			return nil, missingRecaptchaToken
		}
		encArr = append(encArr, payloads.RECAPGENERATE(uid, vid, site.Site, site.Sitekey, *token))
	}

	rawPayload, err := json.Marshal(encArr)
	if err != nil {
		return nil, err
	}

	resObj := px.Payload{
		Value: pxUtils.EncodePayload(string(rawPayload), 50),
		AppID: site.Sitekey,
		Tag:   request.Version,
		UUID:  uid,
		FT:    request.Tag,
		SEQ:   strconv.Itoa(count),
		EN:    "NTA",
		PC:    pxUtils.CreatePC(string(rawPayload), fmt.Sprintf("%s:%s:%s", uid, request.Version, request.Tag)),
		Ua:    request.UA,
	}

	if request.SID != nil {
		resObj.CS = &cs
		resObj.SID = &sid
		resObj.VID = &vid
		if count <= 1 {
			resObj.RSC = int32(count + 1)
		} else {
			resObj.RSC = int32(count + 1)
		}

		if PXHD != nil {
			resObj.PXHD = PXHD
		}

	}

	return &resObj, nil
}

func LOCALPX2(sitein px.SITE, uain string) (*px.Payload, error) {
	uid := uuid.NewV4().String()
	var encArr []interface{}

	site := helpers.ConvertToSite(sitein)

	addReqObj := payloads.PX2{}
	addReqObj.Instantiate(site.Site, uid)
	encArr = append(encArr, addReqObj)

	rawPayload, err := json.Marshal(encArr)
	if err != nil {
		return nil, err
	}
	payload := &px.Payload{
		Value: pxUtils.EncodePayload(string(rawPayload), 50),
		AppID: site.Sitekey,
		UUID:  uid,
		SEQ:   "0",
		EN:    "NTA",
		Tag:   site.Tag,
		FT:    site.Version,
		PC:    pxUtils.CreatePC(string(rawPayload), fmt.Sprintf("%s:%s:%s", uid, site.Version, site.Tag)),
		Ua:    uain,
		RSC:   int32(1),
	}

	return payload, nil
}

func LOCALPX3(_ context.Context, request *px.PXRequest) (*px.Payload, error) {
	var encArr []interface{}

	site := helpers.ConvertToSite(request.Site)
	var sid string
	uid := request.UUID
	vid := strings.SplitN(request.VID, "|", 2)[0]
	cs := request.CS
	PXHD := request.PXHD
	count, err := strconv.Atoi(request.Count)
	if err != nil {
		return nil, err
	}

	if request.SID == nil {
		sid = uid
	} else {
		sid = *request.SID
	}

	switch request.Type {
	case px.PXType_PX3:
		if request.VarObject == "undefined" {
			return nil, noResponseToParse
		}

		addReqObj := payloads.PX3{}
		addReqObj.Instantiate(site.Site, uid, request.VarObject, request.UA)
		encArr = append(encArr, addReqObj)
	case px.PXType_PX4:
		if request.VarObject == "undefined" {
			return nil, noResponseToParse
		}
		px4add := payloads.PX4{}
		px4add.Instantiate(uid, site.Sitekey, request.UA, 2)
		encArr = append(encArr, px4add)
	case px.PXType_PX34:
		if request.VarObject == "undefined" {
			return nil, noResponseToParse
		}
		px3add := payloads.PX3{}
		px3add.Instantiate(site.Site, uid, request.VarObject, request.UA)
		px4add := payloads.PX4{}
		px4add.Instantiate(uid, site.Sitekey, request.UA, 2)
		encArr = append(encArr, px3add, px4add)
	case px.PXType_EVENT:
		encArr = payloads.MixedArrayGenerate(site.Site, uid, sid, vid, request.UA, 4+rand.Intn(10), 2+rand.Intn(2))
	case px.PXType_MOE:
		addReqObj := payloads.MouseOverEvent{}
		addReqObj.Instantiate(sid, uid, site.Sitekey, 2+rand.Intn(3))
		encArr = append(encArr, addReqObj)
	case px.PXType_MME:
		addReqObj := payloads.MouseMoveEvent{}
		addReqObj.Instantiate(sid, uid, 2+rand.Intn(3))
		encArr = append(encArr, addReqObj)
	case px.PXType_RE:
		for _, v := range payloads.RequestEventGenerate(site.Site, uid, sid, vid, 4+rand.Intn(10), 2+rand.Intn(2)) {
			encArr = append(encArr, v)
		}
	case px.PXType_UAE:
		for _, v := range payloads.UserAgentEventGenerate(site.Site, uid, sid, vid, request.UA, 4+rand.Intn(10), 2+rand.Intn(2)) {
			encArr = append(encArr, v)
		}
	case px.PXType_BRE:
		for _, v := range payloads.BrowserRequestEventGenerate(site.Site, uid, sid, vid, 4+rand.Intn(10), 2+rand.Intn(2)) {
			encArr = append(encArr, v)
		}
	case px.PXType_HCAPLOW:
		encArr = payloads.LOWSECHCAPGENERATE(uid, vid, site.Site, site.Sitekey)
	case px.PXType_HCAPHIGH:
		encArr = payloads.HIGHSECHCAPGENERATE(uid, vid, site.Site, site.Sitekey, request.UA)
	case px.PXType_RECAP:
		token := request.RecaptchaToken
		if token == nil {
			return nil, missingRecaptchaToken
		}
		encArr = append(encArr, payloads.RECAPGENERATE(uid, vid, site.Site, site.Sitekey, *token))
	}

	rawPayload, err := json.Marshal(encArr)
	if err != nil {
		return nil, err
	}

	resObj := px.Payload{
		Value: pxUtils.EncodePayload(string(rawPayload), 50),
		AppID: site.Sitekey,
		Tag:   request.Version,
		UUID:  uid,
		FT:    request.Tag,
		SEQ:   strconv.Itoa(count),
		EN:    "NTA",
		PC:    pxUtils.CreatePC(string(rawPayload), fmt.Sprintf("%s:%s:%s", uid, request.Version, request.Tag)),
		Ua:    request.UA,
	}

	if request.SID != nil {
		resObj.CS = &cs
		resObj.SID = &sid
		resObj.VID = &vid
		if count <= 1 {
			resObj.RSC = int32(count + 1)
		} else {
			resObj.RSC = int32(count + 1)
		}

		if PXHD != nil {
			resObj.PXHD = PXHD
		}

	}

	return &resObj, nil
}

func (s Server) GetCookie(ctx context.Context, request *px.CookieRequest) (*px.Cookie, error) {
	panic("implement me")
}

func LocalGetCookie(sitein px.SITE, ua string) (string, error) {
	encpayload, _ := LOCALPX2(sitein, ua)
	client := http.Client{}
	bodystring := []byte(fmt.Sprintf("payload=%s=&appId=%s&tag=%s&uuid=%s&ft=%s&seq=%s&en=%s&pc=%s&rsc=%s", encpayload.Value, encpayload.AppID, encpayload.Tag, encpayload.UUID, encpayload.FT, encpayload.SEQ, encpayload.EN, encpayload.PC, strconv.Itoa(int(encpayload.RSC))))
	req, _ := http.NewRequest("POST", "https://collector-pxu6b0qd2s.px-cloud.net/api/v2/collector", bytes.NewBuffer(bodystring))
	req.Header.Set("User-Agent", ua)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//req.Form = bodyform
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	log.Print(string(body))
	return string(body), nil
}
