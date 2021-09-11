package generator

import (
	"math/rand"
	"strings"
)

type RECAPPAYLOAD struct {
	T string `json:"t"`
	D struct {
		PX70   int    `json:"PX70"`
		PX34   string `json:"PX34"`
		PX610  bool   `json:"PX610"`
		PX754  bool   `json:"PX754"`
		PX755  string `json:"PX755"`
		PX756  string `json:"PX756"`
		PX757  string `json:"PX757"`
		PX850  int    `json:"PX850"`
		PX851  int    `json:"PX851"`
		PX1056 int    `json:"PX1056"`
		PX1038 string `json:"PX1038"`
		PX371  bool   `json:"PX371"`
		PX250  string `json:"PX250"`
		PX788  string `json:"PX788, omitempty"`
		PX708  string `json:"PX708, omitempty"`
		PX1129 bool  `json:"PX1129, omitempty"`
		PX96   string `json:"PX96"`
	} `json:"d"`
}

func InstantiateRECAP(uuid, vid, site, sitekey, captchatype, px250val, token string)RECAPPAYLOAD{
	var payload RECAPPAYLOAD
	payload.T = "PX761"
	tsin := 1000 + rand.Intn(60000)
	payload.D.PX70 = tsin
	endpointstring := strings.Replace("sitekey", "TypeError: Cannot read property '0' of null\n    at It (https://client.perimeterx.net/sitekey/main.min.js:2:14201)\n    at re (https://client.perimeterx.net/sitekey/main.min.js:2:26243)\n    at Object.Un [as PX763] (https://client.perimeterx.net/sitekey/main.min.js:2:25361)\n    at u (https://captcha.px-cdn.net/sitekey/captcha.js?a=c&amp;u=uuid&amp;v=vid&amp;m=0:3:75126)\n    at c (https://captcha.px-cdn.net/sitekey/captcha.js?a=c&amp;u=uuid&amp;v=vid&amp;m=0:3:71250)\n    at i (https://captcha.px-cdn.net/sitekey/captcha.js?a=c&amp;u=uuid&amp;v=vid&amp;m=0:3:70778)\n    at Xo (https://captcha.px-cdn.net/sitekey/captcha.js?a=c&amp;u=uuid&amp;v=vid&amp;m=0:3:69831)", sitekey, -1)
	endpointstring = strings.Replace("uuid", endpointstring, uuid, -1)
	endpointstring = strings.Replace("vid", endpointstring, vid, -1)
	payload.D.PX34 = endpointstring
	payload.D.PX610 = true
	payload.D.PX754 = true
	payload.D.PX755 = token
	payload.D.PX756 = captchatype
	payload.D.PX757 = site
	payload.D.PX850 = 4 + rand.Intn(3)
	payload.D.PX851 = 2000 + rand.Intn(5000)
	payload.D.PX1056 = GenerateTimestamp()
	payload.D.PX1038 = uuid
	payload.D.PX371 = false
	payload.D.PX250 = px250val
	payload.D.PX96 = site
	return payload
}