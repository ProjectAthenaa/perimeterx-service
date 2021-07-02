package payloads

// todo: PX72 add generator for css selector

import (
	"fmt"
	pxutils "github.com/incizzle/perimeterx-utils-go"
	"math/rand"
	"net/url"
	"strconv"
	"strings"
)

var (
	PX176LIST = []string{"PX760", "PX759", "60_sec_rest"}
	PX73LIST = []string{"div","a","ul","h2","span", "input", "img"}
	PX179LIST = []string{"LI3", "DIV1", "DIV2", "DIV5", "A", "#povActive", "SECTION2", "IMG", "DIV4", "UL", "LI3","DIV2","DIV1","DIV2","DIV1","DIV2","DIV1","DIV2","DIV1", "P2", "LI1", "BODY", "IFRAME"}
	PX874LIST = []string{"src", "A", "img"}
	PX956LIST = []string{"PX971", "PX960", "PX957","PX959","PX958","PX977"}
	PX867LIST = []string{"href", "src", "action"}
	PX869LIST = map[string][]string{
			"PX971":{"https://classic.cdn.media.amplience.net/i/siteval/P1234_0411_right?w=300&h=300&img404=404&v=1", "https://cdn.media.amplience.net/i/siteval/4.3-Lt-Gray-Google-Rating", "https://classic.cdn.media.amplience.net/i/siteval/STS2x"},
			"PX960":{"https://content.stylitics.com/images/collage/7d2c8fcca5393898888a5861b350720c5927eaffb9c0a7?size=small", "https://classic.cdn.media.amplience.net/i/siteval/1P145_1000_right?w=300&h=300&img404=404&v=1",},
			"PX957":{"https://content.stylitics.com/images/collage/7d2c8fc8af323ff0e9bb3ef4610a8b178949146e323348?size=small"},
			"PX959":{"http://demandware.net/", "http://fila.com/"},
			"PX958":{"https://siteval.usablenet.com/pt/start", "https://siteval.usablenet.com/pt/switch", "https://cdn.static.amplience.net", "https://content.stylitics.com/images/collage/7d2c8fcca5393898888a5861b350720c5927eaffb9c0a7?size=small", "https://sc-static.net/scevent.min.js"},
			"PX977":{"https://sc-static.net/scevent.min.js", "https://tr.snapchat.com/cm/i", "https://siteval.usablenet.com/pt/start", "https://siteval.usablenet.com/pt/switch"},
		 }
	PX71LIST = []string{"mouseover", "pointerdown", "mousedown", "pointerup"}
	PX827LIST = []string{"keyup","keydown"}
)

type MouseOverEvent struct {
	T string `json:"t"`
	D struct {
		PX96  string         `json:"PX96,omitempty"`
		PX38   string `json:"PX38"`
		PX70   int    `json:"PX70"`
		PX157  string `json:"PX157"`
		PX72   string `json:"PX72"`
		PX34   string `json:"PX34"`
		PX78   int    `json:"PX78"`
		PX79   int    `json:"PX79"`
		PX850  int    `json:"PX850"`
		PX851  int    `json:"PX851"`
		PX1056 int64  `json:"PX1056"`
		PX1038 string `json:"PX1038"`
	} `json:"d"`
}
type MouseMoveEvent struct {
	T string `json:"t"`
	D struct {
		PX96  string         `json:"PX96,omitempty"`
		PX82  string         `json:"PX82"`
		PX176 string         `json:"PX176"`
		PX177 int64          `json:"PX177"`
		PX181 string         `json:"PX181"`
		PX178 int            `json:"PX178"`
		PX179 map[string]int `json:"PX179"`
		PX180 string         `json:"PX180"`
		PX58  []mouseEvent   `json:"PX58"`
		PX410  string        `json:"PX410"`
		PX608  []string      `json:"PX608"`
		PX584  []string      `json:"PX584"`
		PX462  bool          `json:"PX462"`
		PX850  int           `json:"PX850"`
		PX851  int           `json:"PX851"`
		PX1056 int64         `json:"PX1056"`
		PX1038 string        `json:"PX1038"`
	} `json:"d"`
}
type RequestEvent struct{
	T string `json:"t"`
	D struct {
		PX96  string       `json:"PX96,omitempty"`
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
		PX1056 int64       `json:"PX1056"`
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
type UserAgentEvent struct {
	T string `json:"t"`
	D struct {
		PX96   string `json:"PX96,omitempty"`
		PX204  int    `json:"PX204"`
		PX59   string `json:"PX59"`
		PX887  bool   `json:"PX887"`
		PX888  int    `json:"PX888"`
		PX839  int    `json:"PX839"`
		PX919  int    `json:"PX919"`
		PX359  string `json:"PX359"`
		PX357  string `json:"PX357"`
		PX358  string `json:"PX358"`
		PX850  int    `json:"PX850"`
		PX851  int    `json:"PX851"`
		PX1056 int64  `json:"PX1056"`
		PX1038 string `json:"PX1038"`
	} `json:"d"`
}
type BrowserRequestEvent struct {
	T string `json:"t"`
	D struct {
		PX96  string       `json:"PX96,omitempty"`
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
		PX1056 int64       `json:"PX1056"`
		PX1038 string      `json:"PX1038"`
	} `json:"d"`
}
func(this *MouseOverEvent) Instantiate(site, uuid, sitekey string, seq int){
	this.T = "PX297"
	this.D.PX70 = 200000 + rand.Intn(100000)
	this.D.PX38 = "mouseover"
	this.D.PX157 = "true"
	this.D.PX72 = "LI:nth-child(3)>DIV:nth-child(1)>DIV:nth-child(1)>DIV:nth-child(2)>DIV:nth-child(5)>A"
	this.D.PX34 = fmt.Sprintf("TypeError: Cannot read property '0' of null\n    at Tt (https://%s/px/%s/init.js:2:14107)\n    at HTMLBodyElement.fe (https://%s/px/%s/init.js:2:26488)", site, sitekey, site, sitekey)
	this.D.PX78 = rand.Intn(600) - 300
	this.D.PX79 = rand.Intn(600)
	this.D.PX850 = seq
	this.D.PX851 = 10000 + rand.Intn(10000)
	this.D.PX1056 = generateTimestamp()
	this.D.PX1038 = uuid
}
func(this *MouseMoveEvent) Instantiate(site, uuid string, seq int){
	this.T = "PX175"
	this.D.PX82 = strconv.Itoa(300+rand.Intn(600))+"+"+strconv.Itoa(300+rand.Intn(600))
	this.D.PX176 = PX176LIST[rand.Intn(len(PX176LIST))]
	this.D.PX177 = generateTimestamp()
	this.D.PX181 = uuid
	this.D.PX178 = 0
	this.D.PX179 = PX179GENERATE(10)
	this.D.PX180 = site
	px58, scrollarray := PX58GENERATE(10+rand.Intn(10), 3)
	this.D.PX58 = px58
	this.D.PX410 = PX410GENERATE(10)
	this.D.PX608 = PX608GENERATE(23)
	this.D.PX584 = scrollarray
	this.D.PX462 = true
	this.D.PX850 = seq
	this.D.PX851 = 200000 + rand.Intn(200000)
	this.D.PX1056 = generateTimestamp()
	this.D.PX1038 = uuid
}
func(this *RequestEvent) Instantiate(uuid, sid, vid string, performancenow, start int, rawUrl url.URL){
	this.T = "PX811"
	PXCHOICE := PX956LIST[rand.Intn(len(PX956LIST))]
	this.D.PX956 = PXCHOICE
	this.D.PX868 = 1
	this.D.PX869 = strings.Replace("siteval", PX869LIST[PXCHOICE][rand.Intn(len(PX869LIST[PXCHOICE]))], rawUrl.Host, -1)
	this.D.PX870 = false
	this.D.PX928 = "https:"
	this.D.PX812 = rawUrl.Host
	this.D.PX852 = "gray"
	this.D.PX908 = vid
	this.D.PX909 = sid
	this.D.PX1010 = "boarded"
	this.D.PX910 = uuid
	this.D.PX1047 = true
	this.D.PX907 = "1.3.3837"
	this.D.PX918 = "5195230638"
	this.D.PX820 = "PX816"
	this.D.PX55 = strings.Replace("siteval", "https%3A%2F%2Fwww.siteval.com%2Fon%2Fdemandware.store%2FSites-siteval-US-Site%2Fdefault%2FPX-Show%3Furl%3DaHR0cHM6Ly93d3cuaGliYmV0dC5jb20vb24vZGVtYW5kd2FyZS5zdG9yZS9TaXRlcy1IaWJiZXR0LVVTLVNpdGU%253d%26frame%3D1617327309845", rawUrl.Host, -1)
	this.D.PX850 = start
	this.D.PX851 = performancenow
	this.D.PX1056 = generateTimestamp()
	this.D.PX1038 = uuid
	switch PXCHOICE {
	case "PX971":
		this.D.PX813 = performancenow
		this.D.PX911 = true
		this.D.PX877 = 200 + rand.Intn(450)
		this.D.PX906 = "img"
	case "PX957":
		basesuburl, _ := url.Parse(this.D.PX869)
		this.D.PX871 = 1
		this.D.PX874 = PX874LIST[rand.Intn(len(PX874LIST))]
		this.D.PX875 = basesuburl.Host
		this.D.PX1002 = 0
	case "PX958":
		basesuburl, _ := url.Parse(this.D.PX869)
		this.D.PX871 = 1
		this.D.PX867 = PX867LIST[rand.Intn(len(PX867LIST))]
		this.D.PX874 = PX874LIST[rand.Intn(len(PX874LIST))]
		this.D.PX875 = basesuburl.Host
		this.D.PX1002 = rand.Intn(6)
	case "PX959":
		this.D.PX813 = performancenow
		this.D.PX874 = PX874LIST[rand.Intn(len(PX874LIST))]
	case "PX960":
		this.D.PX867 = PX867LIST[rand.Intn(len(PX867LIST))]
		this.D.PX813 = performancenow
	case "PX977":
		this.D.PX746 = "appendChild"
		this.D.PX813 = performancenow
		this.D.PX867 = PX867LIST[rand.Intn(len(PX867LIST))]
	}
}
func(this *UserAgentEvent) Instantiate(uuid, sid, vid, ua string, start int){
	this.T = "PX203"
	if len(sid) == 0{
		sid = uuid
	}
	this.D.PX204 = 1
	this.D.PX59 = ua
	this.D.PX887 = true
	this.D.PX888 = 240000
	this.D.PX839 = 0
	this.D.PX919 = 0
	this.D.PX359 = pxutils.H1(uuid, ua)
	this.D.PX357 = pxutils.H1(vid, ua)
	this.D.PX358 = pxutils.H1(sid, ua)
	this.D.PX850 = start
	this.D.PX851 = 1000
	this.D.PX1056 = generateTimestamp()
	this.D.PX1038 = uuid
}
func(this *BrowserRequestEvent) Instantiate(uuid, sid, vid string, performancenow, start int, rawUrl url.URL){
	this.T = "PX805"
	this.D.PX91 = 3040
	this.D.PX92 = 1080
	this.D.PX269 = 3778
	this.D.PX270 = 1080
	this.D.PX185 = 950 + rand.Intn(120)
	this.D.PX186 = 350 + rand.Intn(50)
	this.D.PX232 = 1087
	this.D.PX231 = 1903
	this.D.PX908 = vid
	this.D.PX909 = sid
	this.D.PX1010 = "boarded"
	this.D.PX910 = uuid
	this.D.PX1047 = true
	this.D.PX907 = "1.3.3837"
	this.D.PX918 = "5195230638"
	this.D.PX820 = "PX816"
	this.D.PX55 = strings.Replace("siteval", "https%3A%2F%2Fwww.siteval.com%2Fon%2Fdemandware.store%2FSites-siteval-US-Site%2Fdefault%2FPX-Show%3Furl%3DaHR0cHM6Ly93d3cuaGliYmV0dC5jb20vb24vZGVtYW5kd2FyZS5zdG9yZS9TaXRlcy1IaWJiZXR0LVVTLVNpdGU%253d%26frame%3D1617327309845", rawUrl.Host, -1)
	this.D.PX850 = start
	this.D.PX851 = performancenow
	this.D.PX1056 = generateTimestamp()
	this.D.PX1038 = uuid
}

func RequestEventGenerate(site, uuid, sid, vid string, length, start int) []RequestEvent {
	rawUrl, _ := url.Parse(site)
	reqEventArrOut := []RequestEvent{}
	performancenow := 300000 + rand.Intn(500000)
	for i := 0; i < length; i++{
		addreqobj := RequestEvent{}
		if i == 0{
			addreqobj.D.PX96 = site
		}
		addreqobj.Instantiate(uuid, sid, vid, performancenow, start, *rawUrl)
		reqEventArrOut = append(reqEventArrOut, addreqobj)
		start++
		performancenow = performancenow + rand.Intn(120)
	}
	return reqEventArrOut
}
func BrowserRequestEventGenerate(site, uuid, sid, vid string, length, start int) []BrowserRequestEvent {
	rawUrl, _ := url.Parse(site)
	reqEventArrOut := []BrowserRequestEvent{}
	performancenow := 300000 + rand.Intn(500000)
	for i := 0; i < length; i++{
		addreqobj := BrowserRequestEvent{}
		if i == 0{
			addreqobj.D.PX96 = site
		}
		addreqobj.Instantiate(uuid, sid, vid, performancenow, start, *rawUrl)
		reqEventArrOut = append(reqEventArrOut, addreqobj)
		start++
		performancenow = performancenow + rand.Intn(120)
	}
	return reqEventArrOut
}
func UserAgentEventGenerate(site, uuid, sid, vid, ua string, length, start int) []UserAgentEvent {
	reqEventArrOut := []UserAgentEvent{}
	performancenow := 300000 + rand.Intn(500000)
	for i := 0; i < length; i++{
		addreqobj := UserAgentEvent{}
		if i == 0{
			addreqobj.D.PX96 = site
		}
		addreqobj.Instantiate(uuid, sid, vid, ua, start)
		reqEventArrOut = append(reqEventArrOut, addreqobj)
		start++
		performancenow = performancenow + rand.Intn(120)
	}
	return reqEventArrOut
}
func MixedArrayGenerate(site, uuid, sid, vid, ua string, length, start int) []interface{} {
	rawUrl, _ := url.Parse(site)
	reqEventArrOut := []interface{}{}
	performancenow := 300000 + rand.Intn(500000)
	for i := 0; i < length; i++{
		choice := rand.Intn(15)
		if choice < 12{
			addreqobj := &RequestEvent{}
			if i == 0{
				addreqobj.D.PX96 = site
			}
			addreqobj.Instantiate(uuid, sid, vid, performancenow, start, *rawUrl)
			reqEventArrOut = append(reqEventArrOut, addreqobj)
		}else if choice < 16{
			addreqobj := &UserAgentEvent{}
			if i == 0{
				addreqobj.D.PX96 = site
			}
			addreqobj.Instantiate(uuid, sid, vid, ua, start)
			reqEventArrOut = append(reqEventArrOut, addreqobj)
		}else{
			addreqobj := &BrowserRequestEvent{}
			if i == 0{
				addreqobj.D.PX96 = site
			}
			addreqobj.Instantiate(uuid, sid, vid, performancenow, start, *rawUrl)
			reqEventArrOut = append(reqEventArrOut, addreqobj)
		}
		start++
		performancenow = performancenow + rand.Intn(120)
	}
	return reqEventArrOut
}

type mouseEvent struct {
	PX71  string `json:"PX71"`
	PX159 string `json:"PX159"`
	PX172 string `json:"PX172,omitempty"`
	PX72  int    `json:"PX72,omitempty"`
	PX73  string `json:"PX73,omitempty"`
	PX74  int    `json:"PX74,omitempty"`
	PX75  int    `json:"PX75,omitempty"`
	PX76  int    `json:"PX76,omitempty"`
	PX77  int    `json:"PX77,omitempty"`
	PX78  int    `json:"PX78,omitempty"`
	PX79  int    `json:"PX79,omitempty"`
	PX70  int    `json:"PX70,omitempty"`
	PX171 int    `json:"PX171,omitempty"`
	PX226 int    `json:"PX226,omitempty"`
	PX227 bool    `json:"PX227,omitempty"`
	PX233 int    `json:"PX233,omitempty"`
	PX856 bool    `json:"PX856,omitempty"`
	PX857 int    `json:"PX857,omitempty"`
	PX858 int    `json:"PX858,omitempty"`
}
func mouseEventArrayGenerate(length, perftimingin, px72 int, SCROLLS *[]string) ([]mouseEvent, int, int){
	var performanceNow int
	eventArrayOut := []mouseEvent{}
	currx := rand.Intn(500)
	curry := rand.Intn(500)
	if perftimingin != 0 {
		performanceNow = perftimingin + rand.Intn(100000)
	}else{
		performanceNow = 200000 + rand.Intn(100000)
	}

	if px72 < 2{
		px72 = 2
	}

	for i := 0; i < length; i++{
		var eventObject mouseEvent
		switch(rand.Intn(6)) {
			case 0:
				eventObject.PX71 = "mousemove"
				eventObject.PX159 = "true"
				eventObject.PX172 = strconv.Itoa(currx) + "," + strconv.Itoa(curry) + "," + strconv.Itoa(performanceNow)
				if rand.Intn(10) > 7{
					lastindex := 100 + rand.Intn(20)
					currx = currx + rand.Intn(60)-30
					curry = curry + rand.Intn(60)-30
					eventObject.PX172 += "|" + strconv.Itoa(currx) + "," + strconv.Itoa(curry) + "," + strconv.Itoa(lastindex)
					if rand.Intn(10) > 7{
						lastindex++
						currx = currx + rand.Intn(60)-30
						curry = curry + rand.Intn(60)-30
						eventObject.PX172 += "|" + strconv.Itoa(currx) + "," + strconv.Itoa(curry) + "," + strconv.Itoa(lastindex)
					}
				}
			case 1:
				eventObject.PX71 = "mouseout"
				eventObject.PX159 = "true"
				eventObject.PX72 = px72
				eventObject.PX73 = PX73LIST[rand.Intn(len(PX73LIST))]
				eventObject.PX74 = rand.Intn(1000)
				eventObject.PX75 = rand.Intn(1000)
				eventObject.PX76 = rand.Intn(1000)
				eventObject.PX77 = rand.Intn(1000)
				eventObject.PX78 = currx
				eventObject.PX79 = curry
				eventObject.PX70 = performanceNow
			case 2:
				eventObject.PX71 = "mouseup"
				eventObject.PX159 = "true"
				eventObject.PX72 = px72
				eventObject.PX73 = PX73LIST[rand.Intn(len(PX73LIST))]
				eventObject.PX74 = rand.Intn(1000)
				eventObject.PX75 = rand.Intn(1000)
				eventObject.PX76 = rand.Intn(1000)
				eventObject.PX77 = rand.Intn(1000)
				eventObject.PX78 = currx
				eventObject.PX79 = curry
				eventObject.PX70 = performanceNow
			case 3:
				eventObject.PX71 = "keyup"
				eventObject.PX159 = "true"
				eventObject.PX72 = px72
				eventObject.PX73 = "body"
				eventObject.PX74 = rand.Intn(1000)
				eventObject.PX75 = rand.Intn(1000)
				eventObject.PX76 = rand.Intn(1000)
				eventObject.PX77 = rand.Intn(1000)
				eventObject.PX78 = currx
				eventObject.PX79 = curry
				eventObject.PX70 = performanceNow
			case 4:
				eventObject.PX71 = "scroll"
				eventObject.PX159 = "true"
				eventObject.PX857 = rand.Intn(600)
				eventObject.PX858 = 0
				eventObject.PX70 = performanceNow
				*SCROLLS = append(*SCROLLS, strconv.Itoa(eventObject.PX857)+",0")
			case 5:
				eventObject.PX71 = "keydown"
				eventObject.PX159 = "true"
				eventObject.PX72 = px72
				eventObject.PX73 = "body"
				eventObject.PX74 = rand.Intn(1000)
				eventObject.PX75 = rand.Intn(1000)
				eventObject.PX76 = rand.Intn(1000)
				eventObject.PX77 = rand.Intn(1000)
				eventObject.PX171 = 60 + rand.Intn(120)
				eventObject.PX226 = 2+rand.Intn(3)
				eventObject.PX227 = true
				eventObject.PX233 = 7 + rand.Intn(3)
				eventObject.PX856 = true
				eventObject.PX70 = performanceNow
		}
		eventArrayOut = append(eventArrayOut, eventObject)
		if rand.Intn(10) > 7{
			px72++
		}
		performanceNow = performanceNow + rand.Intn(120)
		currx = currx + rand.Intn(60)-30
		curry = curry + rand.Intn(60)-30
	}
	return eventArrayOut, performanceNow, px72
}

func PX58GENERATE(length, splits int) ([]mouseEvent, []string){
	totalEventArray := []mouseEvent{}
	timeout := 0
	px72count := 2
	var generatedEventArray []mouseEvent
	scrolleventarr := []string{}
	splitsize := length/splits
	for i := 0; i < splits; i++{
		for j := 0; j > splitsize; j++{
			generatedEventArray, timeout, px72count = mouseEventArrayGenerate(splitsize, timeout, px72count, &scrolleventarr)
			totalEventArray = append(generatedEventArray, totalEventArray...)
		}
	}
	return totalEventArray, scrolleventarr
}
func PX410GENERATE(length int) string{
	var strout string
	timestamp := 200000 + rand.Intn(100000)
	for i := 0; i < length; i++{
		strout += strconv.Itoa(rand.Intn(200)-100) + "," + strconv.Itoa(rand.Intn(200)-100) + strconv.Itoa(timestamp)
		timestamp += rand.Intn(150)
		if rand.Intn(26) > 24{
			timestamp += 100000
		}
	}
	return strout[:len(strout)-1]
}
func PX608GENERATE(length int) []string{
	strarrout := []string{}
	inittimestamp := 200000 + rand.Intn(100000)
	for i := 0; i < length; i++{
		strarrout = append(strarrout, strconv.Itoa(100 + rand.Intn(200))+","+strconv.Itoa(200 + rand.Intn(400))+","+strconv.Itoa(inittimestamp))
		inittimestamp += rand.Intn(150)
		if rand.Intn(26) > 24{
			inittimestamp += 100000
		}
	}
	return strarrout
}
func PX179GENERATE(length int) map[string]int{
	mapout := make(map[string]int)
	for i := 1; i <= length; i++{
		var mapkey string
		for j := 0; j < 2 + rand.Intn(5); j++{
			mapkey += PX176LIST[rand.Intn(len(PX176LIST))]+">"
		}
		mapout[mapkey[:len(mapkey)-1]] = 1
	}
	return mapout
}

type PX582MOUSEEVENT struct{
	PX70  int `json:"PX70"`
	PX71  string `json:"PX71"`
	PX72  int    `json:"PX72"`
	PX159 string `json:"PX159"`
}
func PX582GENERATE(length int)[]PX582MOUSEEVENT {
	outarr := []PX582MOUSEEVENT{}
	px72count := 0
	px70ts := 100 + rand.Intn(1500)
	for i := 0; i < length; i++{
		eventoadd := PX582MOUSEEVENT{
			PX70:  px70ts,
			PX71:  PX71LIST[rand.Intn(len(PX71LIST))],
			PX72:  px72count,
			PX159: "true",
		}
		if eventoadd.PX71 == "mouseout"{
			px72count = px72count + rand.Intn(3) -1
			if px72count < 0{
				px72count = 0
			}
			px70ts = 100 + rand.Intn(1500)
		}
		outarr = append(outarr, eventoadd)
	}
	return outarr
}

type PX827KEYEVENT struct {
	PX70  int    `json:"PX70"`
	PX71  string `json:"PX71"`
	PX159 string `json:"PX159"`
	PX72  int    `json:"PX72"`
}
func PX827GENERATE(length int)[]PX827KEYEVENT {
	outarr := []PX827KEYEVENT{}
	for i := 0; i < length; i++{
		eventoadd := PX827KEYEVENT{
			PX70:  100 + rand.Intn(5000),
			PX71:  PX827LIST[rand.Intn(len(PX827LIST))],
			PX159: "true",
			PX72:  3+rand.Intn(6),
		}
		outarr = append(outarr, eventoadd)
	}
	return outarr
}