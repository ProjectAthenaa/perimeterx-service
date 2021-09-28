package unused
//
//import (
//	"bytes"
//	"encoding/json"
//	"fmt"
//	"github.com/ProjectAthenaa/pxutils"
//	"io"
//	"reflect"
//	"strconv"
//	"strings"
//)
//
//// sends a px payload
//func (session *GenSession) SendPX(payload []entities.PXPayload, path string) (*entities.PXResponse, error) {
//	var finalArray string
//
//	for i, px := range payload {
//		// encoder := json.NewEncoder(px)
//		b, err := json.Marshal(px)
//		if err != nil {
//			return nil, errors.New("failed to create px array")
//		}
//
//		finalArray = finalArray + string(b)
//
//		finalArray = strings.Replace(finalArray, "\\u003c", "<", -1)
//		finalArray = strings.Replace(finalArray, "\\u003e", ">", -1)
//		finalArray = strings.Replace(finalArray, "\\u0026", "&", -1)
//
//		if i != len(payload)-1 {
//			finalArray = finalArray + ","
//		}
//	}
//
//	payloads := "[" + finalArray + "]"
//
//	// fmt.Println("cls", session.Cls)
//	if len(session.Cls) != 0 {
//		sts, _ := strconv.Atoi(session.Sts)
//		px3ckey := strings.Replace(payloads, "ClsAndStsMathsKey", pxutils.SimpleTextEncode(session.Cls[0], sts%10+2), 1)
//		payloads = strings.Replace(px3ckey, "ClsAndStsValue", pxutils.SimpleTextEncode(session.Cls[0], sts%10+1), 1)
//	}
//
//	encoded := pxutils.EncodePayload(payloads, 50)
//	pc := pxutils.CreatePC(payloads, fmt.Sprintf("%s:%s:%d", session.Uuid, session.Site.Tag, session.Site.Ft))
//
//	data := urlValues.Values{"ORDER": {"payload", "appId", "tag", "uuid", "ft", "seq", "en", "cs", "pc", "sid", "vid", "cts", "ci", "rsc"}}
//
//	data.Add("payload", encoded)
//	data.Add("appId", session.Site.AppID)
//	data.Add("tag", session.Site.Tag)
//	data.Add("uuid", session.Uuid)
//	data.Add("ft", strconv.Itoa(session.Site.Ft))
//	data.Add("seq", strconv.Itoa(session.Seq))
//	data.Add("en", "NTA")
//
//	if session.Cs != "" {
//		data.Add("cs", session.Cs)
//	}
//
//	data.Add("pc", pc)
//
//	if session.Sid != "" {
//		data.Add("sid", session.Sid)
//	}
//
//	if session.Vid != "" {
//		data.Add("vid", session.Vid)
//	}
//
//	if session.Cts != "" {
//		data.Add("cts", session.Cts)
//	}
//
//	if session.Ci[1] != "" {
//		data.Add("ci", session.Ci[0])
//	}
//
//	data.Add("rsc", strconv.Itoa(session.Rsc))
//
//	collectorURL := fmt.Sprintf("https://collector-%s.px-cloud.net", strings.ToLower(session.Site.AppID))
//
//	postBody := data.EncodeWithOrder()
//	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", collectorURL, path), bytes.NewBuffer([]byte(postBody)))
//
//	if err != nil {
//		return nil, errors.New("failed to create px request")
//	}
//
//	req.Header = http.Header{
//		"Host":             {req.URL.Host},
//		"content-length":   {},
//		"accept":           {"*/*"},
//		"origin":           {"https://" + session.Site.Www},
//		"sec-ch-ua-mobile": {"?0"},
//		"sec-fetch-site":   {"cross-site"},
//		"sec-fetch-mode":   {"cors"},
//		"sec-ch-ua":        {`" Not;A Brand";v="99", "Google Chrome";v="91", "Chromium";v="91"`},
//		"sec-fetch-dest":   {"empty"},
//		"accept-encoding":  {"gzip, deflate, br"},
//		"content-type":     {"application/x-www-form-urlencoded"},
//		"referer":          {fmt.Sprintf("https://%s/", session.Site.Www)},
//		"accept-language":  {"en-US,en;q=0.9"},
//		"user-agent":       {session.Device.Data.UserAgent},
//		"pragma":           {"no-cache"},
//		"cache-control":    {"no-cache"},
//		//"HEADERORDER": {
//		//	"Host",
//		//	"sec-fetch-site",
//		//
//		//	"sec-ch-ua-mobile",
//		//	"sec-fetch-dest",
//		//
//		//	"accept-language",
//		//
//		//	"sec-fetch-mode",
//		//	"referer",
//		//	"pragma",
//		//	"origin",
//		//
//		//
//		//	"content-type",
//		//	"accept",
//		//	"cache-control",
//		//
//		//	"accept-encoding",
//		//
//		//	"sec-ch-ua",
//		//	"user-agent",
//		//	"content-length",
//		//},
//	}
//
//	order := randomizeHeaders(req.Header)
//	req.Header["HEADERORDER"] = order
//
//	resp, err := session.Client.Do(req)
//	if err != nil {
//		return nil, errors.New("failed to send px request")
//	}
//
//	defer resp.Body.Close()
//
//	bodyText, err := io.ReadAll(resp.Body)
//	if err != nil {
//		return nil, errors.New("failed to parse px response")
//	}
//
//	var response entities.PXResponse
//
//	if err := json.Unmarshal(bodyText, &response); err != nil {
//		return nil, err
//	}
//
//	// set values from Do into session, e.g. session.Cls = Do.Cls
//	v := reflect.ValueOf(response.Do)
//	typeOfS := v.Type()
//	sessionReflect := reflect.ValueOf(session)
//
//	for i := 0; i < v.NumField(); i++ {
//		field := sessionReflect.Elem().FieldByName(typeOfS.Field(i).Name)
//		if field.IsValid() && !v.Field(i).IsZero() {
//			// fmt.Printf("Field: %s\tValue: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
//			field.Set(reflect.ValueOf(v.Field(i).Interface()))
//		}
//	}
//
//	session.PXResponse = response
//
//	return &response, nil
//}
//
