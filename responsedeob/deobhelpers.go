package responsedeob

import (
	"github.com/json-iterator/go"
	"strings"
)

var json = jsoniter.ConfigFastest

type ResponseJSON struct{
	CS		string `json:"cs"`
	SID  	string `json:"sid"`
	VID  	string `json:"vid"`
	CLS0 	string `json:"cls0"`
	CLS1 	string `json:"cls1"`
	WCS  	string `json:"wcs"`
	DRC  	string `json:"drc"`
	CI 		string `json:"ci"`
	CP 		string `json:"cp"`
	CTS 	string `json:"cts"`
	CITOKEN string `json:"citoken"`
	HTMLCITOKEN string `json:"htmlcitoken"`
	STS string `json:"sts"`
	EN		string `json:"en"`
}

type ResponseJSONRaw struct{
	Do []string `json:"do"`
}

func SplitResponse(resstring []byte) (string, ResponseJSON){
	var rp *ResponseJSONRaw
	json.Unmarshal(resstring, &rp)

	var resobj ResponseJSON
	var cookie string

	Iterator:
	for _, val := range rp.Do{
		split := strings.Split(val, "|")
		switch split[0]{
			case "cs":
				resobj.CS = split[1]
			case "sid":
				resobj.SID = split[1]
			case "vid":
				resobj.VID = split[1]
			case "cls":
				resobj.CLS0 = split[1]
				if len(split) >= 3{
					resobj.CLS1 = split[2]
				}
			case "wcs":
				resobj.WCS = split[1]
			case "drc":
				resobj.DRC = split[1]
			case "ci":
				resobj.CI = split[3]
				resobj.CITOKEN = split[2]
				resobj.HTMLCITOKEN = split[4]
			case "cts":
				resobj.CTS = split[1]
			case "cp":
				resobj.CP = split[2]
			case "bake":
				cookie = split[3]
			case "en":
				resobj.EN = split[3]
				break Iterator
		case "sts":
			resobj.STS = split[1]
		}
	}

	return cookie, resobj
}