package helpers

const (
	walmart = "walmart"
	hibbet  = "hibbet"
	ssense  = "ssense"
	snipes  = "snipes"
	onygo   = "onygo"
	solebox = "solebox"
)

var (
	Walmart = Site{walmart, "https://www.walmart.com/", "PXu6b0qd2S", "v6.5.5", "202"}
	Hibbet  = Site{hibbet, "https://www.hibbet.com/", "PXAJDckzHD", "v6.5.5", "202"}
	Ssense  = Site{ssense, "https://www.ssense.com/", "PX58Asv359", "v6.5.5", "202"}
	Snipes  = Site{snipes, "https://www.snipes.com/", "pxszbf5p84", "v6.5.5", "202"}
	Onygo   = Site{onygo, "https://www.onygo.com/", "pxj1n025xg", "v6.5.5", "202"}
	Solebox = Site{solebox, "https://www.solebox.com/", "pxur63h57z", "v6.5.5", "202"}
)

type Site struct {
	Name    string
	Site    string
	Sitekey string
	Tag		string
	Version string
}
