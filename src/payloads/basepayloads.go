package payloads

import (
	pxutils "github.com/incizzle/perimeterx-utils-go"
	"math/rand"
	"strconv"
	"strings"
)

type PX2 struct {
	T string `json:"t"`
	D struct {
		PX371  bool	  `json:"PX371"`
		PX96   string `json:"PX96"`
		PX63   string `json:"PX63"`
		PX191  int    `json:"PX191"`
		PX850  int    `json:"PX850"`
		PX851  int    `json:"PX851"`
		PX1008 int    `json:"PX1008"`
		PX1055 int64  `json:"PX1055"`
		PX1056 int64  `json:"PX1056"`
		PX1038 string `json:"PX1038"`
	} `json:"d"`
}
type PX3 struct {
	T string `json:"t"`
	D struct {
		PX371  bool			 `json:"PX371"`
		PX234  bool          `json:"PX234"`
		PX235  bool          `json:"PX235"`
		PX151  bool          `json:"PX151"`
		PX239  bool          `json:"PX239"`
		PX240  bool          `json:"PX240"`
		PX152  bool          `json:"PX152"`
		PX153  bool          `json:"PX153"`
		PX314  bool          `json:"PX314"`
		PX192  bool          `json:"PX192"`
		PX196  bool          `json:"PX196"`
		PX207  bool          `json:"PX207"`
		PX251  bool          `json:"PX251"`
		PX982  int64         `json:"PX982"`
		PX983  string        `json:"PX983"`
		PX986  string        `json:"PX986"`
		PX985  int           `json:"PX985"`
		PX1033 string        `json:"PX1033"`
		PX1019 string        `json:"PX1019"`
		PX1020 string        `json:"PX1020"`
		PX1021 string        `json:"PX1021"`
		PX1022 string        `json:"PX1022"`
		PX1035 bool          `json:"PX1035"`
		PX1025 bool          `json:"PX1025"`
		PX359  string        `json:"PX359"`
		PX943  string        `json:"PX943"`
		PX357  string        `json:"PX357"`
		PX358  string        `json:"PX358"`
		PX229  int           `json:"PX229"`
		PX230  int           `json:"PX230"`
		PX91   int           `json:"PX91"`
		PX92   int           `json:"PX92"`
		PX269  int           `json:"PX269"`
		PX270  int           `json:"PX270"`
		PX93   string        `json:"PX93"`
		PX185  int           `json:"PX185"`
		PX186  int           `json:"PX186"`
		PX187  int           `json:"PX187"`
		PX188  int           `json:"PX188"`
		PX95   bool          `json:"PX95"`
		PX400  int           `json:"PX400"`
		PX404  string        `json:"PX404"`
		PX90   []string      `json:"PX90"`
		PX190  string        `json:"PX190"`
		PX552  string        `json:"PX552"`
		PX399  string        `json:"PX399"`
		PX549  int           `json:"PX549"`
		PX411  int           `json:"PX411"`
		PX402  int           `json:"PX402"`
		PX548  int           `json:"PX548"`
		PX405  bool          `json:"PX405"`
		PX547  bool          `json:"PX547"`
		PX134  bool          `json:"PX134"`
		PX89   bool          `json:"PX89"`
		PX170  int           `json:"PX170"`
		PX85   []string      `json:"PX85"`
		PX59   string        `json:"PX59"`
		PX61   string        `json:"PX61"`
		PX313  []string      `json:"PX313"`
		PX63   string        `json:"PX63"`
		PX86   bool          `json:"PX86"`
		PX154  int           `json:"PX154"`
		PX133  bool          `json:"PX133"`
		PX88   bool          `json:"PX88"`
		PX169  int           `json:"PX169"`
		PX62   string        `json:"PX62"`
		PX69   string        `json:"PX69"`
		PX64   string        `json:"PX64"`
		PX65   string        `json:"PX65"`
		PX66   string        `json:"PX66"`
		PX60   bool          `json:"PX60"`
		PX87   bool          `json:"PX87"`
		PX821  int64         `json:"PX821"`
		PX822  int           `json:"PX822"`
		PX823  int           `json:"PX823"`
		PX147  bool          `json:"PX147"`
		PX155  string        `json:"PX155"`
		PX236  bool          `json:"PX236"`
		PX194  bool          `json:"PX194"`
		PX195  bool          `json:"PX195"`
		PX237  int           `json:"PX237"`
		PX238  string        `json:"PX238"`
		PX208  string        `json:"PX208"`
		PX218  int           `json:"PX218"`
		PX231  int           `json:"PX231"`
		PX232  int           `json:"PX232"`
		PX254  bool          `json:"PX254"`
		PX295  bool          `json:"PX295"`
		PX268  bool          `json:"PX268"`
		PX166  bool          `json:"PX166"`
		PX138  bool          `json:"PX138"`
		PX143  bool          `json:"PX143"`
		PX714  string        `json:"PX714"`
		PX715  string        `json:"PX715"`
		PX724  string        `json:"PX724"`
		PX725  string        `json:"PX725"`
		PX729  string        `json:"PX729"`
		PX443  bool          `json:"PX443"`
		PX466  bool          `json:"PX466"`
		PX467  bool          `json:"PX467"`
		PX468  bool          `json:"PX468"`
		PX191  int           `json:"PX191"`
		PX94   int           `json:"PX94"`
		PX120  []interface{} `json:"PX120"`
		PX141  bool          `json:"PX141"`
		PX96   string        `json:"PX96"`
		PX55   string        `json:"PX55"`
		PX1065 int           `json:"PX1065"`
		PX850  int           `json:"PX850"`
		PX851  int           `json:"PX851"`
		PX1054 int64         `json:"PX1054"`
		PX1008 int           `json:"PX1008"`
		PX1055 int64         `json:"PX1055"`
		PX1056 int64         `json:"PX1056"`
		PX1038 string        `json:"PX1038"`
	} `json:"d"`
}

func(this *PX2) Instantiate(site, uuid string){
	this.T = "PX2"
	this.D.PX371 = true
	this.D.PX96 = site
	this.D.PX63 = "Win32"
	this.D.PX191 = 0
	this.D.PX850 = 0
	this.D.PX851 = 500 + rand.Intn(100)
	this.D.PX1008 = 3600
	this.D.PX1055 = generateTimestamp()
	this.D.PX1056 = generateTimestamp()
	this.D.PX1038 = uuid
}
func(this *PX3) Instantiate(site, uuid, varobj, ua string){
	this.T = "PX3"
	px2res := PX2RESHANDLER(varobj)
	this.D.PX371 = true
	this.D.PX234 = false
	this.D.PX235 = false
	this.D.PX151 = false
	this.D.PX239 = false
	this.D.PX240 = false
	this.D.PX152 = false
	this.D.PX153 = false
	this.D.PX314 = false
	this.D.PX192 = false
	this.D.PX196 = false
	this.D.PX207 = false
	this.D.PX251 = false
	this.D.PX982 = generateTimestamp()
	this.D.PX983 = strings.Split(px2res["cls"],"|")[0]
	this.D.PX986 = strings.Split(px2res["cls"],"|")[1]
	this.D.PX985, _ = strconv.Atoi(px2res["drc"])
	this.D.PX1033 = "e0eaf10e"
	this.D.PX1019 = "3f2b5a30"
	this.D.PX1020 = "7766a52d"
	this.D.PX1021 = "714b1317"
	this.D.PX1022 = "6a90378d"
	this.D.PX1035 = false
	this.D.PX1025 = false
	this.D.PX359 = pxutils.H1(uuid, ua)
	this.D.PX943 = px2res["wcs"]
	this.D.PX357 = pxutils.H1(strings.Split(px2res["vid"],"|")[0], ua)
	if v, ok := px2res["sid"]; ok{
		this.D.PX358 = pxutils.H1(v, ua)
	}else{
		this.D.PX358 = pxutils.H1(uuid, ua)
	}
	this.D.PX229 = 24
	this.D.PX230 = 24
	this.D.PX91 = 1920
	this.D.PX92 = 1080
	this.D.PX269 = 3778
	this.D.PX270 = 1080
	this.D.PX93 = "3840X1080"
	this.D.PX185 = 950 + rand.Intn(120)
	this.D.PX186 = 350 + rand.Intn(50)
	this.D.PX187 = 0
	this.D.PX188 = 0
	this.D.PX95 = true
	this.D.PX400 = 111
	this.D.PX404 = "144|54|54|180|68"
	this.D.PX90 = []string{"loadTimes", "csi", "app", "runtime"}
	this.D.PX190 = ""
	this.D.PX552 = "false"
	this.D.PX399 = "false"
	this.D.PX549 = 1
	this.D.PX411 = 1
	this.D.PX402 = 1
	this.D.PX548 = 1
	this.D.PX405 = true
	this.D.PX547 = true
	this.D.PX134 = true
	this.D.PX89 = true
	this.D.PX170 = 3
	this.D.PX85 = []string{"Chrome PDF Plugin", "Chrome PDF Viewer", "Native Client"}
	this.D.PX59 = ua
	this.D.PX61 = "en-US"
	this.D.PX313 = []string{"en-US"}
	this.D.PX63 = "Win32"
	this.D.PX86 = true
	this.D.PX154 = 240
	this.D.PX133 = true
	this.D.PX88 = true
	this.D.PX169 = 4 //most likely core count
	this.D.PX62 = "Gecko"
	this.D.PX69 = "20030107"
	this.D.PX64 = strings.SplitN(ua, "/", 2)[1]
	this.D.PX65 = "Netscape"
	this.D.PX66 = "Mozilla"
	this.D.PX60 = true
	this.D.PX87 = true
	this.D.PX821 = 4294705152 //total memory
	this.D.PX822 = 10000000 + rand.Intn(40000000)
	this.D.PX823 = this.D.PX822 - 2000000 - rand.Intn(8000000)
	this.D.PX147 = false
	this.D.PX155 = "Fri Apr 02 2021 00:31:32 GMT-0400 (Eastern Daylight Time)" // timezone
	this.D.PX236 = false
	this.D.PX194 = false
	this.D.PX195 = true
	this.D.PX237 = 0
	this.D.PX238 = "missing"
	this.D.PX208 = "visible"
	this.D.PX218 = 0
	this.D.PX231 = 1087
	this.D.PX232 = 1903
	this.D.PX254 = false
	this.D.PX295 = false
	this.D.PX268 = false
	this.D.PX166 = true
	this.D.PX138 = true
	this.D.PX143 = true
	this.D.PX714 = "64556c77"
	this.D.PX715 = ""
	this.D.PX724 = "10207b2f"
	this.D.PX725 = "10207b2f"
	this.D.PX729 = "90e65465"
	this.D.PX443 = true
	this.D.PX466 = true
	this.D.PX467 = true
	this.D.PX468 = true
	this.D.PX191 = 0
	this.D.PX94 = 1 + rand.Intn(2)
	this.D.PX120 = []interface{}{}
	this.D.PX141 = false
	this.D.PX96 = site
	this.D.PX55 = ""
	this.D.PX1065 = 1
	this.D.PX850 = 1
	this.D.PX851 = 7000 + rand.Intn(1500)
	this.D.PX1054 = 1617355508198
	this.D.PX1008 = 3600
	this.D.PX1055 = 1617355501361
	this.D.PX1056 = 1617355521093
	this.D.PX1038 = uuid
}
