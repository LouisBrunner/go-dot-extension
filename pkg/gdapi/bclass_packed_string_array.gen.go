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

type ptrsForPackedStringArrayList struct {
	ctrFn                               gdc.PtrConstructor
	ctrFromPackedStringArrayFn          gdc.PtrConstructor
	ctrFromArrayFn                      gdc.PtrConstructor
	destructorFn                        gdc.PtrDestructor
	methodSizeFn                        gdc.PtrBuiltInMethod
	methodIsEmptyFn                     gdc.PtrBuiltInMethod
	methodSetFn                         gdc.PtrBuiltInMethod
	methodPushBackFn                    gdc.PtrBuiltInMethod
	methodAppendFn                      gdc.PtrBuiltInMethod
	methodAppendArrayFn                 gdc.PtrBuiltInMethod
	methodRemoveAtFn                    gdc.PtrBuiltInMethod
	methodInsertFn                      gdc.PtrBuiltInMethod
	methodFillFn                        gdc.PtrBuiltInMethod
	methodResizeFn                      gdc.PtrBuiltInMethod
	methodClearFn                       gdc.PtrBuiltInMethod
	methodHasFn                         gdc.PtrBuiltInMethod
	methodReverseFn                     gdc.PtrBuiltInMethod
	methodSliceFn                       gdc.PtrBuiltInMethod
	methodToByteArrayFn                 gdc.PtrBuiltInMethod
	methodSortFn                        gdc.PtrBuiltInMethod
	methodBsearchFn                     gdc.PtrBuiltInMethod
	methodDuplicateFn                   gdc.PtrBuiltInMethod
	methodFindFn                        gdc.PtrBuiltInMethod
	methodRfindFn                       gdc.PtrBuiltInMethod
	methodCountFn                       gdc.PtrBuiltInMethod
	operatorNotFn                       gdc.PtrOperatorEvaluator
	operatorInDictionaryFn              gdc.PtrOperatorEvaluator
	operatorInArrayFn                   gdc.PtrOperatorEvaluator
	operatorEqualPackedStringArrayFn    gdc.PtrOperatorEvaluator
	operatorNotEqualPackedStringArrayFn gdc.PtrOperatorEvaluator
	operatorAddPackedStringArrayFn      gdc.PtrOperatorEvaluator
	toVariantFn                         gdc.TypeFromVariantConstructorFunc
	fromVariantFn                       gdc.VariantFromTypeConstructorFunc
}

var ptrsForPackedStringArray ptrsForPackedStringArrayList

func initPackedStringArrayPtrs(iface gdc.Interface) {
	ptrsForPackedStringArray.ctrFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypePackedStringArray, 0))
	ptrsForPackedStringArray.ctrFromPackedStringArrayFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypePackedStringArray, 1))
	ptrsForPackedStringArray.ctrFromArrayFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypePackedStringArray, 2))
	ptrsForPackedStringArray.destructorFn = ensurePtr(iface.VariantGetPtrDestructor(gdc.VariantTypePackedStringArray))
	{
		methodName := StringNameFromStr("size")
		defer methodName.Destroy()
		ptrsForPackedStringArray.methodSizeFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedStringArray, methodName.AsCPtr(), 3173160232))
	}
	{
		methodName := StringNameFromStr("is_empty")
		defer methodName.Destroy()
		ptrsForPackedStringArray.methodIsEmptyFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedStringArray, methodName.AsCPtr(), 3918633141))
	}
	{
		methodName := StringNameFromStr("set")
		defer methodName.Destroy()
		ptrsForPackedStringArray.methodSetFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedStringArray, methodName.AsCPtr(), 725585539))
	}
	{
		methodName := StringNameFromStr("push_back")
		defer methodName.Destroy()
		ptrsForPackedStringArray.methodPushBackFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedStringArray, methodName.AsCPtr(), 816187996))
	}
	{
		methodName := StringNameFromStr("append")
		defer methodName.Destroy()
		ptrsForPackedStringArray.methodAppendFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedStringArray, methodName.AsCPtr(), 816187996))
	}
	{
		methodName := StringNameFromStr("append_array")
		defer methodName.Destroy()
		ptrsForPackedStringArray.methodAppendArrayFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedStringArray, methodName.AsCPtr(), 1120103966))
	}
	{
		methodName := StringNameFromStr("remove_at")
		defer methodName.Destroy()
		ptrsForPackedStringArray.methodRemoveAtFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedStringArray, methodName.AsCPtr(), 2823966027))
	}
	{
		methodName := StringNameFromStr("insert")
		defer methodName.Destroy()
		ptrsForPackedStringArray.methodInsertFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedStringArray, methodName.AsCPtr(), 2432393153))
	}
	{
		methodName := StringNameFromStr("fill")
		defer methodName.Destroy()
		ptrsForPackedStringArray.methodFillFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedStringArray, methodName.AsCPtr(), 3174917410))
	}
	{
		methodName := StringNameFromStr("resize")
		defer methodName.Destroy()
		ptrsForPackedStringArray.methodResizeFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedStringArray, methodName.AsCPtr(), 848867239))
	}
	{
		methodName := StringNameFromStr("clear")
		defer methodName.Destroy()
		ptrsForPackedStringArray.methodClearFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedStringArray, methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("has")
		defer methodName.Destroy()
		ptrsForPackedStringArray.methodHasFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedStringArray, methodName.AsCPtr(), 2566493496))
	}
	{
		methodName := StringNameFromStr("reverse")
		defer methodName.Destroy()
		ptrsForPackedStringArray.methodReverseFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedStringArray, methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("slice")
		defer methodName.Destroy()
		ptrsForPackedStringArray.methodSliceFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedStringArray, methodName.AsCPtr(), 2094601407))
	}
	{
		methodName := StringNameFromStr("to_byte_array")
		defer methodName.Destroy()
		ptrsForPackedStringArray.methodToByteArrayFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedStringArray, methodName.AsCPtr(), 247621236))
	}
	{
		methodName := StringNameFromStr("sort")
		defer methodName.Destroy()
		ptrsForPackedStringArray.methodSortFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedStringArray, methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("bsearch")
		defer methodName.Destroy()
		ptrsForPackedStringArray.methodBsearchFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedStringArray, methodName.AsCPtr(), 328976671))
	}
	{
		methodName := StringNameFromStr("duplicate")
		defer methodName.Destroy()
		ptrsForPackedStringArray.methodDuplicateFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedStringArray, methodName.AsCPtr(), 2991231410))
	}
	{
		methodName := StringNameFromStr("find")
		defer methodName.Destroy()
		ptrsForPackedStringArray.methodFindFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedStringArray, methodName.AsCPtr(), 1760645412))
	}
	{
		methodName := StringNameFromStr("rfind")
		defer methodName.Destroy()
		ptrsForPackedStringArray.methodRfindFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedStringArray, methodName.AsCPtr(), 1760645412))
	}
	{
		methodName := StringNameFromStr("count")
		defer methodName.Destroy()
		ptrsForPackedStringArray.methodCountFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedStringArray, methodName.AsCPtr(), 2920860731))
	}
	ptrsForPackedStringArray.operatorNotFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, gdc.VariantTypePackedStringArray, gdc.VariantTypeNil))
	ptrsForPackedStringArray.operatorInDictionaryFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypePackedStringArray, gdc.VariantTypeDictionary))
	ptrsForPackedStringArray.operatorInArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypePackedStringArray, gdc.VariantTypeArray))
	ptrsForPackedStringArray.operatorEqualPackedStringArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypePackedStringArray, gdc.VariantTypePackedStringArray))
	ptrsForPackedStringArray.operatorNotEqualPackedStringArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypePackedStringArray, gdc.VariantTypePackedStringArray))
	ptrsForPackedStringArray.operatorAddPackedStringArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpAdd, gdc.VariantTypePackedStringArray, gdc.VariantTypePackedStringArray))
	ptrsForPackedStringArray.toVariantFn = ensurePtr(iface.GetVariantToTypeConstructor(gdc.VariantTypePackedStringArray))
	ptrsForPackedStringArray.fromVariantFn = ensurePtr(iface.GetVariantFromTypeConstructor(gdc.VariantTypePackedStringArray))
}

type PackedStringArray struct {
	//data   *[classSizePackedStringArray]byte
	data   unsafe.Pointer
	iface  gdc.Interface
	pinner runtime.Pinner
}

// Enums

// Constructors
func newPackedStringArray() *PackedStringArray {
	me := &PackedStringArray{
		//data:   new([classSizePackedStringArray]byte),
		data:  giface.MemAlloc(classSizePackedStringArray),
		iface: giface,
	}
	me.pinner.Pin(me)
	me.pinner.Pin(me.AsTypePtr())
	return me
}

func NewPackedStringArray() *PackedStringArray {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	me := newPackedStringArray()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForPackedStringArray.ctrFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{}))
	return me
}

func NewPackedStringArrayFromPackedStringArray(from PackedStringArray) *PackedStringArray {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	me := newPackedStringArray()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForPackedStringArray.ctrFromPackedStringArrayFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr()}))
	return me
}

func NewPackedStringArrayFromArray(from Array) *PackedStringArray {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	me := newPackedStringArray()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForPackedStringArray.ctrFromArrayFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr()}))
	return me
}

// Destructor
func (me *PackedStringArray) Destroy() {
	//me.iface.CallPtrDestructor(ensurePtr(ptrsForPackedStringArray.destructorFn), me.AsTypePtr())
	//me.pinner.Unpin()
}

// Variant support
func (me *Variant) AsPackedStringArray() (*PackedStringArray, error) {
	if me.Type() != gdc.VariantTypePackedStringArray {
		return nil, fmt.Errorf("variant is not a PackedStringArray")
	}
	bclass := newPackedStringArray()
	me.iface.CallTypeFromVariantConstructorFunc(ensurePtr(ptrsForPackedStringArray.toVariantFn), bclass.asUninitialized(), me.AsPtr())
	return bclass, nil
}

func (me *PackedStringArray) AsVariant() *Variant {
	va := newVariant()
	va.inner = me
	me.iface.CallVariantFromTypeConstructorFunc(ensurePtr(ptrsForPackedStringArray.fromVariantFn), va.asUninitialized(), me.AsTypePtr())
	return va
}

// Pointers
func PackedStringArrayFromPtr(ptr gdc.ConstTypePtr) *PackedStringArray {
	me := newPackedStringArray()
	dataCopy(me.data, unsafe.Pointer(ptr), classSizePackedStringArray)
	return me
}

func (me *PackedStringArray) ToTypePtr(ptr gdc.TypePtr) {
	dataCopy(unsafe.Pointer(ptr), me.data, classSizePackedStringArray)
}

func (me *PackedStringArray) Type() gdc.VariantType {
	return gdc.VariantTypePackedStringArray
}

func (me *PackedStringArray) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.data)
}

func (me *PackedStringArray) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.AsTypePtr())
}

func (me *PackedStringArray) asUninitialized() gdc.UninitializedTypePtr {
	return gdc.UninitializedTypePtr(me.AsTypePtr())
}
func (me *PackedStringArray) Get(i int64) string {
	ret := me.iface.PackedStringArrayOperatorIndexConst(me.AsCTypePtr(), gdc.Int(i))
	return StringFromPtr(gdc.ConstTypePtr(ret)).String()
}

// Methods

func (me *PackedStringArray) Size() int64 {
	ret := NewInt()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedStringArray.methodSizeFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *PackedStringArray) IsEmpty() bool {
	ret := NewBool()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedStringArray.methodIsEmptyFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *PackedStringArray) Set(index int64, value String) {
	varg0 := NewIntFromInt(index)
	defer varg0.Destroy()

	args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), value.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedStringArray.methodSetFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedStringArray) PushBack(value String) bool {
	ret := NewBool()
	defer ret.Destroy()

	args := []gdc.ConstTypePtr{value.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedStringArray.methodPushBackFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *PackedStringArray) Append(value String) bool {
	ret := NewBool()
	defer ret.Destroy()

	args := []gdc.ConstTypePtr{value.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedStringArray.methodAppendFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *PackedStringArray) AppendArray(array PackedStringArray) {

	args := []gdc.ConstTypePtr{array.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedStringArray.methodAppendArrayFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedStringArray) RemoveAt(index int64) {
	varg0 := NewIntFromInt(index)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedStringArray.methodRemoveAtFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedStringArray) Insert(at_index int64, value String) int64 {
	ret := NewInt()
	defer ret.Destroy()
	varg0 := NewIntFromInt(at_index)
	defer varg0.Destroy()

	args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), value.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedStringArray.methodInsertFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *PackedStringArray) Fill(value String) {

	args := []gdc.ConstTypePtr{value.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedStringArray.methodFillFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedStringArray) Resize(new_size int64) int64 {
	ret := NewInt()
	defer ret.Destroy()
	varg0 := NewIntFromInt(new_size)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedStringArray.methodResizeFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *PackedStringArray) Clear() {
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedStringArray.methodClearFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedStringArray) Has(value String) bool {
	ret := NewBool()
	defer ret.Destroy()

	args := []gdc.ConstTypePtr{value.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedStringArray.methodHasFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *PackedStringArray) Reverse() {
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedStringArray.methodReverseFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedStringArray) Slice(begin int64, end int64) PackedStringArray {
	ret := NewPackedStringArray()

	varg0 := NewIntFromInt(begin)
	defer varg0.Destroy()
	varg1 := NewIntFromInt(end)
	defer varg1.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedStringArray.methodSliceFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *PackedStringArray) ToByteArray() PackedByteArray {
	ret := NewPackedByteArray()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedStringArray.methodToByteArrayFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *PackedStringArray) Sort() {
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedStringArray.methodSortFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedStringArray) Bsearch(value String, before bool) int64 {
	ret := NewInt()
	defer ret.Destroy()

	varg1 := NewBoolFromBool(before)
	defer varg1.Destroy()
	args := []gdc.ConstTypePtr{value.AsCTypePtr(), varg1.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedStringArray.methodBsearchFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *PackedStringArray) Duplicate() PackedStringArray {
	ret := NewPackedStringArray()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedStringArray.methodDuplicateFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *PackedStringArray) Find(value String, from int64) int64 {
	ret := NewInt()
	defer ret.Destroy()

	varg1 := NewIntFromInt(from)
	defer varg1.Destroy()
	args := []gdc.ConstTypePtr{value.AsCTypePtr(), varg1.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedStringArray.methodFindFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *PackedStringArray) Rfind(value String, from int64) int64 {
	ret := NewInt()
	defer ret.Destroy()

	varg1 := NewIntFromInt(from)
	defer varg1.Destroy()
	args := []gdc.ConstTypePtr{value.AsCTypePtr(), varg1.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedStringArray.methodRfindFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *PackedStringArray) Count(value String) int64 {
	ret := NewInt()
	defer ret.Destroy()

	args := []gdc.ConstTypePtr{value.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedStringArray.methodCountFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

// Operators

func (me *PackedStringArray) EqualVariant(right Variant) bool {
	opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type())
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *PackedStringArray) NotEqualVariant(right Variant) bool {
	opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type())
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *PackedStringArray) Not() bool {
	opPtr := ptrsForPackedStringArray.operatorNotFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), nil, ret.AsTypePtr())
	return ret.Get()
}

func (me *PackedStringArray) InDictionary(right Dictionary) bool {
	opPtr := ptrsForPackedStringArray.operatorInDictionaryFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *PackedStringArray) InArray(right Array) bool {
	opPtr := ptrsForPackedStringArray.operatorInArrayFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *PackedStringArray) EqualPackedStringArray(right PackedStringArray) bool {
	opPtr := ptrsForPackedStringArray.operatorEqualPackedStringArrayFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *PackedStringArray) NotEqualPackedStringArray(right PackedStringArray) bool {
	opPtr := ptrsForPackedStringArray.operatorNotEqualPackedStringArrayFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *PackedStringArray) AddPackedStringArray(right PackedStringArray) PackedStringArray {
	opPtr := ptrsForPackedStringArray.operatorAddPackedStringArrayFn
	ret := NewPackedStringArray()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

// Members
