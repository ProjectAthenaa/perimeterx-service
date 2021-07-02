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

func PX2RESHANDLER(rawres string) map[string]string {
	resjson := struct {
		Do []string `json:"do"`
	}{}
	json.Unmarshal([]byte(rawres), &resjson)
	objOut := make(map[string]string)
	for _, v := range resjson.Do{
		splitstr := strings.SplitN(v, "|", 2)
		objOut[splitstr[0]] = splitstr[1]
	}
	return objOut
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
func LOCALPX3(sitein px.SITE, reqtype px.PXType, response, countin, uuid, ua, captchatoken string) (*px.Payload, error) {
	var encArr []interface{}

	resobj := PX2RESHANDLER(response)

	site := helpers.ConvertToSite(sitein)
	sid := resobj["sid"]
	vid := strings.SplitN(resobj["vid"], "|", 2)[0]
	count, err := strconv.Atoi(countin)
	if err != nil {
		return nil, err
	}

	if sid == "" {
		sid = uuid
	}

	switch reqtype {
	case px.PXType_PX3:
		if response == "" {
			return nil, noResponseToParse
		}
		addReqObj := payloads.PX3{}
		addReqObj.Instantiate(site.Site, uuid, response, ua)
		encArr = append(encArr, addReqObj)
	case px.PXType_PX4:
		if response == "" {
			return nil, noResponseToParse
		}
		px4add := payloads.PX4{}
		px4add.Instantiate(uuid, site.Sitekey, ua, 2)
		encArr = append(encArr, px4add)
	case px.PXType_PX34:
		if response == "" {
			return nil, noResponseToParse
		}
		px3add := payloads.PX3{}
		px3add.Instantiate(site.Site, uuid, response, ua)
		px4add := payloads.PX4{}
		px4add.Instantiate(uuid, site.Sitekey, ua, 2)
		encArr = append(encArr, px3add, px4add)
	case px.PXType_EVENT:
		encArr = payloads.MixedArrayGenerate(site.Site, uuid, sid, vid, ua, 4+rand.Intn(10), 2+rand.Intn(2))
	case px.PXType_MOE:
		addReqObj := payloads.MouseOverEvent{}
		addReqObj.Instantiate(sid, uuid, site.Sitekey, 2+rand.Intn(3))
		encArr = append(encArr, addReqObj)
	case px.PXType_MME:
		addReqObj := payloads.MouseMoveEvent{}
		addReqObj.Instantiate(sid, uuid, 2+rand.Intn(3))
		encArr = append(encArr, addReqObj)
	case px.PXType_RE:
		for _, v := range payloads.RequestEventGenerate(site.Site, uuid, sid, vid, 4+rand.Intn(10), 2+rand.Intn(2)) {
			encArr = append(encArr, v)
		}
	case px.PXType_UAE:
		for _, v := range payloads.UserAgentEventGenerate(site.Site, uuid, sid, vid, ua, 4+rand.Intn(10), 2+rand.Intn(2)) {
			encArr = append(encArr, v)
		}
	case px.PXType_BRE:
		for _, v := range payloads.BrowserRequestEventGenerate(site.Site, uuid, sid, vid, 4+rand.Intn(10), 2+rand.Intn(2)) {
			encArr = append(encArr, v)
		}
	case px.PXType_HCAPLOW:
		encArr = payloads.LOWSECHCAPGENERATE(uuid, vid, site.Site, site.Sitekey)
	case px.PXType_HCAPHIGH:
		encArr = payloads.HIGHSECHCAPGENERATE(uuid, vid, site.Site, site.Sitekey, ua)
	case px.PXType_RECAP:
		token := captchatoken
		if token == "" {
			return nil, missingRecaptchaToken
		}
		encArr = append(encArr, payloads.RECAPGENERATE(uuid, vid, site.Site, site.Sitekey, token))
	}

	rawPayload, err := json.Marshal(encArr)
	if err != nil {
		return nil, err
	}

	resObj := px.Payload{
		Value: pxUtils.EncodePayload(string(rawPayload), 50),
		AppID: site.Sitekey,
		Tag:   site.Version,
		UUID:  uuid,
		FT:    site.Tag,
		SEQ:   strconv.Itoa(count),
		EN:    "NTA",
		PC:    pxUtils.CreatePC(string(rawPayload), fmt.Sprintf("%s:%s:%s", uuid, site.Version, site.Tag)),
		Ua:    ua,
	}
	
	if resobj["sid"] != "" {
		cs := resobj["cs"]
		resObj.CS = &cs
		resObj.SID = &sid
		resObj.VID = &vid
		if count <= 1 {
			resObj.RSC = int32(count + 1)
		} else {
			resObj.RSC = int32(count + 1)
		}

		if val, ok := resobj["pxhd"]; ok {
			resObj.PXHD = &val
		}

	}

	return &resObj, nil
}

func (s Server) GetCookie(ctx context.Context, request *px.CookieRequest) (*px.Cookie, error) {
	panic("implement me")
}

func LocalGetCookie(sitein px.SITE, ua string) (string, error) {

	//todo
	px2encpayload, _ := LOCALPX2(sitein, ua)
	client := http.Client{}
	bodystring := []byte(fmt.Sprintf("payload=%s=&appId=%s&tag=%s&uuid=%s&ft=%s&seq=%s&en=%s&pc=%s&rsc=%s", px2encpayload.Value, px2encpayload.AppID, px2encpayload.Tag, px2encpayload.UUID, px2encpayload.FT, px2encpayload.SEQ, px2encpayload.EN, px2encpayload.PC, strconv.Itoa(int(px2encpayload.RSC))))
	req, _ := http.NewRequest("POST", "https://collector-pxu6b0qd2s.px-cloud.net/api/v2/collector", bytes.NewBuffer(bodystring))
	req.Header.Set("User-Agent", ua)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)

	//todo
	px3encpayload, _ := LOCALPX3(sitein, px.PXType_PX3, string(body), "1", px2encpayload.UUID, ua, "")
	bodystring = []byte(fmt.Sprintf("payload=%s=&appId=%s&tag=%s&uuid=%s&ft=%s&seq=%s&en=%s&pc=%s&rsc=%ssid=%S&vid=%S&cs=%s", px3encpayload.Value, px3encpayload.AppID, px3encpayload.Tag, px3encpayload.UUID, px3encpayload.FT, px3encpayload.SEQ, px3encpayload.EN, px3encpayload.PC, strconv.Itoa(int(px3encpayload.RSC)), px3encpayload.SID, px3encpayload.VID, px3encpayload.CS))
	req, _ = http.NewRequest("POST", "https://collector-pxu6b0qd2s.px-cloud.net/api/v2/collector", bytes.NewBuffer(bodystring))
	req.Header.Set("User-Agent", ua)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, _ = client.Do(req)
	body, _ = ioutil.ReadAll(resp.Body)
	log.Print(string(body))
	return string(body), nil
}
