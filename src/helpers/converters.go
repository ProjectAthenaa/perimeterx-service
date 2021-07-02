package helpers

import px "main/services/protos"

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
