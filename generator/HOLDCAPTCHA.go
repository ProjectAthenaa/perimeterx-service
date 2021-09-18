package generator

import (
	"fmt"
	"github.com/ProjectAthenaa/perimeterx-service/siteconstants"
	"math/rand"
	"strings"
)

type HCAPPAYLOAD struct {
	T string `json:"t"`
	D struct {
		PX70   int               `json:"PX70"`
		PX34   string            `json:"PX34"`
		PX1129 bool              `json:"PX1129"`
		PX610  bool              `json:"PX610"`
		PX1132 string            `json:"PX1132"`
		PX1133 int	             `json:"PX1133"`
		PX179  []string          `json:"PX179"`
		PX827  []PX827KEYEVENT   `json:"PX827"`
		PX582  []PX582MOUSEEVENT `json:"PX582"`
		PX608  []string          `json:"PX608"`
		PX74   int               `json:"PX74"`
		PX75   int               `json:"PX75"`
		PX76   int               `json:"PX76"`
		PX77   int               `json:"PX77"`
		PX91   int               `json:"PX91"`
		PX92   int               `json:"PX92"`
		PX603  int      `json:"PX603"`
		PX604  int      `json:"PX604"`
		PX605  []int    `json:"PX605"`
		PX789  string            `json:"PX789"`
		PX790  int               `json:"PX790"` //ts start
		PX791  int               `json:"PX791"`
		PX792  int      `json:"PX792"`
		PX793  float64  `json:"PX793"`
		PX794  float64  `json:"PX794"`
		PX796  int      `json:"PX796"`//ts end
		PX797  int      `json:"PX797"`
		PX798  int      `json:"PX798"`
		PX799  float64  `json:"PX799"`
		PX800  float64  `json:"PX800"`
		PX801  string   `json:"PX801"`
		PX821  int64    `json:"PX821"`
		PX822  int      `json:"PX822"`
		PX823  int      `json:"PX823"`
		PX803  int      `json:"PX803"`
		PX606  bool     `json:"PX606"`
		PX769  string   `json:"PX769"`
		PX976  bool     `json:"PX976"`
		PX978  bool     `json:"PX978"`
		PX185  int      `json:"PX185"`
		PX186  int      `json:"PX186"`
		PX850  int      `json:"PX850"`
		PX851  int      `json:"PX851"`
		PX1056 int      `json:"PX1056"`
		PX1038 string   `json:"PX1038"`
		PX371  bool     `json:"PX371"`
		PX250  string   `json:"PX250"`
		PX708  string   `json:"PX708"`
		PX96   string   `json:"PX96"`
		PX787  []string `json:"px_787"`
		PX1151 string 	`json:"px_1151"`
		PX1160 bool `json:"px_1160"`
		PX1165 string `json:"px_1165"`
		PX1168 bool `json:"px_1168"`
		PX1171 int `json:"px_1171"`
		PX1182 string `json:"px_1182"`
		PX1188 bool `json:"px_1188"`
	} `json:"d"`
}

func GenerateHCAP(uuid, vid, token, cp, recaptoken string, holdduration int, sitedata *siteconstants.SiteData) []interface{}{
	var outarr []interface{}
	px4obj := InstantiatePX4(uuid, sitedata.AppId)
	outarr = append(outarr, px4obj)

	event1 := InstantiateMouseOverEvent(sitedata.Url, uuid, sitedata.AppId, 4)
	event1.D.PX96 = sitedata.Url
	outarr = append(outarr, event1)

	hcapobj := InstantiateHCAP(uuid, token, cp, sitedata.AppId, holdduration)
	hcapobj.D.PX96 = sitedata.Url
	outarr = append(outarr, hcapobj)

	recapobj := InstantiateRECAP(uuid, vid, sitedata.Url, sitedata.AppId, "pxCaptcha", "PX560", recaptoken)
	outarr = append(outarr, recapobj)

	event2 := InstantiateMouseMoveEvent(uuid, 6)
	event2.D.PX96 = sitedata.Url
	outarr = append(outarr, event2)

	return outarr
}

func InstantiateHCAP(uuid, token, cpval, sitekey string, holdduration int) HCAPPAYLOAD{
	var payload HCAPPAYLOAD
	tstart := 1000 +rand.Intn(1000)
	tend := tstart+holdduration+2

	payload.T = "PX561"
	outappstr := strings.Replace("sitekey", "TypeError: Cannot read property '0' of null\n    at xt (https://client.perimeterx.net/sitekey/main.min.js:2:14047)\\n    at ne (https://client.perimeterx.net/sitekey/main.min.js:2:26105)\\n    at te (https://client.perimeterx.net/sitekey/main.min.js:2:26050)\\n    at c (https://captcha.px-cdn.net/sitekey/captcha.js?a=c&amp;u=uuid&amp;v=&amp;m=0:3:75953)\\n    at u (https://captcha.px-cdn.net/sitekey/captcha.js?a=c&amp;u=uuid&amp;v=&amp;m=0:3:71869)\\n    at a (https://captcha.px-cdn.net/sitekey/captcha.js?a=c&amp;u=uuid&amp;v=&amp;m=0:3:71397)\\n    at Ce (https://captcha.px-cdn.net/sitekey/captcha.js?a=c&amp;u=uuid&amp;v=&amp;m=0:3:70450)", sitekey, -1)
	payload.D.PX34 = strings.Replace("uuid", outappstr, uuid, -1)
	payload.D.PX70 = 1000 + rand.Intn(2000)
	payload.D.PX74 = 200 + rand.Intn(200)
	payload.D.PX75 = 100 + rand.Intn(100)
	payload.D.PX76 =  0
	payload.D.PX77 =  0
	payload.D.PX91 = 1920
	payload.D.PX92 = 1080
	payload.D.PX179 = []string{
		"#sign-in-widget>A1",
		"H1",
		"#sign-in-widget",
		"DIV1>P1",
		"DIV3>DIV1",
		"",
	}
	payload.D.PX185 = 200 + rand.Intn(1500)
	payload.D.PX186 = 200 + rand.Intn(2000)
	payload.D.PX250 = "PX560"
	payload.D.PX371 = true
	payload.D.PX582 = PX582GENERATE(tstart, tend)
	payload.D.PX603 = 1000 +rand.Intn(1000)
	payload.D.PX604 = 2000 +rand.Intn(1000)
	payload.D.PX605 = []int{holdduration+rand.Intn(100)}
	payload.D.PX606 = true
	payload.D.PX608 = PX608HCAPGENERATE(tstart)
	payload.D.PX610 = true
	payload.D.PX708 = "pxhc"
	payload.D.PX769 = token //token
	payload.D.PX787 = []string{
		"getAttribute",
		"className",
		"nodeName",
		"nodeName",
		"nodeName",
		"nodeName",
		"nodeName",
		"nodeName",
		"nodeName",
		"nodeName",
		"nodeName",
		"nodeName",
		"nodeName",
		"nodeName",
		"nodeName",
		"nodeName",
		"nodeName",
		"nodeName",
		"nodeName",
		"nodeName",
		"nodeName",
		"nodeName",
		"scrollIntoViewIfNeeded",
		"scrollIntoViewIfNeeded",
		"scrollIntoViewIfNeeded",
	}
	payload.D.PX789 = "pointerdown"
	payload.D.PX790 = tstart //ts start
	payload.D.PX791 = 300 +rand.Intn(500) //x0
	payload.D.PX792 = 300 +rand.Intn(500) //y0
	payload.D.PX793 = 100 +rand.Float64()*100 //x1
	payload.D.PX794 = 20.0 + rand.Float64()*50 //y1
	payload.D.PX796 =  tend//ts end (+2.5 s)
	payload.D.PX797 = payload.D.PX791 //x0
	payload.D.PX798 = payload.D.PX792 //y0
	payload.D.PX799 = payload.D.PX793 //x1
	payload.D.PX800 = 20.0 + rand.Float64()*50 //y1
	payload.D.PX801 = "pointerup"
	payload.D.PX803 = 400 + rand.Intn(50)
	payload.D.PX821 = 4294662870 //mem limit
	payload.D.PX822 = 30000000 + rand.Intn(60000000) // mem allocated
	payload.D.PX823 = 30000000 + rand.Intn(payload.D.PX822-30000000) //mem used
	payload.D.PX827 = PX827GENERATE(2+rand.Intn(4))
	payload.D.PX850 = 4
	payload.D.PX851 = 2500 + rand.Intn(2000)
	payload.D.PX976 = false
	payload.D.PX978 = false
	payload.D.PX1038 = uuid
	payload.D.PX1056 = GenerateTimestamp()
	payload.D.PX1129 = true
	payload.D.PX1132 = cpval //cpval1
	payload.D.PX1133 = 2 + rand.Intn(25)
	payload.D.PX1151 = "4YCJ4YGQ4YCa4YCG4YCf4YCe4YGQ4YGI4YGD4YGe4YGQ4YCa4YCX4YCT4YCW4YGQ4YGI4YGD4YGe4YGQ4YCG4YCb4YCG4YCe4YCX4YGQ4YGI4YGD4YGe4YGQ4YCf4YCX4YCG4YCT4YGQ4YGI4YGD4YGe4YGQ4YCB4YCG4YCL4YCe4YCX4YGQ4YGI4YGA4YGe4YGQ4YCB4YCR4YCA4YCb4YCC4YCG4YGQ4YGI4YGH4YGe4YGQ4YCQ4YCd4YCW4YCL4YGQ4YGI4YGD4YGe4YGQ4YCW4YCb4YCE4YGQ4YGI4YGL4YGe4YGQ4YCT4YGQ4YGI4YGE4YGe4YGQ4YCB4YCC4YCT4YCc4YGQ4YGI4YGH4YGe4YGQ4YCa4YGD4YGQ4YGI4YGD4YGe4YGQ4YCC4YGQ4YGI4YGB4YGe4YGQ4YCa4YCA4YGQ4YGI4YGD4YCP"
	payload.D.PX1160 =  false
	payload.D.PX1165 =  "visible"
	payload.D.PX1168 =  true
	payload.D.PX1171 =  1
	payload.D.PX1182 =  "PX1183"
	payload.D.PX1188 =  false
	return payload
}

type PX582MOUSEEVENT struct {
	PX70  int    `json:"PX70"`
	PX71  string `json:"PX71"`
	PX72  int    `json:"PX72"`
	PX159 string `json:"PX159"`
}
type PX827KEYEVENT struct {
	PX70  int    `json:"PX70"`
	PX71  string `json:"PX71"`
	PX159 string `json:"PX159"`
	PX72  int    `json:"PX72"`
}
func PX582GENERATE(start, end int) []PX582MOUSEEVENT {
	outarr := []PX582MOUSEEVENT{}
	px72count := 0
	px70ts := start - 2500
	for i := 0; i < 3; i++ {
		eventoadd := PX582MOUSEEVENT{
			PX70:  px70ts,
			PX71:  PX71LIST[rand.Intn(len(PX71LIST))],
			PX72:  px72count,
			PX159: "true",
		}
		if eventoadd.PX71 == "mouseout" {
			px72count = px72count + rand.Intn(3) - 1
			if px72count < 0 {
				px72count = 0
			}
			px70ts = 100 + rand.Intn(700)
		}
		outarr = append(outarr, eventoadd)
	}

	outarr = append(outarr, PX582MOUSEEVENT{
		PX70:  start,
		PX71:  "pointerdown",
		PX72:  px72count,
		PX159: "true",
	})
	outarr = append(outarr, PX582MOUSEEVENT{
		PX70:  start,
		PX71:  "mousedown",
		PX72:  px72count,
		PX159: "true",
	})
	outarr = append(outarr, PX582MOUSEEVENT{
		PX70:  end,
		PX71:  "pointerup",
		PX72:  px72count,
		PX159: "true",
	})

	return outarr
}
func PX827GENERATE(length int) []PX827KEYEVENT {
	outarr := []PX827KEYEVENT{}
	for i := 0; i < length; i++ {
		eventoadd := PX827KEYEVENT{
			PX70:  100 + rand.Intn(5000),
			PX71:  PX827LIST[rand.Intn(len(PX827LIST))],
			PX159: "true",
			PX72:  3 + rand.Intn(6),
		}
		outarr = append(outarr, eventoadd)
	}
	return outarr
}
func PX608HCAPGENERATE(start int) []string{
	var outarr []string
	startx := 50+rand.Intn(250)
	starty := 50+rand.Intn(250)
	startts := start - 2500
	for i := 0; i < 6; i++{
		outarr = append(outarr, fmt.Sprintf("%d,%d,%d", startx, starty, startts))
		startx += rand.Intn(20)
		starty += rand.Intn(20)
		startts += rand.Intn(250)
	}
	return outarr
}