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

type ptrsForPackedByteArrayList struct {
  ctrFn gdc.PtrConstructor
  ctrFromPackedByteArrayFn gdc.PtrConstructor
  ctrFromArrayFn gdc.PtrConstructor
  destructorFn gdc.PtrDestructor
  methodSizeFn gdc.PtrBuiltInMethod
  methodIsEmptyFn gdc.PtrBuiltInMethod
  methodSetFn gdc.PtrBuiltInMethod
  methodPushBackFn gdc.PtrBuiltInMethod
  methodAppendFn gdc.PtrBuiltInMethod
  methodAppendArrayFn gdc.PtrBuiltInMethod
  methodRemoveAtFn gdc.PtrBuiltInMethod
  methodInsertFn gdc.PtrBuiltInMethod
  methodFillFn gdc.PtrBuiltInMethod
  methodResizeFn gdc.PtrBuiltInMethod
  methodClearFn gdc.PtrBuiltInMethod
  methodHasFn gdc.PtrBuiltInMethod
  methodReverseFn gdc.PtrBuiltInMethod
  methodSliceFn gdc.PtrBuiltInMethod
  methodSortFn gdc.PtrBuiltInMethod
  methodBsearchFn gdc.PtrBuiltInMethod
  methodDuplicateFn gdc.PtrBuiltInMethod
  methodFindFn gdc.PtrBuiltInMethod
  methodRfindFn gdc.PtrBuiltInMethod
  methodCountFn gdc.PtrBuiltInMethod
  methodGetStringFromAsciiFn gdc.PtrBuiltInMethod
  methodGetStringFromUtf8Fn gdc.PtrBuiltInMethod
  methodGetStringFromUtf16Fn gdc.PtrBuiltInMethod
  methodGetStringFromUtf32Fn gdc.PtrBuiltInMethod
  methodGetStringFromWcharFn gdc.PtrBuiltInMethod
  methodHexEncodeFn gdc.PtrBuiltInMethod
  methodCompressFn gdc.PtrBuiltInMethod
  methodDecompressFn gdc.PtrBuiltInMethod
  methodDecompressDynamicFn gdc.PtrBuiltInMethod
  methodDecodeU8Fn gdc.PtrBuiltInMethod
  methodDecodeS8Fn gdc.PtrBuiltInMethod
  methodDecodeU16Fn gdc.PtrBuiltInMethod
  methodDecodeS16Fn gdc.PtrBuiltInMethod
  methodDecodeU32Fn gdc.PtrBuiltInMethod
  methodDecodeS32Fn gdc.PtrBuiltInMethod
  methodDecodeU64Fn gdc.PtrBuiltInMethod
  methodDecodeS64Fn gdc.PtrBuiltInMethod
  methodDecodeHalfFn gdc.PtrBuiltInMethod
  methodDecodeFloatFn gdc.PtrBuiltInMethod
  methodDecodeDoubleFn gdc.PtrBuiltInMethod
  methodHasEncodedVarFn gdc.PtrBuiltInMethod
  methodDecodeVarFn gdc.PtrBuiltInMethod
  methodDecodeVarSizeFn gdc.PtrBuiltInMethod
  methodToInt32ArrayFn gdc.PtrBuiltInMethod
  methodToInt64ArrayFn gdc.PtrBuiltInMethod
  methodToFloat32ArrayFn gdc.PtrBuiltInMethod
  methodToFloat64ArrayFn gdc.PtrBuiltInMethod
  methodEncodeU8Fn gdc.PtrBuiltInMethod
  methodEncodeS8Fn gdc.PtrBuiltInMethod
  methodEncodeU16Fn gdc.PtrBuiltInMethod
  methodEncodeS16Fn gdc.PtrBuiltInMethod
  methodEncodeU32Fn gdc.PtrBuiltInMethod
  methodEncodeS32Fn gdc.PtrBuiltInMethod
  methodEncodeU64Fn gdc.PtrBuiltInMethod
  methodEncodeS64Fn gdc.PtrBuiltInMethod
  methodEncodeHalfFn gdc.PtrBuiltInMethod
  methodEncodeFloatFn gdc.PtrBuiltInMethod
  methodEncodeDoubleFn gdc.PtrBuiltInMethod
  methodEncodeVarFn gdc.PtrBuiltInMethod
  operatorNotFn gdc.PtrOperatorEvaluator
  operatorInDictionaryFn gdc.PtrOperatorEvaluator
  operatorInArrayFn gdc.PtrOperatorEvaluator
  operatorEqualPackedByteArrayFn gdc.PtrOperatorEvaluator
  operatorNotEqualPackedByteArrayFn gdc.PtrOperatorEvaluator
  operatorAddPackedByteArrayFn gdc.PtrOperatorEvaluator
  toVariantFn gdc.TypeFromVariantConstructorFunc
  fromVariantFn gdc.VariantFromTypeConstructorFunc
}

var ptrsForPackedByteArray ptrsForPackedByteArrayList

func initPackedByteArrayPtrs(iface gdc.Interface) {
  ptrsForPackedByteArray.ctrFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypePackedByteArray, 0))
  ptrsForPackedByteArray.ctrFromPackedByteArrayFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypePackedByteArray, 1))
  ptrsForPackedByteArray.ctrFromArrayFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypePackedByteArray, 2))
  ptrsForPackedByteArray.destructorFn = ensurePtr(iface.VariantGetPtrDestructor(gdc.VariantTypePackedByteArray))
  {
    methodName := StringNameFromStr("size")
    defer methodName.Destroy()
    ptrsForPackedByteArray.methodSizeFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, methodName.AsCPtr(), 3173160232))
  }
  {
    methodName := StringNameFromStr("is_empty")
    defer methodName.Destroy()
    ptrsForPackedByteArray.methodIsEmptyFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, methodName.AsCPtr(), 3918633141))
  }
  {
    methodName := StringNameFromStr("set")
    defer methodName.Destroy()
    ptrsForPackedByteArray.methodSetFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, methodName.AsCPtr(), 3638975848))
  }
  {
    methodName := StringNameFromStr("push_back")
    defer methodName.Destroy()
    ptrsForPackedByteArray.methodPushBackFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, methodName.AsCPtr(), 694024632))
  }
  {
    methodName := StringNameFromStr("append")
    defer methodName.Destroy()
    ptrsForPackedByteArray.methodAppendFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, methodName.AsCPtr(), 694024632))
  }
  {
    methodName := StringNameFromStr("append_array")
    defer methodName.Destroy()
    ptrsForPackedByteArray.methodAppendArrayFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, methodName.AsCPtr(), 791097111))
  }
  {
    methodName := StringNameFromStr("remove_at")
    defer methodName.Destroy()
    ptrsForPackedByteArray.methodRemoveAtFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, methodName.AsCPtr(), 2823966027))
  }
  {
    methodName := StringNameFromStr("insert")
    defer methodName.Destroy()
    ptrsForPackedByteArray.methodInsertFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, methodName.AsCPtr(), 1487112728))
  }
  {
    methodName := StringNameFromStr("fill")
    defer methodName.Destroy()
    ptrsForPackedByteArray.methodFillFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, methodName.AsCPtr(), 2823966027))
  }
  {
    methodName := StringNameFromStr("resize")
    defer methodName.Destroy()
    ptrsForPackedByteArray.methodResizeFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, methodName.AsCPtr(), 848867239))
  }
  {
    methodName := StringNameFromStr("clear")
    defer methodName.Destroy()
    ptrsForPackedByteArray.methodClearFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("has")
    defer methodName.Destroy()
    ptrsForPackedByteArray.methodHasFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, methodName.AsCPtr(), 931488181))
  }
  {
    methodName := StringNameFromStr("reverse")
    defer methodName.Destroy()
    ptrsForPackedByteArray.methodReverseFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("slice")
    defer methodName.Destroy()
    ptrsForPackedByteArray.methodSliceFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, methodName.AsCPtr(), 2278869132))
  }
  {
    methodName := StringNameFromStr("sort")
    defer methodName.Destroy()
    ptrsForPackedByteArray.methodSortFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("bsearch")
    defer methodName.Destroy()
    ptrsForPackedByteArray.methodBsearchFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, methodName.AsCPtr(), 3380005890))
  }
  {
    methodName := StringNameFromStr("duplicate")
    defer methodName.Destroy()
    ptrsForPackedByteArray.methodDuplicateFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, methodName.AsCPtr(), 851781288))
  }
  {
    methodName := StringNameFromStr("find")
    defer methodName.Destroy()
    ptrsForPackedByteArray.methodFindFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, methodName.AsCPtr(), 2984303840))
  }
  {
    methodName := StringNameFromStr("rfind")
    defer methodName.Destroy()
    ptrsForPackedByteArray.methodRfindFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, methodName.AsCPtr(), 2984303840))
  }
  {
    methodName := StringNameFromStr("count")
    defer methodName.Destroy()
    ptrsForPackedByteArray.methodCountFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, methodName.AsCPtr(), 4103005248))
  }
  {
    methodName := StringNameFromStr("get_string_from_ascii")
    defer methodName.Destroy()
    ptrsForPackedByteArray.methodGetStringFromAsciiFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, methodName.AsCPtr(), 3942272618))
  }
  {
    methodName := StringNameFromStr("get_string_from_utf8")
    defer methodName.Destroy()
    ptrsForPackedByteArray.methodGetStringFromUtf8Fn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, methodName.AsCPtr(), 3942272618))
  }
  {
    methodName := StringNameFromStr("get_string_from_utf16")
    defer methodName.Destroy()
    ptrsForPackedByteArray.methodGetStringFromUtf16Fn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, methodName.AsCPtr(), 3942272618))
  }
  {
    methodName := StringNameFromStr("get_string_from_utf32")
    defer methodName.Destroy()
    ptrsForPackedByteArray.methodGetStringFromUtf32Fn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, methodName.AsCPtr(), 3942272618))
  }
  {
    methodName := StringNameFromStr("get_string_from_wchar")
    defer methodName.Destroy()
    ptrsForPackedByteArray.methodGetStringFromWcharFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, methodName.AsCPtr(), 3942272618))
  }
  {
    methodName := StringNameFromStr("hex_encode")
    defer methodName.Destroy()
    ptrsForPackedByteArray.methodHexEncodeFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, methodName.AsCPtr(), 3942272618))
  }
  {
    methodName := StringNameFromStr("compress")
    defer methodName.Destroy()
    ptrsForPackedByteArray.methodCompressFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, methodName.AsCPtr(), 1845905913))
  }
  {
    methodName := StringNameFromStr("decompress")
    defer methodName.Destroy()
    ptrsForPackedByteArray.methodDecompressFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, methodName.AsCPtr(), 2278869132))
  }
  {
    methodName := StringNameFromStr("decompress_dynamic")
    defer methodName.Destroy()
    ptrsForPackedByteArray.methodDecompressDynamicFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, methodName.AsCPtr(), 2278869132))
  }
  {
    methodName := StringNameFromStr("decode_u8")
    defer methodName.Destroy()
    ptrsForPackedByteArray.methodDecodeU8Fn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, methodName.AsCPtr(), 4103005248))
  }
  {
    methodName := StringNameFromStr("decode_s8")
    defer methodName.Destroy()
    ptrsForPackedByteArray.methodDecodeS8Fn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, methodName.AsCPtr(), 4103005248))
  }
  {
    methodName := StringNameFromStr("decode_u16")
    defer methodName.Destroy()
    ptrsForPackedByteArray.methodDecodeU16Fn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, methodName.AsCPtr(), 4103005248))
  }
  {
    methodName := StringNameFromStr("decode_s16")
    defer methodName.Destroy()
    ptrsForPackedByteArray.methodDecodeS16Fn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, methodName.AsCPtr(), 4103005248))
  }
  {
    methodName := StringNameFromStr("decode_u32")
    defer methodName.Destroy()
    ptrsForPackedByteArray.methodDecodeU32Fn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, methodName.AsCPtr(), 4103005248))
  }
  {
    methodName := StringNameFromStr("decode_s32")
    defer methodName.Destroy()
    ptrsForPackedByteArray.methodDecodeS32Fn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, methodName.AsCPtr(), 4103005248))
  }
  {
    methodName := StringNameFromStr("decode_u64")
    defer methodName.Destroy()
    ptrsForPackedByteArray.methodDecodeU64Fn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, methodName.AsCPtr(), 4103005248))
  }
  {
    methodName := StringNameFromStr("decode_s64")
    defer methodName.Destroy()
    ptrsForPackedByteArray.methodDecodeS64Fn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, methodName.AsCPtr(), 4103005248))
  }
  {
    methodName := StringNameFromStr("decode_half")
    defer methodName.Destroy()
    ptrsForPackedByteArray.methodDecodeHalfFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, methodName.AsCPtr(), 1401583798))
  }
  {
    methodName := StringNameFromStr("decode_float")
    defer methodName.Destroy()
    ptrsForPackedByteArray.methodDecodeFloatFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, methodName.AsCPtr(), 1401583798))
  }
  {
    methodName := StringNameFromStr("decode_double")
    defer methodName.Destroy()
    ptrsForPackedByteArray.methodDecodeDoubleFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, methodName.AsCPtr(), 1401583798))
  }
  {
    methodName := StringNameFromStr("has_encoded_var")
    defer methodName.Destroy()
    ptrsForPackedByteArray.methodHasEncodedVarFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, methodName.AsCPtr(), 2914632957))
  }
  {
    methodName := StringNameFromStr("decode_var")
    defer methodName.Destroy()
    ptrsForPackedByteArray.methodDecodeVarFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, methodName.AsCPtr(), 1740420038))
  }
  {
    methodName := StringNameFromStr("decode_var_size")
    defer methodName.Destroy()
    ptrsForPackedByteArray.methodDecodeVarSizeFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, methodName.AsCPtr(), 954237325))
  }
  {
    methodName := StringNameFromStr("to_int32_array")
    defer methodName.Destroy()
    ptrsForPackedByteArray.methodToInt32ArrayFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, methodName.AsCPtr(), 3158844420))
  }
  {
    methodName := StringNameFromStr("to_int64_array")
    defer methodName.Destroy()
    ptrsForPackedByteArray.methodToInt64ArrayFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, methodName.AsCPtr(), 1961294120))
  }
  {
    methodName := StringNameFromStr("to_float32_array")
    defer methodName.Destroy()
    ptrsForPackedByteArray.methodToFloat32ArrayFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, methodName.AsCPtr(), 3575107827))
  }
  {
    methodName := StringNameFromStr("to_float64_array")
    defer methodName.Destroy()
    ptrsForPackedByteArray.methodToFloat64ArrayFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, methodName.AsCPtr(), 1627308337))
  }
  {
    methodName := StringNameFromStr("encode_u8")
    defer methodName.Destroy()
    ptrsForPackedByteArray.methodEncodeU8Fn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, methodName.AsCPtr(), 3638975848))
  }
  {
    methodName := StringNameFromStr("encode_s8")
    defer methodName.Destroy()
    ptrsForPackedByteArray.methodEncodeS8Fn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, methodName.AsCPtr(), 3638975848))
  }
  {
    methodName := StringNameFromStr("encode_u16")
    defer methodName.Destroy()
    ptrsForPackedByteArray.methodEncodeU16Fn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, methodName.AsCPtr(), 3638975848))
  }
  {
    methodName := StringNameFromStr("encode_s16")
    defer methodName.Destroy()
    ptrsForPackedByteArray.methodEncodeS16Fn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, methodName.AsCPtr(), 3638975848))
  }
  {
    methodName := StringNameFromStr("encode_u32")
    defer methodName.Destroy()
    ptrsForPackedByteArray.methodEncodeU32Fn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, methodName.AsCPtr(), 3638975848))
  }
  {
    methodName := StringNameFromStr("encode_s32")
    defer methodName.Destroy()
    ptrsForPackedByteArray.methodEncodeS32Fn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, methodName.AsCPtr(), 3638975848))
  }
  {
    methodName := StringNameFromStr("encode_u64")
    defer methodName.Destroy()
    ptrsForPackedByteArray.methodEncodeU64Fn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, methodName.AsCPtr(), 3638975848))
  }
  {
    methodName := StringNameFromStr("encode_s64")
    defer methodName.Destroy()
    ptrsForPackedByteArray.methodEncodeS64Fn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, methodName.AsCPtr(), 3638975848))
  }
  {
    methodName := StringNameFromStr("encode_half")
    defer methodName.Destroy()
    ptrsForPackedByteArray.methodEncodeHalfFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, methodName.AsCPtr(), 1113000516))
  }
  {
    methodName := StringNameFromStr("encode_float")
    defer methodName.Destroy()
    ptrsForPackedByteArray.methodEncodeFloatFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, methodName.AsCPtr(), 1113000516))
  }
  {
    methodName := StringNameFromStr("encode_double")
    defer methodName.Destroy()
    ptrsForPackedByteArray.methodEncodeDoubleFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, methodName.AsCPtr(), 1113000516))
  }
  {
    methodName := StringNameFromStr("encode_var")
    defer methodName.Destroy()
    ptrsForPackedByteArray.methodEncodeVarFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedByteArray, methodName.AsCPtr(), 2604460497))
  }
  ptrsForPackedByteArray.operatorNotFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, gdc.VariantTypePackedByteArray, gdc.VariantTypeNil))
  ptrsForPackedByteArray.operatorInDictionaryFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypePackedByteArray, gdc.VariantTypeDictionary))
  ptrsForPackedByteArray.operatorInArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypePackedByteArray, gdc.VariantTypeArray))
  ptrsForPackedByteArray.operatorEqualPackedByteArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypePackedByteArray, gdc.VariantTypePackedByteArray))
  ptrsForPackedByteArray.operatorNotEqualPackedByteArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypePackedByteArray, gdc.VariantTypePackedByteArray))
  ptrsForPackedByteArray.operatorAddPackedByteArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpAdd, gdc.VariantTypePackedByteArray, gdc.VariantTypePackedByteArray))
  ptrsForPackedByteArray.toVariantFn = ensurePtr(iface.GetVariantToTypeConstructor(gdc.VariantTypePackedByteArray))
  ptrsForPackedByteArray.fromVariantFn = ensurePtr(iface.GetVariantFromTypeConstructor(gdc.VariantTypePackedByteArray))
}

type PackedByteArray struct {
  data   *[classSizePackedByteArray]byte
  iface  gdc.Interface
  pinner runtime.Pinner
}

// Enums

// Constructors
func newPackedByteArray() *PackedByteArray {
  me := &PackedByteArray{
    data:   new([classSizePackedByteArray]byte),
    iface:  giface,
  }
  me.pinner.Pin(me)
  me.pinner.Pin(me.AsTypePtr())
  return me
}

func NewPackedByteArray() *PackedByteArray {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newPackedByteArray()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForPackedByteArray.ctrFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{}))
  return me
}

func NewPackedByteArrayFromPackedByteArray(from PackedByteArray, ) *PackedByteArray {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newPackedByteArray()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForPackedByteArray.ctrFromPackedByteArrayFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return me
}

func NewPackedByteArrayFromArray(from Array, ) *PackedByteArray {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newPackedByteArray()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForPackedByteArray.ctrFromArrayFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return me
}

// Destructor
func (me *PackedByteArray) Destroy() {
	me.iface.CallPtrDestructor(ensurePtr(ptrsForPackedByteArray.destructorFn), me.AsTypePtr())
  me.pinner.Unpin()
}

// Variant support
func (me *Variant) AsPackedByteArray() (*PackedByteArray, error) {
	if me.Type() != gdc.VariantTypePackedByteArray {
		return nil, fmt.Errorf("variant is not a PackedByteArray")
	}
  bclass := newPackedByteArray()
	me.iface.CallTypeFromVariantConstructorFunc(ensurePtr(ptrsForPackedByteArray.toVariantFn), bclass.asUninitialized(), me.AsPtr())
	return bclass, nil
}

func (me *PackedByteArray) AsVariant() *Variant {
  va := newVariant()
  va.inner = me
  me.iface.CallVariantFromTypeConstructorFunc(ensurePtr(ptrsForPackedByteArray.fromVariantFn), va.asUninitialized(), me.AsTypePtr())
  return va
}

// Pointers
func PackedByteArrayFromPtr(ptr gdc.ConstTypePtr) *PackedByteArray {
  me := newPackedByteArray()
  dataFromPtr(me.data[:], ptr)
  return me
}

func (me *PackedByteArray) Type() gdc.VariantType {
  return gdc.VariantTypePackedByteArray
}

func (me *PackedByteArray) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(unsafe.Pointer(me.data))
}

func (me *PackedByteArray) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.AsTypePtr())
}

func (me *PackedByteArray) asUninitialized() gdc.UninitializedTypePtr {
  return gdc.UninitializedTypePtr(me.AsTypePtr())
}

// Methods

func (me *PackedByteArray) Size() int64 {
  ret := NewInt()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedByteArray.methodSizeFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedByteArray) IsEmpty() bool {
  ret := NewBool()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedByteArray.methodIsEmptyFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedByteArray) Set(index int64, value int64, )  {
  varg0 := NewIntFromInt(index)
  defer varg0.Destroy()
  varg1 := NewIntFromInt(value)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedByteArray.methodSetFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedByteArray) PushBack(value int64, ) bool {
  ret := NewBool()
  defer ret.Destroy()
  varg0 := NewIntFromInt(value)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedByteArray.methodPushBackFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedByteArray) Append(value int64, ) bool {
  ret := NewBool()
  defer ret.Destroy()
  varg0 := NewIntFromInt(value)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedByteArray.methodAppendFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedByteArray) AppendArray(array PackedByteArray, )  {

  args := []gdc.ConstTypePtr{array.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedByteArray.methodAppendArrayFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedByteArray) RemoveAt(index int64, )  {
  varg0 := NewIntFromInt(index)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedByteArray.methodRemoveAtFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedByteArray) Insert(at_index int64, value int64, ) int64 {
  ret := NewInt()
  defer ret.Destroy()
  varg0 := NewIntFromInt(at_index)
  defer varg0.Destroy()
  varg1 := NewIntFromInt(value)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedByteArray.methodInsertFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedByteArray) Fill(value int64, )  {
  varg0 := NewIntFromInt(value)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedByteArray.methodFillFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedByteArray) Resize(new_size int64, ) int64 {
  ret := NewInt()
  defer ret.Destroy()
  varg0 := NewIntFromInt(new_size)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedByteArray.methodResizeFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedByteArray) Clear()  {
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedByteArray.methodClearFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedByteArray) Has(value int64, ) bool {
  ret := NewBool()
  defer ret.Destroy()
  varg0 := NewIntFromInt(value)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedByteArray.methodHasFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedByteArray) Reverse()  {
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedByteArray.methodReverseFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedByteArray) Slice(begin int64, end int64, ) PackedByteArray {
  ret := NewPackedByteArray()

  varg0 := NewIntFromInt(begin)
  defer varg0.Destroy()
  varg1 := NewIntFromInt(end)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedByteArray.methodSliceFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *PackedByteArray) Sort()  {
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedByteArray.methodSortFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedByteArray) Bsearch(value int64, before bool, ) int64 {
  ret := NewInt()
  defer ret.Destroy()
  varg0 := NewIntFromInt(value)
  defer varg0.Destroy()
  varg1 := NewBoolFromBool(before)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedByteArray.methodBsearchFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedByteArray) Duplicate() PackedByteArray {
  ret := NewPackedByteArray()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedByteArray.methodDuplicateFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *PackedByteArray) Find(value int64, from int64, ) int64 {
  ret := NewInt()
  defer ret.Destroy()
  varg0 := NewIntFromInt(value)
  defer varg0.Destroy()
  varg1 := NewIntFromInt(from)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedByteArray.methodFindFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedByteArray) Rfind(value int64, from int64, ) int64 {
  ret := NewInt()
  defer ret.Destroy()
  varg0 := NewIntFromInt(value)
  defer varg0.Destroy()
  varg1 := NewIntFromInt(from)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedByteArray.methodRfindFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedByteArray) Count(value int64, ) int64 {
  ret := NewInt()
  defer ret.Destroy()
  varg0 := NewIntFromInt(value)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedByteArray.methodCountFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedByteArray) GetStringFromAscii() String {
  ret := NewString()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedByteArray.methodGetStringFromAsciiFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *PackedByteArray) GetStringFromUtf8() String {
  ret := NewString()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedByteArray.methodGetStringFromUtf8Fn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *PackedByteArray) GetStringFromUtf16() String {
  ret := NewString()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedByteArray.methodGetStringFromUtf16Fn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *PackedByteArray) GetStringFromUtf32() String {
  ret := NewString()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedByteArray.methodGetStringFromUtf32Fn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *PackedByteArray) GetStringFromWchar() String {
  ret := NewString()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedByteArray.methodGetStringFromWcharFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *PackedByteArray) HexEncode() String {
  ret := NewString()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedByteArray.methodHexEncodeFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *PackedByteArray) Compress(compression_mode int64, ) PackedByteArray {
  ret := NewPackedByteArray()

  varg0 := NewIntFromInt(compression_mode)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedByteArray.methodCompressFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *PackedByteArray) Decompress(buffer_size int64, compression_mode int64, ) PackedByteArray {
  ret := NewPackedByteArray()

  varg0 := NewIntFromInt(buffer_size)
  defer varg0.Destroy()
  varg1 := NewIntFromInt(compression_mode)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedByteArray.methodDecompressFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *PackedByteArray) DecompressDynamic(max_output_size int64, compression_mode int64, ) PackedByteArray {
  ret := NewPackedByteArray()

  varg0 := NewIntFromInt(max_output_size)
  defer varg0.Destroy()
  varg1 := NewIntFromInt(compression_mode)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedByteArray.methodDecompressDynamicFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *PackedByteArray) DecodeU8(byte_offset int64, ) int64 {
  ret := NewInt()
  defer ret.Destroy()
  varg0 := NewIntFromInt(byte_offset)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedByteArray.methodDecodeU8Fn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedByteArray) DecodeS8(byte_offset int64, ) int64 {
  ret := NewInt()
  defer ret.Destroy()
  varg0 := NewIntFromInt(byte_offset)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedByteArray.methodDecodeS8Fn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedByteArray) DecodeU16(byte_offset int64, ) int64 {
  ret := NewInt()
  defer ret.Destroy()
  varg0 := NewIntFromInt(byte_offset)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedByteArray.methodDecodeU16Fn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedByteArray) DecodeS16(byte_offset int64, ) int64 {
  ret := NewInt()
  defer ret.Destroy()
  varg0 := NewIntFromInt(byte_offset)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedByteArray.methodDecodeS16Fn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedByteArray) DecodeU32(byte_offset int64, ) int64 {
  ret := NewInt()
  defer ret.Destroy()
  varg0 := NewIntFromInt(byte_offset)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedByteArray.methodDecodeU32Fn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedByteArray) DecodeS32(byte_offset int64, ) int64 {
  ret := NewInt()
  defer ret.Destroy()
  varg0 := NewIntFromInt(byte_offset)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedByteArray.methodDecodeS32Fn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedByteArray) DecodeU64(byte_offset int64, ) int64 {
  ret := NewInt()
  defer ret.Destroy()
  varg0 := NewIntFromInt(byte_offset)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedByteArray.methodDecodeU64Fn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedByteArray) DecodeS64(byte_offset int64, ) int64 {
  ret := NewInt()
  defer ret.Destroy()
  varg0 := NewIntFromInt(byte_offset)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedByteArray.methodDecodeS64Fn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedByteArray) DecodeHalf(byte_offset int64, ) float64 {
  ret := NewFloat()
  defer ret.Destroy()
  varg0 := NewIntFromInt(byte_offset)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedByteArray.methodDecodeHalfFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedByteArray) DecodeFloat(byte_offset int64, ) float64 {
  ret := NewFloat()
  defer ret.Destroy()
  varg0 := NewIntFromInt(byte_offset)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedByteArray.methodDecodeFloatFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedByteArray) DecodeDouble(byte_offset int64, ) float64 {
  ret := NewFloat()
  defer ret.Destroy()
  varg0 := NewIntFromInt(byte_offset)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedByteArray.methodDecodeDoubleFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedByteArray) HasEncodedVar(byte_offset int64, allow_objects bool, ) bool {
  ret := NewBool()
  defer ret.Destroy()
  varg0 := NewIntFromInt(byte_offset)
  defer varg0.Destroy()
  varg1 := NewBoolFromBool(allow_objects)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedByteArray.methodHasEncodedVarFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedByteArray) DecodeVar(byte_offset int64, allow_objects bool, ) Variant {
  ret := NewVariant()

  varg0 := NewIntFromInt(byte_offset)
  defer varg0.Destroy()
  varg1 := NewBoolFromBool(allow_objects)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedByteArray.methodDecodeVarFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *PackedByteArray) DecodeVarSize(byte_offset int64, allow_objects bool, ) int64 {
  ret := NewInt()
  defer ret.Destroy()
  varg0 := NewIntFromInt(byte_offset)
  defer varg0.Destroy()
  varg1 := NewBoolFromBool(allow_objects)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedByteArray.methodDecodeVarSizeFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *PackedByteArray) ToInt32Array() PackedInt32Array {
  ret := NewPackedInt32Array()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedByteArray.methodToInt32ArrayFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *PackedByteArray) ToInt64Array() PackedInt64Array {
  ret := NewPackedInt64Array()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedByteArray.methodToInt64ArrayFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *PackedByteArray) ToFloat32Array() PackedFloat32Array {
  ret := NewPackedFloat32Array()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedByteArray.methodToFloat32ArrayFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *PackedByteArray) ToFloat64Array() PackedFloat64Array {
  ret := NewPackedFloat64Array()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedByteArray.methodToFloat64ArrayFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *PackedByteArray) EncodeU8(byte_offset int64, value int64, )  {
  varg0 := NewIntFromInt(byte_offset)
  defer varg0.Destroy()
  varg1 := NewIntFromInt(value)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedByteArray.methodEncodeU8Fn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedByteArray) EncodeS8(byte_offset int64, value int64, )  {
  varg0 := NewIntFromInt(byte_offset)
  defer varg0.Destroy()
  varg1 := NewIntFromInt(value)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedByteArray.methodEncodeS8Fn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedByteArray) EncodeU16(byte_offset int64, value int64, )  {
  varg0 := NewIntFromInt(byte_offset)
  defer varg0.Destroy()
  varg1 := NewIntFromInt(value)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedByteArray.methodEncodeU16Fn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedByteArray) EncodeS16(byte_offset int64, value int64, )  {
  varg0 := NewIntFromInt(byte_offset)
  defer varg0.Destroy()
  varg1 := NewIntFromInt(value)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedByteArray.methodEncodeS16Fn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedByteArray) EncodeU32(byte_offset int64, value int64, )  {
  varg0 := NewIntFromInt(byte_offset)
  defer varg0.Destroy()
  varg1 := NewIntFromInt(value)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedByteArray.methodEncodeU32Fn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedByteArray) EncodeS32(byte_offset int64, value int64, )  {
  varg0 := NewIntFromInt(byte_offset)
  defer varg0.Destroy()
  varg1 := NewIntFromInt(value)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedByteArray.methodEncodeS32Fn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedByteArray) EncodeU64(byte_offset int64, value int64, )  {
  varg0 := NewIntFromInt(byte_offset)
  defer varg0.Destroy()
  varg1 := NewIntFromInt(value)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedByteArray.methodEncodeU64Fn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedByteArray) EncodeS64(byte_offset int64, value int64, )  {
  varg0 := NewIntFromInt(byte_offset)
  defer varg0.Destroy()
  varg1 := NewIntFromInt(value)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedByteArray.methodEncodeS64Fn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedByteArray) EncodeHalf(byte_offset int64, value float64, )  {
  varg0 := NewIntFromInt(byte_offset)
  defer varg0.Destroy()
  varg1 := NewFloatFromFloat32(value)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedByteArray.methodEncodeHalfFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedByteArray) EncodeFloat(byte_offset int64, value float64, )  {
  varg0 := NewIntFromInt(byte_offset)
  defer varg0.Destroy()
  varg1 := NewFloatFromFloat32(value)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedByteArray.methodEncodeFloatFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedByteArray) EncodeDouble(byte_offset int64, value float64, )  {
  varg0 := NewIntFromInt(byte_offset)
  defer varg0.Destroy()
  varg1 := NewFloatFromFloat32(value)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedByteArray.methodEncodeDoubleFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedByteArray) EncodeVar(byte_offset int64, value Variant, allow_objects bool, ) int64 {
  ret := NewInt()
  defer ret.Destroy()
  varg0 := NewIntFromInt(byte_offset)
  defer varg0.Destroy()

  varg2 := NewBoolFromBool(allow_objects)
  defer varg2.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), value.AsCTypePtr(), varg2.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedByteArray.methodEncodeVarFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

// Operators

func (me *PackedByteArray) EqualVariant(right Variant) bool {
  opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type())
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *PackedByteArray) NotEqualVariant(right Variant) bool {
  opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type())
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *PackedByteArray) Not() bool {
  opPtr := ptrsForPackedByteArray.operatorNotFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), nil, ret.AsTypePtr())
  return ret.Get()
}

func (me *PackedByteArray) InDictionary(right Dictionary) bool {
  opPtr := ptrsForPackedByteArray.operatorInDictionaryFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *PackedByteArray) InArray(right Array) bool {
  opPtr := ptrsForPackedByteArray.operatorInArrayFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *PackedByteArray) EqualPackedByteArray(right PackedByteArray) bool {
  opPtr := ptrsForPackedByteArray.operatorEqualPackedByteArrayFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *PackedByteArray) NotEqualPackedByteArray(right PackedByteArray) bool {
  opPtr := ptrsForPackedByteArray.operatorNotEqualPackedByteArrayFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *PackedByteArray) AddPackedByteArray(right PackedByteArray) PackedByteArray {
  opPtr := ptrsForPackedByteArray.operatorAddPackedByteArrayFn
  ret := NewPackedByteArray()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

// Members
