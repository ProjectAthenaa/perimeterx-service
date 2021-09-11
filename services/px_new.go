package services

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/ProjectAthenaa/perimeterx-service/generator"
	"github.com/ProjectAthenaa/perimeterx-service/responsedeob"
	"github.com/ProjectAthenaa/perimeterx-service/siteconstants"
	"github.com/ProjectAthenaa/sonic-core/sonic/antibots/perimeterx"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

//cookie marshalled payload, error

type PXInPayload struct {
	SITE           perimeterx.SITE
	TYPE           perimeterx.PXType
	COOKIE         string
	RESPONSEOBJECT []byte
	TOKEN          string
	RSC            int
}

func ConstructPayload(instruct PXInPayload) (COOKIE string, payload []byte, err error) {
	rand.Seed(time.Now().UnixNano())
	SITE := siteconstants.SiteDataHolder[instruct.SITE]
	TYPE := instruct.TYPE

	var UUID, TOKEN string
	var RSC int
	var resobj responsedeob.ResponseJSON
	if TYPE == perimeterx.PXType_PX3 || TYPE == perimeterx.PXType_PX34 {
		UUID = generator.GenerateUUID()
	} else if TYPE == perimeterx.PXType_HCAPHIGH || TYPE == perimeterx.PXType_RECAP || TYPE == perimeterx.PXType_HCAPLOW {
		COOKIE, resobj = responsedeob.SplitResponse(instruct.RESPONSEOBJECT)
		TOKEN = instruct.TOKEN
	} else {
		COOKIE, resobj = responsedeob.SplitResponse(instruct.RESPONSEOBJECT)
		RSC = instruct.RSC
	}

	switch TYPE {
	case perimeterx.PXType_PX3, perimeterx.PXType_PX34:
		payload, err = generator.GenInit(&SITE, UUID)
		return COOKIE, payload, err
	case perimeterx.PXType_RE, perimeterx.PXType_BRE, perimeterx.PXType_MME, perimeterx.PXType_MOE, perimeterx.PXType_UAE, perimeterx.PXType_EVENT:
		payload, err = generator.GenEvents(&SITE, UUID, resobj, RSC)
		return COOKIE, payload, err
	case perimeterx.PXType_HCAPHIGH, perimeterx.PXType_HCAPLOW:
		payload, err = generator.GenHoldCaptcha(&SITE, UUID, resobj, TOKEN)
		return COOKIE, payload, err
	case perimeterx.PXType_RECAP:
		payload, err = generator.GenReCaptcha(&SITE, UUID, resobj)
		return COOKIE, payload, err
	}

	return "", nil, errors.New("could not create payload")
}

//https://collector-pxu6b0qd2s.px-cloud.net/api/v2/collector
//https://collector-pxu6b0qd2s.px-cloud.net/assets/js/bundle

func SendData(payloadin string, site *siteconstants.SiteData) string {
	res, _ := http.Post(fmt.Sprintf("https://collector-%s.px-cloud.net/api/v2/collector", site.Key), "application/x-www-form-urlencoded", bytes.NewBuffer([]byte(payloadin)))
	body, _ := ioutil.ReadAll(res.Body)
	return string(body)
}
