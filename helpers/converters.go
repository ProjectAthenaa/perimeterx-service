package helpers

import (
	px "github.com/ProjectAthenaa/sonic-core/sonic/antibots/perimeterx"
)

func ConvertToSite(site px.SITE) Site {
	switch site {
	case px.SITE_WALMART:
		return Walmart
	case px.SITE_HIBBET:
		return Hibbet
	case px.SITE_SSENSE:
		return Ssense
	case px.SITE_SNIPES:
		return Snipes
	case px.SITE_ONYGO:
		return Onygo
	case px.SITE_SOLEBOX:
		return Solebox
	default:
		return Site{}
	}
}
