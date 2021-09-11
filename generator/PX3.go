package generator

import (
	"github.com/ProjectAthenaa/perimeterx-service/responsedeob"
	"github.com/ProjectAthenaa/perimeterx-service/siteconstants"
	"github.com/ProjectAthenaa/pxutils"
	"math/rand"
	"strconv"
	"strings"
)

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
		PX982  int           `json:"PX982"`
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

func GeneratePX3(uuid string, resobj responsedeob.ResponseJSON, sitedata *siteconstants.SiteData) []PX3{
	px3obj := InstantiatePX3(uuid)
	px3obj.D.PX983 = resobj.CLS0
	px3obj.D.PX986 = resobj.CLS1
	px3obj.D.PX985, _ = strconv.Atoi(resobj.DRC)
	px3obj.D.PX943 = resobj.WCS
	px3obj.D.PX357 = pxutils.H1(resobj.VID, siteconstants.UA)
	px3obj.D.PX358 = pxutils.H1(resobj.SID, siteconstants.UA)
    px3obj.D.PX96 = sitedata.Url
	return []PX3{px3obj}
}

func InstantiatePX3(uuid string) PX3{
	var payload PX3
	payload.T = "PX3"
	payload.D.PX371 = true
	payload.D.PX234 = false
	payload.D.PX235 = false
	payload.D.PX151 = false
	payload.D.PX239 = false
	payload.D.PX240 = false
	payload.D.PX152 = false
	payload.D.PX153 = false
	payload.D.PX314 = false
	payload.D.PX192 = false
	payload.D.PX196 = false
	payload.D.PX207 = false
	payload.D.PX251 = false
	payload.D.PX982 = GenerateTimestamp()
	payload.D.PX1033 = "e0eaf10e"
	payload.D.PX1019 = "3f2b5a30"
	payload.D.PX1020 = "7766a52d"
	payload.D.PX1021 = "714b1317"
	payload.D.PX1022 = "6a90378d"
	payload.D.PX1035 = false
	payload.D.PX1025 = false
	payload.D.PX359 = pxutils.H1(uuid, siteconstants.UA)
	payload.D.PX229 = 24
	payload.D.PX230 = 24
	payload.D.PX91 = 1920
	payload.D.PX92 = 1080
	payload.D.PX269 = 3778
	payload.D.PX270 = 1080
	payload.D.PX93 = "3840X1080"
	payload.D.PX185 = 950 + rand.Intn(120)
	payload.D.PX186 = 350 + rand.Intn(50)
	payload.D.PX187 = 0
	payload.D.PX188 = 0
	payload.D.PX95 = true
	payload.D.PX400 = 111
	payload.D.PX404 = "144|54|54|180|68"
	payload.D.PX90 = []string{"loadTimes", "csi", "app", "runtime"}
	payload.D.PX190 = ""
	payload.D.PX552 = "false"
	payload.D.PX399 = "false"
	payload.D.PX549 = 1
	payload.D.PX411 = 1
	payload.D.PX402 = 1
	payload.D.PX548 = 1
	payload.D.PX405 = true
	payload.D.PX547 = true
	payload.D.PX134 = true
	payload.D.PX89 = true
	payload.D.PX170 = 3
	payload.D.PX85 = []string{"Chrome PDF Plugin", "Chrome PDF Viewer", "Native Client"}
	payload.D.PX59 = siteconstants.UA
	payload.D.PX61 = "en-US"
	payload.D.PX313 = []string{"en-US"}
	payload.D.PX63 = "Win32"
	payload.D.PX86 = true
	payload.D.PX154 = 240
	payload.D.PX133 = true
	payload.D.PX88 = true
	payload.D.PX169 = 4 //most likely core count
	payload.D.PX62 = "Gecko"
	payload.D.PX69 = "20030107"
	payload.D.PX64 = strings.SplitN(siteconstants.UA, "/", 2)[1]
	payload.D.PX65 = "Netscape"
	payload.D.PX66 = "Mozilla"
	payload.D.PX60 = true
	payload.D.PX87 = true
	payload.D.PX821 = 4294705152 //total memory
	payload.D.PX822 = 10000000 + rand.Intn(40000000)
	payload.D.PX823 = payload.D.PX822 - 2000000 - rand.Intn(8000000)
	payload.D.PX147 = false
	payload.D.PX155 = "Fri Apr 02 2021 00:31:32 GMT-0400 (Eastern Daylight Time)" // timezone
	payload.D.PX236 = false
	payload.D.PX194 = false
	payload.D.PX195 = true
	payload.D.PX237 = 0
	payload.D.PX238 = "missing"
	payload.D.PX208 = "visible"
	payload.D.PX218 = 0
	payload.D.PX231 = 1087
	payload.D.PX232 = 1903
	payload.D.PX254 = false
	payload.D.PX295 = false
	payload.D.PX268 = false
	payload.D.PX166 = true
	payload.D.PX138 = true
	payload.D.PX143 = true
	payload.D.PX714 = "64556c77"
	payload.D.PX715 = ""
	payload.D.PX724 = "10207b2f"
	payload.D.PX725 = "10207b2f"
	payload.D.PX729 = "90e65465"
	payload.D.PX443 = true
	payload.D.PX466 = true
	payload.D.PX467 = true
	payload.D.PX468 = true
	payload.D.PX191 = 0
	payload.D.PX94 = 1 + rand.Intn(2)
	payload.D.PX120 = []interface{}{}
	payload.D.PX141 = false
	payload.D.PX55 = ""
	payload.D.PX1065 = 1
	payload.D.PX850 = 1
	payload.D.PX851 = 7000 + rand.Intn(1500)
	payload.D.PX1054 = 1617355508198
	payload.D.PX1008 = 3600
	payload.D.PX1055 = 1617355501361
	payload.D.PX1056 = 1617355521093
	payload.D.PX1038 = uuid
	return payload
}