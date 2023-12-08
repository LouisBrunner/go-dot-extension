// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Color struct {
  obj gdc.ObjectPtr
}

func (me *Color) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Color) BaseClass() string {
  return "Color"
}

// TODO: needed?
// const (
// // )

var (
  ColorAliceBlue = "Color(0.941176, 0.972549, 1, 1)" // TODO: construct correctly
  ColorAntiqueWhite = "Color(0.980392, 0.921569, 0.843137, 1)" // TODO: construct correctly
  ColorAqua = "Color(0, 1, 1, 1)" // TODO: construct correctly
  ColorAquamarine = "Color(0.498039, 1, 0.831373, 1)" // TODO: construct correctly
  ColorAzure = "Color(0.941176, 1, 1, 1)" // TODO: construct correctly
  ColorBeige = "Color(0.960784, 0.960784, 0.862745, 1)" // TODO: construct correctly
  ColorBisque = "Color(1, 0.894118, 0.768627, 1)" // TODO: construct correctly
  ColorBlack = "Color(0, 0, 0, 1)" // TODO: construct correctly
  ColorBlanchedAlmond = "Color(1, 0.921569, 0.803922, 1)" // TODO: construct correctly
  ColorBlue = "Color(0, 0, 1, 1)" // TODO: construct correctly
  ColorBlueViolet = "Color(0.541176, 0.168627, 0.886275, 1)" // TODO: construct correctly
  ColorBrown = "Color(0.647059, 0.164706, 0.164706, 1)" // TODO: construct correctly
  ColorBurlywood = "Color(0.870588, 0.721569, 0.529412, 1)" // TODO: construct correctly
  ColorCadetBlue = "Color(0.372549, 0.619608, 0.627451, 1)" // TODO: construct correctly
  ColorChartreuse = "Color(0.498039, 1, 0, 1)" // TODO: construct correctly
  ColorChocolate = "Color(0.823529, 0.411765, 0.117647, 1)" // TODO: construct correctly
  ColorCoral = "Color(1, 0.498039, 0.313726, 1)" // TODO: construct correctly
  ColorCornflowerBlue = "Color(0.392157, 0.584314, 0.929412, 1)" // TODO: construct correctly
  ColorCornsilk = "Color(1, 0.972549, 0.862745, 1)" // TODO: construct correctly
  ColorCrimson = "Color(0.862745, 0.0784314, 0.235294, 1)" // TODO: construct correctly
  ColorCyan = "Color(0, 1, 1, 1)" // TODO: construct correctly
  ColorDarkBlue = "Color(0, 0, 0.545098, 1)" // TODO: construct correctly
  ColorDarkCyan = "Color(0, 0.545098, 0.545098, 1)" // TODO: construct correctly
  ColorDarkGoldenrod = "Color(0.721569, 0.52549, 0.0431373, 1)" // TODO: construct correctly
  ColorDarkGray = "Color(0.662745, 0.662745, 0.662745, 1)" // TODO: construct correctly
  ColorDarkGreen = "Color(0, 0.392157, 0, 1)" // TODO: construct correctly
  ColorDarkKhaki = "Color(0.741176, 0.717647, 0.419608, 1)" // TODO: construct correctly
  ColorDarkMagenta = "Color(0.545098, 0, 0.545098, 1)" // TODO: construct correctly
  ColorDarkOliveGreen = "Color(0.333333, 0.419608, 0.184314, 1)" // TODO: construct correctly
  ColorDarkOrange = "Color(1, 0.54902, 0, 1)" // TODO: construct correctly
  ColorDarkOrchid = "Color(0.6, 0.196078, 0.8, 1)" // TODO: construct correctly
  ColorDarkRed = "Color(0.545098, 0, 0, 1)" // TODO: construct correctly
  ColorDarkSalmon = "Color(0.913725, 0.588235, 0.478431, 1)" // TODO: construct correctly
  ColorDarkSeaGreen = "Color(0.560784, 0.737255, 0.560784, 1)" // TODO: construct correctly
  ColorDarkSlateBlue = "Color(0.282353, 0.239216, 0.545098, 1)" // TODO: construct correctly
  ColorDarkSlateGray = "Color(0.184314, 0.309804, 0.309804, 1)" // TODO: construct correctly
  ColorDarkTurquoise = "Color(0, 0.807843, 0.819608, 1)" // TODO: construct correctly
  ColorDarkViolet = "Color(0.580392, 0, 0.827451, 1)" // TODO: construct correctly
  ColorDeepPink = "Color(1, 0.0784314, 0.576471, 1)" // TODO: construct correctly
  ColorDeepSkyBlue = "Color(0, 0.74902, 1, 1)" // TODO: construct correctly
  ColorDimGray = "Color(0.411765, 0.411765, 0.411765, 1)" // TODO: construct correctly
  ColorDodgerBlue = "Color(0.117647, 0.564706, 1, 1)" // TODO: construct correctly
  ColorFirebrick = "Color(0.698039, 0.133333, 0.133333, 1)" // TODO: construct correctly
  ColorFloralWhite = "Color(1, 0.980392, 0.941176, 1)" // TODO: construct correctly
  ColorForestGreen = "Color(0.133333, 0.545098, 0.133333, 1)" // TODO: construct correctly
  ColorFuchsia = "Color(1, 0, 1, 1)" // TODO: construct correctly
  ColorGainsboro = "Color(0.862745, 0.862745, 0.862745, 1)" // TODO: construct correctly
  ColorGhostWhite = "Color(0.972549, 0.972549, 1, 1)" // TODO: construct correctly
  ColorGold = "Color(1, 0.843137, 0, 1)" // TODO: construct correctly
  ColorGoldenrod = "Color(0.854902, 0.647059, 0.12549, 1)" // TODO: construct correctly
  ColorGray = "Color(0.745098, 0.745098, 0.745098, 1)" // TODO: construct correctly
  ColorGreen = "Color(0, 1, 0, 1)" // TODO: construct correctly
  ColorGreenYellow = "Color(0.678431, 1, 0.184314, 1)" // TODO: construct correctly
  ColorHoneydew = "Color(0.941176, 1, 0.941176, 1)" // TODO: construct correctly
  ColorHotPink = "Color(1, 0.411765, 0.705882, 1)" // TODO: construct correctly
  ColorIndianRed = "Color(0.803922, 0.360784, 0.360784, 1)" // TODO: construct correctly
  ColorIndigo = "Color(0.294118, 0, 0.509804, 1)" // TODO: construct correctly
  ColorIvory = "Color(1, 1, 0.941176, 1)" // TODO: construct correctly
  ColorKhaki = "Color(0.941176, 0.901961, 0.54902, 1)" // TODO: construct correctly
  ColorLavender = "Color(0.901961, 0.901961, 0.980392, 1)" // TODO: construct correctly
  ColorLavenderBlush = "Color(1, 0.941176, 0.960784, 1)" // TODO: construct correctly
  ColorLawnGreen = "Color(0.486275, 0.988235, 0, 1)" // TODO: construct correctly
  ColorLemonChiffon = "Color(1, 0.980392, 0.803922, 1)" // TODO: construct correctly
  ColorLightBlue = "Color(0.678431, 0.847059, 0.901961, 1)" // TODO: construct correctly
  ColorLightCoral = "Color(0.941176, 0.501961, 0.501961, 1)" // TODO: construct correctly
  ColorLightCyan = "Color(0.878431, 1, 1, 1)" // TODO: construct correctly
  ColorLightGoldenrod = "Color(0.980392, 0.980392, 0.823529, 1)" // TODO: construct correctly
  ColorLightGray = "Color(0.827451, 0.827451, 0.827451, 1)" // TODO: construct correctly
  ColorLightGreen = "Color(0.564706, 0.933333, 0.564706, 1)" // TODO: construct correctly
  ColorLightPink = "Color(1, 0.713726, 0.756863, 1)" // TODO: construct correctly
  ColorLightSalmon = "Color(1, 0.627451, 0.478431, 1)" // TODO: construct correctly
  ColorLightSeaGreen = "Color(0.12549, 0.698039, 0.666667, 1)" // TODO: construct correctly
  ColorLightSkyBlue = "Color(0.529412, 0.807843, 0.980392, 1)" // TODO: construct correctly
  ColorLightSlateGray = "Color(0.466667, 0.533333, 0.6, 1)" // TODO: construct correctly
  ColorLightSteelBlue = "Color(0.690196, 0.768627, 0.870588, 1)" // TODO: construct correctly
  ColorLightYellow = "Color(1, 1, 0.878431, 1)" // TODO: construct correctly
  ColorLime = "Color(0, 1, 0, 1)" // TODO: construct correctly
  ColorLimeGreen = "Color(0.196078, 0.803922, 0.196078, 1)" // TODO: construct correctly
  ColorLinen = "Color(0.980392, 0.941176, 0.901961, 1)" // TODO: construct correctly
  ColorMagenta = "Color(1, 0, 1, 1)" // TODO: construct correctly
  ColorMaroon = "Color(0.690196, 0.188235, 0.376471, 1)" // TODO: construct correctly
  ColorMediumAquamarine = "Color(0.4, 0.803922, 0.666667, 1)" // TODO: construct correctly
  ColorMediumBlue = "Color(0, 0, 0.803922, 1)" // TODO: construct correctly
  ColorMediumOrchid = "Color(0.729412, 0.333333, 0.827451, 1)" // TODO: construct correctly
  ColorMediumPurple = "Color(0.576471, 0.439216, 0.858824, 1)" // TODO: construct correctly
  ColorMediumSeaGreen = "Color(0.235294, 0.701961, 0.443137, 1)" // TODO: construct correctly
  ColorMediumSlateBlue = "Color(0.482353, 0.407843, 0.933333, 1)" // TODO: construct correctly
  ColorMediumSpringGreen = "Color(0, 0.980392, 0.603922, 1)" // TODO: construct correctly
  ColorMediumTurquoise = "Color(0.282353, 0.819608, 0.8, 1)" // TODO: construct correctly
  ColorMediumVioletRed = "Color(0.780392, 0.0823529, 0.521569, 1)" // TODO: construct correctly
  ColorMidnightBlue = "Color(0.0980392, 0.0980392, 0.439216, 1)" // TODO: construct correctly
  ColorMintCream = "Color(0.960784, 1, 0.980392, 1)" // TODO: construct correctly
  ColorMistyRose = "Color(1, 0.894118, 0.882353, 1)" // TODO: construct correctly
  ColorMoccasin = "Color(1, 0.894118, 0.709804, 1)" // TODO: construct correctly
  ColorNavajoWhite = "Color(1, 0.870588, 0.678431, 1)" // TODO: construct correctly
  ColorNavyBlue = "Color(0, 0, 0.501961, 1)" // TODO: construct correctly
  ColorOldLace = "Color(0.992157, 0.960784, 0.901961, 1)" // TODO: construct correctly
  ColorOlive = "Color(0.501961, 0.501961, 0, 1)" // TODO: construct correctly
  ColorOliveDrab = "Color(0.419608, 0.556863, 0.137255, 1)" // TODO: construct correctly
  ColorOrange = "Color(1, 0.647059, 0, 1)" // TODO: construct correctly
  ColorOrangeRed = "Color(1, 0.270588, 0, 1)" // TODO: construct correctly
  ColorOrchid = "Color(0.854902, 0.439216, 0.839216, 1)" // TODO: construct correctly
  ColorPaleGoldenrod = "Color(0.933333, 0.909804, 0.666667, 1)" // TODO: construct correctly
  ColorPaleGreen = "Color(0.596078, 0.984314, 0.596078, 1)" // TODO: construct correctly
  ColorPaleTurquoise = "Color(0.686275, 0.933333, 0.933333, 1)" // TODO: construct correctly
  ColorPaleVioletRed = "Color(0.858824, 0.439216, 0.576471, 1)" // TODO: construct correctly
  ColorPapayaWhip = "Color(1, 0.937255, 0.835294, 1)" // TODO: construct correctly
  ColorPeachPuff = "Color(1, 0.854902, 0.72549, 1)" // TODO: construct correctly
  ColorPeru = "Color(0.803922, 0.521569, 0.247059, 1)" // TODO: construct correctly
  ColorPink = "Color(1, 0.752941, 0.796078, 1)" // TODO: construct correctly
  ColorPlum = "Color(0.866667, 0.627451, 0.866667, 1)" // TODO: construct correctly
  ColorPowderBlue = "Color(0.690196, 0.878431, 0.901961, 1)" // TODO: construct correctly
  ColorPurple = "Color(0.627451, 0.12549, 0.941176, 1)" // TODO: construct correctly
  ColorRebeccaPurple = "Color(0.4, 0.2, 0.6, 1)" // TODO: construct correctly
  ColorRed = "Color(1, 0, 0, 1)" // TODO: construct correctly
  ColorRosyBrown = "Color(0.737255, 0.560784, 0.560784, 1)" // TODO: construct correctly
  ColorRoyalBlue = "Color(0.254902, 0.411765, 0.882353, 1)" // TODO: construct correctly
  ColorSaddleBrown = "Color(0.545098, 0.270588, 0.0745098, 1)" // TODO: construct correctly
  ColorSalmon = "Color(0.980392, 0.501961, 0.447059, 1)" // TODO: construct correctly
  ColorSandyBrown = "Color(0.956863, 0.643137, 0.376471, 1)" // TODO: construct correctly
  ColorSeaGreen = "Color(0.180392, 0.545098, 0.341176, 1)" // TODO: construct correctly
  ColorSeashell = "Color(1, 0.960784, 0.933333, 1)" // TODO: construct correctly
  ColorSienna = "Color(0.627451, 0.321569, 0.176471, 1)" // TODO: construct correctly
  ColorSilver = "Color(0.752941, 0.752941, 0.752941, 1)" // TODO: construct correctly
  ColorSkyBlue = "Color(0.529412, 0.807843, 0.921569, 1)" // TODO: construct correctly
  ColorSlateBlue = "Color(0.415686, 0.352941, 0.803922, 1)" // TODO: construct correctly
  ColorSlateGray = "Color(0.439216, 0.501961, 0.564706, 1)" // TODO: construct correctly
  ColorSnow = "Color(1, 0.980392, 0.980392, 1)" // TODO: construct correctly
  ColorSpringGreen = "Color(0, 1, 0.498039, 1)" // TODO: construct correctly
  ColorSteelBlue = "Color(0.27451, 0.509804, 0.705882, 1)" // TODO: construct correctly
  ColorTan = "Color(0.823529, 0.705882, 0.54902, 1)" // TODO: construct correctly
  ColorTeal = "Color(0, 0.501961, 0.501961, 1)" // TODO: construct correctly
  ColorThistle = "Color(0.847059, 0.74902, 0.847059, 1)" // TODO: construct correctly
  ColorTomato = "Color(1, 0.388235, 0.278431, 1)" // TODO: construct correctly
  ColorTransparent = "Color(1, 1, 1, 0)" // TODO: construct correctly
  ColorTurquoise = "Color(0.25098, 0.878431, 0.815686, 1)" // TODO: construct correctly
  ColorViolet = "Color(0.933333, 0.509804, 0.933333, 1)" // TODO: construct correctly
  ColorWebGray = "Color(0.501961, 0.501961, 0.501961, 1)" // TODO: construct correctly
  ColorWebGreen = "Color(0, 0.501961, 0, 1)" // TODO: construct correctly
  ColorWebMaroon = "Color(0.501961, 0, 0, 1)" // TODO: construct correctly
  ColorWebPurple = "Color(0.501961, 0, 0.501961, 1)" // TODO: construct correctly
  ColorWheat = "Color(0.960784, 0.870588, 0.701961, 1)" // TODO: construct correctly
  ColorWhite = "Color(1, 1, 1, 1)" // TODO: construct correctly
  ColorWhiteSmoke = "Color(0.960784, 0.960784, 0.960784, 1)" // TODO: construct correctly
  ColorYellow = "Color(1, 1, 0, 1)" // TODO: construct correctly
  ColorYellowGreen = "Color(0.603922, 0.803922, 0.196078, 1)" // TODO: construct correctly
)

func  (me *Color) ToArgb32() { // TODO: return value
  // TODO: implement
}

func  (me *Color) ToAbgr32() { // TODO: return value
  // TODO: implement
}

func  (me *Color) ToRgba32() { // TODO: return value
  // TODO: implement
}

func  (me *Color) ToArgb64() { // TODO: return value
  // TODO: implement
}

func  (me *Color) ToAbgr64() { // TODO: return value
  // TODO: implement
}

func  (me *Color) ToRgba64() { // TODO: return value
  // TODO: implement
}

func  (me *Color) ToHtml(with_alpha bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Color) Clamp(min Color, max Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *Color) Inverted() { // TODO: return value
  // TODO: implement
}

func  (me *Color) Lerp(to Color, weight float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Color) Lightened(amount float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Color) Darkened(amount float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Color) Blend(over Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *Color) GetLuminance() { // TODO: return value
  // TODO: implement
}

func  (me *Color) SrgbToLinear() { // TODO: return value
  // TODO: implement
}

func  (me *Color) LinearToSrgb() { // TODO: return value
  // TODO: implement
}

func  (me *Color) IsEqualApprox(to Color, ) { // TODO: return value
  // TODO: implement
}

func  ColorHex(hex int, ) { // TODO: return value
  // TODO: implement
}

func  ColorHex64(hex int, ) { // TODO: return value
  // TODO: implement
}

func  ColorHtml(rgba String, ) { // TODO: return value
  // TODO: implement
}

func  ColorHtmlIsValid(color String, ) { // TODO: return value
  // TODO: implement
}

func  ColorFromString(str String, default_ Color, ) { // TODO: return value
  // TODO: implement
}

func  ColorFromHsv(h float32, s float32, v float32, alpha float32, ) { // TODO: return value
  // TODO: implement
}

func  ColorFromOkHsl(h float32, s float32, l float32, alpha float32, ) { // TODO: return value
  // TODO: implement
}

func  ColorFromRgbe9995(rgbe int, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
