package unused
//import (
//	"encoding/json"
//
//	"github.com/CicadaAIO/px-api/binpkg"
//
//	//"github.com/smartystreets/goconvey/web/server/parser"
//	"io/ioutil"
//	"math"
//	"math/rand"
//	"os"
//	"strconv"
//	"time"
//
//	"github.com/CicadaAIO/px-api/utils"
//	pxutils "github.com/incizzle/perimeterx-utils-go"
//
//	"github.com/CicadaAIO/px-api/entities"
//	//"github.com/ua-parser/uap-go/uaparser"
//)
//
//// PX3Data is a struct containing the json payload of a px3 bundle
//type PX3Data struct {
//	PX234             bool     `json:"PX234"`             // value:	!!window.spawn
//	PX235             bool     `json:"PX235"`             // value:  !!window.emit
//	PX151             bool     `json:"PX151"`             // value:  !!window.webdriver
//	PX239             bool     `json:"PX239"`             // value:  !!window._Selenium_IDE_Recorder
//	PX240             bool     `json:"PX240"`             // value:  !!document.__webdriver_script_fn
//	PX152             bool     `json:"PX152"`             // value:  !!window.domAutomation || !!window.domAutomationController
//	PX153             bool     `json:"PX153"`             // value:  !!window._phantom || !!window.callPhantom
//	PX314             bool     `json:"PX314"`             // value:  !!window.geb
//	PX192             bool     `json:"PX192"`             // value:  !!window.awesomium
//	PX196             bool     `json:"PX196"`             // value:  toStringIsNative(window.RunPerfTest)
//	PX207             bool     `json:"PX207"`             // value:  !!window.fmget_targets
//	PX251             bool     `json:"PX251"`             // value:  !!window.__nightmare
//	PX982             int64    `json:"PX982"`             // set to sts response
//	PX983             string   `json:"PX983"`             // cls.split("|")[0]
//	ClsAndStsMathsKey string   `json:"ClsAndStsMathsKey"` //  t[xor(t["PX983"], t["PX982"] % 10 + 2)] = xor(t["PX983"], t["PX982"] % 10 + 1);
//	PX986             string   `json:"PX986"`             // cls.split("|")[1]
//	PX985             int      `json:"PX985"`             // drc (string to int)
//	PX1033            string   `json:"PX1033"`            // encode(ie(window.chrome))
//	PX1019            string   `json:"PX1019"`            // hasOwnPropertyCheck(window, "loads of shit"); array for this starts onrendersubtreeactivation ends onelementpainted
//	PX1020            string   `json:"PX1020"`            // hasOwnPropertyCheck(document, c)
//	PX1021            string   `json:"PX1021"`            // hasOwnPropertyCheck(navigator, u)
//	PX1022            string   `json:"PX1022"`            // hasOwnPropertyCheck(location, m)
//	PX1035            bool     `json:"PX1035"`            // delete window.webdriver check (always true)
//	PX1139            bool     `json:"PX1139"`            // delete window.plugins check (always false)
//	PX1025            bool     `json:"PX1025"`            //TODO figure out this shit properly, atm seems to be getOwnPropertyDescriptor check on window.navigator
//	PX359             string   `json:"PX359"`             // some sha hmac of uuid and useragent
//	PX943             string   `json:"PX943"`             // wcs response
//	PX357             string   `json:"PX357"`             // sha hmac of vid and useragent
//	PX358             string   `json:"PX358"`             // sha hmac of sid and useragent
//	PX229             int      `json:"PX229"`             // window.screen && +screen.colorDepth || 0
//	PX230             int      `json:"PX230"`             // screen.pixelDepth
//	PX91              int      `json:"PX91"`              // window.screen.width
//	PX92              int      `json:"PX92"`              // window.screen.height
//	PX269             int      `json:"PX269"`             // window.screen.availWidth
//	PX270             int      `json:"PX270"`             // window.screen.availHeight
//	PX93              string   `json:"PX93"`              // screen.width.toString() + "X" + screen.height.toString()
//	PX185             int      `json:"PX185"`             // window.innerHeight
//	PX186             int      `json:"PX186"`             // window.innerWidth
//	PX187             int      `json:"PX187"`             // window.scrollX || window.pageXOffset || 0
//	PX188             int      `json:"PX188"`             // window.scrollY || window.pageYOffset || 0
//	PX95              bool     `json:"PX95"`              // value: !(0 === window.outerWidth && 0 === window.outerHeight)
//	PX400             int      `json:"PX400"`             // value: "" + !!window.fetch + !!window.performance.toJSON + !!document.createELement
//	PX404             string   `json:"PX404"`             // checks for various properties on window.chrome and appends the length of them + "|"
//	PX90              []string `json:"PX90"`              // Object.keys(window.chrome)
//	PX190             string   `json:"PX190"`             // window.chrome.runtime.id || "" (usually "")
//	PX552             string   `json:"PX552"`             // navigator.webdriver + ""
//	PX399             string   `json:"PX399"`             // navigator.webdriver + ""
//	PX549             int      `json:"PX549"`             // "webdriver" in navigator
//	PX411             int      `json:"PX411"`             // "webdriver" in navigator
//	PX405             bool     `json:"PX405"`             // value:  !!window.caches (NOT STATIC)
//	PX547             bool     `json:"PX547"`             // value:  !!window.caches
//	PX134             bool     `json:"PX134"`             // navigator.plugins check ( effectively "[object PluginArray]" === navigator.plugins.toString()), preset as this is always true
//	PX89              bool     `json:"PX89"`              // navigator.plugins check
//	PX170             int      `json:"PX170"`             // navigator.plugins.length
//	PX85              []string `json:"PX85"`              // navigator.plugins stringify
//	PX59              string   `json:"PX59"`              // useragent
//	PX61              string   `json:"PX61"`              // navigator.language
//	PX313             []string `json:"PX313"`             // navigator.languages
//	PX63              string   `json:"PX63"`              // navigator.platform
//	PX86              bool     `json:"PX86"`              // value: !!(navigator.doNotTrack || null === navigator.doNotTrack || navigator.msDoNotTrack || window.doNotTrack)
//	PX1173            int      `json:"PX1173"`
//	PX154             int      `json:"PX154"`            // (new Date).getTimezoneOffset();
//	PX1157            int      `json:"PX1157"`           // navigator.deviceMemory
//	PX133             bool     `json:"PX133"`            // check for MSMimeTypesCollection in  navigator.mimeTypes
//	PX88              bool     `json:"PX88"`             // check for MSMimeTypesCollection in  navigator.mimeTypes
//	PX169             int      `json:"PX169"`            // navigator.mimeTypes.length
//	PX62              string   `json:"PX62"`             // navigator.product
//	PX69              string   `json:"PX69"`             // navigator.productSub
//	PX64              string   `json:"PX64"`             // navigator.appVersion
//	PX65              string   `json:"PX65"`             // navigator.appName
//	PX66              string   `json:"PX66"`             // navigator.appCodeName
//	PX1144            bool     `json:"PX1144,omitempty"` // "query" === navigator.permissions.query.name, empty on safari
//
//	// performance.memory not supported on firefox or safari
//	PX1152 int    `json:"PX1152,omitempty"` // navigator.connection.downlink
//	PX1153 int    `json:"PX1153,omitempty"` // navigator.connection.rtt
//	PX1154 *bool  `json:"PX1154,omitempty"` // navigator.connection.saveData
//	PX1155 string `json:"PX1155,omitempty"` // navigator.connection.effectiveType
//
//	PX60 bool `json:"PX60"` // "onLine" in navigator && !0 === navigator.onLine
//	PX87 bool `json:"PX87"` // navigator.geolocation + "" == "[object Geolocation]"
//
//	// performance.memory not supported on firefox or safari
//	PX821 int64 `json:"PX821,omitempty"` // window.performance["memory"].jsHeapSizeLimit
//	PX822 int   `json:"PX822,omitempty"` // window.performance["memory"].totalJSHeapSize
//	PX823 int   `json:"PX823,omitempty"` // window.performance["memory"].usedJSHeapSize
//
//	PX147  bool          `json:"PX147"`  //value: !!window.ActiveXObject
//	PX155  string        `json:"PX155"`  // window.Date()
//	PX236  bool          `json:"PX236"`  // value: !!window.Buffer
//	PX194  bool          `json:"PX194"`  // value:  !!window.v8Locale
//	PX195  bool          `json:"PX195"`  // value: !!navigator.sendBeacon
//	PX237  int           `json:"PX237"`  // navigator.msMaxTouchPoints (only on IE 10)
//	PX238  string        `json:"PX238"`  // navigator.msDoNotTrack (IE only again)
//	PX208  string        `json:"PX208"`  // document.visibilityState
//	PX218  int           `json:"PX218"`  // +document.documentMode || 0 (IE / edge only once again)
//	PX231  int           `json:"PX231"`  // +window.outerHeight
//	PX232  int           `json:"PX232"`  // +window.outerWidth
//	PX254  bool          `json:"PX254"`  // value: !!window.showModalDialog
//	PX295  bool          `json:"PX295"`  // if document can create "TouchEvent event"
//	PX268  bool          `json:"PX268"`  // value: !!window.ontouchstart
//	PX166  bool          `json:"PX166"`  // function string check on window.setTimeout
//	PX138  bool          `json:"PX138"`  //TODO function string check on window.openDatabase, figure out different browsers and wether they have this func
//	PX143  bool          `json:"PX143"`  // func string check (window.BatteryManager), chrome, edge and opera
//	PX1142 int           `json:"PX1142"` //TODO figure out what thing is thing.cssFromResourceApi
//	PX1143 int           `json:"PX1143"` // thing.imgFromResourceApi
//	PX1146 int           `json:"PX1146"` // thing.fontFromResourceApi
//	PX1147 int           `json:"PX1147"` // thing.cssFromStyleSheets
//	PX714  string        `json:"PX714"`  //TODO figure out weirdOp weirdOp(window.console.log)
//	PX715  string        `json:"PX715"`  // weirdOp(Object.getOwnPropertyDescriptor(HTMLDocument.prototype, "cookie").get)
//	PX724  string        `json:"PX724"`  // weirdOp(Object.prototype.toString)
//	PX725  string        `json:"PX725"`  // weirdOp(navigator.toString)
//	PX729  string        `json:"PX729"`  // weirdOp2(webdriver shit)
//	PX443  bool          `json:"PX443"`  // value: !!window.isSecureContext
//	PX466  bool          `json:"PX466"`  // value: !!window.Worklet
//	PX467  bool          `json:"PX467"`  // value: !!window.AudioWorklet
//	PX468  bool          `json:"PX468"`  // value: !!window.AudioWorkletNode
//	PX191  int           `json:"PX191"`  // window.self === window.top ? 0 : 1
//	PX94   int           `json:"PX94"`   // window.history && "number" == typeof window.history.length && window.history.length || -1
//	PX120  []interface{} `json:"PX120"`  //TODO document.location.ancestorOrigins some shit done w this
//	PX141  bool          `json:"PX141"`  //value: !!window.onorientationchange
//	PX96   string        `json:"PX96"`   // window.location.href
//	PX55   string        `json:"PX55"`   // encodeURIComponent(document.referrer)
//	PX34   string        `json:"PX34"`
//	PX1065 int           `json:"PX1065"` //TODO check i believe this is number of PX3 requests sent
//	PX850  int           `json:"PX850"`  // number of requests sent from one script instance
//	PX851  int           `json:"PX851"`  // Math.round(window.performance.now())
//	PX1054 int64         `json:"PX1054"` // (new Date).getTime()
//	PX1008 int           `json:"PX1008"` // TODO strange maths shit
//	PX1055 int64         `json:"PX1055"` // (new Date).getTime()
//	PX1056 int64         `json:"PX1056"` //  (new Date).getTime()
//	PX1038 string        `json:"PX1038"` // uuid
//	PX371  bool          `json:"PX371"`  // some script checks (always true)
//
//	// these properties are only set on blocked / captcha pages
//	PX250 string `json:"PX250,omitempty"` // depends on captcha mode (PX560 for pxhc or pxc, I believe that's px hold captcha or px captcha)
//	PX708 string `json:"PX708,omitempty"` // window._pxAction
//}
//
//// read presets, these are values that are never going to change based on device,
//// e.g. PX234, it's value is !!window.webdriver so will be false for every device
//func readFile(path string) string {
//	fi, err := os.Open(path)
//	if err != nil {
//		panic(err)
//	}
//
//	defer func() {
//		if err := fi.Close(); err != nil {
//			panic(err)
//		}
//	}()
//
//	out, err := ioutil.ReadAll(fi)
//	if err != nil {
//		panic(err)
//	}
//
//	return string(out)
//}
//
////var presetPath, _ = filepath.Abs("./bin/px3-presets.json")
//
////var presets = readFile(presetPath)
//
//var presets = binpkg.Px3Presets
//var initialData PX3Data
//
//func init() {
//	if err := json.Unmarshal([]byte(presets), &initialData); err != nil {
//		panic(err)
//	}
//}
//
//// CAN'T DO THIS ON LAMBDA
//
//// GeneratePx3 gens
//func (session *GenSession) GeneratePx3() entities.PXPayload {
//	var d PX3Data
//
//	d = initialData
//
//	d.ClsAndStsMathsKey = "ClsAndStsValue"
//
//	d.PX982, _ = strconv.ParseInt(session.Sts, 10, 64)
//	// if strings.Contains(session.Cls, "|") {
//	if len(session.Cls) == 2 {
//		d.PX983 = session.Cls[0]
//		d.PX986 = session.Cls[1]
//
//	} else {
//		d.PX983 = session.Cls[0] //cls part 1
//		d.PX986 = session.Cls[0] //cls part 2
//	}
//
//	d.PX985, _ = strconv.Atoi(session.PXResponse.Do.Drc)
//
//	d.PX943 = session.Wcs
//	d.PX359 = pxutils.H12(pxutils.H1(session.Uuid, session.Device.Data.UserAgent))
//	d.PX357 = pxutils.H12(pxutils.H1(session.Vid, session.Device.Data.UserAgent))
//	d.PX358 = pxutils.H12(pxutils.H1(session.RawSid, session.Device.Data.UserAgent))
//	//need to remove padding here
//
//	// this.Px359 = pxutils.H12(pxutils.H1(uuid, device.Dev.UserAgent))
//	// this.Px357 = pxutils.H12(pxutils.H1(vid, device.Dev.UserAgent))
//	// this.PX358 = pxutils.H12(pxutils.H1(sid, device.Dev.UserAgent))
//
//	d.PX59 = session.Device.Data.UserAgent
//	d.PX61 = session.Device.Data.Language
//	d.PX63 = session.Device.Data.Platform
//	d.PX64 = session.Device.Data.AppVersion
//	d.PX85 = session.Device.Data.Plugins
//	d.PX91 = session.Device.Data.ScreenLength
//	d.PX92 = session.Device.Data.ScreenHeight
//	d.PX93 = strconv.Itoa(session.Device.Data.ResolutionX) + "X" + strconv.Itoa(session.Device.Data.ResolutionY)
//	d.PX169 = session.Device.Data.MimeTypes
//	d.PX170 = session.Device.Data.PluginsLength
//	d.PX185 = session.Device.Data.InnerHeight
//	d.PX186 = session.Device.Data.InnerWidth
//	d.PX229 = session.Device.Data.ColorDepth
//	d.PX230 = session.Device.Data.PixelDepth
//	d.PX231 = session.Device.Data.OuterHeight
//	d.PX232 = session.Device.Data.OuterWidth
//	d.PX269 = session.Device.Data.AvailWidth
//	d.PX270 = session.Device.Data.AvailHeight
//
//	d.PX400 = 111
//	if session.Site.Name == "Snipes" {
//		// snipes site edits window.fetch so it's no longer native code changing the length of it and subsequently this property
//		d.PX400 = 401
//
//	}
//
//	d.PX404 = "144|54|54|180|68"
//	d.PX90 = []string{"loadTimes", "csi", "app", "runtime"}
//	d.PX552 = "false"
//	d.PX399 = d.PX552
//
//	d.PX549 = 1
//	d.PX411 = d.PX549
//
//	d.PX405 = true
//	d.PX547 = d.PX405
//
//	d.PX134 = true
//	d.PX89 = d.PX134
//
//	d.PX170 = len(session.Device.Data.Plugins)
//	d.PX85 = session.Device.Data.Plugins
//	d.PX59 = session.Device.Data.UserAgent
//	d.PX61 = session.Device.Data.Language
//	d.PX313 = session.Device.Data.Languages
//	d.PX1173 = len(session.Device.Data.Languages)
//
//	d.PX63 = session.Device.Data.Platform
//	d.PX64 = session.Device.Data.AppVersion
//	d.PX65 = session.Device.Data.AppName
//	d.PX66 = session.Device.Data.AppCodeName
//	d.PX154 = 240 //! new york timezone
//	d.PX1157 = session.Device.Data.DeviceMemory
//	d.PX1152 = 10 //TODO make dynamic
//	d.PX1153 = 100
//	d.PX1154 = &[]bool{false}[0]
//	d.PX1155 = "4g"
//
//	d.PX60 = true
//	d.PX87 = true
//
//	// TODO make dynamic / randomised
//	d.PX821 = session.Device.Data.HeapSizeLimit
//	d.PX822 = session.Device.Data.TotalHeapSize
//	d.PX823 = session.Device.Data.UsedHeapSize // TODO fix this, harvester will have got wacky usage amount
//
//	d.PX1144 = true
//	d.PX138 = true
//	d.PX143 = true
//
//	d.PX714 = "64556c77"
//	d.PX715 = "" // trying to retrieve .get causes an error so this should be empty
//	d.PX724 = "10207b2f"
//	d.PX725 = "10207b2f"
//	d.PX729 = "90e65465"
//
//	//! might not need to be chrome only, just doing this to finish px3
//	d.PX443 = true
//	d.PX466 = true
//	d.PX467 = true
//	d.PX468 = true
//	//case "Firefox":
//	//	d.PX1144 = true
//	//	d.PX138 = false
//	//	d.PX143 = false
//	//case "Safari":
//	//	d.PX138 = false
//	//	d.PX143 = false
//	//}
//
//	d.PX147 = false
//	d.PX155 = utils.TimeIn(time.Now(), "America/New_York").Format("Mon Jan 2 2006 15:04:05 GMT-0700 (Eastern Daylight Time)")
//
//	d.PX208 = "visible" // should never be hidden, always want to pretend user tabbed in
//
//	d.PX231 = session.Device.Data.OuterHeight
//	d.PX232 = session.Device.Data.OuterWidth
//
//	d.PX254 = false //TODO, this value is true for desktop safari
//	d.PX295 = false //TODO, could be true for phones
//	d.PX295 = false //TODO, check if this is false for all mobile devices
//	d.PX268 = false //TODO, same as above
//
//	d.PX1173 = 2
//
//	switch session.Site.Name {
//	case "Walmart":
//		d.PX34 = "Error\n    at xe (https://www.walmart.com/px/PXu6b0qd2S/init.js:3:10744)\n    at we (https://www.walmart.com/px/PXu6b0qd2S/init.js:3:1750)"
//		if session.Type == entities.HoldCaptcha {
//			//! CHROME ONLY, if using other browsers must be dynamic
//			d.PX1142 = 2
//			d.PX1143 = 0
//			d.PX1146 = 0
//			d.PX1147 = 0
//		} else {
//			// normal request
//			d.PX1142 = 4
//			d.PX1143 = 22
//			d.PX1146 = 2
//			d.PX1147 = 2
//		}
//
//	case "Snipes":
//	}
//
//	d.PX191 = 0
//	d.PX94 = 4                // number of pages navigated before accessing this page
//	d.PX96 = session.Site.URL // href
//	d.PX55 = ""               // document.referer
//	d.PX1065 = 2              //! always seems to be 2, unsure why
//	d.PX850 = 1               // already sent px2
//
//	d.PX851 = int(math.Round(rand.Float64()*1000 + 1000))
//	// switch device.Data.
//
//	stamp := utils.TimeIn(time.Now(), "America/New_York").UnixNano() / int64(time.Millisecond)
//	start3 := stamp - int64(math.Round(rand.Float64()*10+65))
//
//	d.PX1054 = start3
//	d.PX1008 = 3600
//	d.PX1055 = session.StartTime
//	d.PX1056 = stamp
//	d.PX1038 = session.Uuid
//	d.PX371 = true
//
//	switch session.Site.Name {
//	case "Walmart":
//		d.PX34 = "Error\n    at xe (https://www.walmart.com/px/PXu6b0qd2S/init.js:3:10744)\n    at we (https://www.walmart.com/px/PXu6b0qd2S/init.js:3:1750)"
//	}
//
//	switch session.Type {
//	case entities.NonBlocked:
//		// do nothing
//	case entities.HoldCaptcha:
//		d.PX96 = session.Site.URL + "blocked" // href
//
//		d.PX250 = "PX560"
//		d.PX708 = "pxhc"
//	}
//
//	return entities.PXPayload{
//		T: "PX3",
//		D: d,
//	}
//}
