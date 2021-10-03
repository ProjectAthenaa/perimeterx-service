package generator

import (
	"github.com/ProjectAthenaa/perimeterx-service/siteconstants"
	"math/rand"
)

type PX2 struct {
	T string `json:"t"`
	D struct {
		PX63   string `json:"PX63"`
		PX96   string `json:"PX96"`
		PX191  int    `json:"PX191"`
		PX371  bool	  `json:"PX371"`
		PX850  int    `json:"PX850"`
		PX851  int    `json:"PX851"`
		PX1008 int    `json:"PX1008"`
		PX1038 string `json:"PX1038"`
		PX1055 int  `json:"PX1055"`
		PX1056 int  `json:"PX1056"`
	} `json:"d"`
}

func GeneratePX2(uuid string, sitedata *siteconstants.SiteData) []PX2{
	px2obj := InstantiatePX2(uuid)
	px2obj.D.PX96 = sitedata.Url
	return []PX2{px2obj}
}

func InstantiatePX2(uuid string) PX2{
	var payload PX2
	payload.T = "PX2"
	payload.D.PX371 = true
	payload.D.PX63 = "Win32"
	payload.D.PX191 = 0
	payload.D.PX850 = 0
	payload.D.PX851 = 1000 + rand.Intn(100)
	payload.D.PX1008 = 3600
	payload.D.PX1055 = GenerateTimestamp()
	payload.D.PX1056 = GenerateTimestamp()
	payload.D.PX1038 = uuid
	return payload
}