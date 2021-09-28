package generator

import (
	"fmt"
	"github.com/ProjectAthenaa/perimeterx-service/responsedeob"
	"github.com/ProjectAthenaa/perimeterx-service/siteconstants"
	"github.com/ProjectAthenaa/pxutils"
	jsoniter "github.com/json-iterator/go"
	"log"
	"math/rand"
	"strconv"
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


//px3
//"PX1061": [
//        {
//          "tagName": "INPUT",
//          "id": "",
//          "type": "search",
//          "name": "q",
//          "height": 39.9974365234375,
//          "width": 31.992826461791992,
//          "x": 394.3135070800781,
//          "y": 23.988216400146484
//        }
//      ],

func GenPX3(sitedata *siteconstants.SiteData, UUID string, resobj responsedeob.ResponseJSON) ([]byte, error){
	rp, err := json.Marshal(GeneratePX3(UUID, resobj, sitedata))
	log.Println(string(rp))
	if err != nil{
		return nil, err
	}
	//rp = []byte(`[{"t":"PX3","d":{"PX234":false,"PX235":false,"PX151":false,"PX239":false,"PX240":false,"PX152":false,"PX153":false,"PX314":false,"PX192":false,"PX196":false,"PX207":false,"PX251":false,"PX982":1632805361331,"PX983":"92510302980276122309","PX986":"12682647075947275688","PX985":5072,"PX1033":"e0eaf10e","PX1019":"7d688b1b","PX1020":"7766a52d","PX1021":"180dd7e3","PX1022":"6a90378d","PX1035":true,"PX1139":false,"PX1025":false,"PX359":"5e08db3f07bf65c0bf3b1f53d4cf3d8e","PX943":"c59a3sf49s11in5f0dh0","PX357":"ae4e536d1cc3d1dc118d94664067387b","PX358":"3ed37dd90c2ab1e679fb9fd71ffa4401","PX229":24,"PX230":24,"PX91":3148,"PX92":886,"PX269":3098,"PX270":886,"PX93":"3148X886","PX185":774,"PX186":668,"PX187":0,"PX188":330.3278503417969,"PX95":true,"PX400":539,"PX404":"144|66|66|180|80","PX90":["loadTimes","csi","app","runtime"],"PX190":"","PX552":"false","PX399":"false","PX549":1,"PX411":1,"PX405":true,"PX547":true,"PX134":true,"PX89":true,"PX170":5,"PX85":["PDF Viewer","Chrome PDF Viewer","Chromium PDF Viewer","Microsoft Edge PDF Viewer","WebKit built-in PDF"],"PX1179":true,"PX1180":true,"PX59":"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.61 Safari/537.36","PX61":"en-US","PX313":["en-US"],"PX63":"Win32","PX86":true,"PX154":240,"PX1157":8,"PX1173":1,"PX133":true,"PX88":true,"PX169":2,"PX62":"Gecko","PX69":"20030107","PX64":"5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.61 Safari/537.36","PX65":"Netscape","PX66":"Mozilla","PX1144":true,"PX1152":10,"PX1153":0,"PX1154":false,"PX1155":"4g","PX60":true,"PX87":true,"PX821":4294705152,"PX822":44436501,"PX823":36606713,"PX147":false,"PX155":"Tue Sep 28 2021 01:02:41 GMT-0400","PX236":false,"PX194":false,"PX195":true,"PX237":0,"PX238":"missing","PX208":"visible","PX218":0,"PX231":891,"PX232":1560,"PX254":false,"PX295":false,"PX268":false,"PX166":true,"PX138":true,"PX143":true,"PX1142":2,"PX1143":6,"PX1146":2,"PX1147":1,"PX714":"64556c77","PX715":"","PX724":"10207b2f","PX725":"10207b2f","PX729":"90e65465","PX443":true,"PX466":true,"PX467":true,"PX468":true,"PX191":0,"PX94":4,"PX120":[],"PX141":false,"PX96":"https://www.walmart.com/","PX55":"","PX34":"Error","PX1065":1,"PX850":1,"PX851":6021,"PX1054":1632805361380,"PX1008":3600,"PX1055":1632805356711,"PX1056":1632805362119,"PX1038":"4c3e8640-2019-11ec-ac5c-75b34d7e7cb4","PX371":true}}]`)
	return json.Marshal(&PayloadOut{
		Payload: pxutils.EncodePayload(string(rp), 50),
		AppID:   sitedata.AppId,
		Tag:     sitedata.Tag,
		Uuid:    UUID,
		Ft:      sitedata.Ft,
		Seq:     "1",
		En:      "NTA",
		Pc:      pxutils.CreatePC(string(rp), fmt.Sprintf(`%s:%s:%s`, UUID, sitedata.Tag, sitedata.Ft)),
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