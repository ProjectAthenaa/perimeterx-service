package payloads

import(
	"encoding/json"
	"strings"
	"time"
	//pxutils "github.com/incizzle/perimeterx-utils-go"
)

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
func generateTimestamp() int64{
	return time.Now().UnixNano()/1000000
}