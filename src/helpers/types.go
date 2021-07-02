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
	Walmart = Site{walmart, "https://www.walmart.com/", "PXu6b0qd2S"}
	Hibbet  = Site{hibbet, "https://www.hibbet.com/", "PXAJDckzHD"}
	Ssense  = Site{ssense, "https://www.ssense.com/", "PX58Asv359"}
	Snipes  = Site{snipes, "https://www.snipes.com/", "pxszbf5p84"}
	Onygo   = Site{onygo, "https://www.onygo.com/", "pxj1n025xg"}
	Solebox = Site{solebox, "https://www.solebox.com/", "pxur63h57z"}
)

type Site struct {
	Name    string
	Site    string
	Sitekey string
}
