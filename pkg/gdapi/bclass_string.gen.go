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

type ptrsForStringList struct {
  ctrFn gdc.PtrConstructor
  ctrFromStringFn gdc.PtrConstructor
  ctrFromStringNameFn gdc.PtrConstructor
  ctrFromNodePathFn gdc.PtrConstructor
  destructorFn gdc.PtrDestructor
  methodCasecmpToFn gdc.PtrBuiltInMethod
  methodNocasecmpToFn gdc.PtrBuiltInMethod
  methodNaturalcasecmpToFn gdc.PtrBuiltInMethod
  methodNaturalnocasecmpToFn gdc.PtrBuiltInMethod
  methodLengthFn gdc.PtrBuiltInMethod
  methodSubstrFn gdc.PtrBuiltInMethod
  methodGetSliceFn gdc.PtrBuiltInMethod
  methodGetSlicecFn gdc.PtrBuiltInMethod
  methodGetSliceCountFn gdc.PtrBuiltInMethod
  methodFindFn gdc.PtrBuiltInMethod
  methodCountFn gdc.PtrBuiltInMethod
  methodCountnFn gdc.PtrBuiltInMethod
  methodFindnFn gdc.PtrBuiltInMethod
  methodRfindFn gdc.PtrBuiltInMethod
  methodRfindnFn gdc.PtrBuiltInMethod
  methodMatchFn gdc.PtrBuiltInMethod
  methodMatchnFn gdc.PtrBuiltInMethod
  methodBeginsWithFn gdc.PtrBuiltInMethod
  methodEndsWithFn gdc.PtrBuiltInMethod
  methodIsSubsequenceOfFn gdc.PtrBuiltInMethod
  methodIsSubsequenceOfnFn gdc.PtrBuiltInMethod
  methodBigramsFn gdc.PtrBuiltInMethod
  methodSimilarityFn gdc.PtrBuiltInMethod
  methodFormatFn gdc.PtrBuiltInMethod
  methodReplaceFn gdc.PtrBuiltInMethod
  methodReplacenFn gdc.PtrBuiltInMethod
  methodRepeatFn gdc.PtrBuiltInMethod
  methodReverseFn gdc.PtrBuiltInMethod
  methodInsertFn gdc.PtrBuiltInMethod
  methodEraseFn gdc.PtrBuiltInMethod
  methodCapitalizeFn gdc.PtrBuiltInMethod
  methodToCamelCaseFn gdc.PtrBuiltInMethod
  methodToPascalCaseFn gdc.PtrBuiltInMethod
  methodToSnakeCaseFn gdc.PtrBuiltInMethod
  methodSplitFn gdc.PtrBuiltInMethod
  methodRsplitFn gdc.PtrBuiltInMethod
  methodSplitFloatsFn gdc.PtrBuiltInMethod
  methodJoinFn gdc.PtrBuiltInMethod
  methodToUpperFn gdc.PtrBuiltInMethod
  methodToLowerFn gdc.PtrBuiltInMethod
  methodLeftFn gdc.PtrBuiltInMethod
  methodRightFn gdc.PtrBuiltInMethod
  methodStripEdgesFn gdc.PtrBuiltInMethod
  methodStripEscapesFn gdc.PtrBuiltInMethod
  methodLstripFn gdc.PtrBuiltInMethod
  methodRstripFn gdc.PtrBuiltInMethod
  methodGetExtensionFn gdc.PtrBuiltInMethod
  methodGetBasenameFn gdc.PtrBuiltInMethod
  methodPathJoinFn gdc.PtrBuiltInMethod
  methodUnicodeAtFn gdc.PtrBuiltInMethod
  methodIndentFn gdc.PtrBuiltInMethod
  methodDedentFn gdc.PtrBuiltInMethod
  methodHashFn gdc.PtrBuiltInMethod
  methodMd5TextFn gdc.PtrBuiltInMethod
  methodSha1TextFn gdc.PtrBuiltInMethod
  methodSha256TextFn gdc.PtrBuiltInMethod
  methodMd5BufferFn gdc.PtrBuiltInMethod
  methodSha1BufferFn gdc.PtrBuiltInMethod
  methodSha256BufferFn gdc.PtrBuiltInMethod
  methodIsEmptyFn gdc.PtrBuiltInMethod
  methodContainsFn gdc.PtrBuiltInMethod
  methodIsAbsolutePathFn gdc.PtrBuiltInMethod
  methodIsRelativePathFn gdc.PtrBuiltInMethod
  methodSimplifyPathFn gdc.PtrBuiltInMethod
  methodGetBaseDirFn gdc.PtrBuiltInMethod
  methodGetFileFn gdc.PtrBuiltInMethod
  methodXmlEscapeFn gdc.PtrBuiltInMethod
  methodXmlUnescapeFn gdc.PtrBuiltInMethod
  methodUriEncodeFn gdc.PtrBuiltInMethod
  methodUriDecodeFn gdc.PtrBuiltInMethod
  methodCEscapeFn gdc.PtrBuiltInMethod
  methodCUnescapeFn gdc.PtrBuiltInMethod
  methodJsonEscapeFn gdc.PtrBuiltInMethod
  methodValidateNodeNameFn gdc.PtrBuiltInMethod
  methodValidateFilenameFn gdc.PtrBuiltInMethod
  methodIsValidIdentifierFn gdc.PtrBuiltInMethod
  methodIsValidIntFn gdc.PtrBuiltInMethod
  methodIsValidFloatFn gdc.PtrBuiltInMethod
  methodIsValidHexNumberFn gdc.PtrBuiltInMethod
  methodIsValidHtmlColorFn gdc.PtrBuiltInMethod
  methodIsValidIpAddressFn gdc.PtrBuiltInMethod
  methodIsValidFilenameFn gdc.PtrBuiltInMethod
  methodToIntFn gdc.PtrBuiltInMethod
  methodToFloatFn gdc.PtrBuiltInMethod
  methodHexToIntFn gdc.PtrBuiltInMethod
  methodBinToIntFn gdc.PtrBuiltInMethod
  methodLpadFn gdc.PtrBuiltInMethod
  methodRpadFn gdc.PtrBuiltInMethod
  methodPadDecimalsFn gdc.PtrBuiltInMethod
  methodPadZerosFn gdc.PtrBuiltInMethod
  methodTrimPrefixFn gdc.PtrBuiltInMethod
  methodTrimSuffixFn gdc.PtrBuiltInMethod
  methodToAsciiBufferFn gdc.PtrBuiltInMethod
  methodToUtf8BufferFn gdc.PtrBuiltInMethod
  methodToUtf16BufferFn gdc.PtrBuiltInMethod
  methodToUtf32BufferFn gdc.PtrBuiltInMethod
  methodHexDecodeFn gdc.PtrBuiltInMethod
  methodToWcharBufferFn gdc.PtrBuiltInMethod
  methodNumScientificFn gdc.PtrBuiltInMethod
  methodNumFn gdc.PtrBuiltInMethod
  methodNumInt64Fn gdc.PtrBuiltInMethod
  methodNumUint64Fn gdc.PtrBuiltInMethod
  methodChrFn gdc.PtrBuiltInMethod
  methodHumanizeSizeFn gdc.PtrBuiltInMethod
  operatorNotFn gdc.PtrOperatorEvaluator
  operatorModuleBoolFn gdc.PtrOperatorEvaluator
  operatorModuleIntFn gdc.PtrOperatorEvaluator
  operatorModuleFloat32Fn gdc.PtrOperatorEvaluator
  operatorEqualStringFn gdc.PtrOperatorEvaluator
  operatorNotEqualStringFn gdc.PtrOperatorEvaluator
  operatorLessStringFn gdc.PtrOperatorEvaluator
  operatorLessEqualStringFn gdc.PtrOperatorEvaluator
  operatorGreaterStringFn gdc.PtrOperatorEvaluator
  operatorGreaterEqualStringFn gdc.PtrOperatorEvaluator
  operatorAddStringFn gdc.PtrOperatorEvaluator
  operatorModuleStringFn gdc.PtrOperatorEvaluator
  operatorInStringFn gdc.PtrOperatorEvaluator
  operatorModuleVector2Fn gdc.PtrOperatorEvaluator
  operatorModuleVector2IFn gdc.PtrOperatorEvaluator
  operatorModuleRect2Fn gdc.PtrOperatorEvaluator
  operatorModuleRect2IFn gdc.PtrOperatorEvaluator
  operatorModuleVector3Fn gdc.PtrOperatorEvaluator
  operatorModuleVector3IFn gdc.PtrOperatorEvaluator
  operatorModuleTransform2DFn gdc.PtrOperatorEvaluator
  operatorModuleVector4Fn gdc.PtrOperatorEvaluator
  operatorModuleVector4IFn gdc.PtrOperatorEvaluator
  operatorModulePlaneFn gdc.PtrOperatorEvaluator
  operatorModuleQuaternionFn gdc.PtrOperatorEvaluator
  operatorModuleAABBFn gdc.PtrOperatorEvaluator
  operatorModuleBasisFn gdc.PtrOperatorEvaluator
  operatorModuleTransform3DFn gdc.PtrOperatorEvaluator
  operatorModuleProjectionFn gdc.PtrOperatorEvaluator
  operatorModuleColorFn gdc.PtrOperatorEvaluator
  operatorEqualStringNameFn gdc.PtrOperatorEvaluator
  operatorNotEqualStringNameFn gdc.PtrOperatorEvaluator
  operatorAddStringNameFn gdc.PtrOperatorEvaluator
  operatorModuleStringNameFn gdc.PtrOperatorEvaluator
  operatorInStringNameFn gdc.PtrOperatorEvaluator
  operatorModuleNodePathFn gdc.PtrOperatorEvaluator
  operatorModuleObjectFn gdc.PtrOperatorEvaluator
  operatorInObjectFn gdc.PtrOperatorEvaluator
  operatorModuleCallableFn gdc.PtrOperatorEvaluator
  operatorModuleSignalFn gdc.PtrOperatorEvaluator
  operatorModuleDictionaryFn gdc.PtrOperatorEvaluator
  operatorInDictionaryFn gdc.PtrOperatorEvaluator
  operatorModuleArrayFn gdc.PtrOperatorEvaluator
  operatorInArrayFn gdc.PtrOperatorEvaluator
  operatorModulePackedByteArrayFn gdc.PtrOperatorEvaluator
  operatorModulePackedInt32ArrayFn gdc.PtrOperatorEvaluator
  operatorModulePackedInt64ArrayFn gdc.PtrOperatorEvaluator
  operatorModulePackedFloat32ArrayFn gdc.PtrOperatorEvaluator
  operatorModulePackedFloat64ArrayFn gdc.PtrOperatorEvaluator
  operatorModulePackedStringArrayFn gdc.PtrOperatorEvaluator
  operatorInPackedStringArrayFn gdc.PtrOperatorEvaluator
  operatorModulePackedVector2ArrayFn gdc.PtrOperatorEvaluator
  operatorModulePackedVector3ArrayFn gdc.PtrOperatorEvaluator
  operatorModulePackedColorArrayFn gdc.PtrOperatorEvaluator
  toVariantFn gdc.TypeFromVariantConstructorFunc
  fromVariantFn gdc.VariantFromTypeConstructorFunc
}

var ptrsForString ptrsForStringList

func initStringPtrs(iface gdc.Interface) {
  ptrsForString.ctrFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeString, 0))
  ptrsForString.ctrFromStringFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeString, 1))
  ptrsForString.ctrFromStringNameFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeString, 2))
  ptrsForString.ctrFromNodePathFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeString, 3))
  ptrsForString.destructorFn = ensurePtr(iface.VariantGetPtrDestructor(gdc.VariantTypeString))
  {
    methodName := StringNameFromStr("casecmp_to")
    defer methodName.Destroy()
    ptrsForString.methodCasecmpToFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 2920860731))
  }
  {
    methodName := StringNameFromStr("nocasecmp_to")
    defer methodName.Destroy()
    ptrsForString.methodNocasecmpToFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 2920860731))
  }
  {
    methodName := StringNameFromStr("naturalcasecmp_to")
    defer methodName.Destroy()
    ptrsForString.methodNaturalcasecmpToFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 2920860731))
  }
  {
    methodName := StringNameFromStr("naturalnocasecmp_to")
    defer methodName.Destroy()
    ptrsForString.methodNaturalnocasecmpToFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 2920860731))
  }
  {
    methodName := StringNameFromStr("length")
    defer methodName.Destroy()
    ptrsForString.methodLengthFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 3173160232))
  }
  {
    methodName := StringNameFromStr("substr")
    defer methodName.Destroy()
    ptrsForString.methodSubstrFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 787537301))
  }
  {
    methodName := StringNameFromStr("get_slice")
    defer methodName.Destroy()
    ptrsForString.methodGetSliceFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 3535100402))
  }
  {
    methodName := StringNameFromStr("get_slicec")
    defer methodName.Destroy()
    ptrsForString.methodGetSlicecFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 787537301))
  }
  {
    methodName := StringNameFromStr("get_slice_count")
    defer methodName.Destroy()
    ptrsForString.methodGetSliceCountFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 2920860731))
  }
  {
    methodName := StringNameFromStr("find")
    defer methodName.Destroy()
    ptrsForString.methodFindFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 1760645412))
  }
  {
    methodName := StringNameFromStr("count")
    defer methodName.Destroy()
    ptrsForString.methodCountFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 2343087891))
  }
  {
    methodName := StringNameFromStr("countn")
    defer methodName.Destroy()
    ptrsForString.methodCountnFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 2343087891))
  }
  {
    methodName := StringNameFromStr("findn")
    defer methodName.Destroy()
    ptrsForString.methodFindnFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 1760645412))
  }
  {
    methodName := StringNameFromStr("rfind")
    defer methodName.Destroy()
    ptrsForString.methodRfindFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 1760645412))
  }
  {
    methodName := StringNameFromStr("rfindn")
    defer methodName.Destroy()
    ptrsForString.methodRfindnFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 1760645412))
  }
  {
    methodName := StringNameFromStr("match")
    defer methodName.Destroy()
    ptrsForString.methodMatchFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 2566493496))
  }
  {
    methodName := StringNameFromStr("matchn")
    defer methodName.Destroy()
    ptrsForString.methodMatchnFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 2566493496))
  }
  {
    methodName := StringNameFromStr("begins_with")
    defer methodName.Destroy()
    ptrsForString.methodBeginsWithFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 2566493496))
  }
  {
    methodName := StringNameFromStr("ends_with")
    defer methodName.Destroy()
    ptrsForString.methodEndsWithFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 2566493496))
  }
  {
    methodName := StringNameFromStr("is_subsequence_of")
    defer methodName.Destroy()
    ptrsForString.methodIsSubsequenceOfFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 2566493496))
  }
  {
    methodName := StringNameFromStr("is_subsequence_ofn")
    defer methodName.Destroy()
    ptrsForString.methodIsSubsequenceOfnFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 2566493496))
  }
  {
    methodName := StringNameFromStr("bigrams")
    defer methodName.Destroy()
    ptrsForString.methodBigramsFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 747180633))
  }
  {
    methodName := StringNameFromStr("similarity")
    defer methodName.Destroy()
    ptrsForString.methodSimilarityFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 2697460964))
  }
  {
    methodName := StringNameFromStr("format")
    defer methodName.Destroy()
    ptrsForString.methodFormatFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 3212199029))
  }
  {
    methodName := StringNameFromStr("replace")
    defer methodName.Destroy()
    ptrsForString.methodReplaceFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 1340436205))
  }
  {
    methodName := StringNameFromStr("replacen")
    defer methodName.Destroy()
    ptrsForString.methodReplacenFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 1340436205))
  }
  {
    methodName := StringNameFromStr("repeat")
    defer methodName.Destroy()
    ptrsForString.methodRepeatFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 2162347432))
  }
  {
    methodName := StringNameFromStr("reverse")
    defer methodName.Destroy()
    ptrsForString.methodReverseFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 3942272618))
  }
  {
    methodName := StringNameFromStr("insert")
    defer methodName.Destroy()
    ptrsForString.methodInsertFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 248737229))
  }
  {
    methodName := StringNameFromStr("erase")
    defer methodName.Destroy()
    ptrsForString.methodEraseFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 787537301))
  }
  {
    methodName := StringNameFromStr("capitalize")
    defer methodName.Destroy()
    ptrsForString.methodCapitalizeFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 3942272618))
  }
  {
    methodName := StringNameFromStr("to_camel_case")
    defer methodName.Destroy()
    ptrsForString.methodToCamelCaseFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 3942272618))
  }
  {
    methodName := StringNameFromStr("to_pascal_case")
    defer methodName.Destroy()
    ptrsForString.methodToPascalCaseFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 3942272618))
  }
  {
    methodName := StringNameFromStr("to_snake_case")
    defer methodName.Destroy()
    ptrsForString.methodToSnakeCaseFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 3942272618))
  }
  {
    methodName := StringNameFromStr("split")
    defer methodName.Destroy()
    ptrsForString.methodSplitFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 1252735785))
  }
  {
    methodName := StringNameFromStr("rsplit")
    defer methodName.Destroy()
    ptrsForString.methodRsplitFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 1252735785))
  }
  {
    methodName := StringNameFromStr("split_floats")
    defer methodName.Destroy()
    ptrsForString.methodSplitFloatsFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 2092079095))
  }
  {
    methodName := StringNameFromStr("join")
    defer methodName.Destroy()
    ptrsForString.methodJoinFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 3595973238))
  }
  {
    methodName := StringNameFromStr("to_upper")
    defer methodName.Destroy()
    ptrsForString.methodToUpperFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 3942272618))
  }
  {
    methodName := StringNameFromStr("to_lower")
    defer methodName.Destroy()
    ptrsForString.methodToLowerFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 3942272618))
  }
  {
    methodName := StringNameFromStr("left")
    defer methodName.Destroy()
    ptrsForString.methodLeftFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 2162347432))
  }
  {
    methodName := StringNameFromStr("right")
    defer methodName.Destroy()
    ptrsForString.methodRightFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 2162347432))
  }
  {
    methodName := StringNameFromStr("strip_edges")
    defer methodName.Destroy()
    ptrsForString.methodStripEdgesFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 907855311))
  }
  {
    methodName := StringNameFromStr("strip_escapes")
    defer methodName.Destroy()
    ptrsForString.methodStripEscapesFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 3942272618))
  }
  {
    methodName := StringNameFromStr("lstrip")
    defer methodName.Destroy()
    ptrsForString.methodLstripFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 3134094431))
  }
  {
    methodName := StringNameFromStr("rstrip")
    defer methodName.Destroy()
    ptrsForString.methodRstripFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 3134094431))
  }
  {
    methodName := StringNameFromStr("get_extension")
    defer methodName.Destroy()
    ptrsForString.methodGetExtensionFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 3942272618))
  }
  {
    methodName := StringNameFromStr("get_basename")
    defer methodName.Destroy()
    ptrsForString.methodGetBasenameFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 3942272618))
  }
  {
    methodName := StringNameFromStr("path_join")
    defer methodName.Destroy()
    ptrsForString.methodPathJoinFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 3134094431))
  }
  {
    methodName := StringNameFromStr("unicode_at")
    defer methodName.Destroy()
    ptrsForString.methodUnicodeAtFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 4103005248))
  }
  {
    methodName := StringNameFromStr("indent")
    defer methodName.Destroy()
    ptrsForString.methodIndentFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 3134094431))
  }
  {
    methodName := StringNameFromStr("dedent")
    defer methodName.Destroy()
    ptrsForString.methodDedentFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 3942272618))
  }
  {
    methodName := StringNameFromStr("hash")
    defer methodName.Destroy()
    ptrsForString.methodHashFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 3173160232))
  }
  {
    methodName := StringNameFromStr("md5_text")
    defer methodName.Destroy()
    ptrsForString.methodMd5TextFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 3942272618))
  }
  {
    methodName := StringNameFromStr("sha1_text")
    defer methodName.Destroy()
    ptrsForString.methodSha1TextFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 3942272618))
  }
  {
    methodName := StringNameFromStr("sha256_text")
    defer methodName.Destroy()
    ptrsForString.methodSha256TextFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 3942272618))
  }
  {
    methodName := StringNameFromStr("md5_buffer")
    defer methodName.Destroy()
    ptrsForString.methodMd5BufferFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 247621236))
  }
  {
    methodName := StringNameFromStr("sha1_buffer")
    defer methodName.Destroy()
    ptrsForString.methodSha1BufferFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 247621236))
  }
  {
    methodName := StringNameFromStr("sha256_buffer")
    defer methodName.Destroy()
    ptrsForString.methodSha256BufferFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 247621236))
  }
  {
    methodName := StringNameFromStr("is_empty")
    defer methodName.Destroy()
    ptrsForString.methodIsEmptyFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 3918633141))
  }
  {
    methodName := StringNameFromStr("contains")
    defer methodName.Destroy()
    ptrsForString.methodContainsFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 2566493496))
  }
  {
    methodName := StringNameFromStr("is_absolute_path")
    defer methodName.Destroy()
    ptrsForString.methodIsAbsolutePathFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 3918633141))
  }
  {
    methodName := StringNameFromStr("is_relative_path")
    defer methodName.Destroy()
    ptrsForString.methodIsRelativePathFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 3918633141))
  }
  {
    methodName := StringNameFromStr("simplify_path")
    defer methodName.Destroy()
    ptrsForString.methodSimplifyPathFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 3942272618))
  }
  {
    methodName := StringNameFromStr("get_base_dir")
    defer methodName.Destroy()
    ptrsForString.methodGetBaseDirFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 3942272618))
  }
  {
    methodName := StringNameFromStr("get_file")
    defer methodName.Destroy()
    ptrsForString.methodGetFileFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 3942272618))
  }
  {
    methodName := StringNameFromStr("xml_escape")
    defer methodName.Destroy()
    ptrsForString.methodXmlEscapeFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 3429816538))
  }
  {
    methodName := StringNameFromStr("xml_unescape")
    defer methodName.Destroy()
    ptrsForString.methodXmlUnescapeFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 3942272618))
  }
  {
    methodName := StringNameFromStr("uri_encode")
    defer methodName.Destroy()
    ptrsForString.methodUriEncodeFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 3942272618))
  }
  {
    methodName := StringNameFromStr("uri_decode")
    defer methodName.Destroy()
    ptrsForString.methodUriDecodeFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 3942272618))
  }
  {
    methodName := StringNameFromStr("c_escape")
    defer methodName.Destroy()
    ptrsForString.methodCEscapeFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 3942272618))
  }
  {
    methodName := StringNameFromStr("c_unescape")
    defer methodName.Destroy()
    ptrsForString.methodCUnescapeFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 3942272618))
  }
  {
    methodName := StringNameFromStr("json_escape")
    defer methodName.Destroy()
    ptrsForString.methodJsonEscapeFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 3942272618))
  }
  {
    methodName := StringNameFromStr("validate_node_name")
    defer methodName.Destroy()
    ptrsForString.methodValidateNodeNameFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 3942272618))
  }
  {
    methodName := StringNameFromStr("validate_filename")
    defer methodName.Destroy()
    ptrsForString.methodValidateFilenameFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 3942272618))
  }
  {
    methodName := StringNameFromStr("is_valid_identifier")
    defer methodName.Destroy()
    ptrsForString.methodIsValidIdentifierFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 3918633141))
  }
  {
    methodName := StringNameFromStr("is_valid_int")
    defer methodName.Destroy()
    ptrsForString.methodIsValidIntFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 3918633141))
  }
  {
    methodName := StringNameFromStr("is_valid_float")
    defer methodName.Destroy()
    ptrsForString.methodIsValidFloatFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 3918633141))
  }
  {
    methodName := StringNameFromStr("is_valid_hex_number")
    defer methodName.Destroy()
    ptrsForString.methodIsValidHexNumberFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 593672999))
  }
  {
    methodName := StringNameFromStr("is_valid_html_color")
    defer methodName.Destroy()
    ptrsForString.methodIsValidHtmlColorFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 3918633141))
  }
  {
    methodName := StringNameFromStr("is_valid_ip_address")
    defer methodName.Destroy()
    ptrsForString.methodIsValidIpAddressFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 3918633141))
  }
  {
    methodName := StringNameFromStr("is_valid_filename")
    defer methodName.Destroy()
    ptrsForString.methodIsValidFilenameFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 3918633141))
  }
  {
    methodName := StringNameFromStr("to_int")
    defer methodName.Destroy()
    ptrsForString.methodToIntFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 3173160232))
  }
  {
    methodName := StringNameFromStr("to_float")
    defer methodName.Destroy()
    ptrsForString.methodToFloatFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 466405837))
  }
  {
    methodName := StringNameFromStr("hex_to_int")
    defer methodName.Destroy()
    ptrsForString.methodHexToIntFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 3173160232))
  }
  {
    methodName := StringNameFromStr("bin_to_int")
    defer methodName.Destroy()
    ptrsForString.methodBinToIntFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 3173160232))
  }
  {
    methodName := StringNameFromStr("lpad")
    defer methodName.Destroy()
    ptrsForString.methodLpadFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 248737229))
  }
  {
    methodName := StringNameFromStr("rpad")
    defer methodName.Destroy()
    ptrsForString.methodRpadFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 248737229))
  }
  {
    methodName := StringNameFromStr("pad_decimals")
    defer methodName.Destroy()
    ptrsForString.methodPadDecimalsFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 2162347432))
  }
  {
    methodName := StringNameFromStr("pad_zeros")
    defer methodName.Destroy()
    ptrsForString.methodPadZerosFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 2162347432))
  }
  {
    methodName := StringNameFromStr("trim_prefix")
    defer methodName.Destroy()
    ptrsForString.methodTrimPrefixFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 3134094431))
  }
  {
    methodName := StringNameFromStr("trim_suffix")
    defer methodName.Destroy()
    ptrsForString.methodTrimSuffixFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 3134094431))
  }
  {
    methodName := StringNameFromStr("to_ascii_buffer")
    defer methodName.Destroy()
    ptrsForString.methodToAsciiBufferFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 247621236))
  }
  {
    methodName := StringNameFromStr("to_utf8_buffer")
    defer methodName.Destroy()
    ptrsForString.methodToUtf8BufferFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 247621236))
  }
  {
    methodName := StringNameFromStr("to_utf16_buffer")
    defer methodName.Destroy()
    ptrsForString.methodToUtf16BufferFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 247621236))
  }
  {
    methodName := StringNameFromStr("to_utf32_buffer")
    defer methodName.Destroy()
    ptrsForString.methodToUtf32BufferFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 247621236))
  }
  {
    methodName := StringNameFromStr("hex_decode")
    defer methodName.Destroy()
    ptrsForString.methodHexDecodeFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 247621236))
  }
  {
    methodName := StringNameFromStr("to_wchar_buffer")
    defer methodName.Destroy()
    ptrsForString.methodToWcharBufferFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 247621236))
  }
  {
    methodName := StringNameFromStr("num_scientific")
    defer methodName.Destroy()
    ptrsForString.methodNumScientificFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 2710373411))
  }
  {
    methodName := StringNameFromStr("num")
    defer methodName.Destroy()
    ptrsForString.methodNumFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 1555901022))
  }
  {
    methodName := StringNameFromStr("num_int64")
    defer methodName.Destroy()
    ptrsForString.methodNumInt64Fn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 2111271071))
  }
  {
    methodName := StringNameFromStr("num_uint64")
    defer methodName.Destroy()
    ptrsForString.methodNumUint64Fn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 2111271071))
  }
  {
    methodName := StringNameFromStr("chr")
    defer methodName.Destroy()
    ptrsForString.methodChrFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 897497541))
  }
  {
    methodName := StringNameFromStr("humanize_size")
    defer methodName.Destroy()
    ptrsForString.methodHumanizeSizeFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeString, methodName.AsCPtr(), 897497541))
  }
  ptrsForString.operatorNotFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, gdc.VariantTypeString, gdc.VariantTypeNil))
  ptrsForString.operatorModuleBoolFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeString, gdc.VariantTypeBool))
  ptrsForString.operatorModuleIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeString, gdc.VariantTypeInt))
  ptrsForString.operatorModuleFloat32Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeString, gdc.VariantTypeFloat))
  ptrsForString.operatorEqualStringFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeString, gdc.VariantTypeString))
  ptrsForString.operatorNotEqualStringFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeString, gdc.VariantTypeString))
  ptrsForString.operatorLessStringFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpLess, gdc.VariantTypeString, gdc.VariantTypeString))
  ptrsForString.operatorLessEqualStringFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpLessEqual, gdc.VariantTypeString, gdc.VariantTypeString))
  ptrsForString.operatorGreaterStringFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpGreater, gdc.VariantTypeString, gdc.VariantTypeString))
  ptrsForString.operatorGreaterEqualStringFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpGreaterEqual, gdc.VariantTypeString, gdc.VariantTypeString))
  ptrsForString.operatorAddStringFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpAdd, gdc.VariantTypeString, gdc.VariantTypeString))
  ptrsForString.operatorModuleStringFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeString, gdc.VariantTypeString))
  ptrsForString.operatorInStringFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeString, gdc.VariantTypeString))
  ptrsForString.operatorModuleVector2Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeString, gdc.VariantTypeVector2))
  ptrsForString.operatorModuleVector2IFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeString, gdc.VariantTypeVector2I))
  ptrsForString.operatorModuleRect2Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeString, gdc.VariantTypeRect2))
  ptrsForString.operatorModuleRect2IFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeString, gdc.VariantTypeRect2I))
  ptrsForString.operatorModuleVector3Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeString, gdc.VariantTypeVector3))
  ptrsForString.operatorModuleVector3IFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeString, gdc.VariantTypeVector3I))
  ptrsForString.operatorModuleTransform2DFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeString, gdc.VariantTypeTransform2D))
  ptrsForString.operatorModuleVector4Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeString, gdc.VariantTypeVector4))
  ptrsForString.operatorModuleVector4IFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeString, gdc.VariantTypeVector4I))
  ptrsForString.operatorModulePlaneFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeString, gdc.VariantTypePlane))
  ptrsForString.operatorModuleQuaternionFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeString, gdc.VariantTypeQuaternion))
  ptrsForString.operatorModuleAABBFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeString, gdc.VariantTypeAABB))
  ptrsForString.operatorModuleBasisFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeString, gdc.VariantTypeBasis))
  ptrsForString.operatorModuleTransform3DFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeString, gdc.VariantTypeTransform3D))
  ptrsForString.operatorModuleProjectionFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeString, gdc.VariantTypeProjection))
  ptrsForString.operatorModuleColorFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeString, gdc.VariantTypeColor))
  ptrsForString.operatorEqualStringNameFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeString, gdc.VariantTypeStringName))
  ptrsForString.operatorNotEqualStringNameFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeString, gdc.VariantTypeStringName))
  ptrsForString.operatorAddStringNameFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpAdd, gdc.VariantTypeString, gdc.VariantTypeStringName))
  ptrsForString.operatorModuleStringNameFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeString, gdc.VariantTypeStringName))
  ptrsForString.operatorInStringNameFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeString, gdc.VariantTypeStringName))
  ptrsForString.operatorModuleNodePathFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeString, gdc.VariantTypeNodePath))
  ptrsForString.operatorModuleObjectFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeString, gdc.VariantTypeObject))
  ptrsForString.operatorInObjectFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeString, gdc.VariantTypeObject))
  ptrsForString.operatorModuleCallableFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeString, gdc.VariantTypeCallable))
  ptrsForString.operatorModuleSignalFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeString, gdc.VariantTypeSignal))
  ptrsForString.operatorModuleDictionaryFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeString, gdc.VariantTypeDictionary))
  ptrsForString.operatorInDictionaryFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeString, gdc.VariantTypeDictionary))
  ptrsForString.operatorModuleArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeString, gdc.VariantTypeArray))
  ptrsForString.operatorInArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeString, gdc.VariantTypeArray))
  ptrsForString.operatorModulePackedByteArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeString, gdc.VariantTypePackedByteArray))
  ptrsForString.operatorModulePackedInt32ArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeString, gdc.VariantTypePackedInt32Array))
  ptrsForString.operatorModulePackedInt64ArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeString, gdc.VariantTypePackedInt64Array))
  ptrsForString.operatorModulePackedFloat32ArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeString, gdc.VariantTypePackedFloat32Array))
  ptrsForString.operatorModulePackedFloat64ArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeString, gdc.VariantTypePackedFloat64Array))
  ptrsForString.operatorModulePackedStringArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeString, gdc.VariantTypePackedStringArray))
  ptrsForString.operatorInPackedStringArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeString, gdc.VariantTypePackedStringArray))
  ptrsForString.operatorModulePackedVector2ArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeString, gdc.VariantTypePackedVector2Array))
  ptrsForString.operatorModulePackedVector3ArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeString, gdc.VariantTypePackedVector3Array))
  ptrsForString.operatorModulePackedColorArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeString, gdc.VariantTypePackedColorArray))
  ptrsForString.toVariantFn = ensurePtr(iface.GetVariantToTypeConstructor(gdc.VariantTypeString))
  ptrsForString.fromVariantFn = ensurePtr(iface.GetVariantFromTypeConstructor(gdc.VariantTypeString))
}

type String struct {
  data   *[classSizeString]byte
  iface  gdc.Interface
  pinner runtime.Pinner
}

// Enums

// Constructors
func newString() *String {
  me := &String{
    data:   new([classSizeString]byte),
    iface:  giface,
  }
  me.pinner.Pin(me)
  me.pinner.Pin(me.AsTypePtr())
  return me
}

func NewString() *String {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newString()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForString.ctrFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{}))
  return me
}

func NewStringFromString(from String, ) *String {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newString()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForString.ctrFromStringFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return me
}

func NewStringFromStringName(from StringName, ) *String {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newString()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForString.ctrFromStringNameFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return me
}

func NewStringFromNodePath(from NodePath, ) *String {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newString()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForString.ctrFromNodePathFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return me
}

// Destructor
func (me *String) Destroy() {
	me.iface.CallPtrDestructor(ensurePtr(ptrsForString.destructorFn), me.AsTypePtr())
  me.pinner.Unpin()
}

// Variant support
func (me *Variant) AsString() (*String, error) {
	if me.Type() != gdc.VariantTypeString {
		return nil, fmt.Errorf("variant is not a String")
	}
  bclass := newString()
	me.iface.CallTypeFromVariantConstructorFunc(ensurePtr(ptrsForString.toVariantFn), bclass.asUninitialized(), me.AsPtr())
	return bclass, nil
}

func (me *String) AsVariant() *Variant {
  va := newVariant()
  va.inner = me
  me.iface.CallVariantFromTypeConstructorFunc(ensurePtr(ptrsForString.fromVariantFn), va.asUninitialized(), me.AsTypePtr())
  return va
}

// Pointers
func StringFromPtr(ptr gdc.ConstTypePtr) *String {
  me := newString()
  dataFromPtr(me.data[:], ptr)
  return me
}

func (me *String) Type() gdc.VariantType {
  return gdc.VariantTypeString
}

func (me *String) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(unsafe.Pointer(me.data))
}

func (me *String) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.AsTypePtr())
}

func (me *String) asUninitialized() gdc.UninitializedTypePtr {
  return gdc.UninitializedTypePtr(me.AsTypePtr())
}

// Methods

func (me *String) CasecmpTo(to String, ) int64 {
  ret := NewInt()
  defer ret.Destroy()

  args := []gdc.ConstTypePtr{to.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodCasecmpToFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *String) NocasecmpTo(to String, ) int64 {
  ret := NewInt()
  defer ret.Destroy()

  args := []gdc.ConstTypePtr{to.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodNocasecmpToFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *String) NaturalcasecmpTo(to String, ) int64 {
  ret := NewInt()
  defer ret.Destroy()

  args := []gdc.ConstTypePtr{to.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodNaturalcasecmpToFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *String) NaturalnocasecmpTo(to String, ) int64 {
  ret := NewInt()
  defer ret.Destroy()

  args := []gdc.ConstTypePtr{to.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodNaturalnocasecmpToFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *String) Length() int64 {
  ret := NewInt()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodLengthFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *String) Substr(from int64, len_ int64, ) String {
  ret := NewString()

  varg0 := NewIntFromInt(from)
  defer varg0.Destroy()
  varg1 := NewIntFromInt(len_)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodSubstrFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *String) GetSlice(delimiter String, slice int64, ) String {
  ret := NewString()


  varg1 := NewIntFromInt(slice)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{delimiter.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodGetSliceFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *String) GetSlicec(delimiter int64, slice int64, ) String {
  ret := NewString()

  varg0 := NewIntFromInt(delimiter)
  defer varg0.Destroy()
  varg1 := NewIntFromInt(slice)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodGetSlicecFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *String) GetSliceCount(delimiter String, ) int64 {
  ret := NewInt()
  defer ret.Destroy()

  args := []gdc.ConstTypePtr{delimiter.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodGetSliceCountFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *String) Find(what String, from int64, ) int64 {
  ret := NewInt()
  defer ret.Destroy()

  varg1 := NewIntFromInt(from)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{what.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodFindFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *String) Count(what String, from int64, to int64, ) int64 {
  ret := NewInt()
  defer ret.Destroy()

  varg1 := NewIntFromInt(from)
  defer varg1.Destroy()
  varg2 := NewIntFromInt(to)
  defer varg2.Destroy()
  args := []gdc.ConstTypePtr{what.AsCTypePtr(), varg1.AsCTypePtr(), varg2.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodCountFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *String) Countn(what String, from int64, to int64, ) int64 {
  ret := NewInt()
  defer ret.Destroy()

  varg1 := NewIntFromInt(from)
  defer varg1.Destroy()
  varg2 := NewIntFromInt(to)
  defer varg2.Destroy()
  args := []gdc.ConstTypePtr{what.AsCTypePtr(), varg1.AsCTypePtr(), varg2.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodCountnFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *String) Findn(what String, from int64, ) int64 {
  ret := NewInt()
  defer ret.Destroy()

  varg1 := NewIntFromInt(from)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{what.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodFindnFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *String) Rfind(what String, from int64, ) int64 {
  ret := NewInt()
  defer ret.Destroy()

  varg1 := NewIntFromInt(from)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{what.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodRfindFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *String) Rfindn(what String, from int64, ) int64 {
  ret := NewInt()
  defer ret.Destroy()

  varg1 := NewIntFromInt(from)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{what.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodRfindnFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *String) Match(expr String, ) bool {
  ret := NewBool()
  defer ret.Destroy()

  args := []gdc.ConstTypePtr{expr.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodMatchFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *String) Matchn(expr String, ) bool {
  ret := NewBool()
  defer ret.Destroy()

  args := []gdc.ConstTypePtr{expr.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodMatchnFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *String) BeginsWith(text String, ) bool {
  ret := NewBool()
  defer ret.Destroy()

  args := []gdc.ConstTypePtr{text.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodBeginsWithFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *String) EndsWith(text String, ) bool {
  ret := NewBool()
  defer ret.Destroy()

  args := []gdc.ConstTypePtr{text.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodEndsWithFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *String) IsSubsequenceOf(text String, ) bool {
  ret := NewBool()
  defer ret.Destroy()

  args := []gdc.ConstTypePtr{text.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodIsSubsequenceOfFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *String) IsSubsequenceOfn(text String, ) bool {
  ret := NewBool()
  defer ret.Destroy()

  args := []gdc.ConstTypePtr{text.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodIsSubsequenceOfnFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *String) Bigrams() PackedStringArray {
  ret := NewPackedStringArray()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodBigramsFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *String) Similarity(text String, ) float64 {
  ret := NewFloat()
  defer ret.Destroy()

  args := []gdc.ConstTypePtr{text.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodSimilarityFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *String) Format(values Variant, placeholder String, ) String {
  ret := NewString()



  args := []gdc.ConstTypePtr{values.AsCTypePtr(), placeholder.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodFormatFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *String) Replace(what String, forwhat String, ) String {
  ret := NewString()



  args := []gdc.ConstTypePtr{what.AsCTypePtr(), forwhat.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodReplaceFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *String) Replacen(what String, forwhat String, ) String {
  ret := NewString()



  args := []gdc.ConstTypePtr{what.AsCTypePtr(), forwhat.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodReplacenFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *String) Repeat(count int64, ) String {
  ret := NewString()

  varg0 := NewIntFromInt(count)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodRepeatFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *String) Reverse() String {
  ret := NewString()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodReverseFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *String) Insert(position int64, what String, ) String {
  ret := NewString()

  varg0 := NewIntFromInt(position)
  defer varg0.Destroy()

  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), what.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodInsertFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *String) Erase(position int64, chars int64, ) String {
  ret := NewString()

  varg0 := NewIntFromInt(position)
  defer varg0.Destroy()
  varg1 := NewIntFromInt(chars)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodEraseFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *String) Capitalize() String {
  ret := NewString()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodCapitalizeFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *String) ToCamelCase() String {
  ret := NewString()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodToCamelCaseFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *String) ToPascalCase() String {
  ret := NewString()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodToPascalCaseFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *String) ToSnakeCase() String {
  ret := NewString()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodToSnakeCaseFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *String) Split(delimiter String, allow_empty bool, maxsplit int64, ) PackedStringArray {
  ret := NewPackedStringArray()


  varg1 := NewBoolFromBool(allow_empty)
  defer varg1.Destroy()
  varg2 := NewIntFromInt(maxsplit)
  defer varg2.Destroy()
  args := []gdc.ConstTypePtr{delimiter.AsCTypePtr(), varg1.AsCTypePtr(), varg2.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodSplitFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *String) Rsplit(delimiter String, allow_empty bool, maxsplit int64, ) PackedStringArray {
  ret := NewPackedStringArray()


  varg1 := NewBoolFromBool(allow_empty)
  defer varg1.Destroy()
  varg2 := NewIntFromInt(maxsplit)
  defer varg2.Destroy()
  args := []gdc.ConstTypePtr{delimiter.AsCTypePtr(), varg1.AsCTypePtr(), varg2.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodRsplitFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *String) SplitFloats(delimiter String, allow_empty bool, ) PackedFloat64Array {
  ret := NewPackedFloat64Array()


  varg1 := NewBoolFromBool(allow_empty)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{delimiter.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodSplitFloatsFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *String) Join(parts PackedStringArray, ) String {
  ret := NewString()


  args := []gdc.ConstTypePtr{parts.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodJoinFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *String) ToUpper() String {
  ret := NewString()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodToUpperFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *String) ToLower() String {
  ret := NewString()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodToLowerFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *String) Left(length int64, ) String {
  ret := NewString()

  varg0 := NewIntFromInt(length)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodLeftFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *String) Right(length int64, ) String {
  ret := NewString()

  varg0 := NewIntFromInt(length)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodRightFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *String) StripEdges(left bool, right bool, ) String {
  ret := NewString()

  varg0 := NewBoolFromBool(left)
  defer varg0.Destroy()
  varg1 := NewBoolFromBool(right)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodStripEdgesFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *String) StripEscapes() String {
  ret := NewString()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodStripEscapesFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *String) Lstrip(chars String, ) String {
  ret := NewString()


  args := []gdc.ConstTypePtr{chars.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodLstripFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *String) Rstrip(chars String, ) String {
  ret := NewString()


  args := []gdc.ConstTypePtr{chars.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodRstripFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *String) GetExtension() String {
  ret := NewString()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodGetExtensionFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *String) GetBasename() String {
  ret := NewString()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodGetBasenameFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *String) PathJoin(file String, ) String {
  ret := NewString()


  args := []gdc.ConstTypePtr{file.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodPathJoinFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *String) UnicodeAt(at int64, ) int64 {
  ret := NewInt()
  defer ret.Destroy()
  varg0 := NewIntFromInt(at)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodUnicodeAtFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *String) Indent(prefix String, ) String {
  ret := NewString()


  args := []gdc.ConstTypePtr{prefix.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodIndentFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *String) Dedent() String {
  ret := NewString()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodDedentFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *String) Hash() int64 {
  ret := NewInt()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodHashFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *String) Md5Text() String {
  ret := NewString()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodMd5TextFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *String) Sha1Text() String {
  ret := NewString()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodSha1TextFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *String) Sha256Text() String {
  ret := NewString()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodSha256TextFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *String) Md5Buffer() PackedByteArray {
  ret := NewPackedByteArray()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodMd5BufferFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *String) Sha1Buffer() PackedByteArray {
  ret := NewPackedByteArray()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodSha1BufferFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *String) Sha256Buffer() PackedByteArray {
  ret := NewPackedByteArray()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodSha256BufferFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *String) IsEmpty() bool {
  ret := NewBool()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodIsEmptyFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *String) Contains(what String, ) bool {
  ret := NewBool()
  defer ret.Destroy()

  args := []gdc.ConstTypePtr{what.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodContainsFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *String) IsAbsolutePath() bool {
  ret := NewBool()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodIsAbsolutePathFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *String) IsRelativePath() bool {
  ret := NewBool()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodIsRelativePathFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *String) SimplifyPath() String {
  ret := NewString()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodSimplifyPathFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *String) GetBaseDir() String {
  ret := NewString()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodGetBaseDirFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *String) GetFile() String {
  ret := NewString()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodGetFileFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *String) XmlEscape(escape_quotes bool, ) String {
  ret := NewString()

  varg0 := NewBoolFromBool(escape_quotes)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodXmlEscapeFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *String) XmlUnescape() String {
  ret := NewString()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodXmlUnescapeFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *String) UriEncode() String {
  ret := NewString()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodUriEncodeFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *String) UriDecode() String {
  ret := NewString()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodUriDecodeFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *String) CEscape() String {
  ret := NewString()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodCEscapeFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *String) CUnescape() String {
  ret := NewString()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodCUnescapeFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *String) JsonEscape() String {
  ret := NewString()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodJsonEscapeFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *String) ValidateNodeName() String {
  ret := NewString()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodValidateNodeNameFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *String) ValidateFilename() String {
  ret := NewString()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodValidateFilenameFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *String) IsValidIdentifier() bool {
  ret := NewBool()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodIsValidIdentifierFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *String) IsValidInt() bool {
  ret := NewBool()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodIsValidIntFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *String) IsValidFloat() bool {
  ret := NewBool()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodIsValidFloatFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *String) IsValidHexNumber(with_prefix bool, ) bool {
  ret := NewBool()
  defer ret.Destroy()
  varg0 := NewBoolFromBool(with_prefix)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodIsValidHexNumberFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *String) IsValidHtmlColor() bool {
  ret := NewBool()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodIsValidHtmlColorFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *String) IsValidIpAddress() bool {
  ret := NewBool()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodIsValidIpAddressFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *String) IsValidFilename() bool {
  ret := NewBool()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodIsValidFilenameFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *String) ToInt() int64 {
  ret := NewInt()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodToIntFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *String) ToFloat() float64 {
  ret := NewFloat()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodToFloatFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *String) HexToInt() int64 {
  ret := NewInt()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodHexToIntFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *String) BinToInt() int64 {
  ret := NewInt()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodBinToIntFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *String) Lpad(min_length int64, character String, ) String {
  ret := NewString()

  varg0 := NewIntFromInt(min_length)
  defer varg0.Destroy()

  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), character.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodLpadFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *String) Rpad(min_length int64, character String, ) String {
  ret := NewString()

  varg0 := NewIntFromInt(min_length)
  defer varg0.Destroy()

  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), character.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodRpadFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *String) PadDecimals(digits int64, ) String {
  ret := NewString()

  varg0 := NewIntFromInt(digits)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodPadDecimalsFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *String) PadZeros(digits int64, ) String {
  ret := NewString()

  varg0 := NewIntFromInt(digits)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodPadZerosFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *String) TrimPrefix(prefix String, ) String {
  ret := NewString()


  args := []gdc.ConstTypePtr{prefix.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodTrimPrefixFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *String) TrimSuffix(suffix String, ) String {
  ret := NewString()


  args := []gdc.ConstTypePtr{suffix.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodTrimSuffixFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *String) ToAsciiBuffer() PackedByteArray {
  ret := NewPackedByteArray()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodToAsciiBufferFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *String) ToUtf8Buffer() PackedByteArray {
  ret := NewPackedByteArray()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodToUtf8BufferFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *String) ToUtf16Buffer() PackedByteArray {
  ret := NewPackedByteArray()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodToUtf16BufferFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *String) ToUtf32Buffer() PackedByteArray {
  ret := NewPackedByteArray()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodToUtf32BufferFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *String) HexDecode() PackedByteArray {
  ret := NewPackedByteArray()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodHexDecodeFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *String) ToWcharBuffer() PackedByteArray {
  ret := NewPackedByteArray()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodToWcharBufferFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func StringNumScientific(number float64, ) String {
  ret := NewString()

  varg0 := NewFloatFromFloat32(number)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodNumScientificFn), nil, unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func StringNum(number float64, decimals int64, ) String {
  ret := NewString()

  varg0 := NewFloatFromFloat32(number)
  defer varg0.Destroy()
  varg1 := NewIntFromInt(decimals)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodNumFn), nil, unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func StringNumInt64(number int64, base int64, capitalize_hex bool, ) String {
  ret := NewString()

  varg0 := NewIntFromInt(number)
  defer varg0.Destroy()
  varg1 := NewIntFromInt(base)
  defer varg1.Destroy()
  varg2 := NewBoolFromBool(capitalize_hex)
  defer varg2.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), varg2.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodNumInt64Fn), nil, unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func StringNumUint64(number int64, base int64, capitalize_hex bool, ) String {
  ret := NewString()

  varg0 := NewIntFromInt(number)
  defer varg0.Destroy()
  varg1 := NewIntFromInt(base)
  defer varg1.Destroy()
  varg2 := NewBoolFromBool(capitalize_hex)
  defer varg2.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), varg2.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodNumUint64Fn), nil, unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func StringChr(char int64, ) String {
  ret := NewString()

  varg0 := NewIntFromInt(char)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodChrFn), nil, unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func StringHumanizeSize(size int64, ) String {
  ret := NewString()

  varg0 := NewIntFromInt(size)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForString.methodHumanizeSizeFn), nil, unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

// Operators

func (me *String) EqualVariant(right Variant) bool {
  opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type())
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *String) NotEqualVariant(right Variant) bool {
  opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type())
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *String) ModuleVariant(right Variant) String {
  opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type())
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *String) Not() bool {
  opPtr := ptrsForString.operatorNotFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), nil, ret.AsTypePtr())
  return ret.Get()
}

func (me *String) ModuleBool(rightArg bool) String {
  right := NewBoolFromBool(rightArg)
  defer right.Destroy()

  opPtr := ptrsForString.operatorModuleBoolFn
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *String) ModuleInt(rightArg int64) String {
  right := NewIntFromInt(rightArg)
  defer right.Destroy()

  opPtr := ptrsForString.operatorModuleIntFn
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *String) ModuleFloat32(rightArg float64) String {
  right := NewFloatFromFloat32(rightArg)
  defer right.Destroy()

  opPtr := ptrsForString.operatorModuleFloat32Fn
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *String) EqualString(right String) bool {
  opPtr := ptrsForString.operatorEqualStringFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *String) NotEqualString(right String) bool {
  opPtr := ptrsForString.operatorNotEqualStringFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *String) LessString(right String) bool {
  opPtr := ptrsForString.operatorLessStringFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *String) LessEqualString(right String) bool {
  opPtr := ptrsForString.operatorLessEqualStringFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *String) GreaterString(right String) bool {
  opPtr := ptrsForString.operatorGreaterStringFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *String) GreaterEqualString(right String) bool {
  opPtr := ptrsForString.operatorGreaterEqualStringFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *String) AddString(right String) String {
  opPtr := ptrsForString.operatorAddStringFn
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *String) ModuleString(right String) String {
  opPtr := ptrsForString.operatorModuleStringFn
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *String) InString(right String) bool {
  opPtr := ptrsForString.operatorInStringFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *String) ModuleVector2(right Vector2) String {
  opPtr := ptrsForString.operatorModuleVector2Fn
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *String) ModuleVector2I(right Vector2i) String {
  opPtr := ptrsForString.operatorModuleVector2IFn
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *String) ModuleRect2(right Rect2) String {
  opPtr := ptrsForString.operatorModuleRect2Fn
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *String) ModuleRect2I(right Rect2i) String {
  opPtr := ptrsForString.operatorModuleRect2IFn
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *String) ModuleVector3(right Vector3) String {
  opPtr := ptrsForString.operatorModuleVector3Fn
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *String) ModuleVector3I(right Vector3i) String {
  opPtr := ptrsForString.operatorModuleVector3IFn
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *String) ModuleTransform2D(right Transform2D) String {
  opPtr := ptrsForString.operatorModuleTransform2DFn
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *String) ModuleVector4(right Vector4) String {
  opPtr := ptrsForString.operatorModuleVector4Fn
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *String) ModuleVector4I(right Vector4i) String {
  opPtr := ptrsForString.operatorModuleVector4IFn
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *String) ModulePlane(right Plane) String {
  opPtr := ptrsForString.operatorModulePlaneFn
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *String) ModuleQuaternion(right Quaternion) String {
  opPtr := ptrsForString.operatorModuleQuaternionFn
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *String) ModuleAABB(right AABB) String {
  opPtr := ptrsForString.operatorModuleAABBFn
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *String) ModuleBasis(right Basis) String {
  opPtr := ptrsForString.operatorModuleBasisFn
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *String) ModuleTransform3D(right Transform3D) String {
  opPtr := ptrsForString.operatorModuleTransform3DFn
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *String) ModuleProjection(right Projection) String {
  opPtr := ptrsForString.operatorModuleProjectionFn
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *String) ModuleColor(right Color) String {
  opPtr := ptrsForString.operatorModuleColorFn
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *String) EqualStringName(right StringName) bool {
  opPtr := ptrsForString.operatorEqualStringNameFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *String) NotEqualStringName(right StringName) bool {
  opPtr := ptrsForString.operatorNotEqualStringNameFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *String) AddStringName(right StringName) String {
  opPtr := ptrsForString.operatorAddStringNameFn
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *String) ModuleStringName(right StringName) String {
  opPtr := ptrsForString.operatorModuleStringNameFn
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *String) InStringName(right StringName) bool {
  opPtr := ptrsForString.operatorInStringNameFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *String) ModuleNodePath(right NodePath) String {
  opPtr := ptrsForString.operatorModuleNodePathFn
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *String) ModuleObject(right Object) String {
  opPtr := ptrsForString.operatorModuleObjectFn
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *String) InObject(right Object) bool {
  opPtr := ptrsForString.operatorInObjectFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *String) ModuleCallable(right Callable) String {
  opPtr := ptrsForString.operatorModuleCallableFn
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *String) ModuleSignal(right Signal) String {
  opPtr := ptrsForString.operatorModuleSignalFn
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *String) ModuleDictionary(right Dictionary) String {
  opPtr := ptrsForString.operatorModuleDictionaryFn
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *String) InDictionary(right Dictionary) bool {
  opPtr := ptrsForString.operatorInDictionaryFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *String) ModuleArray(right Array) String {
  opPtr := ptrsForString.operatorModuleArrayFn
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *String) InArray(right Array) bool {
  opPtr := ptrsForString.operatorInArrayFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *String) ModulePackedByteArray(right PackedByteArray) String {
  opPtr := ptrsForString.operatorModulePackedByteArrayFn
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *String) ModulePackedInt32Array(right PackedInt32Array) String {
  opPtr := ptrsForString.operatorModulePackedInt32ArrayFn
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *String) ModulePackedInt64Array(right PackedInt64Array) String {
  opPtr := ptrsForString.operatorModulePackedInt64ArrayFn
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *String) ModulePackedFloat32Array(right PackedFloat32Array) String {
  opPtr := ptrsForString.operatorModulePackedFloat32ArrayFn
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *String) ModulePackedFloat64Array(right PackedFloat64Array) String {
  opPtr := ptrsForString.operatorModulePackedFloat64ArrayFn
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *String) ModulePackedStringArray(right PackedStringArray) String {
  opPtr := ptrsForString.operatorModulePackedStringArrayFn
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *String) InPackedStringArray(right PackedStringArray) bool {
  opPtr := ptrsForString.operatorInPackedStringArrayFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *String) ModulePackedVector2Array(right PackedVector2Array) String {
  opPtr := ptrsForString.operatorModulePackedVector2ArrayFn
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *String) ModulePackedVector3Array(right PackedVector3Array) String {
  opPtr := ptrsForString.operatorModulePackedVector3ArrayFn
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *String) ModulePackedColorArray(right PackedColorArray) String {
  opPtr := ptrsForString.operatorModulePackedColorArrayFn
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

// Members
