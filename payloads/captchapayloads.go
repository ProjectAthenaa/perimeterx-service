package payloads

import (
	"math/rand"
	"strings"
)

var (
	PX280LIST = []string{"AMD Radeon R9 M370X OpenGL Engine", "ANGLE (NVIDIA GeForce GTX 1660 SUPER Direct3D11 vs_5_0 ps_5_0)", "ANGLE (NVIDIA GeForce RTX 2080 Direct3D11 vs_5_0 ps_5_0)", "Intel(R) Iris(TM) Plus Graphics OpenGL Engine", "ANGLE (NVIDIA GeForce GTX 1060 3GB Direct3D11 vs_5_0 ps_5_0)", "ANGLE (NVIDIA GeForce GTX 1070 Direct3D11 vs_5_0 ps_5_0)", "ANGLE (NVIDIA GeForce GTX 1080 Direct3D11 vs_5_0 ps_5_0)", "ANGLE (AMD Radeon RX 5600 XT Direct3D11 vs_5_0 ps_5_0)", "ANGLE (Radeon RX 570 Series Direct3D11 vs_5_0 ps_5_0)", "ANGLE (Radeon RX 580 Series Direct3D11 vs_5_0 ps_5_0)", "Google SwiftShader"}
	PX176CAPTCHALIST = []string{"LI3", "DIV1", "DIV2", "DIV5", "A", "#povActive", "SECTION2", "IMG", "DIV4", "UL", "LI3","DIV2","DIV1","DIV2","DIV1","DIV2","DIV1","DIV2","DIV1", "P2", "LI1", "BODY", "IFRAME","#px-captcha"}
)

type RECAPPAYLOAD struct {
	T string `json:"t"`
	D struct {
		PX70   int    `json:"PX70"`
		PX34   string `json:"PX34"`
		PX610  bool   `json:"PX610"`
		PX754  bool   `json:"PX754"`
		PX755  string `json:"PX755"`
		PX756  string `json:"PX756"`
		PX757  string `json:"PX757"`
		PX850  int    `json:"PX850"`
		PX851  int    `json:"PX851"`
		PX1056 int64  `json:"PX1056"`
		PX1038 string `json:"PX1038"`
		PX371  bool   `json:"PX371"`
		PX250  string `json:"PX250"`
		PX788  string `json:"PX788, omitempty"`
		PX708  string `json:"PX708, omitempty"`
		PX1129 bool  `json:"PX1129, omitempty"`
		PX96   string `json:"PX96"`
	} `json:"d"`
}
type HCAPPAYLOAD struct {
	T string `json:"t"`
	D struct {
		PX70   int               `json:"PX70"`
		PX34   string            `json:"PX34"`
		PX1129 bool              `json:"PX1129"`
		PX610  bool              `json:"PX610"`
		PX1132 string            `json:"PX1132"`
		PX1133 float64           `json:"PX1133"`
		PX179  []string          `json:"PX179"`
		PX827  []PX827KEYEVENT   `json:"PX827"`
		PX582  []PX582MOUSEEVENT `json:"PX582"`
		PX608  []string          `json:"PX608"`
		PX74   int               `json:"PX74"`
		PX75   int               `json:"PX75"`
		PX76   int               `json:"PX76"`
		PX77   int               `json:"PX77"`
		PX91   int               `json:"PX91"`
		PX92   int               `json:"PX92"`
		PX789  string            `json:"PX789"`
		PX790  int               `json:"PX790"`
		PX791  int               `json:"PX791"`
		PX792  int      `json:"PX792"`
		PX793  float64  `json:"PX793"`
		PX794  float64  `json:"PX794"`
		PX821  int64    `json:"PX821"`
		PX822  int      `json:"PX822"`
		PX823  int      `json:"PX823"`
		PX801  string   `json:"PX801"`
		PX796  int      `json:"PX796"`
		PX797  int      `json:"PX797"`
		PX798  int      `json:"PX798"`
		PX799  float64  `json:"PX799"`
		PX800  float64  `json:"PX800"`
		PX603  int      `json:"PX603"`
		PX604  int      `json:"PX604"`
		PX605  []int    `json:"PX605"`
		PX803  int      `json:"PX803"`
		PX606  bool     `json:"PX606"`
		PX769  string   `json:"PX769"`
		PX976  bool     `json:"PX976"`
		PX978  bool     `json:"PX978"`
		PX989  bool     `json:"PX989"`
		PX185  int      `json:"PX185"`
		PX186  int      `json:"PX186"`
		PX850  int      `json:"PX850"`
		PX851  int      `json:"PX851"`
		PX1056 int64    `json:"PX1056"`
		PX1038 string   `json:"PX1038"`
		PX371  bool     `json:"PX371"`
		PX250  string   `json:"PX250"`
		PX708  string   `json:"PX708"`
		PX96   string   `json:"PX96"`
	} `json:"d"`
}

type PX4 struct {
	T string `json:"t"`
	D struct {
		PX31   string        `json:"PX31"`
		PX32   string        `json:"PX32"`
		PX274  string        `json:"PX274"`
		PX275  bool          `json:"PX275"`
		PX441  []interface{} `json:"PX441"`
		PX276  string        `json:"PX276"`
		PX440  []interface{} `json:"PX440"`
		PX210  string        `json:"PX210"`
		PX209  string        `json:"PX209"`
		PX277  string        `json:"PX277"`
		PX281  []string      `json:"PX281"`
		PX282  []interface{} `json:"PX282"`
		PX280  string        `json:"PX280"`
		PX279  string        `json:"PX279"`
		PX278  string        `json:"PX278"`
		PX33   int           `json:"PX33"`
		PX248  []interface{} `json:"PX248"`
		PX249  []string      `json:"PX249"`
		PX57   string        `json:"PX57"`
		PX264  int           `json:"PX264"`
		PX266  string        `json:"PX266"`
		PX265  bool          `json:"PX265"`
		PX364  []string      `json:"PX364"`
		PX286  int           `json:"PX286"`
		PX287  int           `json:"PX287"`
		PX288  bool          `json:"PX288"`
		PX289  bool          `json:"PX289"`
		PX290  bool          `json:"PX290"`
		PX291  bool          `json:"PX291"`
		PX292  string        `json:"PX292"`
		PX293  bool          `json:"PX293"`
		PX229  int           `json:"PX229"`
		PX230  int           `json:"PX230"`
		PX91   int           `json:"PX91"`
		PX92   int           `json:"PX92"`
		PX269  int           `json:"PX269"`
		PX270  int           `json:"PX270"`
		PX93   string        `json:"PX93"`
		PX59   string        `json:"PX59"`
		PX61   string        `json:"PX61"`
		PX313  []string      `json:"PX313"`
		PX63   string        `json:"PX63"`
		PX86   bool          `json:"PX86"`
		PX154  int           `json:"PX154"`
		PX312  string        `json:"PX312"`
		PX311  string        `json:"PX311"`
		PX310  string        `json:"PX310"`
		PX401  int           `json:"PX401"`
		PX409  int           `json:"PX409"`
		PX1131 []interface{} `json:"PX1131"`
		PX850  int           `json:"PX850"`
		PX851  int           `json:"PX851"`
		PX1056 int64         `json:"PX1056"`
		PX1038 string        `json:"PX1038"`
	} `json:"d"`
}

func (this *RECAPPAYLOAD) Instantiate(uuid, vid, site, sitekey, captchatype, px250val string, tsin int){
	this.T = "PX761"
	if tsin == 0{
		tsin = 1000 + rand.Intn(60000)
	}
	this.D.PX70 = tsin
	endpointstring := strings.Replace("sitekey", "TypeError: Cannot read property '0' of null\n    at It (https://client.perimeterx.net/sitekey/main.min.js:2:14201)\n    at re (https://client.perimeterx.net/sitekey/main.min.js:2:26243)\n    at Object.Un [as PX763] (https://client.perimeterx.net/sitekey/main.min.js:2:25361)\n    at u (https://captcha.px-cdn.net/sitekey/captcha.js?a=c&amp;u=uuid&amp;v=vid&amp;m=0:3:75126)\n    at c (https://captcha.px-cdn.net/sitekey/captcha.js?a=c&amp;u=uuid&amp;v=vid&amp;m=0:3:71250)\n    at i (https://captcha.px-cdn.net/sitekey/captcha.js?a=c&amp;u=uuid&amp;v=vid&amp;m=0:3:70778)\n    at Xo (https://captcha.px-cdn.net/sitekey/captcha.js?a=c&amp;u=uuid&amp;v=vid&amp;m=0:3:69831)", sitekey, -1)
	endpointstring = strings.Replace("uuid", endpointstring, uuid, -1)
	endpointstring = strings.Replace("vid", endpointstring, vid, -1)
	this.D.PX34 = endpointstring
	this.D.PX610 = true
	this.D.PX754 = true
	this.D.PX755 = "" //token
	this.D.PX756 = captchatype
	this.D.PX757 = site
	this.D.PX850 = 4 + rand.Intn(3)
	this.D.PX851 = 2000 + rand.Intn(5000)
	this.D.PX1056 = generateTimestamp()
	this.D.PX1038 = uuid
	this.D.PX371 = false
	this.D.PX250 = px250val
	this.D.PX96 = site
}
func (this *RECAPPAYLOAD) Settoken(token string){
	this.D.PX755 = token //token
}

func (this *HCAPPAYLOAD) Instantiate(uuid, site, sitekey string, start int){
	this.T = "PX561"
	this.D.PX70 = 1000 + rand.Intn(2000)
	outappstr := strings.Replace("sitekey", "TypeError: Cannot read property '0' of null\n    at xt (https://client.perimeterx.net/sitekey/main.min.js:2:14047)\\n    at ne (https://client.perimeterx.net/sitekey/main.min.js:2:26105)\\n    at te (https://client.perimeterx.net/sitekey/main.min.js:2:26050)\\n    at c (https://captcha.px-cdn.net/sitekey/captcha.js?a=c&amp;u=uuid&amp;v=&amp;m=0:3:75953)\\n    at u (https://captcha.px-cdn.net/sitekey/captcha.js?a=c&amp;u=uuid&amp;v=&amp;m=0:3:71869)\\n    at a (https://captcha.px-cdn.net/sitekey/captcha.js?a=c&amp;u=uuid&amp;v=&amp;m=0:3:71397)\\n    at Ce (https://captcha.px-cdn.net/sitekey/captcha.js?a=c&amp;u=uuid&amp;v=&amp;m=0:3:70450)", sitekey, -1)
	this.D.PX34 = strings.Replace("uuid", outappstr, uuid, -1)
	this.D.PX1129 = true
	this.D.PX610 = true
	this.D.PX1132 = "9d26298c1ff79d3ed0e8f31a5bb9c218b9f104634f8b3b5138ff67b4c2ccdacc"
	this.D.PX1133 = 2.0 + rand.Float64()*100
	this.D.PX179 = PX179CAPTCHAGENERATE(5)
	this.D.PX582 = PX582GENERATE(5)
	this.D.PX827 = PX827GENERATE(2+rand.Intn(4))
	this.D.PX608 = PX608GENERATE(1+rand.Intn(2))
	this.D.PX74 = 200 + rand.Intn(200)
	this.D.PX75 = 100 + rand.Intn(100)
	this.D.PX76 =  0
	this.D.PX77 =  0
	this.D.PX91 = 1920
	this.D.PX92 = 1080
	this.D.PX789 = "pointerdown"
	this.D.PX790 = 600 +rand.Intn(600) //ts start
	this.D.PX791 = 300 +rand.Intn(500) //x0
	this.D.PX792 = 300 +rand.Intn(500) //y0
	this.D.PX793 = 100 +rand.Float64()*100 //x1
	this.D.PX794 = 20.0 + rand.Float64()*50 //y1
	this.D.PX821 = 4294662870 //mem limit
	this.D.PX822 = 30000000 + rand.Intn(60000000) // mem allocated
	this.D.PX823 = 30000000 + rand.Intn(this.D.PX822-30000000) //mem used
	this.D.PX801 = "pointerup"
	this.D.PX796 = this.D.PX790 + 2500 + rand.Intn(250) //ts end (+2.5 s)
	this.D.PX797 = this.D.PX791 //x0
	this.D.PX798 = this.D.PX792 //y0
	this.D.PX799 = this.D.PX793 //x1
	this.D.PX800 = 20.0 + rand.Float64()*50 //y1
	this.D.PX603 = 1000 +rand.Intn(1000)
	this.D.PX604 = 2000 +rand.Intn(1000)
	this.D.PX605 = []int{1400+rand.Intn(600)}
	this.D.PX803 = 400 + rand.Intn(50)
	this.D.PX606 = true
	this.D.PX769 = "" //token
	this.D.PX976 = false
	this.D.PX978 = false
	this.D.PX989 = false
	this.D.PX185 = 200 + rand.Intn(1500)
	this.D.PX186 = 200 + rand.Intn(2000)
	this.D.PX850 = start
	this.D.PX851 = 2500 + rand.Intn(2000)
	this.D.PX1056 = generateTimestamp()
	this.D.PX1038 = uuid
	this.D.PX371 = false
	this.D.PX250 = "PX560"
	this.D.PX708 = "pxhc"
	this.D.PX96 = site
}
func (this *HCAPPAYLOAD) Settoken(token string){
	this.D.PX769 = token
}

func (this *PX4) Instantiate(uuid, sitekey, ua string, start int){
	this.T = "PX4"
	this.D.PX275 = true
	this.D.PX31  = "126.86972438948578"
	this.D.PX32  = "2dce8c55c6897067fdf0c76ddf6e6d50"
	this.D.PX274 = "014ad7475b97c7ac09e35f7c3293434c"
	this.D.PX210 = "WebKit WebGL"
	this.D.PX209 = "WebKit"
	this.D.PX277 = "WebGL 1.0 (OpenGL ES 2.0 Chromium)"
	this.D.PX441 = []interface{}{}
	this.D.PX276 = "" // ?
	this.D.PX440 = []interface{}{}
	this.D.PX281 = []string{"ANGLE_instanced_arrays", "EXT_blend_minmax", "EXT_color_buffer_half_float", "EXT_disjoint_timer_query", "EXT_float_blend", "EXT_frag_depth", "EXT_shader_texture_lod", "EXT_texture_compression_bptc", "EXT_texture_compression_rgtc", "EXT_texture_filter_anisotropic", "WEBKIT_EXT_texture_filter_anisotropic", "EXT_sRGB", "KHR_parallel_shader_compile", "OES_element_index_uint", "OES_fbo_render_mipmap", "OES_standard_derivatives", "OES_texture_float", "OES_texture_float_linear", "OES_texture_half_float", "OES_texture_half_float_linear", "OES_vertex_array_object", "WEBGL_color_buffer_float", "WEBGL_compressed_texture_s3tc", "WEBKIT_WEBGL_compressed_texture_s3tc", "WEBGL_compressed_texture_s3tc_srgb", "WEBGL_debug_renderer_info", "WEBGL_debug_shaders", "WEBGL_depth_texture", "WEBKIT_WEBGL_depth_texture", "WEBGL_draw_buffers", "WEBGL_lose_context", "WEBKIT_WEBGL_lose_context", "WEBGL_multi_draw"}
	this.D.PX282 = []interface{}{"[1, 1]", "[1, 1024]", 8, "yes", 8, 24, 8, 16, 32, 16384, 1024, 16384, 16, 16384, 30, 16, 16, 4095, "[32767, 32767]", "no_fp", 23, 127, 127, 23, 127, 127, 23, 127, 127, 23, 127, 127, 23, 127, 127, 23, 127, 127, 23, 127, 127, 23, 127, 127, 23, 127, 127, 23, 127, 127, 23, 127, 127, 23, 127, 127}
	this.D.PX280 = PX280LIST[rand.Intn(len(PX280LIST))]
	this.D.PX279 = "Google Inc."
	this.D.PX278 = "WebGL GLSL ES 1.0 (OpenGL ES GLSL ES 1.0 Chromium)"
	this.D.PX33 = 10 + rand.Intn(50)
	this.D.PX248 = []interface{}{}
	this.D.PX249 = []string{"_pxAppId", "_pxJsClientSrc", "_pxFirstPartyEnabled", "_pxVid", "_pxUuid", "_pxHostUrl", "_pxToggleOpenForm", "_pxSubmitForm", "_pxItemSelected", "_pxAction", "_pxMobile", "_uR63h57Zhandler", "_pxInit"}
	this.D.PX57 = strings.Replace("sitekey", "TypeError: Cannot read property '0' of null at Me (https://client.perimeterx.net/sitekey/main.min.js:2:14650) at ke (https://client.perimeterx.net/sitekey/main.min.js:2:14812) at na (https://client.perimeterx.net/sitekey/main.min.js:4:12916) at https://client.perimeterx.net/sitekey/main.min.js:4:14660 at OfflineAudioContext.Vc.n.oncomplete (https://client.perimeterx.net/sitekey/main.min.js:4:6567)", sitekey, -1)
	this.D.PX264 = 33
	this.D.PX266 = "e420a9ce957195c29bde4c763dbc111e"
	this.D.PX265 = true
	this.D.PX364 = []string{"Chrome PDF Plugin::Portable Document Format::application/x-google-chrome-pdf~pdf", "Chrome PDF Viewer::::application/pdf~pdf", "Native Client::::application/x-nacl~::application/x-pnacl~"}
	this.D.PX286 = 1
	this.D.PX287 = 12
	this.D.PX288 = true
	this.D.PX289 = true
	this.D.PX290 = true
	this.D.PX291 = false
	this.D.PX292 = "missing"
	this.D.PX293 = true
	this.D.PX229 = 24
	this.D.PX230 = 24
	this.D.PX91 = 3840
	this.D.PX92 = 1080
	this.D.PX269 = 3778
	this.D.PX270 = 1080
	this.D.PX93 = "3840X1080"
	this.D.PX59 = ua
	this.D.PX61 = "en-US"
	this.D.PX313 = []string{"en-US"}
	this.D.PX63 =  "Win32"
	this.D.PX86 = true
	this.D.PX154 = 240
	this.D.PX312 = "d41d8cd98f00b204e9800998ecf8427e"
	this.D.PX311 = "fd7149bbfb316699ef918fa7bb7510a8"
	this.D.PX310 = "fd7149bbfb316699ef918fa7bb7510a8"
	this.D.PX401 = 13969
	this.D.PX409 = 2
	this.D.PX1131 = []interface{}{}
	this.D.PX850 = start
	this.D.PX851 = 100 + rand.Intn(14000)
	this.D.PX1056 = generateTimestamp()
	this.D.PX1038 = uuid
}

func PX179CAPTCHAGENERATE(length int) []string{
	mapout := []string{}
	for i := 1; i <= length; i++{
		var mapkey string
		for j := 0; j < 2 + rand.Intn(5); j++{
			mapkey += PX176CAPTCHALIST[rand.Intn(len(PX176CAPTCHALIST))]+">"
		}
		mapout = append(mapout, mapkey[:len(mapkey)-1])
	}
	found := false
	for _, a := range mapout {
		if a == "#px-captcha" {
			found = true
		}
	}
	if !found{
		mapout = append(mapout, "#px-captcha")
	}
	return mapout
}

func RECAPGENERATE(uuid, vid, site, sitekey, token string) RECAPPAYLOAD {
	recapobj := RECAPPAYLOAD{}
	recapobj.Instantiate(uuid, vid, site, sitekey, "reCaptcha", "PX557", 0)
	recapobj.Settoken(token)
	recapobj.D.PX788 = "c"
	return recapobj
}
func LOWSECHCAPGENERATE(uuid, vid, site, sitekey string) []interface{}{
	outarr := []interface{}{}
	hcapobj := HCAPPAYLOAD{}
	hcapobj.Instantiate(uuid, site, sitekey, 5)
	outarr = append(outarr, hcapobj)
	recapobj := RECAPPAYLOAD{}
	recapobj.Instantiate(uuid, vid, site, sitekey, "pxCaptcha", "PX560", hcapobj.D.PX70)
	recapobj.Settoken(uuid)
	recapobj.D.PX708 = "pxhc"
	recapobj.D.PX1129 = true
	recapobj.D.PX754 = false
	recapobj.D.PX850 = hcapobj.D.PX850 + 1
	recapobj.D.PX851 = hcapobj.D.PX851
	outarr = append(outarr, recapobj)
	return outarr
}
func HIGHSECHCAPGENERATE(uuid, vid, site, sitekey, ua string) []interface{}{
	outarr := []interface{}{}
	mouseoverobj := MouseOverEvent{}
	mouseoverobj.Instantiate(site, uuid, sitekey, 2)
	outarr = append(outarr, mouseoverobj)
	px4obj := PX4{}
	px4obj.Instantiate(uuid, sitekey, ua, 3)
	outarr = append(outarr, px4obj)
	mousemoveobj := MouseMoveEvent{}
	mousemoveobj.Instantiate(site, uuid, 4)
	outarr = append(outarr, mousemoveobj)
	hcapobj := HCAPPAYLOAD{}
	hcapobj.Instantiate(uuid, site, sitekey, 5)
	outarr = append(outarr, hcapobj)
	recapobj := RECAPPAYLOAD{}
	recapobj.Instantiate(uuid, vid, site, sitekey, "pxCaptcha", "PX560", hcapobj.D.PX70)
	recapobj.Settoken(uuid)
	recapobj.D.PX708 = "pxhc"
	recapobj.D.PX1129 = true
	recapobj.D.PX754 = false
	recapobj.D.PX850 = hcapobj.D.PX850 + 1
	recapobj.D.PX851 = hcapobj.D.PX851
	outarr = append(outarr, recapobj)
	return outarr
}