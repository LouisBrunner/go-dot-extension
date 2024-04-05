// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "fmt"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused imports
var _ = fmt.Sprintf("")

type ptrsForColorList struct {
  ctrFn gdc.PtrConstructor
  ctrFromColorFn gdc.PtrConstructor
  ctrFromColorFloat32Fn gdc.PtrConstructor
  ctrFromFloat32Float32Float32Fn gdc.PtrConstructor
  ctrFromFloat32Float32Float32Float32Fn gdc.PtrConstructor
  ctrFromStringFn gdc.PtrConstructor
  ctrFromStringFloat32Fn gdc.PtrConstructor
  methodToArgb32Fn gdc.PtrBuiltInMethod
  methodToAbgr32Fn gdc.PtrBuiltInMethod
  methodToRgba32Fn gdc.PtrBuiltInMethod
  methodToArgb64Fn gdc.PtrBuiltInMethod
  methodToAbgr64Fn gdc.PtrBuiltInMethod
  methodToRgba64Fn gdc.PtrBuiltInMethod
  methodToHtmlFn gdc.PtrBuiltInMethod
  methodClampFn gdc.PtrBuiltInMethod
  methodInvertedFn gdc.PtrBuiltInMethod
  methodLerpFn gdc.PtrBuiltInMethod
  methodLightenedFn gdc.PtrBuiltInMethod
  methodDarkenedFn gdc.PtrBuiltInMethod
  methodBlendFn gdc.PtrBuiltInMethod
  methodGetLuminanceFn gdc.PtrBuiltInMethod
  methodSrgbToLinearFn gdc.PtrBuiltInMethod
  methodLinearToSrgbFn gdc.PtrBuiltInMethod
  methodIsEqualApproxFn gdc.PtrBuiltInMethod
  methodHexFn gdc.PtrBuiltInMethod
  methodHex64Fn gdc.PtrBuiltInMethod
  methodHtmlFn gdc.PtrBuiltInMethod
  methodHtmlIsValidFn gdc.PtrBuiltInMethod
  methodFromStringFn gdc.PtrBuiltInMethod
  methodFromHsvFn gdc.PtrBuiltInMethod
  methodFromOkHslFn gdc.PtrBuiltInMethod
  methodFromRgbe9995Fn gdc.PtrBuiltInMethod
  operatorNegateFn gdc.PtrOperatorEvaluator
  operatorPositiveFn gdc.PtrOperatorEvaluator
  operatorNotFn gdc.PtrOperatorEvaluator
  operatorMultiplyIntFn gdc.PtrOperatorEvaluator
  operatorDivideIntFn gdc.PtrOperatorEvaluator
  operatorMultiplyFloat32Fn gdc.PtrOperatorEvaluator
  operatorDivideFloat32Fn gdc.PtrOperatorEvaluator
  operatorEqualColorFn gdc.PtrOperatorEvaluator
  operatorNotEqualColorFn gdc.PtrOperatorEvaluator
  operatorAddColorFn gdc.PtrOperatorEvaluator
  operatorSubtractColorFn gdc.PtrOperatorEvaluator
  operatorMultiplyColorFn gdc.PtrOperatorEvaluator
  operatorDivideColorFn gdc.PtrOperatorEvaluator
  operatorInDictionaryFn gdc.PtrOperatorEvaluator
  operatorInArrayFn gdc.PtrOperatorEvaluator
  operatorInPackedColorArrayFn gdc.PtrOperatorEvaluator
  memberrGetterFn gdc.PtrGetter
  memberrSetterFn gdc.PtrSetter
  membergGetterFn gdc.PtrGetter
  membergSetterFn gdc.PtrSetter
  memberbGetterFn gdc.PtrGetter
  memberbSetterFn gdc.PtrSetter
  memberaGetterFn gdc.PtrGetter
  memberaSetterFn gdc.PtrSetter
  memberr8GetterFn gdc.PtrGetter
  memberr8SetterFn gdc.PtrSetter
  memberg8GetterFn gdc.PtrGetter
  memberg8SetterFn gdc.PtrSetter
  memberb8GetterFn gdc.PtrGetter
  memberb8SetterFn gdc.PtrSetter
  membera8GetterFn gdc.PtrGetter
  membera8SetterFn gdc.PtrSetter
  memberhGetterFn gdc.PtrGetter
  memberhSetterFn gdc.PtrSetter
  membersGetterFn gdc.PtrGetter
  membersSetterFn gdc.PtrSetter
  membervGetterFn gdc.PtrGetter
  membervSetterFn gdc.PtrSetter
  toVariantFn gdc.TypeFromVariantConstructorFunc
  fromVariantFn gdc.VariantFromTypeConstructorFunc
}

var ptrsForColor ptrsForColorList

func initColorPtrs(iface gdc.Interface) {
  ptrsForColor.ctrFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeColor, 0))
  ptrsForColor.ctrFromColorFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeColor, 1))
  ptrsForColor.ctrFromColorFloat32Fn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeColor, 2))
  ptrsForColor.ctrFromFloat32Float32Float32Fn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeColor, 3))
  ptrsForColor.ctrFromFloat32Float32Float32Float32Fn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeColor, 4))
  ptrsForColor.ctrFromStringFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeColor, 5))
  ptrsForColor.ctrFromStringFloat32Fn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeColor, 6))
  {
    methodName := StringNameFromStr("to_argb32")
    defer methodName.Destroy()
    ptrsForColor.methodToArgb32Fn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeColor, methodName.AsCPtr(), 3173160232))
  }
  {
    methodName := StringNameFromStr("to_abgr32")
    defer methodName.Destroy()
    ptrsForColor.methodToAbgr32Fn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeColor, methodName.AsCPtr(), 3173160232))
  }
  {
    methodName := StringNameFromStr("to_rgba32")
    defer methodName.Destroy()
    ptrsForColor.methodToRgba32Fn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeColor, methodName.AsCPtr(), 3173160232))
  }
  {
    methodName := StringNameFromStr("to_argb64")
    defer methodName.Destroy()
    ptrsForColor.methodToArgb64Fn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeColor, methodName.AsCPtr(), 3173160232))
  }
  {
    methodName := StringNameFromStr("to_abgr64")
    defer methodName.Destroy()
    ptrsForColor.methodToAbgr64Fn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeColor, methodName.AsCPtr(), 3173160232))
  }
  {
    methodName := StringNameFromStr("to_rgba64")
    defer methodName.Destroy()
    ptrsForColor.methodToRgba64Fn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeColor, methodName.AsCPtr(), 3173160232))
  }
  {
    methodName := StringNameFromStr("to_html")
    defer methodName.Destroy()
    ptrsForColor.methodToHtmlFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeColor, methodName.AsCPtr(), 3429816538))
  }
  {
    methodName := StringNameFromStr("clamp")
    defer methodName.Destroy()
    ptrsForColor.methodClampFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeColor, methodName.AsCPtr(), 105651410))
  }
  {
    methodName := StringNameFromStr("inverted")
    defer methodName.Destroy()
    ptrsForColor.methodInvertedFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeColor, methodName.AsCPtr(), 3334027602))
  }
  {
    methodName := StringNameFromStr("lerp")
    defer methodName.Destroy()
    ptrsForColor.methodLerpFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeColor, methodName.AsCPtr(), 402949615))
  }
  {
    methodName := StringNameFromStr("lightened")
    defer methodName.Destroy()
    ptrsForColor.methodLightenedFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeColor, methodName.AsCPtr(), 1466039168))
  }
  {
    methodName := StringNameFromStr("darkened")
    defer methodName.Destroy()
    ptrsForColor.methodDarkenedFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeColor, methodName.AsCPtr(), 1466039168))
  }
  {
    methodName := StringNameFromStr("blend")
    defer methodName.Destroy()
    ptrsForColor.methodBlendFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeColor, methodName.AsCPtr(), 3803690977))
  }
  {
    methodName := StringNameFromStr("get_luminance")
    defer methodName.Destroy()
    ptrsForColor.methodGetLuminanceFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeColor, methodName.AsCPtr(), 466405837))
  }
  {
    methodName := StringNameFromStr("srgb_to_linear")
    defer methodName.Destroy()
    ptrsForColor.methodSrgbToLinearFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeColor, methodName.AsCPtr(), 3334027602))
  }
  {
    methodName := StringNameFromStr("linear_to_srgb")
    defer methodName.Destroy()
    ptrsForColor.methodLinearToSrgbFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeColor, methodName.AsCPtr(), 3334027602))
  }
  {
    methodName := StringNameFromStr("is_equal_approx")
    defer methodName.Destroy()
    ptrsForColor.methodIsEqualApproxFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeColor, methodName.AsCPtr(), 3167426256))
  }
  {
    methodName := StringNameFromStr("hex")
    defer methodName.Destroy()
    ptrsForColor.methodHexFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeColor, methodName.AsCPtr(), 351421375))
  }
  {
    methodName := StringNameFromStr("hex64")
    defer methodName.Destroy()
    ptrsForColor.methodHex64Fn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeColor, methodName.AsCPtr(), 351421375))
  }
  {
    methodName := StringNameFromStr("html")
    defer methodName.Destroy()
    ptrsForColor.methodHtmlFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeColor, methodName.AsCPtr(), 2500054655))
  }
  {
    methodName := StringNameFromStr("html_is_valid")
    defer methodName.Destroy()
    ptrsForColor.methodHtmlIsValidFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeColor, methodName.AsCPtr(), 2942997125))
  }
  {
    methodName := StringNameFromStr("from_string")
    defer methodName.Destroy()
    ptrsForColor.methodFromStringFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeColor, methodName.AsCPtr(), 3755044230))
  }
  {
    methodName := StringNameFromStr("from_hsv")
    defer methodName.Destroy()
    ptrsForColor.methodFromHsvFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeColor, methodName.AsCPtr(), 1573799446))
  }
  {
    methodName := StringNameFromStr("from_ok_hsl")
    defer methodName.Destroy()
    ptrsForColor.methodFromOkHslFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeColor, methodName.AsCPtr(), 1573799446))
  }
  {
    methodName := StringNameFromStr("from_rgbe9995")
    defer methodName.Destroy()
    ptrsForColor.methodFromRgbe9995Fn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeColor, methodName.AsCPtr(), 351421375))
  }
  ptrsForColor.operatorNegateFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNegate, gdc.VariantTypeColor, gdc.VariantTypeNil))
  ptrsForColor.operatorPositiveFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpPositive, gdc.VariantTypeColor, gdc.VariantTypeNil))
  ptrsForColor.operatorNotFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, gdc.VariantTypeColor, gdc.VariantTypeNil))
  ptrsForColor.operatorMultiplyIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeColor, gdc.VariantTypeInt))
  ptrsForColor.operatorDivideIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpDivide, gdc.VariantTypeColor, gdc.VariantTypeInt))
  ptrsForColor.operatorMultiplyFloat32Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeColor, gdc.VariantTypeFloat))
  ptrsForColor.operatorDivideFloat32Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpDivide, gdc.VariantTypeColor, gdc.VariantTypeFloat))
  ptrsForColor.operatorEqualColorFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeColor, gdc.VariantTypeColor))
  ptrsForColor.operatorNotEqualColorFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeColor, gdc.VariantTypeColor))
  ptrsForColor.operatorAddColorFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpAdd, gdc.VariantTypeColor, gdc.VariantTypeColor))
  ptrsForColor.operatorSubtractColorFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpSubtract, gdc.VariantTypeColor, gdc.VariantTypeColor))
  ptrsForColor.operatorMultiplyColorFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeColor, gdc.VariantTypeColor))
  ptrsForColor.operatorDivideColorFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpDivide, gdc.VariantTypeColor, gdc.VariantTypeColor))
  ptrsForColor.operatorInDictionaryFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeColor, gdc.VariantTypeDictionary))
  ptrsForColor.operatorInArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeColor, gdc.VariantTypeArray))
  ptrsForColor.operatorInPackedColorArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeColor, gdc.VariantTypePackedColorArray))
  {
    memberName := StringNameFromStr("r")
    defer memberName.Destroy()
    ptrsForColor.memberrGetterFn = ensurePtr(iface.VariantGetPtrGetter(gdc.VariantTypeColor, memberName.AsCPtr()))
    ptrsForColor.memberrSetterFn = ensurePtr(iface.VariantGetPtrSetter(gdc.VariantTypeColor, memberName.AsCPtr()))
  }
  {
    memberName := StringNameFromStr("g")
    defer memberName.Destroy()
    ptrsForColor.membergGetterFn = ensurePtr(iface.VariantGetPtrGetter(gdc.VariantTypeColor, memberName.AsCPtr()))
    ptrsForColor.membergSetterFn = ensurePtr(iface.VariantGetPtrSetter(gdc.VariantTypeColor, memberName.AsCPtr()))
  }
  {
    memberName := StringNameFromStr("b")
    defer memberName.Destroy()
    ptrsForColor.memberbGetterFn = ensurePtr(iface.VariantGetPtrGetter(gdc.VariantTypeColor, memberName.AsCPtr()))
    ptrsForColor.memberbSetterFn = ensurePtr(iface.VariantGetPtrSetter(gdc.VariantTypeColor, memberName.AsCPtr()))
  }
  {
    memberName := StringNameFromStr("a")
    defer memberName.Destroy()
    ptrsForColor.memberaGetterFn = ensurePtr(iface.VariantGetPtrGetter(gdc.VariantTypeColor, memberName.AsCPtr()))
    ptrsForColor.memberaSetterFn = ensurePtr(iface.VariantGetPtrSetter(gdc.VariantTypeColor, memberName.AsCPtr()))
  }
  {
    memberName := StringNameFromStr("r8")
    defer memberName.Destroy()
    ptrsForColor.memberr8GetterFn = ensurePtr(iface.VariantGetPtrGetter(gdc.VariantTypeColor, memberName.AsCPtr()))
    ptrsForColor.memberr8SetterFn = ensurePtr(iface.VariantGetPtrSetter(gdc.VariantTypeColor, memberName.AsCPtr()))
  }
  {
    memberName := StringNameFromStr("g8")
    defer memberName.Destroy()
    ptrsForColor.memberg8GetterFn = ensurePtr(iface.VariantGetPtrGetter(gdc.VariantTypeColor, memberName.AsCPtr()))
    ptrsForColor.memberg8SetterFn = ensurePtr(iface.VariantGetPtrSetter(gdc.VariantTypeColor, memberName.AsCPtr()))
  }
  {
    memberName := StringNameFromStr("b8")
    defer memberName.Destroy()
    ptrsForColor.memberb8GetterFn = ensurePtr(iface.VariantGetPtrGetter(gdc.VariantTypeColor, memberName.AsCPtr()))
    ptrsForColor.memberb8SetterFn = ensurePtr(iface.VariantGetPtrSetter(gdc.VariantTypeColor, memberName.AsCPtr()))
  }
  {
    memberName := StringNameFromStr("a8")
    defer memberName.Destroy()
    ptrsForColor.membera8GetterFn = ensurePtr(iface.VariantGetPtrGetter(gdc.VariantTypeColor, memberName.AsCPtr()))
    ptrsForColor.membera8SetterFn = ensurePtr(iface.VariantGetPtrSetter(gdc.VariantTypeColor, memberName.AsCPtr()))
  }
  {
    memberName := StringNameFromStr("h")
    defer memberName.Destroy()
    ptrsForColor.memberhGetterFn = ensurePtr(iface.VariantGetPtrGetter(gdc.VariantTypeColor, memberName.AsCPtr()))
    ptrsForColor.memberhSetterFn = ensurePtr(iface.VariantGetPtrSetter(gdc.VariantTypeColor, memberName.AsCPtr()))
  }
  {
    memberName := StringNameFromStr("s")
    defer memberName.Destroy()
    ptrsForColor.membersGetterFn = ensurePtr(iface.VariantGetPtrGetter(gdc.VariantTypeColor, memberName.AsCPtr()))
    ptrsForColor.membersSetterFn = ensurePtr(iface.VariantGetPtrSetter(gdc.VariantTypeColor, memberName.AsCPtr()))
  }
  {
    memberName := StringNameFromStr("v")
    defer memberName.Destroy()
    ptrsForColor.membervGetterFn = ensurePtr(iface.VariantGetPtrGetter(gdc.VariantTypeColor, memberName.AsCPtr()))
    ptrsForColor.membervSetterFn = ensurePtr(iface.VariantGetPtrSetter(gdc.VariantTypeColor, memberName.AsCPtr()))
  }
  ptrsForColor.toVariantFn = ensurePtr(iface.GetVariantToTypeConstructor(gdc.VariantTypeColor))
  ptrsForColor.fromVariantFn = ensurePtr(iface.GetVariantFromTypeConstructor(gdc.VariantTypeColor))
}

type Color struct {
  data   *[classSizeColor]byte
  iface  gdc.Interface
  pinner runtime.Pinner
}

// Constants

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

// Enums

// Constructors
func newColor() *Color {
  me := &Color{
    data:   new([classSizeColor]byte),
    iface:  giface,
  }
  me.pinner.Pin(me)
  me.pinner.Pin(me.AsTypePtr())
  return me
}

func NewColor() *Color {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newColor()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForColor.ctrFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{}))
  return me
}

func NewColorFromColor(from Color, ) *Color {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newColor()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForColor.ctrFromColorFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return me
}

func NewColorFromColorFloat32(from Color, alpha float64, ) *Color {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  pinner.Pin(&alpha)
  me := newColor()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForColor.ctrFromColorFloat32Fn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), gdc.ConstTypePtr(&alpha), }))
  return me
}

func NewColorFromFloat32Float32Float32(r float64, g float64, b float64, ) *Color {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  pinner.Pin(&r)
  pinner.Pin(&g)
  pinner.Pin(&b)
  me := newColor()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForColor.ctrFromFloat32Float32Float32Fn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{gdc.ConstTypePtr(&r), gdc.ConstTypePtr(&g), gdc.ConstTypePtr(&b), }))
  return me
}

func NewColorFromFloat32Float32Float32Float32(r float64, g float64, b float64, a float64, ) *Color {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  pinner.Pin(&r)
  pinner.Pin(&g)
  pinner.Pin(&b)
  pinner.Pin(&a)
  me := newColor()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForColor.ctrFromFloat32Float32Float32Float32Fn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{gdc.ConstTypePtr(&r), gdc.ConstTypePtr(&g), gdc.ConstTypePtr(&b), gdc.ConstTypePtr(&a), }))
  return me
}

func NewColorFromString(code String, ) *Color {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newColor()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForColor.ctrFromStringFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{code.AsCTypePtr(), }))
  return me
}

func NewColorFromStringFloat32(code String, alpha float64, ) *Color {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  pinner.Pin(&alpha)
  me := newColor()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForColor.ctrFromStringFloat32Fn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{code.AsCTypePtr(), gdc.ConstTypePtr(&alpha), }))
  return me
}

// Destructor
func (me *Color) Destroy() {
  me.pinner.Unpin()
}

// Variant support
func (me *Variant) AsColor() (*Color, error) {
	if me.Type() != gdc.VariantTypeColor {
		return nil, fmt.Errorf("variant is not a Color")
	}
  bclass := newColor()
	me.iface.CallTypeFromVariantConstructorFunc(ensurePtr(ptrsForColor.toVariantFn), bclass.asUninitialized(), me.AsPtr())
	return bclass, nil
}

func (me *Color) AsVariant() *Variant {
  va := newVariant()
  va.inner = me
  me.iface.CallVariantFromTypeConstructorFunc(ensurePtr(ptrsForColor.fromVariantFn), va.asUninitialized(), me.AsTypePtr())
  return va
}

// Pointers
func ColorFromPtr(ptr gdc.ConstTypePtr) *Color {
  me := newColor()
  dataFromPtr(me.data[:], ptr)
  return me
}

func (me *Color) Type() gdc.VariantType {
  return gdc.VariantTypeColor
}

func (me *Color) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(unsafe.Pointer(me.data))
}

func (me *Color) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.AsTypePtr())
}

func (me *Color) asUninitialized() gdc.UninitializedTypePtr {
  return gdc.UninitializedTypePtr(me.AsTypePtr())
}

// Methods

func (me *Color) ToArgb32() int64 {
  ret := NewInt()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForColor.methodToArgb32Fn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Color) ToAbgr32() int64 {
  ret := NewInt()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForColor.methodToAbgr32Fn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Color) ToRgba32() int64 {
  ret := NewInt()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForColor.methodToRgba32Fn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Color) ToArgb64() int64 {
  ret := NewInt()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForColor.methodToArgb64Fn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Color) ToAbgr64() int64 {
  ret := NewInt()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForColor.methodToAbgr64Fn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Color) ToRgba64() int64 {
  ret := NewInt()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForColor.methodToRgba64Fn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Color) ToHtml(with_alpha bool, ) String {
  ret := NewString()

  varg0 := NewBoolFromBool(with_alpha)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForColor.methodToHtmlFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Color) Clamp(min Color, max Color, ) Color {
  ret := NewColor()



  args := []gdc.ConstTypePtr{min.AsCTypePtr(), max.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForColor.methodClampFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Color) Inverted() Color {
  ret := NewColor()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForColor.methodInvertedFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Color) Lerp(to Color, weight float64, ) Color {
  ret := NewColor()


  varg1 := NewFloatFromFloat32(weight)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{to.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForColor.methodLerpFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Color) Lightened(amount float64, ) Color {
  ret := NewColor()

  varg0 := NewFloatFromFloat32(amount)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForColor.methodLightenedFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Color) Darkened(amount float64, ) Color {
  ret := NewColor()

  varg0 := NewFloatFromFloat32(amount)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForColor.methodDarkenedFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Color) Blend(over Color, ) Color {
  ret := NewColor()


  args := []gdc.ConstTypePtr{over.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForColor.methodBlendFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Color) GetLuminance() float64 {
  ret := NewFloat()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForColor.methodGetLuminanceFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Color) SrgbToLinear() Color {
  ret := NewColor()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForColor.methodSrgbToLinearFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Color) LinearToSrgb() Color {
  ret := NewColor()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForColor.methodLinearToSrgbFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Color) IsEqualApprox(to Color, ) bool {
  ret := NewBool()
  defer ret.Destroy()

  args := []gdc.ConstTypePtr{to.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForColor.methodIsEqualApproxFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func ColorHex(hex int64, ) Color {
  ret := NewColor()

  varg0 := NewIntFromInt(hex)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForColor.methodHexFn), nil, unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func ColorHex64(hex int64, ) Color {
  ret := NewColor()

  varg0 := NewIntFromInt(hex)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForColor.methodHex64Fn), nil, unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func ColorHtml(rgba String, ) Color {
  ret := NewColor()


  args := []gdc.ConstTypePtr{rgba.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForColor.methodHtmlFn), nil, unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func ColorHtmlIsValid(color String, ) bool {
  ret := NewBool()
  defer ret.Destroy()

  args := []gdc.ConstTypePtr{color.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForColor.methodHtmlIsValidFn), nil, unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func ColorFromString(str String, default_ Color, ) Color {
  ret := NewColor()



  args := []gdc.ConstTypePtr{str.AsCTypePtr(), default_.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForColor.methodFromStringFn), nil, unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func ColorFromHsv(h float64, s float64, v float64, alpha float64, ) Color {
  ret := NewColor()

  varg0 := NewFloatFromFloat32(h)
  defer varg0.Destroy()
  varg1 := NewFloatFromFloat32(s)
  defer varg1.Destroy()
  varg2 := NewFloatFromFloat32(v)
  defer varg2.Destroy()
  varg3 := NewFloatFromFloat32(alpha)
  defer varg3.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), varg2.AsCTypePtr(), varg3.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForColor.methodFromHsvFn), nil, unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func ColorFromOkHsl(h float64, s float64, l float64, alpha float64, ) Color {
  ret := NewColor()

  varg0 := NewFloatFromFloat32(h)
  defer varg0.Destroy()
  varg1 := NewFloatFromFloat32(s)
  defer varg1.Destroy()
  varg2 := NewFloatFromFloat32(l)
  defer varg2.Destroy()
  varg3 := NewFloatFromFloat32(alpha)
  defer varg3.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), varg2.AsCTypePtr(), varg3.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForColor.methodFromOkHslFn), nil, unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func ColorFromRgbe9995(rgbe int64, ) Color {
  ret := NewColor()

  varg0 := NewIntFromInt(rgbe)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForColor.methodFromRgbe9995Fn), nil, unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

// Operators

func (me *Color) EqualVariant(right Variant) bool {
  opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type())
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Color) NotEqualVariant(right Variant) bool {
  opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type())
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Color) Negate() Color {
  opPtr := ptrsForColor.operatorNegateFn
  ret := NewColor()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), nil, ret.AsTypePtr())
  return *ret
}

func (me *Color) Positive() Color {
  opPtr := ptrsForColor.operatorPositiveFn
  ret := NewColor()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), nil, ret.AsTypePtr())
  return *ret
}

func (me *Color) Not() bool {
  opPtr := ptrsForColor.operatorNotFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), nil, ret.AsTypePtr())
  return ret.Get()
}

func (me *Color) MultiplyInt(rightArg int64) Color {
  right := NewIntFromInt(rightArg)
  defer right.Destroy()

  opPtr := ptrsForColor.operatorMultiplyIntFn
  ret := NewColor()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Color) DivideInt(rightArg int64) Color {
  right := NewIntFromInt(rightArg)
  defer right.Destroy()

  opPtr := ptrsForColor.operatorDivideIntFn
  ret := NewColor()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Color) MultiplyFloat32(rightArg float64) Color {
  right := NewFloatFromFloat32(rightArg)
  defer right.Destroy()

  opPtr := ptrsForColor.operatorMultiplyFloat32Fn
  ret := NewColor()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Color) DivideFloat32(rightArg float64) Color {
  right := NewFloatFromFloat32(rightArg)
  defer right.Destroy()

  opPtr := ptrsForColor.operatorDivideFloat32Fn
  ret := NewColor()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Color) EqualColor(right Color) bool {
  opPtr := ptrsForColor.operatorEqualColorFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Color) NotEqualColor(right Color) bool {
  opPtr := ptrsForColor.operatorNotEqualColorFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Color) AddColor(right Color) Color {
  opPtr := ptrsForColor.operatorAddColorFn
  ret := NewColor()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Color) SubtractColor(right Color) Color {
  opPtr := ptrsForColor.operatorSubtractColorFn
  ret := NewColor()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Color) MultiplyColor(right Color) Color {
  opPtr := ptrsForColor.operatorMultiplyColorFn
  ret := NewColor()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Color) DivideColor(right Color) Color {
  opPtr := ptrsForColor.operatorDivideColorFn
  ret := NewColor()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *Color) InDictionary(right Dictionary) bool {
  opPtr := ptrsForColor.operatorInDictionaryFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Color) InArray(right Array) bool {
  opPtr := ptrsForColor.operatorInArrayFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Color) InPackedColorArray(right PackedColorArray) bool {
  opPtr := ptrsForColor.operatorInPackedColorArrayFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

// Members

func (me *Color) R() float64 {
  ret := NewFloat()
  me.iface.CallPtrGetter(ensurePtr(ptrsForColor.memberrGetterFn), me.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Color) SetR(value float64) {
  valueV := NewFloatFromFloat32(value)
  defer valueV.Destroy()
  me.iface.CallPtrSetter(ensurePtr(ptrsForColor.memberrSetterFn), me.AsTypePtr(), valueV.AsCTypePtr())
}

func (me *Color) G() float64 {
  ret := NewFloat()
  me.iface.CallPtrGetter(ensurePtr(ptrsForColor.membergGetterFn), me.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Color) SetG(value float64) {
  valueV := NewFloatFromFloat32(value)
  defer valueV.Destroy()
  me.iface.CallPtrSetter(ensurePtr(ptrsForColor.membergSetterFn), me.AsTypePtr(), valueV.AsCTypePtr())
}

func (me *Color) B() float64 {
  ret := NewFloat()
  me.iface.CallPtrGetter(ensurePtr(ptrsForColor.memberbGetterFn), me.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Color) SetB(value float64) {
  valueV := NewFloatFromFloat32(value)
  defer valueV.Destroy()
  me.iface.CallPtrSetter(ensurePtr(ptrsForColor.memberbSetterFn), me.AsTypePtr(), valueV.AsCTypePtr())
}

func (me *Color) A() float64 {
  ret := NewFloat()
  me.iface.CallPtrGetter(ensurePtr(ptrsForColor.memberaGetterFn), me.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Color) SetA(value float64) {
  valueV := NewFloatFromFloat32(value)
  defer valueV.Destroy()
  me.iface.CallPtrSetter(ensurePtr(ptrsForColor.memberaSetterFn), me.AsTypePtr(), valueV.AsCTypePtr())
}

func (me *Color) R8() int64 {
  ret := NewInt()
  me.iface.CallPtrGetter(ensurePtr(ptrsForColor.memberr8GetterFn), me.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Color) SetR8(value int64) {
  valueV := NewIntFromInt(value)
  defer valueV.Destroy()
  me.iface.CallPtrSetter(ensurePtr(ptrsForColor.memberr8SetterFn), me.AsTypePtr(), valueV.AsCTypePtr())
}

func (me *Color) G8() int64 {
  ret := NewInt()
  me.iface.CallPtrGetter(ensurePtr(ptrsForColor.memberg8GetterFn), me.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Color) SetG8(value int64) {
  valueV := NewIntFromInt(value)
  defer valueV.Destroy()
  me.iface.CallPtrSetter(ensurePtr(ptrsForColor.memberg8SetterFn), me.AsTypePtr(), valueV.AsCTypePtr())
}

func (me *Color) B8() int64 {
  ret := NewInt()
  me.iface.CallPtrGetter(ensurePtr(ptrsForColor.memberb8GetterFn), me.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Color) SetB8(value int64) {
  valueV := NewIntFromInt(value)
  defer valueV.Destroy()
  me.iface.CallPtrSetter(ensurePtr(ptrsForColor.memberb8SetterFn), me.AsTypePtr(), valueV.AsCTypePtr())
}

func (me *Color) A8() int64 {
  ret := NewInt()
  me.iface.CallPtrGetter(ensurePtr(ptrsForColor.membera8GetterFn), me.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Color) SetA8(value int64) {
  valueV := NewIntFromInt(value)
  defer valueV.Destroy()
  me.iface.CallPtrSetter(ensurePtr(ptrsForColor.membera8SetterFn), me.AsTypePtr(), valueV.AsCTypePtr())
}

func (me *Color) H() float64 {
  ret := NewFloat()
  me.iface.CallPtrGetter(ensurePtr(ptrsForColor.memberhGetterFn), me.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Color) SetH(value float64) {
  valueV := NewFloatFromFloat32(value)
  defer valueV.Destroy()
  me.iface.CallPtrSetter(ensurePtr(ptrsForColor.memberhSetterFn), me.AsTypePtr(), valueV.AsCTypePtr())
}

func (me *Color) S() float64 {
  ret := NewFloat()
  me.iface.CallPtrGetter(ensurePtr(ptrsForColor.membersGetterFn), me.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Color) SetS(value float64) {
  valueV := NewFloatFromFloat32(value)
  defer valueV.Destroy()
  me.iface.CallPtrSetter(ensurePtr(ptrsForColor.membersSetterFn), me.AsTypePtr(), valueV.AsCTypePtr())
}

func (me *Color) V() float64 {
  ret := NewFloat()
  me.iface.CallPtrGetter(ensurePtr(ptrsForColor.membervGetterFn), me.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Color) SetV(value float64) {
  valueV := NewFloatFromFloat32(value)
  defer valueV.Destroy()
  me.iface.CallPtrSetter(ensurePtr(ptrsForColor.membervSetterFn), me.AsTypePtr(), valueV.AsCTypePtr())
}
