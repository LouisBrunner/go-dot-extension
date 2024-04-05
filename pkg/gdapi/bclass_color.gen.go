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
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("ALICE_BLUE")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value ALICE_BLUE: " + err.Error())
    }
    ColorAliceBlue = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("ANTIQUE_WHITE")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value ANTIQUE_WHITE: " + err.Error())
    }
    ColorAntiqueWhite = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("AQUA")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value AQUA: " + err.Error())
    }
    ColorAqua = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("AQUAMARINE")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value AQUAMARINE: " + err.Error())
    }
    ColorAquamarine = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("AZURE")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value AZURE: " + err.Error())
    }
    ColorAzure = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("BEIGE")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value BEIGE: " + err.Error())
    }
    ColorBeige = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("BISQUE")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value BISQUE: " + err.Error())
    }
    ColorBisque = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("BLACK")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value BLACK: " + err.Error())
    }
    ColorBlack = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("BLANCHED_ALMOND")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value BLANCHED_ALMOND: " + err.Error())
    }
    ColorBlanchedAlmond = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("BLUE")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value BLUE: " + err.Error())
    }
    ColorBlue = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("BLUE_VIOLET")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value BLUE_VIOLET: " + err.Error())
    }
    ColorBlueViolet = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("BROWN")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value BROWN: " + err.Error())
    }
    ColorBrown = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("BURLYWOOD")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value BURLYWOOD: " + err.Error())
    }
    ColorBurlywood = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("CADET_BLUE")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value CADET_BLUE: " + err.Error())
    }
    ColorCadetBlue = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("CHARTREUSE")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value CHARTREUSE: " + err.Error())
    }
    ColorChartreuse = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("CHOCOLATE")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value CHOCOLATE: " + err.Error())
    }
    ColorChocolate = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("CORAL")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value CORAL: " + err.Error())
    }
    ColorCoral = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("CORNFLOWER_BLUE")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value CORNFLOWER_BLUE: " + err.Error())
    }
    ColorCornflowerBlue = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("CORNSILK")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value CORNSILK: " + err.Error())
    }
    ColorCornsilk = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("CRIMSON")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value CRIMSON: " + err.Error())
    }
    ColorCrimson = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("CYAN")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value CYAN: " + err.Error())
    }
    ColorCyan = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("DARK_BLUE")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value DARK_BLUE: " + err.Error())
    }
    ColorDarkBlue = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("DARK_CYAN")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value DARK_CYAN: " + err.Error())
    }
    ColorDarkCyan = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("DARK_GOLDENROD")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value DARK_GOLDENROD: " + err.Error())
    }
    ColorDarkGoldenrod = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("DARK_GRAY")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value DARK_GRAY: " + err.Error())
    }
    ColorDarkGray = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("DARK_GREEN")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value DARK_GREEN: " + err.Error())
    }
    ColorDarkGreen = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("DARK_KHAKI")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value DARK_KHAKI: " + err.Error())
    }
    ColorDarkKhaki = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("DARK_MAGENTA")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value DARK_MAGENTA: " + err.Error())
    }
    ColorDarkMagenta = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("DARK_OLIVE_GREEN")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value DARK_OLIVE_GREEN: " + err.Error())
    }
    ColorDarkOliveGreen = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("DARK_ORANGE")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value DARK_ORANGE: " + err.Error())
    }
    ColorDarkOrange = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("DARK_ORCHID")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value DARK_ORCHID: " + err.Error())
    }
    ColorDarkOrchid = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("DARK_RED")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value DARK_RED: " + err.Error())
    }
    ColorDarkRed = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("DARK_SALMON")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value DARK_SALMON: " + err.Error())
    }
    ColorDarkSalmon = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("DARK_SEA_GREEN")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value DARK_SEA_GREEN: " + err.Error())
    }
    ColorDarkSeaGreen = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("DARK_SLATE_BLUE")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value DARK_SLATE_BLUE: " + err.Error())
    }
    ColorDarkSlateBlue = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("DARK_SLATE_GRAY")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value DARK_SLATE_GRAY: " + err.Error())
    }
    ColorDarkSlateGray = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("DARK_TURQUOISE")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value DARK_TURQUOISE: " + err.Error())
    }
    ColorDarkTurquoise = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("DARK_VIOLET")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value DARK_VIOLET: " + err.Error())
    }
    ColorDarkViolet = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("DEEP_PINK")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value DEEP_PINK: " + err.Error())
    }
    ColorDeepPink = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("DEEP_SKY_BLUE")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value DEEP_SKY_BLUE: " + err.Error())
    }
    ColorDeepSkyBlue = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("DIM_GRAY")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value DIM_GRAY: " + err.Error())
    }
    ColorDimGray = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("DODGER_BLUE")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value DODGER_BLUE: " + err.Error())
    }
    ColorDodgerBlue = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("FIREBRICK")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value FIREBRICK: " + err.Error())
    }
    ColorFirebrick = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("FLORAL_WHITE")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value FLORAL_WHITE: " + err.Error())
    }
    ColorFloralWhite = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("FOREST_GREEN")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value FOREST_GREEN: " + err.Error())
    }
    ColorForestGreen = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("FUCHSIA")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value FUCHSIA: " + err.Error())
    }
    ColorFuchsia = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("GAINSBORO")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value GAINSBORO: " + err.Error())
    }
    ColorGainsboro = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("GHOST_WHITE")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value GHOST_WHITE: " + err.Error())
    }
    ColorGhostWhite = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("GOLD")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value GOLD: " + err.Error())
    }
    ColorGold = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("GOLDENROD")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value GOLDENROD: " + err.Error())
    }
    ColorGoldenrod = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("GRAY")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value GRAY: " + err.Error())
    }
    ColorGray = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("GREEN")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value GREEN: " + err.Error())
    }
    ColorGreen = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("GREEN_YELLOW")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value GREEN_YELLOW: " + err.Error())
    }
    ColorGreenYellow = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("HONEYDEW")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value HONEYDEW: " + err.Error())
    }
    ColorHoneydew = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("HOT_PINK")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value HOT_PINK: " + err.Error())
    }
    ColorHotPink = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("INDIAN_RED")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value INDIAN_RED: " + err.Error())
    }
    ColorIndianRed = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("INDIGO")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value INDIGO: " + err.Error())
    }
    ColorIndigo = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("IVORY")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value IVORY: " + err.Error())
    }
    ColorIvory = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("KHAKI")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value KHAKI: " + err.Error())
    }
    ColorKhaki = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("LAVENDER")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value LAVENDER: " + err.Error())
    }
    ColorLavender = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("LAVENDER_BLUSH")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value LAVENDER_BLUSH: " + err.Error())
    }
    ColorLavenderBlush = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("LAWN_GREEN")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value LAWN_GREEN: " + err.Error())
    }
    ColorLawnGreen = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("LEMON_CHIFFON")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value LEMON_CHIFFON: " + err.Error())
    }
    ColorLemonChiffon = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("LIGHT_BLUE")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value LIGHT_BLUE: " + err.Error())
    }
    ColorLightBlue = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("LIGHT_CORAL")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value LIGHT_CORAL: " + err.Error())
    }
    ColorLightCoral = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("LIGHT_CYAN")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value LIGHT_CYAN: " + err.Error())
    }
    ColorLightCyan = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("LIGHT_GOLDENROD")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value LIGHT_GOLDENROD: " + err.Error())
    }
    ColorLightGoldenrod = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("LIGHT_GRAY")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value LIGHT_GRAY: " + err.Error())
    }
    ColorLightGray = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("LIGHT_GREEN")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value LIGHT_GREEN: " + err.Error())
    }
    ColorLightGreen = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("LIGHT_PINK")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value LIGHT_PINK: " + err.Error())
    }
    ColorLightPink = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("LIGHT_SALMON")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value LIGHT_SALMON: " + err.Error())
    }
    ColorLightSalmon = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("LIGHT_SEA_GREEN")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value LIGHT_SEA_GREEN: " + err.Error())
    }
    ColorLightSeaGreen = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("LIGHT_SKY_BLUE")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value LIGHT_SKY_BLUE: " + err.Error())
    }
    ColorLightSkyBlue = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("LIGHT_SLATE_GRAY")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value LIGHT_SLATE_GRAY: " + err.Error())
    }
    ColorLightSlateGray = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("LIGHT_STEEL_BLUE")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value LIGHT_STEEL_BLUE: " + err.Error())
    }
    ColorLightSteelBlue = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("LIGHT_YELLOW")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value LIGHT_YELLOW: " + err.Error())
    }
    ColorLightYellow = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("LIME")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value LIME: " + err.Error())
    }
    ColorLime = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("LIME_GREEN")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value LIME_GREEN: " + err.Error())
    }
    ColorLimeGreen = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("LINEN")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value LINEN: " + err.Error())
    }
    ColorLinen = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("MAGENTA")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value MAGENTA: " + err.Error())
    }
    ColorMagenta = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("MAROON")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value MAROON: " + err.Error())
    }
    ColorMaroon = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("MEDIUM_AQUAMARINE")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value MEDIUM_AQUAMARINE: " + err.Error())
    }
    ColorMediumAquamarine = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("MEDIUM_BLUE")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value MEDIUM_BLUE: " + err.Error())
    }
    ColorMediumBlue = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("MEDIUM_ORCHID")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value MEDIUM_ORCHID: " + err.Error())
    }
    ColorMediumOrchid = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("MEDIUM_PURPLE")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value MEDIUM_PURPLE: " + err.Error())
    }
    ColorMediumPurple = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("MEDIUM_SEA_GREEN")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value MEDIUM_SEA_GREEN: " + err.Error())
    }
    ColorMediumSeaGreen = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("MEDIUM_SLATE_BLUE")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value MEDIUM_SLATE_BLUE: " + err.Error())
    }
    ColorMediumSlateBlue = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("MEDIUM_SPRING_GREEN")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value MEDIUM_SPRING_GREEN: " + err.Error())
    }
    ColorMediumSpringGreen = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("MEDIUM_TURQUOISE")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value MEDIUM_TURQUOISE: " + err.Error())
    }
    ColorMediumTurquoise = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("MEDIUM_VIOLET_RED")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value MEDIUM_VIOLET_RED: " + err.Error())
    }
    ColorMediumVioletRed = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("MIDNIGHT_BLUE")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value MIDNIGHT_BLUE: " + err.Error())
    }
    ColorMidnightBlue = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("MINT_CREAM")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value MINT_CREAM: " + err.Error())
    }
    ColorMintCream = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("MISTY_ROSE")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value MISTY_ROSE: " + err.Error())
    }
    ColorMistyRose = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("MOCCASIN")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value MOCCASIN: " + err.Error())
    }
    ColorMoccasin = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("NAVAJO_WHITE")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value NAVAJO_WHITE: " + err.Error())
    }
    ColorNavajoWhite = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("NAVY_BLUE")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value NAVY_BLUE: " + err.Error())
    }
    ColorNavyBlue = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("OLD_LACE")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value OLD_LACE: " + err.Error())
    }
    ColorOldLace = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("OLIVE")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value OLIVE: " + err.Error())
    }
    ColorOlive = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("OLIVE_DRAB")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value OLIVE_DRAB: " + err.Error())
    }
    ColorOliveDrab = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("ORANGE")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value ORANGE: " + err.Error())
    }
    ColorOrange = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("ORANGE_RED")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value ORANGE_RED: " + err.Error())
    }
    ColorOrangeRed = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("ORCHID")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value ORCHID: " + err.Error())
    }
    ColorOrchid = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("PALE_GOLDENROD")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value PALE_GOLDENROD: " + err.Error())
    }
    ColorPaleGoldenrod = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("PALE_GREEN")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value PALE_GREEN: " + err.Error())
    }
    ColorPaleGreen = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("PALE_TURQUOISE")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value PALE_TURQUOISE: " + err.Error())
    }
    ColorPaleTurquoise = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("PALE_VIOLET_RED")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value PALE_VIOLET_RED: " + err.Error())
    }
    ColorPaleVioletRed = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("PAPAYA_WHIP")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value PAPAYA_WHIP: " + err.Error())
    }
    ColorPapayaWhip = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("PEACH_PUFF")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value PEACH_PUFF: " + err.Error())
    }
    ColorPeachPuff = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("PERU")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value PERU: " + err.Error())
    }
    ColorPeru = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("PINK")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value PINK: " + err.Error())
    }
    ColorPink = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("PLUM")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value PLUM: " + err.Error())
    }
    ColorPlum = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("POWDER_BLUE")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value POWDER_BLUE: " + err.Error())
    }
    ColorPowderBlue = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("PURPLE")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value PURPLE: " + err.Error())
    }
    ColorPurple = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("REBECCA_PURPLE")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value REBECCA_PURPLE: " + err.Error())
    }
    ColorRebeccaPurple = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("RED")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value RED: " + err.Error())
    }
    ColorRed = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("ROSY_BROWN")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value ROSY_BROWN: " + err.Error())
    }
    ColorRosyBrown = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("ROYAL_BLUE")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value ROYAL_BLUE: " + err.Error())
    }
    ColorRoyalBlue = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("SADDLE_BROWN")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value SADDLE_BROWN: " + err.Error())
    }
    ColorSaddleBrown = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("SALMON")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value SALMON: " + err.Error())
    }
    ColorSalmon = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("SANDY_BROWN")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value SANDY_BROWN: " + err.Error())
    }
    ColorSandyBrown = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("SEA_GREEN")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value SEA_GREEN: " + err.Error())
    }
    ColorSeaGreen = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("SEASHELL")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value SEASHELL: " + err.Error())
    }
    ColorSeashell = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("SIENNA")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value SIENNA: " + err.Error())
    }
    ColorSienna = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("SILVER")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value SILVER: " + err.Error())
    }
    ColorSilver = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("SKY_BLUE")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value SKY_BLUE: " + err.Error())
    }
    ColorSkyBlue = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("SLATE_BLUE")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value SLATE_BLUE: " + err.Error())
    }
    ColorSlateBlue = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("SLATE_GRAY")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value SLATE_GRAY: " + err.Error())
    }
    ColorSlateGray = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("SNOW")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value SNOW: " + err.Error())
    }
    ColorSnow = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("SPRING_GREEN")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value SPRING_GREEN: " + err.Error())
    }
    ColorSpringGreen = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("STEEL_BLUE")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value STEEL_BLUE: " + err.Error())
    }
    ColorSteelBlue = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("TAN")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value TAN: " + err.Error())
    }
    ColorTan = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("TEAL")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value TEAL: " + err.Error())
    }
    ColorTeal = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("THISTLE")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value THISTLE: " + err.Error())
    }
    ColorThistle = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("TOMATO")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value TOMATO: " + err.Error())
    }
    ColorTomato = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("TRANSPARENT")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value TRANSPARENT: " + err.Error())
    }
    ColorTransparent = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("TURQUOISE")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value TURQUOISE: " + err.Error())
    }
    ColorTurquoise = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("VIOLET")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value VIOLET: " + err.Error())
    }
    ColorViolet = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("WEB_GRAY")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value WEB_GRAY: " + err.Error())
    }
    ColorWebGray = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("WEB_GREEN")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value WEB_GREEN: " + err.Error())
    }
    ColorWebGreen = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("WEB_MAROON")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value WEB_MAROON: " + err.Error())
    }
    ColorWebMaroon = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("WEB_PURPLE")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value WEB_PURPLE: " + err.Error())
    }
    ColorWebPurple = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("WHEAT")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value WHEAT: " + err.Error())
    }
    ColorWheat = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("WHITE")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value WHITE: " + err.Error())
    }
    ColorWhite = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("WHITE_SMOKE")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value WHITE_SMOKE: " + err.Error())
    }
    ColorWhiteSmoke = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("YELLOW")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value YELLOW: " + err.Error())
    }
    ColorYellow = *cnst
  }
  {
    va := *newVariant()
    defer va.Destroy()
    name := StringNameFromStr("YELLOW_GREEN")
    defer name.Destroy()
    iface.VariantGetConstantValue(gdc.VariantTypeColor, name.AsCPtr(), va.asUninitialized())
    cnst, err := va.AsColor()
    if err != nil {
      panic("Failed to get constant value YELLOW_GREEN: " + err.Error())
    }
    ColorYellowGreen = *cnst
  }
}

type Color struct {
  data   *[classSizeColor]byte
  iface  gdc.Interface
  pinner runtime.Pinner
}

// Constants

var (
  ColorAliceBlue Color
  ColorAntiqueWhite Color
  ColorAqua Color
  ColorAquamarine Color
  ColorAzure Color
  ColorBeige Color
  ColorBisque Color
  ColorBlack Color
  ColorBlanchedAlmond Color
  ColorBlue Color
  ColorBlueViolet Color
  ColorBrown Color
  ColorBurlywood Color
  ColorCadetBlue Color
  ColorChartreuse Color
  ColorChocolate Color
  ColorCoral Color
  ColorCornflowerBlue Color
  ColorCornsilk Color
  ColorCrimson Color
  ColorCyan Color
  ColorDarkBlue Color
  ColorDarkCyan Color
  ColorDarkGoldenrod Color
  ColorDarkGray Color
  ColorDarkGreen Color
  ColorDarkKhaki Color
  ColorDarkMagenta Color
  ColorDarkOliveGreen Color
  ColorDarkOrange Color
  ColorDarkOrchid Color
  ColorDarkRed Color
  ColorDarkSalmon Color
  ColorDarkSeaGreen Color
  ColorDarkSlateBlue Color
  ColorDarkSlateGray Color
  ColorDarkTurquoise Color
  ColorDarkViolet Color
  ColorDeepPink Color
  ColorDeepSkyBlue Color
  ColorDimGray Color
  ColorDodgerBlue Color
  ColorFirebrick Color
  ColorFloralWhite Color
  ColorForestGreen Color
  ColorFuchsia Color
  ColorGainsboro Color
  ColorGhostWhite Color
  ColorGold Color
  ColorGoldenrod Color
  ColorGray Color
  ColorGreen Color
  ColorGreenYellow Color
  ColorHoneydew Color
  ColorHotPink Color
  ColorIndianRed Color
  ColorIndigo Color
  ColorIvory Color
  ColorKhaki Color
  ColorLavender Color
  ColorLavenderBlush Color
  ColorLawnGreen Color
  ColorLemonChiffon Color
  ColorLightBlue Color
  ColorLightCoral Color
  ColorLightCyan Color
  ColorLightGoldenrod Color
  ColorLightGray Color
  ColorLightGreen Color
  ColorLightPink Color
  ColorLightSalmon Color
  ColorLightSeaGreen Color
  ColorLightSkyBlue Color
  ColorLightSlateGray Color
  ColorLightSteelBlue Color
  ColorLightYellow Color
  ColorLime Color
  ColorLimeGreen Color
  ColorLinen Color
  ColorMagenta Color
  ColorMaroon Color
  ColorMediumAquamarine Color
  ColorMediumBlue Color
  ColorMediumOrchid Color
  ColorMediumPurple Color
  ColorMediumSeaGreen Color
  ColorMediumSlateBlue Color
  ColorMediumSpringGreen Color
  ColorMediumTurquoise Color
  ColorMediumVioletRed Color
  ColorMidnightBlue Color
  ColorMintCream Color
  ColorMistyRose Color
  ColorMoccasin Color
  ColorNavajoWhite Color
  ColorNavyBlue Color
  ColorOldLace Color
  ColorOlive Color
  ColorOliveDrab Color
  ColorOrange Color
  ColorOrangeRed Color
  ColorOrchid Color
  ColorPaleGoldenrod Color
  ColorPaleGreen Color
  ColorPaleTurquoise Color
  ColorPaleVioletRed Color
  ColorPapayaWhip Color
  ColorPeachPuff Color
  ColorPeru Color
  ColorPink Color
  ColorPlum Color
  ColorPowderBlue Color
  ColorPurple Color
  ColorRebeccaPurple Color
  ColorRed Color
  ColorRosyBrown Color
  ColorRoyalBlue Color
  ColorSaddleBrown Color
  ColorSalmon Color
  ColorSandyBrown Color
  ColorSeaGreen Color
  ColorSeashell Color
  ColorSienna Color
  ColorSilver Color
  ColorSkyBlue Color
  ColorSlateBlue Color
  ColorSlateGray Color
  ColorSnow Color
  ColorSpringGreen Color
  ColorSteelBlue Color
  ColorTan Color
  ColorTeal Color
  ColorThistle Color
  ColorTomato Color
  ColorTransparent Color
  ColorTurquoise Color
  ColorViolet Color
  ColorWebGray Color
  ColorWebGreen Color
  ColorWebMaroon Color
  ColorWebPurple Color
  ColorWheat Color
  ColorWhite Color
  ColorWhiteSmoke Color
  ColorYellow Color
  ColorYellowGreen Color
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
