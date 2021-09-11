package generator

import (
	"github.com/ProjectAthenaa/perimeterx-service/siteconstants"
	"math/rand"
	"strings"
)

var (
	PX280LIST = []string{"AMD Radeon R9 M370X OpenGL Engine", "ANGLE (NVIDIA GeForce GTX 1660 SUPER Direct3D11 vs_5_0 ps_5_0)", "ANGLE (NVIDIA GeForce RTX 2080 Direct3D11 vs_5_0 ps_5_0)", "Intel(R) Iris(TM) Plus Graphics OpenGL Engine", "ANGLE (NVIDIA GeForce GTX 1060 3GB Direct3D11 vs_5_0 ps_5_0)", "ANGLE (NVIDIA GeForce GTX 1070 Direct3D11 vs_5_0 ps_5_0)", "ANGLE (NVIDIA GeForce GTX 1080 Direct3D11 vs_5_0 ps_5_0)", "ANGLE (AMD Radeon RX 5600 XT Direct3D11 vs_5_0 ps_5_0)", "ANGLE (Radeon RX 570 Series Direct3D11 vs_5_0 ps_5_0)", "ANGLE (Radeon RX 580 Series Direct3D11 vs_5_0 ps_5_0)", "Google SwiftShader"}
)

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
		PX1056 int           `json:"PX1056"`
		PX1038 string        `json:"PX1038"`
	} `json:"d"`
}

func InstantiatePX4(uuid, sitekey string) PX4{
	var payload PX4
	payload.T = "PX4"
	payload.D.PX275 = true
	payload.D.PX31  = "126.86972438948578"
	payload.D.PX32  = "2dce8c55c6897067fdf0c76ddf6e6d50"
	payload.D.PX274 = "014ad7475b97c7ac09e35f7c3293434c"
	payload.D.PX210 = "WebKit WebGL"
	payload.D.PX209 = "WebKit"
	payload.D.PX277 = "WebGL 1.0 (OpenGL ES 2.0 Chromium)"
	payload.D.PX441 = []interface{}{}
	payload.D.PX276 = "" // ?
	payload.D.PX440 = []interface{}{}
	payload.D.PX281 = []string{"ANGLE_instanced_arrays", "EXT_blend_minmax", "EXT_color_buffer_half_float", "EXT_disjoint_timer_query", "EXT_float_blend", "EXT_frag_depth", "EXT_shader_texture_lod", "EXT_texture_compression_bptc", "EXT_texture_compression_rgtc", "EXT_texture_filter_anisotropic", "WEBKIT_EXT_texture_filter_anisotropic", "EXT_sRGB", "KHR_parallel_shader_compile", "OES_element_index_uint", "OES_fbo_render_mipmap", "OES_standard_derivatives", "OES_texture_float", "OES_texture_float_linear", "OES_texture_half_float", "OES_texture_half_float_linear", "OES_vertex_array_object", "WEBGL_color_buffer_float", "WEBGL_compressed_texture_s3tc", "WEBKIT_WEBGL_compressed_texture_s3tc", "WEBGL_compressed_texture_s3tc_srgb", "WEBGL_debug_renderer_info", "WEBGL_debug_shaders", "WEBGL_depth_texture", "WEBKIT_WEBGL_depth_texture", "WEBGL_draw_buffers", "WEBGL_lose_context", "WEBKIT_WEBGL_lose_context", "WEBGL_multi_draw"}
	payload.D.PX282 = []interface{}{"[1, 1]", "[1, 1024]", 8, "yes", 8, 24, 8, 16, 32, 16384, 1024, 16384, 16, 16384, 30, 16, 16, 4095, "[32767, 32767]", "no_fp", 23, 127, 127, 23, 127, 127, 23, 127, 127, 23, 127, 127, 23, 127, 127, 23, 127, 127, 23, 127, 127, 23, 127, 127, 23, 127, 127, 23, 127, 127, 23, 127, 127, 23, 127, 127}
	payload.D.PX280 = PX280LIST[rand.Intn(len(PX280LIST))]
	payload.D.PX279 = "Google Inc."
	payload.D.PX278 = "WebGL GLSL ES 1.0 (OpenGL ES GLSL ES 1.0 Chromium)"
	payload.D.PX33 = 10 + rand.Intn(50)
	payload.D.PX248 = []interface{}{}
	payload.D.PX249 = []string{"_pxAppId", "_pxJsClientSrc", "_pxFirstPartyEnabled", "_pxVid", "_pxUuid", "_pxHostUrl", "_pxToggleOpenForm", "_pxSubmitForm", "_pxItemSelected", "_pxAction", "_pxMobile", "_uR63h57Zhandler", "_pxInit"}
	payload.D.PX57 = strings.Replace("sitekey", "TypeError: Cannot read property '0' of null at Me (https://client.perimeterx.net/sitekey/main.min.js:2:14650) at ke (https://client.perimeterx.net/sitekey/main.min.js:2:14812) at na (https://client.perimeterx.net/sitekey/main.min.js:4:12916) at https://client.perimeterx.net/sitekey/main.min.js:4:14660 at OfflineAudioContext.Vc.n.oncomplete (https://client.perimeterx.net/sitekey/main.min.js:4:6567)", sitekey, -1)
	payload.D.PX264 = 33
	payload.D.PX266 = "e420a9ce957195c29bde4c763dbc111e"
	payload.D.PX265 = true
	payload.D.PX364 = []string{"Chrome PDF Plugin::Portable Document Format::application/x-google-chrome-pdf~pdf", "Chrome PDF Viewer::::application/pdf~pdf", "Native Client::::application/x-nacl~::application/x-pnacl~"}
	payload.D.PX286 = 1
	payload.D.PX287 = 12
	payload.D.PX288 = true
	payload.D.PX289 = true
	payload.D.PX290 = true
	payload.D.PX291 = false
	payload.D.PX292 = "missing"
	payload.D.PX293 = true
	payload.D.PX229 = 24
	payload.D.PX230 = 24
	payload.D.PX91 = 3840
	payload.D.PX92 = 1080
	payload.D.PX269 = 3778
	payload.D.PX270 = 1080
	payload.D.PX93 = "3840X1080"
	payload.D.PX59 = siteconstants.UA
	payload.D.PX61 = "en-US"
	payload.D.PX313 = []string{"en-US"}
	payload.D.PX63 =  "Win32"
	payload.D.PX86 = true
	payload.D.PX154 = 240
	payload.D.PX312 = "d41d8cd98f00b204e9800998ecf8427e"
	payload.D.PX311 = "fd7149bbfb316699ef918fa7bb7510a8"
	payload.D.PX310 = "fd7149bbfb316699ef918fa7bb7510a8"
	payload.D.PX401 = 13969
	payload.D.PX409 = 2
	payload.D.PX1131 = []interface{}{}
	payload.D.PX850 = 3
	payload.D.PX851 = 100 + rand.Intn(14000)
	payload.D.PX1056 = GenerateTimestamp()
	payload.D.PX1038 = uuid
	return payload
}