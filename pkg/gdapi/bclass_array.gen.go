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

type ptrsForArrayList struct {
	ctrFn                              gdc.PtrConstructor
	ctrFromArrayFn                     gdc.PtrConstructor
	ctrFromArrayIntStringNameVariantFn gdc.PtrConstructor
	ctrFromPackedByteArrayFn           gdc.PtrConstructor
	ctrFromPackedInt32ArrayFn          gdc.PtrConstructor
	ctrFromPackedInt64ArrayFn          gdc.PtrConstructor
	ctrFromPackedFloat32ArrayFn        gdc.PtrConstructor
	ctrFromPackedFloat64ArrayFn        gdc.PtrConstructor
	ctrFromPackedStringArrayFn         gdc.PtrConstructor
	ctrFromPackedVector2ArrayFn        gdc.PtrConstructor
	ctrFromPackedVector3ArrayFn        gdc.PtrConstructor
	ctrFromPackedColorArrayFn          gdc.PtrConstructor
	destructorFn                       gdc.PtrDestructor
	methodSizeFn                       gdc.PtrBuiltInMethod
	methodIsEmptyFn                    gdc.PtrBuiltInMethod
	methodClearFn                      gdc.PtrBuiltInMethod
	methodHashFn                       gdc.PtrBuiltInMethod
	methodAssignFn                     gdc.PtrBuiltInMethod
	methodPushBackFn                   gdc.PtrBuiltInMethod
	methodPushFrontFn                  gdc.PtrBuiltInMethod
	methodAppendFn                     gdc.PtrBuiltInMethod
	methodAppendArrayFn                gdc.PtrBuiltInMethod
	methodResizeFn                     gdc.PtrBuiltInMethod
	methodInsertFn                     gdc.PtrBuiltInMethod
	methodRemoveAtFn                   gdc.PtrBuiltInMethod
	methodFillFn                       gdc.PtrBuiltInMethod
	methodEraseFn                      gdc.PtrBuiltInMethod
	methodFrontFn                      gdc.PtrBuiltInMethod
	methodBackFn                       gdc.PtrBuiltInMethod
	methodPickRandomFn                 gdc.PtrBuiltInMethod
	methodFindFn                       gdc.PtrBuiltInMethod
	methodRfindFn                      gdc.PtrBuiltInMethod
	methodCountFn                      gdc.PtrBuiltInMethod
	methodHasFn                        gdc.PtrBuiltInMethod
	methodPopBackFn                    gdc.PtrBuiltInMethod
	methodPopFrontFn                   gdc.PtrBuiltInMethod
	methodPopAtFn                      gdc.PtrBuiltInMethod
	methodSortFn                       gdc.PtrBuiltInMethod
	methodSortCustomFn                 gdc.PtrBuiltInMethod
	methodShuffleFn                    gdc.PtrBuiltInMethod
	methodBsearchFn                    gdc.PtrBuiltInMethod
	methodBsearchCustomFn              gdc.PtrBuiltInMethod
	methodReverseFn                    gdc.PtrBuiltInMethod
	methodDuplicateFn                  gdc.PtrBuiltInMethod
	methodSliceFn                      gdc.PtrBuiltInMethod
	methodFilterFn                     gdc.PtrBuiltInMethod
	methodMapFn                        gdc.PtrBuiltInMethod
	methodReduceFn                     gdc.PtrBuiltInMethod
	methodAnyFn                        gdc.PtrBuiltInMethod
	methodAllFn                        gdc.PtrBuiltInMethod
	methodMaxFn                        gdc.PtrBuiltInMethod
	methodMinFn                        gdc.PtrBuiltInMethod
	methodIsTypedFn                    gdc.PtrBuiltInMethod
	methodIsSameTypedFn                gdc.PtrBuiltInMethod
	methodGetTypedBuiltinFn            gdc.PtrBuiltInMethod
	methodGetTypedClassNameFn          gdc.PtrBuiltInMethod
	methodGetTypedScriptFn             gdc.PtrBuiltInMethod
	methodMakeReadOnlyFn               gdc.PtrBuiltInMethod
	methodIsReadOnlyFn                 gdc.PtrBuiltInMethod
	operatorNotFn                      gdc.PtrOperatorEvaluator
	operatorInDictionaryFn             gdc.PtrOperatorEvaluator
	operatorEqualArrayFn               gdc.PtrOperatorEvaluator
	operatorNotEqualArrayFn            gdc.PtrOperatorEvaluator
	operatorLessArrayFn                gdc.PtrOperatorEvaluator
	operatorLessEqualArrayFn           gdc.PtrOperatorEvaluator
	operatorGreaterArrayFn             gdc.PtrOperatorEvaluator
	operatorGreaterEqualArrayFn        gdc.PtrOperatorEvaluator
	operatorAddArrayFn                 gdc.PtrOperatorEvaluator
	operatorInArrayFn                  gdc.PtrOperatorEvaluator
	toVariantFn                        gdc.TypeFromVariantConstructorFunc
	fromVariantFn                      gdc.VariantFromTypeConstructorFunc
}

var ptrsForArray ptrsForArrayList

func initArrayPtrs(iface gdc.Interface) {
	ptrsForArray.ctrFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeArray, 0))
	ptrsForArray.ctrFromArrayFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeArray, 1))
	ptrsForArray.ctrFromArrayIntStringNameVariantFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeArray, 2))
	ptrsForArray.ctrFromPackedByteArrayFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeArray, 3))
	ptrsForArray.ctrFromPackedInt32ArrayFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeArray, 4))
	ptrsForArray.ctrFromPackedInt64ArrayFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeArray, 5))
	ptrsForArray.ctrFromPackedFloat32ArrayFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeArray, 6))
	ptrsForArray.ctrFromPackedFloat64ArrayFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeArray, 7))
	ptrsForArray.ctrFromPackedStringArrayFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeArray, 8))
	ptrsForArray.ctrFromPackedVector2ArrayFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeArray, 9))
	ptrsForArray.ctrFromPackedVector3ArrayFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeArray, 10))
	ptrsForArray.ctrFromPackedColorArrayFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeArray, 11))
	ptrsForArray.destructorFn = ensurePtr(iface.VariantGetPtrDestructor(gdc.VariantTypeArray))
	{
		methodName := StringNameFromStr("size")
		defer methodName.Destroy()
		ptrsForArray.methodSizeFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, methodName.AsCPtr(), 3173160232))
	}
	{
		methodName := StringNameFromStr("is_empty")
		defer methodName.Destroy()
		ptrsForArray.methodIsEmptyFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, methodName.AsCPtr(), 3918633141))
	}
	{
		methodName := StringNameFromStr("clear")
		defer methodName.Destroy()
		ptrsForArray.methodClearFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("hash")
		defer methodName.Destroy()
		ptrsForArray.methodHashFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, methodName.AsCPtr(), 3173160232))
	}
	{
		methodName := StringNameFromStr("assign")
		defer methodName.Destroy()
		ptrsForArray.methodAssignFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, methodName.AsCPtr(), 2307260970))
	}
	{
		methodName := StringNameFromStr("push_back")
		defer methodName.Destroy()
		ptrsForArray.methodPushBackFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, methodName.AsCPtr(), 3316032543))
	}
	{
		methodName := StringNameFromStr("push_front")
		defer methodName.Destroy()
		ptrsForArray.methodPushFrontFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, methodName.AsCPtr(), 3316032543))
	}
	{
		methodName := StringNameFromStr("append")
		defer methodName.Destroy()
		ptrsForArray.methodAppendFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, methodName.AsCPtr(), 3316032543))
	}
	{
		methodName := StringNameFromStr("append_array")
		defer methodName.Destroy()
		ptrsForArray.methodAppendArrayFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, methodName.AsCPtr(), 2307260970))
	}
	{
		methodName := StringNameFromStr("resize")
		defer methodName.Destroy()
		ptrsForArray.methodResizeFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, methodName.AsCPtr(), 848867239))
	}
	{
		methodName := StringNameFromStr("insert")
		defer methodName.Destroy()
		ptrsForArray.methodInsertFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, methodName.AsCPtr(), 3176316662))
	}
	{
		methodName := StringNameFromStr("remove_at")
		defer methodName.Destroy()
		ptrsForArray.methodRemoveAtFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, methodName.AsCPtr(), 2823966027))
	}
	{
		methodName := StringNameFromStr("fill")
		defer methodName.Destroy()
		ptrsForArray.methodFillFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, methodName.AsCPtr(), 3316032543))
	}
	{
		methodName := StringNameFromStr("erase")
		defer methodName.Destroy()
		ptrsForArray.methodEraseFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, methodName.AsCPtr(), 3316032543))
	}
	{
		methodName := StringNameFromStr("front")
		defer methodName.Destroy()
		ptrsForArray.methodFrontFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, methodName.AsCPtr(), 1460142086))
	}
	{
		methodName := StringNameFromStr("back")
		defer methodName.Destroy()
		ptrsForArray.methodBackFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, methodName.AsCPtr(), 1460142086))
	}
	{
		methodName := StringNameFromStr("pick_random")
		defer methodName.Destroy()
		ptrsForArray.methodPickRandomFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, methodName.AsCPtr(), 1460142086))
	}
	{
		methodName := StringNameFromStr("find")
		defer methodName.Destroy()
		ptrsForArray.methodFindFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, methodName.AsCPtr(), 2336346817))
	}
	{
		methodName := StringNameFromStr("rfind")
		defer methodName.Destroy()
		ptrsForArray.methodRfindFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, methodName.AsCPtr(), 2336346817))
	}
	{
		methodName := StringNameFromStr("count")
		defer methodName.Destroy()
		ptrsForArray.methodCountFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, methodName.AsCPtr(), 1481661226))
	}
	{
		methodName := StringNameFromStr("has")
		defer methodName.Destroy()
		ptrsForArray.methodHasFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, methodName.AsCPtr(), 3680194679))
	}
	{
		methodName := StringNameFromStr("pop_back")
		defer methodName.Destroy()
		ptrsForArray.methodPopBackFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, methodName.AsCPtr(), 1321915136))
	}
	{
		methodName := StringNameFromStr("pop_front")
		defer methodName.Destroy()
		ptrsForArray.methodPopFrontFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, methodName.AsCPtr(), 1321915136))
	}
	{
		methodName := StringNameFromStr("pop_at")
		defer methodName.Destroy()
		ptrsForArray.methodPopAtFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, methodName.AsCPtr(), 3518259424))
	}
	{
		methodName := StringNameFromStr("sort")
		defer methodName.Destroy()
		ptrsForArray.methodSortFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("sort_custom")
		defer methodName.Destroy()
		ptrsForArray.methodSortCustomFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, methodName.AsCPtr(), 3470848906))
	}
	{
		methodName := StringNameFromStr("shuffle")
		defer methodName.Destroy()
		ptrsForArray.methodShuffleFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("bsearch")
		defer methodName.Destroy()
		ptrsForArray.methodBsearchFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, methodName.AsCPtr(), 3372222236))
	}
	{
		methodName := StringNameFromStr("bsearch_custom")
		defer methodName.Destroy()
		ptrsForArray.methodBsearchCustomFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, methodName.AsCPtr(), 161317131))
	}
	{
		methodName := StringNameFromStr("reverse")
		defer methodName.Destroy()
		ptrsForArray.methodReverseFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("duplicate")
		defer methodName.Destroy()
		ptrsForArray.methodDuplicateFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, methodName.AsCPtr(), 636440122))
	}
	{
		methodName := StringNameFromStr("slice")
		defer methodName.Destroy()
		ptrsForArray.methodSliceFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, methodName.AsCPtr(), 1393718243))
	}
	{
		methodName := StringNameFromStr("filter")
		defer methodName.Destroy()
		ptrsForArray.methodFilterFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, methodName.AsCPtr(), 4075186556))
	}
	{
		methodName := StringNameFromStr("map")
		defer methodName.Destroy()
		ptrsForArray.methodMapFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, methodName.AsCPtr(), 4075186556))
	}
	{
		methodName := StringNameFromStr("reduce")
		defer methodName.Destroy()
		ptrsForArray.methodReduceFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, methodName.AsCPtr(), 4272450342))
	}
	{
		methodName := StringNameFromStr("any")
		defer methodName.Destroy()
		ptrsForArray.methodAnyFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, methodName.AsCPtr(), 4129521963))
	}
	{
		methodName := StringNameFromStr("all")
		defer methodName.Destroy()
		ptrsForArray.methodAllFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, methodName.AsCPtr(), 4129521963))
	}
	{
		methodName := StringNameFromStr("max")
		defer methodName.Destroy()
		ptrsForArray.methodMaxFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, methodName.AsCPtr(), 1460142086))
	}
	{
		methodName := StringNameFromStr("min")
		defer methodName.Destroy()
		ptrsForArray.methodMinFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, methodName.AsCPtr(), 1460142086))
	}
	{
		methodName := StringNameFromStr("is_typed")
		defer methodName.Destroy()
		ptrsForArray.methodIsTypedFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, methodName.AsCPtr(), 3918633141))
	}
	{
		methodName := StringNameFromStr("is_same_typed")
		defer methodName.Destroy()
		ptrsForArray.methodIsSameTypedFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, methodName.AsCPtr(), 2988181878))
	}
	{
		methodName := StringNameFromStr("get_typed_builtin")
		defer methodName.Destroy()
		ptrsForArray.methodGetTypedBuiltinFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, methodName.AsCPtr(), 3173160232))
	}
	{
		methodName := StringNameFromStr("get_typed_class_name")
		defer methodName.Destroy()
		ptrsForArray.methodGetTypedClassNameFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, methodName.AsCPtr(), 1825232092))
	}
	{
		methodName := StringNameFromStr("get_typed_script")
		defer methodName.Destroy()
		ptrsForArray.methodGetTypedScriptFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, methodName.AsCPtr(), 1460142086))
	}
	{
		methodName := StringNameFromStr("make_read_only")
		defer methodName.Destroy()
		ptrsForArray.methodMakeReadOnlyFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("is_read_only")
		defer methodName.Destroy()
		ptrsForArray.methodIsReadOnlyFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, methodName.AsCPtr(), 3918633141))
	}
	ptrsForArray.operatorNotFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, gdc.VariantTypeArray, gdc.VariantTypeNil))
	ptrsForArray.operatorInDictionaryFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeArray, gdc.VariantTypeDictionary))
	ptrsForArray.operatorEqualArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeArray, gdc.VariantTypeArray))
	ptrsForArray.operatorNotEqualArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeArray, gdc.VariantTypeArray))
	ptrsForArray.operatorLessArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpLess, gdc.VariantTypeArray, gdc.VariantTypeArray))
	ptrsForArray.operatorLessEqualArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpLessEqual, gdc.VariantTypeArray, gdc.VariantTypeArray))
	ptrsForArray.operatorGreaterArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpGreater, gdc.VariantTypeArray, gdc.VariantTypeArray))
	ptrsForArray.operatorGreaterEqualArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpGreaterEqual, gdc.VariantTypeArray, gdc.VariantTypeArray))
	ptrsForArray.operatorAddArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpAdd, gdc.VariantTypeArray, gdc.VariantTypeArray))
	ptrsForArray.operatorInArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeArray, gdc.VariantTypeArray))
	ptrsForArray.toVariantFn = ensurePtr(iface.GetVariantToTypeConstructor(gdc.VariantTypeArray))
	ptrsForArray.fromVariantFn = ensurePtr(iface.GetVariantFromTypeConstructor(gdc.VariantTypeArray))
}

type Array struct {
	data   *[classSizeArray]byte
	iface  gdc.Interface
	pinner runtime.Pinner
}

// Enums

// Constructors
func newArray() *Array {
	me := &Array{
		data:  new([classSizeArray]byte),
		iface: giface,
	}
	me.pinner.Pin(me)
	me.pinner.Pin(me.AsTypePtr())
	return me
}

func NewArray() *Array {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	me := newArray()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForArray.ctrFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{}))
	return me
}

func NewArrayFromArray(from Array) *Array {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	me := newArray()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForArray.ctrFromArrayFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr()}))
	return me
}

func NewArrayFromArrayIntStringNameVariant(base Array, type_ int64, class_name StringName, script Variant) *Array {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	pinner.Pin(&type_)
	me := newArray()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForArray.ctrFromArrayIntStringNameVariantFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{base.AsCTypePtr(), gdc.ConstTypePtr(&type_), class_name.AsCTypePtr(), script.AsCTypePtr()}))
	return me
}

func NewArrayFromPackedByteArray(from PackedByteArray) *Array {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	me := newArray()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForArray.ctrFromPackedByteArrayFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr()}))
	return me
}

func NewArrayFromPackedInt32Array(from PackedInt32Array) *Array {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	me := newArray()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForArray.ctrFromPackedInt32ArrayFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr()}))
	return me
}

func NewArrayFromPackedInt64Array(from PackedInt64Array) *Array {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	me := newArray()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForArray.ctrFromPackedInt64ArrayFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr()}))
	return me
}

func NewArrayFromPackedFloat32Array(from PackedFloat32Array) *Array {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	me := newArray()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForArray.ctrFromPackedFloat32ArrayFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr()}))
	return me
}

func NewArrayFromPackedFloat64Array(from PackedFloat64Array) *Array {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	me := newArray()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForArray.ctrFromPackedFloat64ArrayFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr()}))
	return me
}

func NewArrayFromPackedStringArray(from PackedStringArray) *Array {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	me := newArray()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForArray.ctrFromPackedStringArrayFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr()}))
	return me
}

func NewArrayFromPackedVector2Array(from PackedVector2Array) *Array {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	me := newArray()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForArray.ctrFromPackedVector2ArrayFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr()}))
	return me
}

func NewArrayFromPackedVector3Array(from PackedVector3Array) *Array {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	me := newArray()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForArray.ctrFromPackedVector3ArrayFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr()}))
	return me
}

func NewArrayFromPackedColorArray(from PackedColorArray) *Array {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	me := newArray()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForArray.ctrFromPackedColorArrayFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr()}))
	return me
}

// Destructor
func (me *Array) Destroy() {
	me.iface.CallPtrDestructor(ensurePtr(ptrsForArray.destructorFn), me.AsTypePtr())
	me.pinner.Unpin()
}

// Variant support
func (me *Variant) AsArray() (*Array, error) {
	if me.Type() != gdc.VariantTypeArray {
		return nil, fmt.Errorf("variant is not a Array")
	}
	bclass := newArray()
	me.iface.CallTypeFromVariantConstructorFunc(ensurePtr(ptrsForArray.toVariantFn), bclass.asUninitialized(), me.AsPtr())
	return bclass, nil
}

func (me *Array) AsVariant() *Variant {
	va := newVariant()
	va.inner = me
	me.iface.CallVariantFromTypeConstructorFunc(ensurePtr(ptrsForArray.fromVariantFn), va.asUninitialized(), me.AsTypePtr())
	return va
}

// Pointers
func ArrayFromPtr(ptr gdc.ConstTypePtr) *Array {
	me := newArray()
	dataFromPtr(me.data[:], ptr)
	return me
}

func (me *Array) ToTypePtr(ptr gdc.TypePtr) {
	dataToPtr(ptr, me.data[:])
}

func (me *Array) Type() gdc.VariantType {
	return gdc.VariantTypeArray
}

func (me *Array) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(unsafe.Pointer(me.data))
}

func (me *Array) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.AsTypePtr())
}

func (me *Array) asUninitialized() gdc.UninitializedTypePtr {
	return gdc.UninitializedTypePtr(me.AsTypePtr())
}

// Methods

func (me *Array) Size() int64 {
	ret := NewInt()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForArray.methodSizeFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Array) IsEmpty() bool {
	ret := NewBool()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForArray.methodIsEmptyFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Array) Clear() {
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForArray.methodClearFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *Array) Hash() int64 {
	ret := NewInt()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForArray.methodHashFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Array) Assign(array Array) {

	args := []gdc.ConstTypePtr{array.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForArray.methodAssignFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *Array) PushBack(value Variant) {

	args := []gdc.ConstTypePtr{value.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForArray.methodPushBackFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *Array) PushFront(value Variant) {

	args := []gdc.ConstTypePtr{value.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForArray.methodPushFrontFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *Array) Append(value Variant) {

	args := []gdc.ConstTypePtr{value.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForArray.methodAppendFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *Array) AppendArray(array Array) {

	args := []gdc.ConstTypePtr{array.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForArray.methodAppendArrayFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *Array) Resize(size int64) int64 {
	ret := NewInt()
	defer ret.Destroy()
	varg0 := NewIntFromInt(size)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForArray.methodResizeFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Array) Insert(position int64, value Variant) int64 {
	ret := NewInt()
	defer ret.Destroy()
	varg0 := NewIntFromInt(position)
	defer varg0.Destroy()

	args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), value.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForArray.methodInsertFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Array) RemoveAt(position int64) {
	varg0 := NewIntFromInt(position)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForArray.methodRemoveAtFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *Array) Fill(value Variant) {

	args := []gdc.ConstTypePtr{value.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForArray.methodFillFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *Array) Erase(value Variant) {

	args := []gdc.ConstTypePtr{value.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForArray.methodEraseFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *Array) Front() Variant {
	ret := NewVariant()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForArray.methodFrontFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Array) Back() Variant {
	ret := NewVariant()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForArray.methodBackFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Array) PickRandom() Variant {
	ret := NewVariant()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForArray.methodPickRandomFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Array) Find(what Variant, from int64) int64 {
	ret := NewInt()
	defer ret.Destroy()

	varg1 := NewIntFromInt(from)
	defer varg1.Destroy()
	args := []gdc.ConstTypePtr{what.AsCTypePtr(), varg1.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForArray.methodFindFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Array) Rfind(what Variant, from int64) int64 {
	ret := NewInt()
	defer ret.Destroy()

	varg1 := NewIntFromInt(from)
	defer varg1.Destroy()
	args := []gdc.ConstTypePtr{what.AsCTypePtr(), varg1.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForArray.methodRfindFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Array) Count(value Variant) int64 {
	ret := NewInt()
	defer ret.Destroy()

	args := []gdc.ConstTypePtr{value.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForArray.methodCountFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Array) Has(value Variant) bool {
	ret := NewBool()
	defer ret.Destroy()

	args := []gdc.ConstTypePtr{value.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForArray.methodHasFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Array) PopBack() Variant {
	ret := NewVariant()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForArray.methodPopBackFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Array) PopFront() Variant {
	ret := NewVariant()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForArray.methodPopFrontFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Array) PopAt(position int64) Variant {
	ret := NewVariant()

	varg0 := NewIntFromInt(position)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForArray.methodPopAtFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Array) Sort() {
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForArray.methodSortFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *Array) SortCustom(func_ Callable) {

	args := []gdc.ConstTypePtr{func_.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForArray.methodSortCustomFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *Array) Shuffle() {
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForArray.methodShuffleFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *Array) Bsearch(value Variant, before bool) int64 {
	ret := NewInt()
	defer ret.Destroy()

	varg1 := NewBoolFromBool(before)
	defer varg1.Destroy()
	args := []gdc.ConstTypePtr{value.AsCTypePtr(), varg1.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForArray.methodBsearchFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Array) BsearchCustom(value Variant, func_ Callable, before bool) int64 {
	ret := NewInt()
	defer ret.Destroy()

	varg2 := NewBoolFromBool(before)
	defer varg2.Destroy()
	args := []gdc.ConstTypePtr{value.AsCTypePtr(), func_.AsCTypePtr(), varg2.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForArray.methodBsearchCustomFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Array) Reverse() {
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForArray.methodReverseFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *Array) Duplicate(deep bool) Array {
	ret := NewArray()

	varg0 := NewBoolFromBool(deep)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForArray.methodDuplicateFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Array) Slice(begin int64, end int64, step int64, deep bool) Array {
	ret := NewArray()

	varg0 := NewIntFromInt(begin)
	defer varg0.Destroy()
	varg1 := NewIntFromInt(end)
	defer varg1.Destroy()
	varg2 := NewIntFromInt(step)
	defer varg2.Destroy()
	varg3 := NewBoolFromBool(deep)
	defer varg3.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), varg2.AsCTypePtr(), varg3.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForArray.methodSliceFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Array) Filter(method Callable) Array {
	ret := NewArray()

	args := []gdc.ConstTypePtr{method.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForArray.methodFilterFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Array) Map(method Callable) Array {
	ret := NewArray()

	args := []gdc.ConstTypePtr{method.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForArray.methodMapFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Array) Reduce(method Callable, accum Variant) Variant {
	ret := NewVariant()

	args := []gdc.ConstTypePtr{method.AsCTypePtr(), accum.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForArray.methodReduceFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Array) Any(method Callable) bool {
	ret := NewBool()
	defer ret.Destroy()

	args := []gdc.ConstTypePtr{method.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForArray.methodAnyFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Array) All(method Callable) bool {
	ret := NewBool()
	defer ret.Destroy()

	args := []gdc.ConstTypePtr{method.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForArray.methodAllFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Array) Max() Variant {
	ret := NewVariant()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForArray.methodMaxFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Array) Min() Variant {
	ret := NewVariant()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForArray.methodMinFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Array) IsTyped() bool {
	ret := NewBool()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForArray.methodIsTypedFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Array) IsSameTyped(array Array) bool {
	ret := NewBool()
	defer ret.Destroy()

	args := []gdc.ConstTypePtr{array.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForArray.methodIsSameTypedFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Array) GetTypedBuiltin() int64 {
	ret := NewInt()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForArray.methodGetTypedBuiltinFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *Array) GetTypedClassName() StringName {
	ret := NewStringName()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForArray.methodGetTypedClassNameFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Array) GetTypedScript() Variant {
	ret := NewVariant()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForArray.methodGetTypedScriptFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *Array) MakeReadOnly() {
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForArray.methodMakeReadOnlyFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *Array) IsReadOnly() bool {
	ret := NewBool()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForArray.methodIsReadOnlyFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

// Operators

func (me *Array) EqualVariant(right Variant) bool {
	opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type())
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Array) NotEqualVariant(right Variant) bool {
	opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type())
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Array) Not() bool {
	opPtr := ptrsForArray.operatorNotFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), nil, ret.AsTypePtr())
	return ret.Get()
}

func (me *Array) InDictionary(right Dictionary) bool {
	opPtr := ptrsForArray.operatorInDictionaryFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Array) EqualArray(right Array) bool {
	opPtr := ptrsForArray.operatorEqualArrayFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Array) NotEqualArray(right Array) bool {
	opPtr := ptrsForArray.operatorNotEqualArrayFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Array) LessArray(right Array) bool {
	opPtr := ptrsForArray.operatorLessArrayFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Array) LessEqualArray(right Array) bool {
	opPtr := ptrsForArray.operatorLessEqualArrayFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Array) GreaterArray(right Array) bool {
	opPtr := ptrsForArray.operatorGreaterArrayFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Array) GreaterEqualArray(right Array) bool {
	opPtr := ptrsForArray.operatorGreaterEqualArrayFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Array) AddArray(right Array) Array {
	opPtr := ptrsForArray.operatorAddArrayFn
	ret := NewArray()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Array) InArray(right Array) bool {
	opPtr := ptrsForArray.operatorInArrayFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

// Members
