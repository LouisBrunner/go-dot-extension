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

type ptrsForStringNameList struct {
	ctrFn                              gdc.PtrConstructor
	ctrFromStringNameFn                gdc.PtrConstructor
	ctrFromStringFn                    gdc.PtrConstructor
	destructorFn                       gdc.PtrDestructor
	methodCasecmpToFn                  gdc.PtrBuiltInMethod
	methodNocasecmpToFn                gdc.PtrBuiltInMethod
	methodNaturalcasecmpToFn           gdc.PtrBuiltInMethod
	methodNaturalnocasecmpToFn         gdc.PtrBuiltInMethod
	methodLengthFn                     gdc.PtrBuiltInMethod
	methodSubstrFn                     gdc.PtrBuiltInMethod
	methodGetSliceFn                   gdc.PtrBuiltInMethod
	methodGetSlicecFn                  gdc.PtrBuiltInMethod
	methodGetSliceCountFn              gdc.PtrBuiltInMethod
	methodFindFn                       gdc.PtrBuiltInMethod
	methodCountFn                      gdc.PtrBuiltInMethod
	methodCountnFn                     gdc.PtrBuiltInMethod
	methodFindnFn                      gdc.PtrBuiltInMethod
	methodRfindFn                      gdc.PtrBuiltInMethod
	methodRfindnFn                     gdc.PtrBuiltInMethod
	methodMatchFn                      gdc.PtrBuiltInMethod
	methodMatchnFn                     gdc.PtrBuiltInMethod
	methodBeginsWithFn                 gdc.PtrBuiltInMethod
	methodEndsWithFn                   gdc.PtrBuiltInMethod
	methodIsSubsequenceOfFn            gdc.PtrBuiltInMethod
	methodIsSubsequenceOfnFn           gdc.PtrBuiltInMethod
	methodBigramsFn                    gdc.PtrBuiltInMethod
	methodSimilarityFn                 gdc.PtrBuiltInMethod
	methodFormatFn                     gdc.PtrBuiltInMethod
	methodReplaceFn                    gdc.PtrBuiltInMethod
	methodReplacenFn                   gdc.PtrBuiltInMethod
	methodRepeatFn                     gdc.PtrBuiltInMethod
	methodReverseFn                    gdc.PtrBuiltInMethod
	methodInsertFn                     gdc.PtrBuiltInMethod
	methodEraseFn                      gdc.PtrBuiltInMethod
	methodCapitalizeFn                 gdc.PtrBuiltInMethod
	methodToCamelCaseFn                gdc.PtrBuiltInMethod
	methodToPascalCaseFn               gdc.PtrBuiltInMethod
	methodToSnakeCaseFn                gdc.PtrBuiltInMethod
	methodSplitFn                      gdc.PtrBuiltInMethod
	methodRsplitFn                     gdc.PtrBuiltInMethod
	methodSplitFloatsFn                gdc.PtrBuiltInMethod
	methodJoinFn                       gdc.PtrBuiltInMethod
	methodToUpperFn                    gdc.PtrBuiltInMethod
	methodToLowerFn                    gdc.PtrBuiltInMethod
	methodLeftFn                       gdc.PtrBuiltInMethod
	methodRightFn                      gdc.PtrBuiltInMethod
	methodStripEdgesFn                 gdc.PtrBuiltInMethod
	methodStripEscapesFn               gdc.PtrBuiltInMethod
	methodLstripFn                     gdc.PtrBuiltInMethod
	methodRstripFn                     gdc.PtrBuiltInMethod
	methodGetExtensionFn               gdc.PtrBuiltInMethod
	methodGetBasenameFn                gdc.PtrBuiltInMethod
	methodPathJoinFn                   gdc.PtrBuiltInMethod
	methodUnicodeAtFn                  gdc.PtrBuiltInMethod
	methodIndentFn                     gdc.PtrBuiltInMethod
	methodDedentFn                     gdc.PtrBuiltInMethod
	methodMd5TextFn                    gdc.PtrBuiltInMethod
	methodSha1TextFn                   gdc.PtrBuiltInMethod
	methodSha256TextFn                 gdc.PtrBuiltInMethod
	methodMd5BufferFn                  gdc.PtrBuiltInMethod
	methodSha1BufferFn                 gdc.PtrBuiltInMethod
	methodSha256BufferFn               gdc.PtrBuiltInMethod
	methodIsEmptyFn                    gdc.PtrBuiltInMethod
	methodContainsFn                   gdc.PtrBuiltInMethod
	methodIsAbsolutePathFn             gdc.PtrBuiltInMethod
	methodIsRelativePathFn             gdc.PtrBuiltInMethod
	methodSimplifyPathFn               gdc.PtrBuiltInMethod
	methodGetBaseDirFn                 gdc.PtrBuiltInMethod
	methodGetFileFn                    gdc.PtrBuiltInMethod
	methodXmlEscapeFn                  gdc.PtrBuiltInMethod
	methodXmlUnescapeFn                gdc.PtrBuiltInMethod
	methodUriEncodeFn                  gdc.PtrBuiltInMethod
	methodUriDecodeFn                  gdc.PtrBuiltInMethod
	methodCEscapeFn                    gdc.PtrBuiltInMethod
	methodCUnescapeFn                  gdc.PtrBuiltInMethod
	methodJsonEscapeFn                 gdc.PtrBuiltInMethod
	methodValidateNodeNameFn           gdc.PtrBuiltInMethod
	methodValidateFilenameFn           gdc.PtrBuiltInMethod
	methodIsValidIdentifierFn          gdc.PtrBuiltInMethod
	methodIsValidIntFn                 gdc.PtrBuiltInMethod
	methodIsValidFloatFn               gdc.PtrBuiltInMethod
	methodIsValidHexNumberFn           gdc.PtrBuiltInMethod
	methodIsValidHtmlColorFn           gdc.PtrBuiltInMethod
	methodIsValidIpAddressFn           gdc.PtrBuiltInMethod
	methodIsValidFilenameFn            gdc.PtrBuiltInMethod
	methodToIntFn                      gdc.PtrBuiltInMethod
	methodToFloatFn                    gdc.PtrBuiltInMethod
	methodHexToIntFn                   gdc.PtrBuiltInMethod
	methodBinToIntFn                   gdc.PtrBuiltInMethod
	methodLpadFn                       gdc.PtrBuiltInMethod
	methodRpadFn                       gdc.PtrBuiltInMethod
	methodPadDecimalsFn                gdc.PtrBuiltInMethod
	methodPadZerosFn                   gdc.PtrBuiltInMethod
	methodTrimPrefixFn                 gdc.PtrBuiltInMethod
	methodTrimSuffixFn                 gdc.PtrBuiltInMethod
	methodToAsciiBufferFn              gdc.PtrBuiltInMethod
	methodToUtf8BufferFn               gdc.PtrBuiltInMethod
	methodToUtf16BufferFn              gdc.PtrBuiltInMethod
	methodToUtf32BufferFn              gdc.PtrBuiltInMethod
	methodHexDecodeFn                  gdc.PtrBuiltInMethod
	methodToWcharBufferFn              gdc.PtrBuiltInMethod
	methodHashFn                       gdc.PtrBuiltInMethod
	operatorNotFn                      gdc.PtrOperatorEvaluator
	operatorModuleBoolFn               gdc.PtrOperatorEvaluator
	operatorModuleIntFn                gdc.PtrOperatorEvaluator
	operatorModuleFloat32Fn            gdc.PtrOperatorEvaluator
	operatorEqualStringFn              gdc.PtrOperatorEvaluator
	operatorNotEqualStringFn           gdc.PtrOperatorEvaluator
	operatorAddStringFn                gdc.PtrOperatorEvaluator
	operatorModuleStringFn             gdc.PtrOperatorEvaluator
	operatorInStringFn                 gdc.PtrOperatorEvaluator
	operatorModuleVector2Fn            gdc.PtrOperatorEvaluator
	operatorModuleVector2IFn           gdc.PtrOperatorEvaluator
	operatorModuleRect2Fn              gdc.PtrOperatorEvaluator
	operatorModuleRect2IFn             gdc.PtrOperatorEvaluator
	operatorModuleVector3Fn            gdc.PtrOperatorEvaluator
	operatorModuleVector3IFn           gdc.PtrOperatorEvaluator
	operatorModuleTransform2DFn        gdc.PtrOperatorEvaluator
	operatorModuleVector4Fn            gdc.PtrOperatorEvaluator
	operatorModuleVector4IFn           gdc.PtrOperatorEvaluator
	operatorModulePlaneFn              gdc.PtrOperatorEvaluator
	operatorModuleQuaternionFn         gdc.PtrOperatorEvaluator
	operatorModuleAABBFn               gdc.PtrOperatorEvaluator
	operatorModuleBasisFn              gdc.PtrOperatorEvaluator
	operatorModuleTransform3DFn        gdc.PtrOperatorEvaluator
	operatorModuleProjectionFn         gdc.PtrOperatorEvaluator
	operatorModuleColorFn              gdc.PtrOperatorEvaluator
	operatorEqualStringNameFn          gdc.PtrOperatorEvaluator
	operatorNotEqualStringNameFn       gdc.PtrOperatorEvaluator
	operatorLessStringNameFn           gdc.PtrOperatorEvaluator
	operatorLessEqualStringNameFn      gdc.PtrOperatorEvaluator
	operatorGreaterStringNameFn        gdc.PtrOperatorEvaluator
	operatorGreaterEqualStringNameFn   gdc.PtrOperatorEvaluator
	operatorAddStringNameFn            gdc.PtrOperatorEvaluator
	operatorModuleStringNameFn         gdc.PtrOperatorEvaluator
	operatorInStringNameFn             gdc.PtrOperatorEvaluator
	operatorModuleNodePathFn           gdc.PtrOperatorEvaluator
	operatorModuleObjectFn             gdc.PtrOperatorEvaluator
	operatorInObjectFn                 gdc.PtrOperatorEvaluator
	operatorModuleCallableFn           gdc.PtrOperatorEvaluator
	operatorModuleSignalFn             gdc.PtrOperatorEvaluator
	operatorModuleDictionaryFn         gdc.PtrOperatorEvaluator
	operatorInDictionaryFn             gdc.PtrOperatorEvaluator
	operatorModuleArrayFn              gdc.PtrOperatorEvaluator
	operatorInArrayFn                  gdc.PtrOperatorEvaluator
	operatorModulePackedByteArrayFn    gdc.PtrOperatorEvaluator
	operatorModulePackedInt32ArrayFn   gdc.PtrOperatorEvaluator
	operatorModulePackedInt64ArrayFn   gdc.PtrOperatorEvaluator
	operatorModulePackedFloat32ArrayFn gdc.PtrOperatorEvaluator
	operatorModulePackedFloat64ArrayFn gdc.PtrOperatorEvaluator
	operatorModulePackedStringArrayFn  gdc.PtrOperatorEvaluator
	operatorInPackedStringArrayFn      gdc.PtrOperatorEvaluator
	operatorModulePackedVector2ArrayFn gdc.PtrOperatorEvaluator
	operatorModulePackedVector3ArrayFn gdc.PtrOperatorEvaluator
	operatorModulePackedColorArrayFn   gdc.PtrOperatorEvaluator
	toVariantFn                        gdc.TypeFromVariantConstructorFunc
	fromVariantFn                      gdc.VariantFromTypeConstructorFunc
}

var ptrsForStringName ptrsForStringNameList

func initStringNamePtrs(iface gdc.Interface) {
	ptrsForStringName.ctrFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeStringName, 0))
	ptrsForStringName.ctrFromStringNameFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeStringName, 1))
	ptrsForStringName.ctrFromStringFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeStringName, 2))
	ptrsForStringName.destructorFn = ensurePtr(iface.VariantGetPtrDestructor(gdc.VariantTypeStringName))
	{
		methodName := StringNameFromStr("casecmp_to")
		defer methodName.Destroy()
		ptrsForStringName.methodCasecmpToFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 2920860731))
	}
	{
		methodName := StringNameFromStr("nocasecmp_to")
		defer methodName.Destroy()
		ptrsForStringName.methodNocasecmpToFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 2920860731))
	}
	{
		methodName := StringNameFromStr("naturalcasecmp_to")
		defer methodName.Destroy()
		ptrsForStringName.methodNaturalcasecmpToFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 2920860731))
	}
	{
		methodName := StringNameFromStr("naturalnocasecmp_to")
		defer methodName.Destroy()
		ptrsForStringName.methodNaturalnocasecmpToFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 2920860731))
	}
	{
		methodName := StringNameFromStr("length")
		defer methodName.Destroy()
		ptrsForStringName.methodLengthFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 3173160232))
	}
	{
		methodName := StringNameFromStr("substr")
		defer methodName.Destroy()
		ptrsForStringName.methodSubstrFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 787537301))
	}
	{
		methodName := StringNameFromStr("get_slice")
		defer methodName.Destroy()
		ptrsForStringName.methodGetSliceFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 3535100402))
	}
	{
		methodName := StringNameFromStr("get_slicec")
		defer methodName.Destroy()
		ptrsForStringName.methodGetSlicecFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 787537301))
	}
	{
		methodName := StringNameFromStr("get_slice_count")
		defer methodName.Destroy()
		ptrsForStringName.methodGetSliceCountFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 2920860731))
	}
	{
		methodName := StringNameFromStr("find")
		defer methodName.Destroy()
		ptrsForStringName.methodFindFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 1760645412))
	}
	{
		methodName := StringNameFromStr("count")
		defer methodName.Destroy()
		ptrsForStringName.methodCountFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 2343087891))
	}
	{
		methodName := StringNameFromStr("countn")
		defer methodName.Destroy()
		ptrsForStringName.methodCountnFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 2343087891))
	}
	{
		methodName := StringNameFromStr("findn")
		defer methodName.Destroy()
		ptrsForStringName.methodFindnFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 1760645412))
	}
	{
		methodName := StringNameFromStr("rfind")
		defer methodName.Destroy()
		ptrsForStringName.methodRfindFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 1760645412))
	}
	{
		methodName := StringNameFromStr("rfindn")
		defer methodName.Destroy()
		ptrsForStringName.methodRfindnFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 1760645412))
	}
	{
		methodName := StringNameFromStr("match")
		defer methodName.Destroy()
		ptrsForStringName.methodMatchFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 2566493496))
	}
	{
		methodName := StringNameFromStr("matchn")
		defer methodName.Destroy()
		ptrsForStringName.methodMatchnFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 2566493496))
	}
	{
		methodName := StringNameFromStr("begins_with")
		defer methodName.Destroy()
		ptrsForStringName.methodBeginsWithFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 2566493496))
	}
	{
		methodName := StringNameFromStr("ends_with")
		defer methodName.Destroy()
		ptrsForStringName.methodEndsWithFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 2566493496))
	}
	{
		methodName := StringNameFromStr("is_subsequence_of")
		defer methodName.Destroy()
		ptrsForStringName.methodIsSubsequenceOfFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 2566493496))
	}
	{
		methodName := StringNameFromStr("is_subsequence_ofn")
		defer methodName.Destroy()
		ptrsForStringName.methodIsSubsequenceOfnFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 2566493496))
	}
	{
		methodName := StringNameFromStr("bigrams")
		defer methodName.Destroy()
		ptrsForStringName.methodBigramsFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 747180633))
	}
	{
		methodName := StringNameFromStr("similarity")
		defer methodName.Destroy()
		ptrsForStringName.methodSimilarityFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 2697460964))
	}
	{
		methodName := StringNameFromStr("format")
		defer methodName.Destroy()
		ptrsForStringName.methodFormatFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 3212199029))
	}
	{
		methodName := StringNameFromStr("replace")
		defer methodName.Destroy()
		ptrsForStringName.methodReplaceFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 1340436205))
	}
	{
		methodName := StringNameFromStr("replacen")
		defer methodName.Destroy()
		ptrsForStringName.methodReplacenFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 1340436205))
	}
	{
		methodName := StringNameFromStr("repeat")
		defer methodName.Destroy()
		ptrsForStringName.methodRepeatFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 2162347432))
	}
	{
		methodName := StringNameFromStr("reverse")
		defer methodName.Destroy()
		ptrsForStringName.methodReverseFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 3942272618))
	}
	{
		methodName := StringNameFromStr("insert")
		defer methodName.Destroy()
		ptrsForStringName.methodInsertFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 248737229))
	}
	{
		methodName := StringNameFromStr("erase")
		defer methodName.Destroy()
		ptrsForStringName.methodEraseFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 787537301))
	}
	{
		methodName := StringNameFromStr("capitalize")
		defer methodName.Destroy()
		ptrsForStringName.methodCapitalizeFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 3942272618))
	}
	{
		methodName := StringNameFromStr("to_camel_case")
		defer methodName.Destroy()
		ptrsForStringName.methodToCamelCaseFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 3942272618))
	}
	{
		methodName := StringNameFromStr("to_pascal_case")
		defer methodName.Destroy()
		ptrsForStringName.methodToPascalCaseFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 3942272618))
	}
	{
		methodName := StringNameFromStr("to_snake_case")
		defer methodName.Destroy()
		ptrsForStringName.methodToSnakeCaseFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 3942272618))
	}
	{
		methodName := StringNameFromStr("split")
		defer methodName.Destroy()
		ptrsForStringName.methodSplitFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 1252735785))
	}
	{
		methodName := StringNameFromStr("rsplit")
		defer methodName.Destroy()
		ptrsForStringName.methodRsplitFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 1252735785))
	}
	{
		methodName := StringNameFromStr("split_floats")
		defer methodName.Destroy()
		ptrsForStringName.methodSplitFloatsFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 2092079095))
	}
	{
		methodName := StringNameFromStr("join")
		defer methodName.Destroy()
		ptrsForStringName.methodJoinFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 3595973238))
	}
	{
		methodName := StringNameFromStr("to_upper")
		defer methodName.Destroy()
		ptrsForStringName.methodToUpperFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 3942272618))
	}
	{
		methodName := StringNameFromStr("to_lower")
		defer methodName.Destroy()
		ptrsForStringName.methodToLowerFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 3942272618))
	}
	{
		methodName := StringNameFromStr("left")
		defer methodName.Destroy()
		ptrsForStringName.methodLeftFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 2162347432))
	}
	{
		methodName := StringNameFromStr("right")
		defer methodName.Destroy()
		ptrsForStringName.methodRightFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 2162347432))
	}
	{
		methodName := StringNameFromStr("strip_edges")
		defer methodName.Destroy()
		ptrsForStringName.methodStripEdgesFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 907855311))
	}
	{
		methodName := StringNameFromStr("strip_escapes")
		defer methodName.Destroy()
		ptrsForStringName.methodStripEscapesFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 3942272618))
	}
	{
		methodName := StringNameFromStr("lstrip")
		defer methodName.Destroy()
		ptrsForStringName.methodLstripFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 3134094431))
	}
	{
		methodName := StringNameFromStr("rstrip")
		defer methodName.Destroy()
		ptrsForStringName.methodRstripFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 3134094431))
	}
	{
		methodName := StringNameFromStr("get_extension")
		defer methodName.Destroy()
		ptrsForStringName.methodGetExtensionFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 3942272618))
	}
	{
		methodName := StringNameFromStr("get_basename")
		defer methodName.Destroy()
		ptrsForStringName.methodGetBasenameFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 3942272618))
	}
	{
		methodName := StringNameFromStr("path_join")
		defer methodName.Destroy()
		ptrsForStringName.methodPathJoinFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 3134094431))
	}
	{
		methodName := StringNameFromStr("unicode_at")
		defer methodName.Destroy()
		ptrsForStringName.methodUnicodeAtFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 4103005248))
	}
	{
		methodName := StringNameFromStr("indent")
		defer methodName.Destroy()
		ptrsForStringName.methodIndentFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 3134094431))
	}
	{
		methodName := StringNameFromStr("dedent")
		defer methodName.Destroy()
		ptrsForStringName.methodDedentFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 3942272618))
	}
	{
		methodName := StringNameFromStr("md5_text")
		defer methodName.Destroy()
		ptrsForStringName.methodMd5TextFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 3942272618))
	}
	{
		methodName := StringNameFromStr("sha1_text")
		defer methodName.Destroy()
		ptrsForStringName.methodSha1TextFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 3942272618))
	}
	{
		methodName := StringNameFromStr("sha256_text")
		defer methodName.Destroy()
		ptrsForStringName.methodSha256TextFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 3942272618))
	}
	{
		methodName := StringNameFromStr("md5_buffer")
		defer methodName.Destroy()
		ptrsForStringName.methodMd5BufferFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 247621236))
	}
	{
		methodName := StringNameFromStr("sha1_buffer")
		defer methodName.Destroy()
		ptrsForStringName.methodSha1BufferFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 247621236))
	}
	{
		methodName := StringNameFromStr("sha256_buffer")
		defer methodName.Destroy()
		ptrsForStringName.methodSha256BufferFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 247621236))
	}
	{
		methodName := StringNameFromStr("is_empty")
		defer methodName.Destroy()
		ptrsForStringName.methodIsEmptyFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 3918633141))
	}
	{
		methodName := StringNameFromStr("contains")
		defer methodName.Destroy()
		ptrsForStringName.methodContainsFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 2566493496))
	}
	{
		methodName := StringNameFromStr("is_absolute_path")
		defer methodName.Destroy()
		ptrsForStringName.methodIsAbsolutePathFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 3918633141))
	}
	{
		methodName := StringNameFromStr("is_relative_path")
		defer methodName.Destroy()
		ptrsForStringName.methodIsRelativePathFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 3918633141))
	}
	{
		methodName := StringNameFromStr("simplify_path")
		defer methodName.Destroy()
		ptrsForStringName.methodSimplifyPathFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 3942272618))
	}
	{
		methodName := StringNameFromStr("get_base_dir")
		defer methodName.Destroy()
		ptrsForStringName.methodGetBaseDirFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 3942272618))
	}
	{
		methodName := StringNameFromStr("get_file")
		defer methodName.Destroy()
		ptrsForStringName.methodGetFileFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 3942272618))
	}
	{
		methodName := StringNameFromStr("xml_escape")
		defer methodName.Destroy()
		ptrsForStringName.methodXmlEscapeFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 3429816538))
	}
	{
		methodName := StringNameFromStr("xml_unescape")
		defer methodName.Destroy()
		ptrsForStringName.methodXmlUnescapeFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 3942272618))
	}
	{
		methodName := StringNameFromStr("uri_encode")
		defer methodName.Destroy()
		ptrsForStringName.methodUriEncodeFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 3942272618))
	}
	{
		methodName := StringNameFromStr("uri_decode")
		defer methodName.Destroy()
		ptrsForStringName.methodUriDecodeFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 3942272618))
	}
	{
		methodName := StringNameFromStr("c_escape")
		defer methodName.Destroy()
		ptrsForStringName.methodCEscapeFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 3942272618))
	}
	{
		methodName := StringNameFromStr("c_unescape")
		defer methodName.Destroy()
		ptrsForStringName.methodCUnescapeFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 3942272618))
	}
	{
		methodName := StringNameFromStr("json_escape")
		defer methodName.Destroy()
		ptrsForStringName.methodJsonEscapeFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 3942272618))
	}
	{
		methodName := StringNameFromStr("validate_node_name")
		defer methodName.Destroy()
		ptrsForStringName.methodValidateNodeNameFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 3942272618))
	}
	{
		methodName := StringNameFromStr("validate_filename")
		defer methodName.Destroy()
		ptrsForStringName.methodValidateFilenameFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 3942272618))
	}
	{
		methodName := StringNameFromStr("is_valid_identifier")
		defer methodName.Destroy()
		ptrsForStringName.methodIsValidIdentifierFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 3918633141))
	}
	{
		methodName := StringNameFromStr("is_valid_int")
		defer methodName.Destroy()
		ptrsForStringName.methodIsValidIntFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 3918633141))
	}
	{
		methodName := StringNameFromStr("is_valid_float")
		defer methodName.Destroy()
		ptrsForStringName.methodIsValidFloatFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 3918633141))
	}
	{
		methodName := StringNameFromStr("is_valid_hex_number")
		defer methodName.Destroy()
		ptrsForStringName.methodIsValidHexNumberFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 593672999))
	}
	{
		methodName := StringNameFromStr("is_valid_html_color")
		defer methodName.Destroy()
		ptrsForStringName.methodIsValidHtmlColorFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 3918633141))
	}
	{
		methodName := StringNameFromStr("is_valid_ip_address")
		defer methodName.Destroy()
		ptrsForStringName.methodIsValidIpAddressFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 3918633141))
	}
	{
		methodName := StringNameFromStr("is_valid_filename")
		defer methodName.Destroy()
		ptrsForStringName.methodIsValidFilenameFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 3918633141))
	}
	{
		methodName := StringNameFromStr("to_int")
		defer methodName.Destroy()
		ptrsForStringName.methodToIntFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 3173160232))
	}
	{
		methodName := StringNameFromStr("to_float")
		defer methodName.Destroy()
		ptrsForStringName.methodToFloatFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 466405837))
	}
	{
		methodName := StringNameFromStr("hex_to_int")
		defer methodName.Destroy()
		ptrsForStringName.methodHexToIntFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 3173160232))
	}
	{
		methodName := StringNameFromStr("bin_to_int")
		defer methodName.Destroy()
		ptrsForStringName.methodBinToIntFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 3173160232))
	}
	{
		methodName := StringNameFromStr("lpad")
		defer methodName.Destroy()
		ptrsForStringName.methodLpadFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 248737229))
	}
	{
		methodName := StringNameFromStr("rpad")
		defer methodName.Destroy()
		ptrsForStringName.methodRpadFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 248737229))
	}
	{
		methodName := StringNameFromStr("pad_decimals")
		defer methodName.Destroy()
		ptrsForStringName.methodPadDecimalsFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 2162347432))
	}
	{
		methodName := StringNameFromStr("pad_zeros")
		defer methodName.Destroy()
		ptrsForStringName.methodPadZerosFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 2162347432))
	}
	{
		methodName := StringNameFromStr("trim_prefix")
		defer methodName.Destroy()
		ptrsForStringName.methodTrimPrefixFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 3134094431))
	}
	{
		methodName := StringNameFromStr("trim_suffix")
		defer methodName.Destroy()
		ptrsForStringName.methodTrimSuffixFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 3134094431))
	}
	{
		methodName := StringNameFromStr("to_ascii_buffer")
		defer methodName.Destroy()
		ptrsForStringName.methodToAsciiBufferFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 247621236))
	}
	{
		methodName := StringNameFromStr("to_utf8_buffer")
		defer methodName.Destroy()
		ptrsForStringName.methodToUtf8BufferFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 247621236))
	}
	{
		methodName := StringNameFromStr("to_utf16_buffer")
		defer methodName.Destroy()
		ptrsForStringName.methodToUtf16BufferFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 247621236))
	}
	{
		methodName := StringNameFromStr("to_utf32_buffer")
		defer methodName.Destroy()
		ptrsForStringName.methodToUtf32BufferFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 247621236))
	}
	{
		methodName := StringNameFromStr("hex_decode")
		defer methodName.Destroy()
		ptrsForStringName.methodHexDecodeFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 247621236))
	}
	{
		methodName := StringNameFromStr("to_wchar_buffer")
		defer methodName.Destroy()
		ptrsForStringName.methodToWcharBufferFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 247621236))
	}
	{
		methodName := StringNameFromStr("hash")
		defer methodName.Destroy()
		ptrsForStringName.methodHashFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, methodName.AsCPtr(), 3173160232))
	}
	ptrsForStringName.operatorNotFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, gdc.VariantTypeStringName, gdc.VariantTypeNil))
	ptrsForStringName.operatorModuleBoolFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeStringName, gdc.VariantTypeBool))
	ptrsForStringName.operatorModuleIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeStringName, gdc.VariantTypeInt))
	ptrsForStringName.operatorModuleFloat32Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeStringName, gdc.VariantTypeFloat))
	ptrsForStringName.operatorEqualStringFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeStringName, gdc.VariantTypeString))
	ptrsForStringName.operatorNotEqualStringFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeStringName, gdc.VariantTypeString))
	ptrsForStringName.operatorAddStringFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpAdd, gdc.VariantTypeStringName, gdc.VariantTypeString))
	ptrsForStringName.operatorModuleStringFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeStringName, gdc.VariantTypeString))
	ptrsForStringName.operatorInStringFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeStringName, gdc.VariantTypeString))
	ptrsForStringName.operatorModuleVector2Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeStringName, gdc.VariantTypeVector2))
	ptrsForStringName.operatorModuleVector2IFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeStringName, gdc.VariantTypeVector2I))
	ptrsForStringName.operatorModuleRect2Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeStringName, gdc.VariantTypeRect2))
	ptrsForStringName.operatorModuleRect2IFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeStringName, gdc.VariantTypeRect2I))
	ptrsForStringName.operatorModuleVector3Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeStringName, gdc.VariantTypeVector3))
	ptrsForStringName.operatorModuleVector3IFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeStringName, gdc.VariantTypeVector3I))
	ptrsForStringName.operatorModuleTransform2DFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeStringName, gdc.VariantTypeTransform2D))
	ptrsForStringName.operatorModuleVector4Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeStringName, gdc.VariantTypeVector4))
	ptrsForStringName.operatorModuleVector4IFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeStringName, gdc.VariantTypeVector4I))
	ptrsForStringName.operatorModulePlaneFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeStringName, gdc.VariantTypePlane))
	ptrsForStringName.operatorModuleQuaternionFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeStringName, gdc.VariantTypeQuaternion))
	ptrsForStringName.operatorModuleAABBFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeStringName, gdc.VariantTypeAABB))
	ptrsForStringName.operatorModuleBasisFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeStringName, gdc.VariantTypeBasis))
	ptrsForStringName.operatorModuleTransform3DFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeStringName, gdc.VariantTypeTransform3D))
	ptrsForStringName.operatorModuleProjectionFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeStringName, gdc.VariantTypeProjection))
	ptrsForStringName.operatorModuleColorFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeStringName, gdc.VariantTypeColor))
	ptrsForStringName.operatorEqualStringNameFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeStringName, gdc.VariantTypeStringName))
	ptrsForStringName.operatorNotEqualStringNameFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeStringName, gdc.VariantTypeStringName))
	ptrsForStringName.operatorLessStringNameFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpLess, gdc.VariantTypeStringName, gdc.VariantTypeStringName))
	ptrsForStringName.operatorLessEqualStringNameFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpLessEqual, gdc.VariantTypeStringName, gdc.VariantTypeStringName))
	ptrsForStringName.operatorGreaterStringNameFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpGreater, gdc.VariantTypeStringName, gdc.VariantTypeStringName))
	ptrsForStringName.operatorGreaterEqualStringNameFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpGreaterEqual, gdc.VariantTypeStringName, gdc.VariantTypeStringName))
	ptrsForStringName.operatorAddStringNameFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpAdd, gdc.VariantTypeStringName, gdc.VariantTypeStringName))
	ptrsForStringName.operatorModuleStringNameFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeStringName, gdc.VariantTypeStringName))
	ptrsForStringName.operatorInStringNameFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeStringName, gdc.VariantTypeStringName))
	ptrsForStringName.operatorModuleNodePathFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeStringName, gdc.VariantTypeNodePath))
	ptrsForStringName.operatorModuleObjectFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeStringName, gdc.VariantTypeObject))
	ptrsForStringName.operatorInObjectFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeStringName, gdc.VariantTypeObject))
	ptrsForStringName.operatorModuleCallableFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeStringName, gdc.VariantTypeCallable))
	ptrsForStringName.operatorModuleSignalFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeStringName, gdc.VariantTypeSignal))
	ptrsForStringName.operatorModuleDictionaryFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeStringName, gdc.VariantTypeDictionary))
	ptrsForStringName.operatorInDictionaryFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeStringName, gdc.VariantTypeDictionary))
	ptrsForStringName.operatorModuleArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeStringName, gdc.VariantTypeArray))
	ptrsForStringName.operatorInArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeStringName, gdc.VariantTypeArray))
	ptrsForStringName.operatorModulePackedByteArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeStringName, gdc.VariantTypePackedByteArray))
	ptrsForStringName.operatorModulePackedInt32ArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeStringName, gdc.VariantTypePackedInt32Array))
	ptrsForStringName.operatorModulePackedInt64ArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeStringName, gdc.VariantTypePackedInt64Array))
	ptrsForStringName.operatorModulePackedFloat32ArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeStringName, gdc.VariantTypePackedFloat32Array))
	ptrsForStringName.operatorModulePackedFloat64ArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeStringName, gdc.VariantTypePackedFloat64Array))
	ptrsForStringName.operatorModulePackedStringArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeStringName, gdc.VariantTypePackedStringArray))
	ptrsForStringName.operatorInPackedStringArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeStringName, gdc.VariantTypePackedStringArray))
	ptrsForStringName.operatorModulePackedVector2ArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeStringName, gdc.VariantTypePackedVector2Array))
	ptrsForStringName.operatorModulePackedVector3ArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeStringName, gdc.VariantTypePackedVector3Array))
	ptrsForStringName.operatorModulePackedColorArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeStringName, gdc.VariantTypePackedColorArray))
	ptrsForStringName.toVariantFn = ensurePtr(iface.GetVariantToTypeConstructor(gdc.VariantTypeStringName))
	ptrsForStringName.fromVariantFn = ensurePtr(iface.GetVariantFromTypeConstructor(gdc.VariantTypeStringName))
}

type StringName struct {
	//data   *[classSizeStringName]byte
	data   unsafe.Pointer
	iface  gdc.Interface
	pinner runtime.Pinner
}

// Enums

// Constructors
func newStringName() *StringName {
	me := &StringName{
		//data:   new([classSizeStringName]byte),
		data:  giface.MemAlloc(classSizeStringName),
		iface: giface,
	}
	me.pinner.Pin(me)
	me.pinner.Pin(me.AsTypePtr())
	return me
}

func NewStringName() *StringName {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	me := newStringName()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForStringName.ctrFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{}))
	return me
}

func NewStringNameFromStringName(from StringName) *StringName {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	me := newStringName()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForStringName.ctrFromStringNameFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr()}))
	return me
}

func NewStringNameFromString(from String) *StringName {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	me := newStringName()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForStringName.ctrFromStringFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr()}))
	return me
}

// Destructor
func (me *StringName) Destroy() {
	//me.iface.CallPtrDestructor(ensurePtr(ptrsForStringName.destructorFn), me.AsTypePtr())
	//me.pinner.Unpin()
}

// Variant support
func (me *Variant) AsStringName() (*StringName, error) {
	if me.Type() != gdc.VariantTypeStringName {
		return nil, fmt.Errorf("variant is not a StringName")
	}
	bclass := newStringName()
	me.iface.CallTypeFromVariantConstructorFunc(ensurePtr(ptrsForStringName.toVariantFn), bclass.asUninitialized(), me.AsPtr())
	return bclass, nil
}

func (me *StringName) AsVariant() *Variant {
	va := newVariant()
	va.inner = me
	me.iface.CallVariantFromTypeConstructorFunc(ensurePtr(ptrsForStringName.fromVariantFn), va.asUninitialized(), me.AsTypePtr())
	return va
}

// Pointers
func StringNameFromPtr(ptr gdc.ConstTypePtr) *StringName {
	me := newStringName()
	dataCopy(me.data, unsafe.Pointer(ptr), classSizeStringName)
	return me
}

func (me *StringName) ToTypePtr(ptr gdc.TypePtr) {
	dataCopy(unsafe.Pointer(ptr), me.data, classSizeStringName)
}

func (me *StringName) Type() gdc.VariantType {
	return gdc.VariantTypeStringName
}

func (me *StringName) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.data)
}

func (me *StringName) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.AsTypePtr())
}

func (me *StringName) asUninitialized() gdc.UninitializedTypePtr {
	return gdc.UninitializedTypePtr(me.AsTypePtr())
}

// Methods

func (me *StringName) CasecmpTo(to String) int64 {
	ret := NewInt()
	defer ret.Destroy()

	args := []gdc.ConstTypePtr{to.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodCasecmpToFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *StringName) NocasecmpTo(to String) int64 {
	ret := NewInt()
	defer ret.Destroy()

	args := []gdc.ConstTypePtr{to.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodNocasecmpToFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *StringName) NaturalcasecmpTo(to String) int64 {
	ret := NewInt()
	defer ret.Destroy()

	args := []gdc.ConstTypePtr{to.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodNaturalcasecmpToFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *StringName) NaturalnocasecmpTo(to String) int64 {
	ret := NewInt()
	defer ret.Destroy()

	args := []gdc.ConstTypePtr{to.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodNaturalnocasecmpToFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *StringName) Length() int64 {
	ret := NewInt()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodLengthFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *StringName) Substr(from int64, len_ int64) String {
	ret := NewString()

	varg0 := NewIntFromInt(from)
	defer varg0.Destroy()
	varg1 := NewIntFromInt(len_)
	defer varg1.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodSubstrFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *StringName) GetSlice(delimiter String, slice int64) String {
	ret := NewString()

	varg1 := NewIntFromInt(slice)
	defer varg1.Destroy()
	args := []gdc.ConstTypePtr{delimiter.AsCTypePtr(), varg1.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodGetSliceFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *StringName) GetSlicec(delimiter int64, slice int64) String {
	ret := NewString()

	varg0 := NewIntFromInt(delimiter)
	defer varg0.Destroy()
	varg1 := NewIntFromInt(slice)
	defer varg1.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodGetSlicecFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *StringName) GetSliceCount(delimiter String) int64 {
	ret := NewInt()
	defer ret.Destroy()

	args := []gdc.ConstTypePtr{delimiter.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodGetSliceCountFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *StringName) Find(what String, from int64) int64 {
	ret := NewInt()
	defer ret.Destroy()

	varg1 := NewIntFromInt(from)
	defer varg1.Destroy()
	args := []gdc.ConstTypePtr{what.AsCTypePtr(), varg1.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodFindFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *StringName) Count(what String, from int64, to int64) int64 {
	ret := NewInt()
	defer ret.Destroy()

	varg1 := NewIntFromInt(from)
	defer varg1.Destroy()
	varg2 := NewIntFromInt(to)
	defer varg2.Destroy()
	args := []gdc.ConstTypePtr{what.AsCTypePtr(), varg1.AsCTypePtr(), varg2.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodCountFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *StringName) Countn(what String, from int64, to int64) int64 {
	ret := NewInt()
	defer ret.Destroy()

	varg1 := NewIntFromInt(from)
	defer varg1.Destroy()
	varg2 := NewIntFromInt(to)
	defer varg2.Destroy()
	args := []gdc.ConstTypePtr{what.AsCTypePtr(), varg1.AsCTypePtr(), varg2.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodCountnFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *StringName) Findn(what String, from int64) int64 {
	ret := NewInt()
	defer ret.Destroy()

	varg1 := NewIntFromInt(from)
	defer varg1.Destroy()
	args := []gdc.ConstTypePtr{what.AsCTypePtr(), varg1.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodFindnFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *StringName) Rfind(what String, from int64) int64 {
	ret := NewInt()
	defer ret.Destroy()

	varg1 := NewIntFromInt(from)
	defer varg1.Destroy()
	args := []gdc.ConstTypePtr{what.AsCTypePtr(), varg1.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodRfindFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *StringName) Rfindn(what String, from int64) int64 {
	ret := NewInt()
	defer ret.Destroy()

	varg1 := NewIntFromInt(from)
	defer varg1.Destroy()
	args := []gdc.ConstTypePtr{what.AsCTypePtr(), varg1.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodRfindnFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *StringName) Match(expr String) bool {
	ret := NewBool()
	defer ret.Destroy()

	args := []gdc.ConstTypePtr{expr.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodMatchFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *StringName) Matchn(expr String) bool {
	ret := NewBool()
	defer ret.Destroy()

	args := []gdc.ConstTypePtr{expr.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodMatchnFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *StringName) BeginsWith(text String) bool {
	ret := NewBool()
	defer ret.Destroy()

	args := []gdc.ConstTypePtr{text.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodBeginsWithFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *StringName) EndsWith(text String) bool {
	ret := NewBool()
	defer ret.Destroy()

	args := []gdc.ConstTypePtr{text.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodEndsWithFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *StringName) IsSubsequenceOf(text String) bool {
	ret := NewBool()
	defer ret.Destroy()

	args := []gdc.ConstTypePtr{text.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodIsSubsequenceOfFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *StringName) IsSubsequenceOfn(text String) bool {
	ret := NewBool()
	defer ret.Destroy()

	args := []gdc.ConstTypePtr{text.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodIsSubsequenceOfnFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *StringName) Bigrams() PackedStringArray {
	ret := NewPackedStringArray()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodBigramsFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *StringName) Similarity(text String) float64 {
	ret := NewFloat()
	defer ret.Destroy()

	args := []gdc.ConstTypePtr{text.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodSimilarityFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *StringName) Format(values Variant, placeholder String) String {
	ret := NewString()

	args := []gdc.ConstTypePtr{values.AsCTypePtr(), placeholder.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodFormatFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *StringName) Replace(what String, forwhat String) String {
	ret := NewString()

	args := []gdc.ConstTypePtr{what.AsCTypePtr(), forwhat.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodReplaceFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *StringName) Replacen(what String, forwhat String) String {
	ret := NewString()

	args := []gdc.ConstTypePtr{what.AsCTypePtr(), forwhat.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodReplacenFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *StringName) Repeat(count int64) String {
	ret := NewString()

	varg0 := NewIntFromInt(count)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodRepeatFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *StringName) Reverse() String {
	ret := NewString()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodReverseFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *StringName) Insert(position int64, what String) String {
	ret := NewString()

	varg0 := NewIntFromInt(position)
	defer varg0.Destroy()

	args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), what.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodInsertFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *StringName) Erase(position int64, chars int64) String {
	ret := NewString()

	varg0 := NewIntFromInt(position)
	defer varg0.Destroy()
	varg1 := NewIntFromInt(chars)
	defer varg1.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodEraseFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *StringName) Capitalize() String {
	ret := NewString()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodCapitalizeFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *StringName) ToCamelCase() String {
	ret := NewString()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodToCamelCaseFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *StringName) ToPascalCase() String {
	ret := NewString()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodToPascalCaseFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *StringName) ToSnakeCase() String {
	ret := NewString()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodToSnakeCaseFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *StringName) Split(delimiter String, allow_empty bool, maxsplit int64) PackedStringArray {
	ret := NewPackedStringArray()

	varg1 := NewBoolFromBool(allow_empty)
	defer varg1.Destroy()
	varg2 := NewIntFromInt(maxsplit)
	defer varg2.Destroy()
	args := []gdc.ConstTypePtr{delimiter.AsCTypePtr(), varg1.AsCTypePtr(), varg2.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodSplitFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *StringName) Rsplit(delimiter String, allow_empty bool, maxsplit int64) PackedStringArray {
	ret := NewPackedStringArray()

	varg1 := NewBoolFromBool(allow_empty)
	defer varg1.Destroy()
	varg2 := NewIntFromInt(maxsplit)
	defer varg2.Destroy()
	args := []gdc.ConstTypePtr{delimiter.AsCTypePtr(), varg1.AsCTypePtr(), varg2.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodRsplitFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *StringName) SplitFloats(delimiter String, allow_empty bool) PackedFloat64Array {
	ret := NewPackedFloat64Array()

	varg1 := NewBoolFromBool(allow_empty)
	defer varg1.Destroy()
	args := []gdc.ConstTypePtr{delimiter.AsCTypePtr(), varg1.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodSplitFloatsFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *StringName) Join(parts PackedStringArray) String {
	ret := NewString()

	args := []gdc.ConstTypePtr{parts.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodJoinFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *StringName) ToUpper() String {
	ret := NewString()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodToUpperFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *StringName) ToLower() String {
	ret := NewString()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodToLowerFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *StringName) Left(length int64) String {
	ret := NewString()

	varg0 := NewIntFromInt(length)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodLeftFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *StringName) Right(length int64) String {
	ret := NewString()

	varg0 := NewIntFromInt(length)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodRightFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *StringName) StripEdges(left bool, right bool) String {
	ret := NewString()

	varg0 := NewBoolFromBool(left)
	defer varg0.Destroy()
	varg1 := NewBoolFromBool(right)
	defer varg1.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodStripEdgesFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *StringName) StripEscapes() String {
	ret := NewString()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodStripEscapesFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *StringName) Lstrip(chars String) String {
	ret := NewString()

	args := []gdc.ConstTypePtr{chars.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodLstripFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *StringName) Rstrip(chars String) String {
	ret := NewString()

	args := []gdc.ConstTypePtr{chars.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodRstripFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *StringName) GetExtension() String {
	ret := NewString()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodGetExtensionFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *StringName) GetBasename() String {
	ret := NewString()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodGetBasenameFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *StringName) PathJoin(file String) String {
	ret := NewString()

	args := []gdc.ConstTypePtr{file.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodPathJoinFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *StringName) UnicodeAt(at int64) int64 {
	ret := NewInt()
	defer ret.Destroy()
	varg0 := NewIntFromInt(at)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodUnicodeAtFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *StringName) Indent(prefix String) String {
	ret := NewString()

	args := []gdc.ConstTypePtr{prefix.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodIndentFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *StringName) Dedent() String {
	ret := NewString()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodDedentFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *StringName) Md5Text() String {
	ret := NewString()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodMd5TextFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *StringName) Sha1Text() String {
	ret := NewString()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodSha1TextFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *StringName) Sha256Text() String {
	ret := NewString()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodSha256TextFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *StringName) Md5Buffer() PackedByteArray {
	ret := NewPackedByteArray()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodMd5BufferFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *StringName) Sha1Buffer() PackedByteArray {
	ret := NewPackedByteArray()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodSha1BufferFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *StringName) Sha256Buffer() PackedByteArray {
	ret := NewPackedByteArray()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodSha256BufferFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *StringName) IsEmpty() bool {
	ret := NewBool()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodIsEmptyFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *StringName) Contains(what String) bool {
	ret := NewBool()
	defer ret.Destroy()

	args := []gdc.ConstTypePtr{what.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodContainsFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *StringName) IsAbsolutePath() bool {
	ret := NewBool()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodIsAbsolutePathFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *StringName) IsRelativePath() bool {
	ret := NewBool()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodIsRelativePathFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *StringName) SimplifyPath() String {
	ret := NewString()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodSimplifyPathFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *StringName) GetBaseDir() String {
	ret := NewString()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodGetBaseDirFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *StringName) GetFile() String {
	ret := NewString()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodGetFileFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *StringName) XmlEscape(escape_quotes bool) String {
	ret := NewString()

	varg0 := NewBoolFromBool(escape_quotes)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodXmlEscapeFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *StringName) XmlUnescape() String {
	ret := NewString()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodXmlUnescapeFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *StringName) UriEncode() String {
	ret := NewString()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodUriEncodeFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *StringName) UriDecode() String {
	ret := NewString()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodUriDecodeFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *StringName) CEscape() String {
	ret := NewString()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodCEscapeFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *StringName) CUnescape() String {
	ret := NewString()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodCUnescapeFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *StringName) JsonEscape() String {
	ret := NewString()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodJsonEscapeFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *StringName) ValidateNodeName() String {
	ret := NewString()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodValidateNodeNameFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *StringName) ValidateFilename() String {
	ret := NewString()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodValidateFilenameFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *StringName) IsValidIdentifier() bool {
	ret := NewBool()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodIsValidIdentifierFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *StringName) IsValidInt() bool {
	ret := NewBool()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodIsValidIntFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *StringName) IsValidFloat() bool {
	ret := NewBool()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodIsValidFloatFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *StringName) IsValidHexNumber(with_prefix bool) bool {
	ret := NewBool()
	defer ret.Destroy()
	varg0 := NewBoolFromBool(with_prefix)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodIsValidHexNumberFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *StringName) IsValidHtmlColor() bool {
	ret := NewBool()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodIsValidHtmlColorFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *StringName) IsValidIpAddress() bool {
	ret := NewBool()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodIsValidIpAddressFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *StringName) IsValidFilename() bool {
	ret := NewBool()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodIsValidFilenameFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *StringName) ToInt() int64 {
	ret := NewInt()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodToIntFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *StringName) ToFloat() float64 {
	ret := NewFloat()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodToFloatFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *StringName) HexToInt() int64 {
	ret := NewInt()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodHexToIntFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *StringName) BinToInt() int64 {
	ret := NewInt()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodBinToIntFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *StringName) Lpad(min_length int64, character String) String {
	ret := NewString()

	varg0 := NewIntFromInt(min_length)
	defer varg0.Destroy()

	args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), character.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodLpadFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *StringName) Rpad(min_length int64, character String) String {
	ret := NewString()

	varg0 := NewIntFromInt(min_length)
	defer varg0.Destroy()

	args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), character.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodRpadFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *StringName) PadDecimals(digits int64) String {
	ret := NewString()

	varg0 := NewIntFromInt(digits)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodPadDecimalsFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *StringName) PadZeros(digits int64) String {
	ret := NewString()

	varg0 := NewIntFromInt(digits)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodPadZerosFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *StringName) TrimPrefix(prefix String) String {
	ret := NewString()

	args := []gdc.ConstTypePtr{prefix.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodTrimPrefixFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *StringName) TrimSuffix(suffix String) String {
	ret := NewString()

	args := []gdc.ConstTypePtr{suffix.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodTrimSuffixFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *StringName) ToAsciiBuffer() PackedByteArray {
	ret := NewPackedByteArray()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodToAsciiBufferFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *StringName) ToUtf8Buffer() PackedByteArray {
	ret := NewPackedByteArray()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodToUtf8BufferFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *StringName) ToUtf16Buffer() PackedByteArray {
	ret := NewPackedByteArray()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodToUtf16BufferFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *StringName) ToUtf32Buffer() PackedByteArray {
	ret := NewPackedByteArray()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodToUtf32BufferFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *StringName) HexDecode() PackedByteArray {
	ret := NewPackedByteArray()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodHexDecodeFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *StringName) ToWcharBuffer() PackedByteArray {
	ret := NewPackedByteArray()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodToWcharBufferFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *StringName) Hash() int64 {
	ret := NewInt()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForStringName.methodHashFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

// Operators

func (me *StringName) EqualVariant(right Variant) bool {
	opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type())
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *StringName) NotEqualVariant(right Variant) bool {
	opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type())
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *StringName) ModuleVariant(right Variant) String {
	opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type())
	ret := NewString()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *StringName) Not() bool {
	opPtr := ptrsForStringName.operatorNotFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), nil, ret.AsTypePtr())
	return ret.Get()
}

func (me *StringName) ModuleBool(rightArg bool) String {
	right := NewBoolFromBool(rightArg)
	defer right.Destroy()

	opPtr := ptrsForStringName.operatorModuleBoolFn
	ret := NewString()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *StringName) ModuleInt(rightArg int64) String {
	right := NewIntFromInt(rightArg)
	defer right.Destroy()

	opPtr := ptrsForStringName.operatorModuleIntFn
	ret := NewString()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *StringName) ModuleFloat32(rightArg float64) String {
	right := NewFloatFromFloat32(rightArg)
	defer right.Destroy()

	opPtr := ptrsForStringName.operatorModuleFloat32Fn
	ret := NewString()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *StringName) EqualString(right String) bool {
	opPtr := ptrsForStringName.operatorEqualStringFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *StringName) NotEqualString(right String) bool {
	opPtr := ptrsForStringName.operatorNotEqualStringFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *StringName) AddString(right String) String {
	opPtr := ptrsForStringName.operatorAddStringFn
	ret := NewString()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *StringName) ModuleString(right String) String {
	opPtr := ptrsForStringName.operatorModuleStringFn
	ret := NewString()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *StringName) InString(right String) bool {
	opPtr := ptrsForStringName.operatorInStringFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *StringName) ModuleVector2(right Vector2) String {
	opPtr := ptrsForStringName.operatorModuleVector2Fn
	ret := NewString()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *StringName) ModuleVector2I(right Vector2i) String {
	opPtr := ptrsForStringName.operatorModuleVector2IFn
	ret := NewString()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *StringName) ModuleRect2(right Rect2) String {
	opPtr := ptrsForStringName.operatorModuleRect2Fn
	ret := NewString()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *StringName) ModuleRect2I(right Rect2i) String {
	opPtr := ptrsForStringName.operatorModuleRect2IFn
	ret := NewString()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *StringName) ModuleVector3(right Vector3) String {
	opPtr := ptrsForStringName.operatorModuleVector3Fn
	ret := NewString()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *StringName) ModuleVector3I(right Vector3i) String {
	opPtr := ptrsForStringName.operatorModuleVector3IFn
	ret := NewString()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *StringName) ModuleTransform2D(right Transform2D) String {
	opPtr := ptrsForStringName.operatorModuleTransform2DFn
	ret := NewString()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *StringName) ModuleVector4(right Vector4) String {
	opPtr := ptrsForStringName.operatorModuleVector4Fn
	ret := NewString()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *StringName) ModuleVector4I(right Vector4i) String {
	opPtr := ptrsForStringName.operatorModuleVector4IFn
	ret := NewString()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *StringName) ModulePlane(right Plane) String {
	opPtr := ptrsForStringName.operatorModulePlaneFn
	ret := NewString()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *StringName) ModuleQuaternion(right Quaternion) String {
	opPtr := ptrsForStringName.operatorModuleQuaternionFn
	ret := NewString()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *StringName) ModuleAABB(right AABB) String {
	opPtr := ptrsForStringName.operatorModuleAABBFn
	ret := NewString()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *StringName) ModuleBasis(right Basis) String {
	opPtr := ptrsForStringName.operatorModuleBasisFn
	ret := NewString()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *StringName) ModuleTransform3D(right Transform3D) String {
	opPtr := ptrsForStringName.operatorModuleTransform3DFn
	ret := NewString()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *StringName) ModuleProjection(right Projection) String {
	opPtr := ptrsForStringName.operatorModuleProjectionFn
	ret := NewString()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *StringName) ModuleColor(right Color) String {
	opPtr := ptrsForStringName.operatorModuleColorFn
	ret := NewString()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *StringName) EqualStringName(right StringName) bool {
	opPtr := ptrsForStringName.operatorEqualStringNameFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *StringName) NotEqualStringName(right StringName) bool {
	opPtr := ptrsForStringName.operatorNotEqualStringNameFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *StringName) LessStringName(right StringName) bool {
	opPtr := ptrsForStringName.operatorLessStringNameFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *StringName) LessEqualStringName(right StringName) bool {
	opPtr := ptrsForStringName.operatorLessEqualStringNameFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *StringName) GreaterStringName(right StringName) bool {
	opPtr := ptrsForStringName.operatorGreaterStringNameFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *StringName) GreaterEqualStringName(right StringName) bool {
	opPtr := ptrsForStringName.operatorGreaterEqualStringNameFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *StringName) AddStringName(right StringName) String {
	opPtr := ptrsForStringName.operatorAddStringNameFn
	ret := NewString()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *StringName) ModuleStringName(right StringName) String {
	opPtr := ptrsForStringName.operatorModuleStringNameFn
	ret := NewString()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *StringName) InStringName(right StringName) bool {
	opPtr := ptrsForStringName.operatorInStringNameFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *StringName) ModuleNodePath(right NodePath) String {
	opPtr := ptrsForStringName.operatorModuleNodePathFn
	ret := NewString()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *StringName) ModuleObject(right Object) String {
	opPtr := ptrsForStringName.operatorModuleObjectFn
	ret := NewString()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *StringName) InObject(right Object) bool {
	opPtr := ptrsForStringName.operatorInObjectFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *StringName) ModuleCallable(right Callable) String {
	opPtr := ptrsForStringName.operatorModuleCallableFn
	ret := NewString()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *StringName) ModuleSignal(right Signal) String {
	opPtr := ptrsForStringName.operatorModuleSignalFn
	ret := NewString()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *StringName) ModuleDictionary(right Dictionary) String {
	opPtr := ptrsForStringName.operatorModuleDictionaryFn
	ret := NewString()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *StringName) InDictionary(right Dictionary) bool {
	opPtr := ptrsForStringName.operatorInDictionaryFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *StringName) ModuleArray(right Array) String {
	opPtr := ptrsForStringName.operatorModuleArrayFn
	ret := NewString()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *StringName) InArray(right Array) bool {
	opPtr := ptrsForStringName.operatorInArrayFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *StringName) ModulePackedByteArray(right PackedByteArray) String {
	opPtr := ptrsForStringName.operatorModulePackedByteArrayFn
	ret := NewString()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *StringName) ModulePackedInt32Array(right PackedInt32Array) String {
	opPtr := ptrsForStringName.operatorModulePackedInt32ArrayFn
	ret := NewString()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *StringName) ModulePackedInt64Array(right PackedInt64Array) String {
	opPtr := ptrsForStringName.operatorModulePackedInt64ArrayFn
	ret := NewString()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *StringName) ModulePackedFloat32Array(right PackedFloat32Array) String {
	opPtr := ptrsForStringName.operatorModulePackedFloat32ArrayFn
	ret := NewString()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *StringName) ModulePackedFloat64Array(right PackedFloat64Array) String {
	opPtr := ptrsForStringName.operatorModulePackedFloat64ArrayFn
	ret := NewString()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *StringName) ModulePackedStringArray(right PackedStringArray) String {
	opPtr := ptrsForStringName.operatorModulePackedStringArrayFn
	ret := NewString()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *StringName) InPackedStringArray(right PackedStringArray) bool {
	opPtr := ptrsForStringName.operatorInPackedStringArrayFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *StringName) ModulePackedVector2Array(right PackedVector2Array) String {
	opPtr := ptrsForStringName.operatorModulePackedVector2ArrayFn
	ret := NewString()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *StringName) ModulePackedVector3Array(right PackedVector3Array) String {
	opPtr := ptrsForStringName.operatorModulePackedVector3ArrayFn
	ret := NewString()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *StringName) ModulePackedColorArray(right PackedColorArray) String {
	opPtr := ptrsForStringName.operatorModulePackedColorArrayFn
	ret := NewString()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

// Members
