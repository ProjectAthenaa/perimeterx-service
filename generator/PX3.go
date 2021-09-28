package generator

import (
	"github.com/ProjectAthenaa/perimeterx-service/responsedeob"
	"github.com/ProjectAthenaa/perimeterx-service/siteconstants"
	"github.com/ProjectAthenaa/pxutils"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type PX3 struct {
	T string `json:"t"`
	D struct {
		PX34   string        `json:"PX34,omitempty"`
		PX55   string        `json:"PX55,omitempty"`
		PX59   string        `json:"PX59,omitempty"`
		PX60   bool          `json:"PX60,omitempty"`
		PX61   string        `json:"PX61,omitempty"`
		PX62   string        `json:"PX62,omitempty"`
		PX63   string        `json:"PX63,omitempty"`
		PX64   string        `json:"PX64,omitempty"`
		PX65   string        `json:"PX65,omitempty"`
		PX66   string        `json:"PX66,omitempty"`
		PX68   bool          `json:"PX68,omitempty"`
		PX69   string        `json:"PX69,omitempty"`
		PX85   []string      `json:"PX85,omitempty"`
		//PX86   bool          `json:"PX86,omitempty"`
		PX87   bool          `json:"PX87,omitempty"`
		PX88   bool          `json:"PX88,omitempty"`
		PX89   bool          `json:"PX89,omitempty"`
		PX90   []string      `json:"PX90,omitempty"`
		PX91   int           `json:"PX91,omitempty"`
		PX92   int           `json:"PX92,omitempty"`
		PX93   string        `json:"PX93,omitempty"`
		PX94   int           `json:"PX94,omitempty"`
		PX95   bool          `json:"PX95,omitempty"`
		PX96   string        `json:"PX96,omitempty"`
		PX120  []interface{} `json:"PX120,omitempty"`
		PX133  bool          `json:"PX133,omitempty"`
		PX134  bool          `json:"PX134,omitempty"`
		PX135  bool          `json:"PX135,omitempty"`
		PX138  bool          `json:"PX138,omitempty"`
		PX139  bool          `json:"PX139,omitempty"`
		PX141  bool          `json:"PX141,omitempty"`
		PX143  bool          `json:"PX143,omitempty"`
		PX144  bool          `json:"PX144,omitempty"`
		PX145  bool          `json:"PX145,omitempty"`
		PX146  bool          `json:"PX146,omitempty"`
		PX147  bool          `json:"PX147,omitempty"`
		PX148  bool          `json:"PX148,omitempty"`
		PX149  bool          `json:"PX149,omitempty"`
		PX150  bool          `json:"PX150,omitempty"`
		PX151  bool          `json:"PX151,omitempty"`
		PX152  bool          `json:"PX152,omitempty"`
		PX153  bool          `json:"PX153,omitempty"`
		PX154  int           `json:"PX154,omitempty"`
		PX155  string        `json:"PX155,omitempty"`
		PX163  bool          `json:"PX163,omitempty"`
		PX166  bool          `json:"PX166,omitempty"`
		PX167  bool          `json:"PX167,omitempty"`
		PX169  int           `json:"PX169,omitempty"`
		PX170  int           `json:"PX170,omitempty"`
		PX184  bool          `json:"PX184,omitempty"`
		PX185  int           `json:"PX185,omitempty"`
		PX186  int           `json:"PX186,omitempty"`
		PX187  int           `json:"PX187,omitempty"`
		PX188  int           `json:"PX188,omitempty"`
		PX190  string        `json:"PX190,omitempty"`
		PX191  int           `json:"PX191,omitempty"`
		PX192  bool          `json:"PX192,omitempty"`
		PX194  bool          `json:"PX194,omitempty"`
		PX195  bool          `json:"PX195,omitempty"`
		PX196  bool          `json:"PX196,omitempty"`
		PX207  bool          `json:"PX207,omitempty"`
		PX208  string        `json:"PX208,omitempty"`
		PX218  int           `json:"PX218,omitempty"`
		PX229  int           `json:"PX229,omitempty"`
		PX230  int           `json:"PX230,omitempty"`
		PX231  int           `json:"PX231,omitempty"`
		PX232  int           `json:"PX232,omitempty"`
		PX234  bool          `json:"PX234,omitempty"`
		PX235  bool          `json:"PX235,omitempty"`
		PX236  bool          `json:"PX236,omitempty"`
		PX237  int           `json:"PX237,omitempty"`
		PX238  string        `json:"PX238,omitempty"`
		PX239  bool          `json:"PX239,omitempty"`
		PX240  bool          `json:"PX240,omitempty"`
		PX247  int           `json:"PX247,omitempty"`
		PX251  bool          `json:"PX251,omitempty"`
		PX254  bool          `json:"PX254,omitempty"`
		PX268  bool          `json:"PX268,omitempty"`
		PX269  int           `json:"PX269,omitempty"`
		PX270  int           `json:"PX270,omitempty"`
		PX295  bool          `json:"PX295,omitempty"`
		PX313  []string      `json:"PX313,omitempty"`
		PX314  bool          `json:"PX314,omitempty"`
		PX357  string        `json:"PX357,omitempty"`
		PX358  string        `json:"PX358,omitempty"`
		PX359  string        `json:"PX359,omitempty"`
		PX371  bool          `json:"PX371,omitempty"`
		PX397  bool          `json:"PX397,omitempty"`
		PX399  string        `json:"PX399,omitempty"`
		PX400  int           `json:"PX400,omitempty"`
		PX402  int           `json:"PX402,omitempty"`
		PX404  string        `json:"PX404,omitempty"`
		PX405  bool          `json:"PX405,omitempty"`
		PX411  int           `json:"PX411,omitempty"`
		PX443  bool          `json:"PX443,omitempty"`
		PX466  bool          `json:"PX466,omitempty"`
		PX467  bool          `json:"PX467,omitempty"`
		PX468  bool          `json:"PX468,omitempty"`
		PX547  bool          `json:"PX547,omitempty"`
		PX548  int           `json:"PX548,omitempty"`
		PX549  int           `json:"PX549,omitempty"`
		PX552  string        `json:"PX552,omitempty"`
		PX714  string        `json:"PX714,omitempty"`
		PX715  string        `json:"PX715,omitempty"`
		PX716  string        `json:"PX716,omitempty"`
		PX717  string        `json:"PX717,omitempty"`
		PX722  string        `json:"PX722,omitempty"`
		PX723  string        `json:"PX723,omitempty"`
		PX724  string        `json:"PX724,omitempty"`
		PX725  string        `json:"PX725,omitempty"`
		PX726  string        `json:"PX726,omitempty"`
		PX727  string        `json:"PX727,omitempty"`
		PX729  string        `json:"PX729,omitempty"`
		PX821  int64         `json:"PX821,omitempty"`
		PX822  int           `json:"PX822,omitempty"`
		PX823  int           `json:"PX823,omitempty"`
		PX850  int           `json:"PX850,omitempty"`
		PX851  int           `json:"PX851,omitempty"`
		PX943  string        `json:"PX943,omitempty"`
		PX982  string        `json:"PX982,omitempty"`
		PX983  string        `json:"PX983,omitempty"`
		PX985  int           `json:"PX985,omitempty"`
		PX986  string        `json:"PX986,omitempty"`
		PX1008 int           `json:"PX1008,omitempty"`
		PX1019 string        `json:"PX1019,omitempty"`
		PX1020 string        `json:"PX1020,omitempty"`
		PX1021 string        `json:"PX1021,omitempty"`
		PX1022 string        `json:"PX1022,omitempty"`
		PX1025 bool          `json:"PX1025,omitempty"`
		PX1033 string        `json:"PX1033,omitempty"`
		PX1035 bool          `json:"PX1035,omitempty"`
		PX1038 string        `json:"PX1038,omitempty"`
		PX1054 int64         `json:"PX1054,omitempty"`
		PX1055 int64         `json:"PX1055,omitempty"`
		PX1056 int64         `json:"PX1056,omitempty"`
		PX1065 int           `json:"PX1065,omitempty"`
		PX1139 bool          `json:"PX1139,omitempty"`
		PX1142 int           `json:"PX1142,omitempty"`
		PX1143 int           `json:"PX1143,omitempty"`
		PX1144 bool          `json:"PX1144,omitempty"`
		PX1146 int           `json:"PX1146,omitempty"`
		PX1147 int           `json:"PX1147,omitempty"`
		PX1152 int           `json:"PX1152,omitempty"`
		PX1153 int           `json:"PX1153,omitempty"`
		PX1154 bool          `json:"PX1154,omitempty"`
		PX1155 string        `json:"PX1155,omitempty"`
		PX1157 int           `json:"PX1157,omitempty"`
		PX1173 int           `json:"PX1173,omitempty"`
		PX1179 bool          `json:"PX1179,omitempty"`
		PX1180 bool          `json:"PX1180,omitempty"`
	} `json:"d"`
}

func GeneratePX3(uuid string, resobj responsedeob.ResponseJSON, sitedata *siteconstants.SiteData) []PX3 {
	px3obj := InstantiatePX3()
	px3obj.D.PX96 = sitedata.Url
	px3obj.D.PX357 = pxutils.H12(pxutils.H1(resobj.VID, siteconstants.UA))
	px3obj.D.PX358 = pxutils.H12(pxutils.H1(resobj.SID, siteconstants.UA))
	px3obj.D.PX359 = pxutils.H12(pxutils.H1(uuid, siteconstants.UA))
	px3obj.D.PX943 = resobj.WCS
	px3obj.D.PX982 = resobj.STS
	px3obj.D.PX983 = resobj.CLS0
	px3obj.D.PX985, _ = strconv.Atoi(resobj.DRC)
	px3obj.D.PX986 = resobj.CLS1
	px3obj.D.PX1038 = uuid
	return []PX3{px3obj}
}

func InstantiatePX3() PX3 {
	var payload PX3
	payload.T = "PX3"
	payload.D.PX34 = "Error"
	payload.D.PX55 = ""
	payload.D.PX59 = siteconstants.UA
	payload.D.PX60 = true
	payload.D.PX61 = "en-US"
	payload.D.PX62 = "Gecko"
	//payload.D.PX63 = "Win32"
	payload.D.PX64 = strings.SplitN(siteconstants.UA, "/", 2)[1]
	payload.D.PX65 = "Netscape"
	payload.D.PX66 = "Mozilla"
	payload.D.PX68 = true
	payload.D.PX69 = "20030107"
	payload.D.PX85 = []string{
		"PDF Viewer",
		"Chrome PDF Viewer",
		"Chromium PDF Viewer",
		"Microsoft Edge PDF Viewer",
		"WebKit built-in PDF",
	}
	payload.D.PX87 = true
	payload.D.PX88 = true
	payload.D.PX89 = true
	payload.D.PX90 = []string{"loadTimes", "csi", "app", "runtime"}
	payload.D.PX91 = 3148
	payload.D.PX92 = 886
	payload.D.PX93 = "3840X1080"
	payload.D.PX94 = 1 + rand.Intn(3)
	payload.D.PX95 = true
	payload.D.PX120 = []interface{}{}
	payload.D.PX133 = true
	payload.D.PX134 = true
	payload.D.PX138 = true
	payload.D.PX141 = false
	//payload.D.PX143 = true
	payload.D.PX144 = true
	payload.D.PX145 = true
	payload.D.PX146 = true
	payload.D.PX147 = false
	payload.D.PX148 = false
	payload.D.PX149 = true
	payload.D.PX150 = true
	payload.D.PX151 = false
	payload.D.PX152 = false
	payload.D.PX153 = false
	//payload.D.PX154 = 240
	payload.D.PX155 = time.Now().Local().Format("Mon Jan 02 2006 15:04:05 GMT-0700")
	//payload.D.PX155 = "Fri Apr 02 2021 00:31:32 GMT-0400 (Eastern Daylight Time)" // timezone
	payload.D.PX163 = false
	payload.D.PX166 = true
	payload.D.PX167 = true
	payload.D.PX169 = 2
	payload.D.PX170 = 5
	payload.D.PX184 = true
	payload.D.PX185 = 774
	payload.D.PX186 = 668
	payload.D.PX187 = 0
	payload.D.PX188 = 0
	payload.D.PX190 = ""
	payload.D.PX191 = 0
	payload.D.PX192 = false
	payload.D.PX194 = false
	payload.D.PX195 = true
	payload.D.PX196 = false
	payload.D.PX207 = false
	payload.D.PX208 = "visible"
	payload.D.PX218 = 0
	payload.D.PX229 = 24
	payload.D.PX230 = 24
	payload.D.PX231 = 891
	payload.D.PX232 = 1560
	payload.D.PX234 = false
	payload.D.PX235 = false
	payload.D.PX236 = false
	payload.D.PX237 = 0
	payload.D.PX238 = "missing"
	payload.D.PX239 = false
	payload.D.PX240 = false
	payload.D.PX247 = 0
	payload.D.PX251 = false
	payload.D.PX254 = false
	payload.D.PX268 = false
	payload.D.PX295 = false
	payload.D.PX269 = 3778
	payload.D.PX270 = 886
	payload.D.PX313 = []string{"en-US"}
	payload.D.PX314 = false
	payload.D.PX371 = true
	payload.D.PX397 = false
	payload.D.PX399 = "false"
	payload.D.PX400 = 539
	//payload.D.PX402 = 1
	payload.D.PX404 = "144|66|66|180|80"
	payload.D.PX405 = true
	payload.D.PX411 = 1
	payload.D.PX443 = true
	payload.D.PX466 = true
	payload.D.PX467 = true
	payload.D.PX468 = true
	payload.D.PX547 = true
	//payload.D.PX548 = 1
	payload.D.PX549 = 1
	payload.D.PX552 = "false"
	payload.D.PX714 = "64556c77"
	payload.D.PX715 = ""
	payload.D.PX716 = "9f762773"
	payload.D.PX717 = "dae10548"
	payload.D.PX722 = "a3d12c4"
	payload.D.PX723 = "a3d12c4"
	payload.D.PX724 = "10207b2f"
	payload.D.PX725 = "10207b2f"
	payload.D.PX726 = "82002457"
	payload.D.PX727 = ""
	payload.D.PX729 = "90e65465"
	payload.D.PX821 = 4294705152 //total memory
	payload.D.PX822 = 10000000 + rand.Intn(40000000)
	payload.D.PX823 = payload.D.PX822 - 2000000 - rand.Intn(8000000)
	payload.D.PX850 = 1
	payload.D.PX851 = 7000 + rand.Intn(1500)
	payload.D.PX1008 = 3600
	payload.D.PX1019 = "7d688b1b"
	payload.D.PX1020 = "7766a52d"
	payload.D.PX1021 = "180dd7e3"
	payload.D.PX1022 = "6a90378d"
	payload.D.PX1025 = false
	payload.D.PX1033 = "e0eaf10e"
	payload.D.PX1035 = false
	payload.D.PX1054 = time.Now().UnixMilli()
	payload.D.PX1055 = time.Now().UnixMilli()-1000
	payload.D.PX1056 = time.Now().UnixMilli()+1000
	payload.D.PX1065 = 1
	payload.D.PX1139 = false
	payload.D.PX1142 = 2
	payload.D.PX1143 = 7
	payload.D.PX1144 = true
	payload.D.PX1146 = 2
	payload.D.PX1147 = 1
	payload.D.PX1152 = 10
	payload.D.PX1153 = 0
	payload.D.PX1154 = false
	payload.D.PX1155 = "4g"
	payload.D.PX1157 = 8
	payload.D.PX1173 = 1
	payload.D.PX1179 = true
	payload.D.PX1180 = true
	return payload
}
