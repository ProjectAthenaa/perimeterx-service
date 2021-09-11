package generator

import (
	"github.com/ProjectAthenaa/perimeterx/siteconstants"
	"math/rand"
	"net/url"
	"strings"
)

//todo: THESE ARE OUTDATE, DO NOT IMPLEMENT

type OldRequestEvent struct {
	T string `json:"t"`
	D struct {
		PX96   string      `json:"PX96,omitempty"`
		PX956  string      `json:"PX956"`
		PX868  int         `json:"PX868"`
		PX869  string      `json:"PX869"`
		PX870  bool        `json:"PX870"`
		PX928  string      `json:"PX928"`
		PX812  string      `json:"PX812"`
		PX852  string      `json:"PX852"`
		PX908  string      `json:"PX908"`
		PX909  string      `json:"PX909"`
		PX1010 string      `json:"PX1010"`
		PX910  string      `json:"PX910"`
		PX1047 bool        `json:"PX1047"`
		PX907  string      `json:"PX907"`
		PX918  string      `json:"PX918"`
		PX820  string      `json:"PX820"`
		PX930  interface{} `json:"PX930"`
		PX55   string      `json:"PX55"`
		PX850  int         `json:"PX850"`
		PX851  int         `json:"PX851"`
		PX1056 int         `json:"PX1056"`
		PX1038 string      `json:"PX1038"`
		PX871  int         `json:"PX871,omitempty"`
		PX867  string      `json:"PX867,omitempty"`
		PX875  string      `json:"PX875,omitempty"`
		PX874  string      `json:"PX874,omitempty"`
		PX1002 int         `json:"PX1002,omitempty"`
		PX746  string      `json:"PX746,omitempty"`
		PX813  int         `json:"PX813,omitempty"`
		PX911  bool        `json:"PX911,omitempty"`
		PX877  int         `json:"PX877,omitempty"`
		PX906  string      `json:"PX906,omitempty"`
	} `json:"d,omitempty"`
}
type BrowserOldRequestEvent struct {
	T string `json:"t"`
	D struct {
		PX96   string      `json:"PX96,omitempty"`
		PX91   int         `json:"PX91"`
		PX92   int         `json:"PX92"`
		PX269  int         `json:"PX269"`
		PX270  int         `json:"PX270"`
		PX186  int         `json:"PX186"`
		PX185  int         `json:"PX185"`
		PX232  int         `json:"PX232"`
		PX231  int         `json:"PX231"`
		PX908  string      `json:"PX908"`
		PX909  string      `json:"PX909"`
		PX1010 string      `json:"PX1010"`
		PX910  string      `json:"PX910"`
		PX1047 bool        `json:"PX1047"`
		PX907  string      `json:"PX907"`
		PX918  string      `json:"PX918"`
		PX820  string      `json:"PX820"`
		PX930  interface{} `json:"PX930"`
		PX55   string      `json:"PX55"`
		PX850  int         `json:"PX850"`
		PX851  int         `json:"PX851"`
		PX1056 int         `json:"PX1056"`
		PX1038 string      `json:"PX1038"`
	} `json:"d"`
}

//these are outdated. should no tbe used
func OldRequestEventGenerate(sid, vid, uuid string, sitedata *siteconstants.SiteData) []OldRequestEvent {
	reqEventArrOut := []OldRequestEvent{}
	performancenow := 300000 + rand.Intn(500000)
	addreqobj := OldRequestEvent{}
	addreqobj.D.PX96 = sitedata.Url
	InstantiateOldRequestEvent(sid, vid, uuid, sitedata.Host)
	reqEventArrOut = append(reqEventArrOut, addreqobj)
	performancenow = performancenow + rand.Intn(120)
	return reqEventArrOut
}
func browserOldRequestEventGenerate(sid, vid, uuid string, sitedata *siteconstants.SiteData) []BrowserOldRequestEvent {
	reqEventArrOut := []BrowserOldRequestEvent{}
	performancenow := 300000 + rand.Intn(500000)
	addreqobj := BrowserOldRequestEvent{}
	addreqobj.D.PX96 = sitedata.Url
	InstantiateBrowserOldRequestEvent(sid, vid, uuid, sitedata.Host)
	reqEventArrOut = append(reqEventArrOut, addreqobj)
	performancenow = performancenow + rand.Intn(120)
	return reqEventArrOut
}
func mixedArrayGenerate(sid, vid, uuid string, sitedata *siteconstants.SiteData) []interface{} {
	reqEventArrOut := []interface{}{}
	performancenow := 300000 + rand.Intn(500000)
	for i := 0; i < 1+rand.Intn(5); i++ {
		choice := rand.Intn(15)
		if choice < 12 {
			addreqobj := &OldRequestEvent{}
			if i == 0 {
				addreqobj.D.PX96 = sitedata.Url
			}
			InstantiateOldRequestEvent(sid, vid, uuid, sitedata.Host)
			reqEventArrOut = append(reqEventArrOut, addreqobj)
		} else if choice < 16 {
			addreqobj := &UserAgentEvent{}
			if i == 0 {
				addreqobj.D.PX96 = sitedata.Url
			}
			InstantiateUserAgentEvent(sid, vid, uuid)
			reqEventArrOut = append(reqEventArrOut, addreqobj)
		} else {
			addreqobj := &BrowserOldRequestEvent{}
			if i == 0 {
				addreqobj.D.PX96 = sitedata.Url
			}
			InstantiateBrowserOldRequestEvent(sid, vid, uuid, sitedata.Host)
			reqEventArrOut = append(reqEventArrOut, addreqobj)
		}
		performancenow = performancenow + rand.Intn(120)
	}
	return reqEventArrOut
}


func InstantiateOldRequestEvent(sid, vid, uuid, site string) OldRequestEvent {
	var payload OldRequestEvent
	payload.T = "PX811"
	PXCHOICE := PX956LIST[rand.Intn(len(PX956LIST))]
	payload.D.PX956 = PXCHOICE
	payload.D.PX868 = 1
	payload.D.PX869 = strings.Replace("siteval", PX869LIST[PXCHOICE][rand.Intn(len(PX869LIST[PXCHOICE]))], site, -1)
	payload.D.PX870 = false
	payload.D.PX928 = "https:"
	payload.D.PX812 = site
	payload.D.PX852 = "gray"
	payload.D.PX908 = vid
	payload.D.PX909 = sid
	payload.D.PX1010 = "boarded"
	payload.D.PX910 = uuid
	payload.D.PX1047 = true
	payload.D.PX907 = "1.3.3837"
	payload.D.PX918 = "5195230638"
	payload.D.PX820 = "PX816"
	payload.D.PX55 = strings.Replace("siteval", "https%3A%2F%2Fwww.siteval.com%2Fon%2Fdemandware.store%2FSites-siteval-US-Site%2Fdefault%2FPX-Show%3Furl%3DaHR0cHM6Ly93d3cuaGliYmV0dC5jb20vb24vZGVtYW5kd2FyZS5zdG9yZS9TaXRlcy1IaWJiZXR0LVVTLVNpdGU%253d%26frame%3D1617327309845", site, -1)
	count := 2 + rand.Intn(10)
	performancenow := 15000*count + rand.Intn(30000*count)
	payload.D.PX850 = count
	payload.D.PX851 = performancenow
	payload.D.PX1056 = GenerateTimestamp()
	payload.D.PX1038 = uuid
	switch PXCHOICE {
	case "PX971":
		payload.D.PX813 = performancenow
		payload.D.PX911 = true
		payload.D.PX877 = 200 + rand.Intn(450)
		payload.D.PX906 = "img"
	case "PX957":
		basesuburl, _ := url.Parse(payload.D.PX869)
		payload.D.PX871 = 1
		payload.D.PX874 = PX874LIST[rand.Intn(len(PX874LIST))]
		payload.D.PX875 = basesuburl.Host
		payload.D.PX1002 = 0
	case "PX958":
		basesuburl, _ := url.Parse(payload.D.PX869)
		payload.D.PX871 = 1
		payload.D.PX867 = PX867LIST[rand.Intn(len(PX867LIST))]
		payload.D.PX874 = PX874LIST[rand.Intn(len(PX874LIST))]
		payload.D.PX875 = basesuburl.Host
		payload.D.PX1002 = rand.Intn(6)
	case "PX959":
		payload.D.PX813 = performancenow
		payload.D.PX874 = PX874LIST[rand.Intn(len(PX874LIST))]
	case "PX960":
		payload.D.PX867 = PX867LIST[rand.Intn(len(PX867LIST))]
		payload.D.PX813 = performancenow
	case "PX977":
		payload.D.PX746 = "appendChild"
		payload.D.PX813 = performancenow
		payload.D.PX867 = PX867LIST[rand.Intn(len(PX867LIST))]
	}
	return payload
}
func InstantiateBrowserOldRequestEvent(sid, vid, uuid, site string) BrowserOldRequestEvent {
	var payload BrowserOldRequestEvent
	payload.T = "PX805"
	payload.D.PX91 = 3040
	payload.D.PX92 = 1080
	payload.D.PX269 = 3778
	payload.D.PX270 = 1080
	payload.D.PX185 = 950 + rand.Intn(120)
	payload.D.PX186 = 350 + rand.Intn(50)
	payload.D.PX232 = 1087
	payload.D.PX231 = 1903
	payload.D.PX908 = vid
	payload.D.PX909 = sid
	payload.D.PX1010 = "boarded"
	payload.D.PX910 = uuid
	payload.D.PX1047 = true
	payload.D.PX907 = "1.3.3837"
	payload.D.PX918 = "5195230638"
	payload.D.PX820 = "PX816"
	payload.D.PX55 = strings.Replace("siteval", "https%3A%2F%2Fwww.siteval.com%2Fon%2Fdemandware.store%2FSites-siteval-US-Site%2Fdefault%2FPX-Show%3Furl%3DaHR0cHM6Ly93d3cuaGliYmV0dC5jb20vb24vZGVtYW5kd2FyZS5zdG9yZS9TaXRlcy1IaWJiZXR0LVVTLVNpdGU%253d%26frame%3D1617327309845", site, -1)
	count := 2 + rand.Intn(10)
	payload.D.PX850 = count
	payload.D.PX851 = 15000*count + rand.Intn(30000*count)
	payload.D.PX1056 = GenerateTimestamp()
	payload.D.PX1038 = uuid
	return payload
}
