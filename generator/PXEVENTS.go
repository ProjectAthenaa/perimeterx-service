package generator

import (
	"fmt"
	"github.com/ProjectAthenaa/perimeterx-service/siteconstants"
	"github.com/ProjectAthenaa/pxutils"
	"math/rand"
	"strconv"
	"strings"
)

var (
	CSSBASES = []string{"#povActive", "BODY", "#header-tabs"}
	CSSBODIES = []string{"A", "LI3", "DIV1", "DIV2", "DIV5",  "SECTION2", "IMG", "DIV4", "UL", "LI3", "DIV2", "DIV1", "DIV2", "DIV1", "DIV2", "DIV1", "DIV2", "DIV1", "P2", "LI1",  "IFRAME"}
	PX176LIST = []string{"PX760", "PX759", "60_sec_rest"}
	PX73LIST  = []string{"div", "a", "ul", "h2", "span", "input", "img"}
	PX874LIST = []string{"src", "A", "img"}
	PX956LIST = []string{"PX971", "PX960", "PX957", "PX959", "PX958", "PX977"}
	PX867LIST = []string{"href", "src", "action"}
	PX869LIST = map[string][]string{
		"PX971": {"https://classic.cdn.media.amplience.net/i/siteval/P1234_0411_right?w=300&h=300&img404=404&v=1", "https://cdn.media.amplience.net/i/siteval/4.3-Lt-Gray-Google-Rating", "https://classic.cdn.media.amplience.net/i/siteval/STS2x"},
		"PX960": {"https://content.stylitics.com/images/collage/7d2c8fcca5393898888a5861b350720c5927eaffb9c0a7?size=small", "https://classic.cdn.media.amplience.net/i/siteval/1P145_1000_right?w=300&h=300&img404=404&v=1"},
		"PX957": {"https://content.stylitics.com/images/collage/7d2c8fc8af323ff0e9bb3ef4610a8b178949146e323348?size=small"},
		"PX959": {"http://demandware.net/", "http://fila.com/"},
		"PX958": {"https://siteval.usablenet.com/pt/start", "https://siteval.usablenet.com/pt/switch", "https://cdn.static.amplience.net", "https://content.stylitics.com/images/collage/7d2c8fcca5393898888a5861b350720c5927eaffb9c0a7?size=small", "https://sc-static.net/scevent.min.js"},
		"PX977": {"https://sc-static.net/scevent.min.js", "https://tr.snapchat.com/cm/i", "https://siteval.usablenet.com/pt/start", "https://siteval.usablenet.com/pt/switch"},
	}
	PX71LIST  = []string{"mouseover", "pointerdown", "mousedown", "pointerup"}
	PX827LIST = []string{"keyup", "keydown"}
)

type UserAgentEvent struct {
	T string `json:"t"`
	D struct {
		PX59   string `json:"PX59"`
		PX96   string `json:"PX96,omitempty"`
		PX204  int    `json:"PX204"`
		PX887  bool   `json:"PX887"`
		PX888  int    `json:"PX888"`
		PX839  int    `json:"PX839"`
		PX919  int    `json:"PX919"`
		PX359  string `json:"PX359"`
		PX357  string `json:"PX357"`
		PX358  string `json:"PX358"`
		PX850  int    `json:"PX850"`
		PX851  int    `json:"PX851"`
		PX1056 int    `json:"PX1056"`
		PX1038 string `json:"PX1038"`
	} `json:"d"`
}
type MouseOverEvent struct {
	T string `json:"t"`
	D struct {
		PX96   string `json:"PX96,omitempty"`
		PX38   string `json:"PX38"`
		PX70   int    `json:"PX70"`
		PX157  string `json:"PX157"`
		PX72   string `json:"PX72"`
		PX34   string `json:"PX34"`
		PX78   int    `json:"PX78"`
		PX79   int    `json:"PX79"`
		PX850  int    `json:"PX850"`
		PX851  int    `json:"PX851"`
		PX1056 int    `json:"PX1056"`
		PX1038 string `json:"PX1038"`
	} `json:"d"`
}
type MouseMoveEvent struct {
	T string `json:"t"`
	D struct {
		PX96   string         `json:"PX96,omitempty"`
		PX82   string         `json:"PX82"`
		PX176  string         `json:"PX176"`
		PX177  int            `json:"PX177"`
		PX181  string         `json:"PX181"`
		PX178  int            `json:"PX178"`
		PX179  map[string]int `json:"PX179"`
		PX180  string         `json:"PX180"`
		PX58   []mouseEvent   `json:"PX58"`
		PX410  string         `json:"PX410"`
		PX608  []string       `json:"PX608"`
		PX584  []string       `json:"PX584"`
		PX462  bool           `json:"PX462"`
		PX850  int            `json:"PX850"`
		PX851  int            `json:"PX851"`
		PX1056 int            `json:"PX1056"`
		PX1038 string         `json:"PX1038"`
	} `json:"d"`
}
type RequestEvent struct{
	T string `json:"t"`
	D struct {
		PX96    string       `json:"PX96,omitempty"`
		PX371   bool         `json:"PX371,omitempty"`
		PX614   bool         `json:"PX614,omitempty"`
	} `json:"d"`
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
	PX227 bool   `json:"PX227,omitempty"`
	PX233 int    `json:"PX233,omitempty"`
	PX856 bool   `json:"PX856,omitempty"`
	PX857 int    `json:"PX857,omitempty"`
	PX858 int    `json:"PX858,omitempty"`
}

func GenerateMouseOverEvent(sid, vid, uuid string, sitedata *siteconstants.SiteData, seq int) []MouseOverEvent{
	var reqEventArrOut []MouseOverEvent
	addreqobj := InstantiateMouseOverEvent(sid, vid, uuid, seq)
	addreqobj.D.PX96 = sitedata.Url
	reqEventArrOut = append(reqEventArrOut, addreqobj)
	return reqEventArrOut
}
func GenerateMouseMoveEvent(uuid string, sitedata *siteconstants.SiteData, seq int) []MouseMoveEvent{
	var reqEventArrOut []MouseMoveEvent
	addreqobj := InstantiateMouseMoveEvent(uuid, seq)
	addreqobj.D.PX180 = sitedata.Url
	reqEventArrOut = append(reqEventArrOut, addreqobj)
	return reqEventArrOut
}
func GenerateRequestEvent(sitedata *siteconstants.SiteData) []RequestEvent{
	var reqEventArrOut []RequestEvent
	addreqobj := InstantiateRequestEvent()
	addreqobj.D.PX96 = sitedata.Url
	reqEventArrOut = append(reqEventArrOut, addreqobj)
	return reqEventArrOut
}
func GenerateUserAgentEvent(sid, vid, uuid string, sitedata *siteconstants.SiteData) []UserAgentEvent {
	var reqEventArrOut []UserAgentEvent
	addreqobj := InstantiateUserAgentEvent(sid, vid, uuid)
	addreqobj.D.PX96 = sitedata.Url
	reqEventArrOut = append(reqEventArrOut, addreqobj)
	return reqEventArrOut
}

func InstantiateMouseOverEvent(site, uuid, sitekey string, seq int) MouseOverEvent {
	var payload MouseOverEvent
	payload.T = "PX297"
	payload.D.PX70 = 200000 + rand.Intn(100000)
	payload.D.PX38 = "mouseover"
	payload.D.PX157 = "true"
	payload.D.PX72 = htmlCssGenerate()
	payload.D.PX34 = fmt.Sprintf("TypeError: Cannot read property '0' of null\n    at Tt (https://%s/px/%s/init.js:2:14107)\n    at HTMLBodyElement.fe (https://%s/px/%s/init.js:2:26488)", site, sitekey, site, sitekey)
	payload.D.PX78 = rand.Intn(600) - 300
	payload.D.PX79 = rand.Intn(600)
	payload.D.PX850 = seq
	payload.D.PX851 = 10000 + rand.Intn(10000)
	payload.D.PX1056 = GenerateTimestamp()
	payload.D.PX1038 = uuid
	return payload
}
func InstantiateMouseMoveEvent(uuid string, seq int) MouseMoveEvent {
	var payload MouseMoveEvent
	payload.T = "PX175"
	payload.D.PX82 = strconv.Itoa(300+rand.Intn(600)) + "+" + strconv.Itoa(300+rand.Intn(600))
	payload.D.PX176 = PX176LIST[rand.Intn(len(PX176LIST))]
	payload.D.PX177 = GenerateTimestamp()
	payload.D.PX181 = uuid
	payload.D.PX178 = 0
	payload.D.PX179 = PX179GENERATE(10)
	px58, scrollarray := PX58GENERATE(10+rand.Intn(10), 3)
	payload.D.PX58 = px58
	payload.D.PX410 = PX410GENERATE(10)
	payload.D.PX608 = PX608GENERATE(23)
	payload.D.PX584 = scrollarray
	payload.D.PX462 = true
	payload.D.PX850 = seq
	payload.D.PX851 = 200000 + rand.Intn(200000)
	payload.D.PX1056 = GenerateTimestamp()
	payload.D.PX1038 = uuid
	return payload
}
func InstantiateRequestEvent() RequestEvent{
	var payload RequestEvent
	payload.D.PX371	= true
	payload.D.PX614	= true
	return payload
}
func InstantiateUserAgentEvent(sid, vid, uuid string) UserAgentEvent {
	var payload UserAgentEvent
	payload.T = "PX203"
	payload.D.PX204 = 1
	payload.D.PX59 = siteconstants.UA
	payload.D.PX887 = true
	payload.D.PX888 = 240000
	payload.D.PX839 = 0
	payload.D.PX919 = 0
	payload.D.PX359 = pxutils.H1(uuid, siteconstants.UA)
	payload.D.PX357 = pxutils.H1(vid, siteconstants.UA)
	payload.D.PX358 = pxutils.H1(sid, siteconstants.UA)
	count := 2 + rand.Intn(10)
	payload.D.PX850 = count
	payload.D.PX851 = 15000*count + rand.Intn(30000*count)
	payload.D.PX1056 = GenerateTimestamp()
	payload.D.PX1038 = uuid
	return payload
}

func mouseEventArrayGenerate(length, perftimingin, px72 int, SCROLLS *[]string) ([]mouseEvent, int, int) {
	var performanceNow int
	eventArrayOut := []mouseEvent{}
	currx := rand.Intn(500)
	curry := rand.Intn(500)
	if perftimingin != 0 {
		performanceNow = perftimingin + rand.Intn(100000)
	} else {
		performanceNow = 200000 + rand.Intn(100000)
	}

	if px72 < 2 {
		px72 = 2
	}

	for i := 0; i < length; i++ {
		var eventObject mouseEvent
		switch rand.Intn(6) {
		case 0:
			eventObject.PX71 = "mousemove"
			eventObject.PX159 = "true"
			eventObject.PX172 = strconv.Itoa(currx) + "," + strconv.Itoa(curry) + "," + strconv.Itoa(performanceNow)
			if rand.Intn(10) > 7 {
				lastindex := 100 + rand.Intn(20)
				currx = currx + rand.Intn(60) - 30
				curry = curry + rand.Intn(60) - 30
				eventObject.PX172 += "|" + strconv.Itoa(currx) + "," + strconv.Itoa(curry) + "," + strconv.Itoa(lastindex)
				if rand.Intn(10) > 7 {
					lastindex++
					currx = currx + rand.Intn(60) - 30
					curry = curry + rand.Intn(60) - 30
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
			eventObject.PX226 = 2 + rand.Intn(3)
			eventObject.PX227 = true
			eventObject.PX233 = 7 + rand.Intn(3)
			eventObject.PX856 = true
			eventObject.PX70 = performanceNow
		}
		eventArrayOut = append(eventArrayOut, eventObject)
		if rand.Intn(10) > 7 {
			px72++
		}
		performanceNow = performanceNow + rand.Intn(120)
		currx = currx + rand.Intn(60) - 30
		curry = curry + rand.Intn(60) - 30
	}
	return eventArrayOut, performanceNow, px72
}
func htmlCssGenerate() string{
	var stringout strings.Builder
	stringout.WriteString(CSSBASES[rand.Intn(len(CSSBASES))])
	for i := 1+rand.Intn(9); i > 0; i -- {
		stringout.WriteString(">"+CSSBODIES[rand.Intn(len(CSSBODIES))])
	}
	return stringout.String()
}

func PX58GENERATE(length, splits int) ([]mouseEvent, []string) {
	totalEventArray := []mouseEvent{}
	timeout := 0
	px72count := 2
	var generatedEventArray []mouseEvent
	scrolleventarr := []string{}
	splitsize := length / splits
	for i := 0; i < splits; i++ {
		for j := 0; j > splitsize; j++ {
			generatedEventArray, timeout, px72count = mouseEventArrayGenerate(splitsize, timeout, px72count, &scrolleventarr)
			totalEventArray = append(generatedEventArray, totalEventArray...)
		}
	}
	return totalEventArray, scrolleventarr
}
func PX410GENERATE(length int) string {
	var strout string
	timestamp := 200000 + rand.Intn(100000)
	for i := 0; i < length; i++ {
		strout += strconv.Itoa(rand.Intn(200)-100) + "," + strconv.Itoa(rand.Intn(200)-100) + strconv.Itoa(timestamp)
		timestamp += rand.Intn(150)
		if rand.Intn(26) > 24 {
			timestamp += 100000
		}
	}
	return strout[:len(strout)-1]
}
func PX608GENERATE(length int) []string {
	strarrout := []string{}
	inittimestamp := 200000 + rand.Intn(100000)
	for i := 0; i < length; i++ {
		strarrout = append(strarrout, strconv.Itoa(100+rand.Intn(200))+","+strconv.Itoa(200+rand.Intn(400))+","+strconv.Itoa(inittimestamp))
		inittimestamp += rand.Intn(150)
		if rand.Intn(26) > 24 {
			inittimestamp += 100000
		}
	}
	return strarrout
}
func PX179GENERATE(length int) map[string]int {
	mapout := make(map[string]int)
	for i := 1; i <= length; i++ {
		var mapkey string
		for j := 0; j < 2+rand.Intn(5); j++ {
			mapkey += PX176LIST[rand.Intn(len(PX176LIST))] + ">"
		}
		mapout[mapkey[:len(mapkey)-1]] = 1
	}
	return mapout
}