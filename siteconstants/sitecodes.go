package siteconstants

import "github.com/ProjectAthenaa/sonic-core/sonic/antibots/perimeterx"

type SiteData struct {
	Tag		string
	Ft		string
	Url 	string
	Host 	string
	AppId	string
}

var SiteDataHolder = map[perimeterx.SITE]SiteData{
	perimeterx.SITE_WALMART: {
		Tag: "v6.7.9",
		Ft: "221",
		Url: "https://www.walmart.com/",
		Host: "walmart.com",
		AppId: "pxu6b0qd2s",
	},
}