package responsedeob

import (
	"github.com/json-iterator/go"
	"strings"
)

var json = jsoniter.ConfigFastest

type ResponseJSON struct{
	SID  	string `json:"sid"`
	VID  	string `json:"vid"`
	CLS0 	string `json:"cls0"`
	CLS1 	string `json:"cls1"`
	WCS  	string `json:"wcs"`
	DRC  	string `json:"drc"`
	CI 		string `json:"ci"`
	CP 		string `json:"cp"`
	CITOKEN string `json:"citoken"`
	HTMLCITOKEN string `json:"htmlcitoken"`
}

type ResponseJSONRaw struct{
	Do []string `json:"do"`
}

func SplitResponse(resstring []byte) (string, ResponseJSON){
	var rp ResponseJSONRaw
	json.Unmarshal(resstring, rp)

	var resobj ResponseJSON
	var cookie string

	Iterator:
	for _, val := range rp.Do{
		split := strings.Split(val, "|")
		switch split[0]{
			case "sid":
				resobj.SID = split[1]
			case "vid":
				resobj.VID = split[1]
			case "cls":
				resobj.CLS0 = split[1]
				resobj.CLS1 = split[2]
			case "wcs":
				resobj.WCS = split[1]
			case "drc":
				resobj.DRC = split[1]
			case "ci":
				resobj.CI = split[3]
				resobj.CITOKEN = split[2]
				resobj.HTMLCITOKEN = split[4]
			case "cp":
				resobj.CP = split[2]
			case "en", "bake":
				cookie = split[3]
				break Iterator
		}
	}

	return cookie, resobj
}