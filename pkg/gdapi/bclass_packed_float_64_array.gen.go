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

type ptrsForPackedFloat64ArrayList struct {
	ctrFn                                gdc.PtrConstructor
	ctrFromPackedFloat64ArrayFn          gdc.PtrConstructor
	ctrFromArrayFn                       gdc.PtrConstructor
	destructorFn                         gdc.PtrDestructor
	methodSizeFn                         gdc.PtrBuiltInMethod
	methodIsEmptyFn                      gdc.PtrBuiltInMethod
	methodSetFn                          gdc.PtrBuiltInMethod
	methodPushBackFn                     gdc.PtrBuiltInMethod
	methodAppendFn                       gdc.PtrBuiltInMethod
	methodAppendArrayFn                  gdc.PtrBuiltInMethod
	methodRemoveAtFn                     gdc.PtrBuiltInMethod
	methodInsertFn                       gdc.PtrBuiltInMethod
	methodFillFn                         gdc.PtrBuiltInMethod
	methodResizeFn                       gdc.PtrBuiltInMethod
	methodClearFn                        gdc.PtrBuiltInMethod
	methodHasFn                          gdc.PtrBuiltInMethod
	methodReverseFn                      gdc.PtrBuiltInMethod
	methodSliceFn                        gdc.PtrBuiltInMethod
	methodToByteArrayFn                  gdc.PtrBuiltInMethod
	methodSortFn                         gdc.PtrBuiltInMethod
	methodBsearchFn                      gdc.PtrBuiltInMethod
	methodDuplicateFn                    gdc.PtrBuiltInMethod
	methodFindFn                         gdc.PtrBuiltInMethod
	methodRfindFn                        gdc.PtrBuiltInMethod
	methodCountFn                        gdc.PtrBuiltInMethod
	operatorNotFn                        gdc.PtrOperatorEvaluator
	operatorInDictionaryFn               gdc.PtrOperatorEvaluator
	operatorInArrayFn                    gdc.PtrOperatorEvaluator
	operatorEqualPackedFloat64ArrayFn    gdc.PtrOperatorEvaluator
	operatorNotEqualPackedFloat64ArrayFn gdc.PtrOperatorEvaluator
	operatorAddPackedFloat64ArrayFn      gdc.PtrOperatorEvaluator
	toVariantFn                          gdc.TypeFromVariantConstructorFunc
	fromVariantFn                        gdc.VariantFromTypeConstructorFunc
}

var ptrsForPackedFloat64Array ptrsForPackedFloat64ArrayList

func initPackedFloat64ArrayPtrs(iface gdc.Interface) {
	ptrsForPackedFloat64Array.ctrFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypePackedFloat64Array, 0))
	ptrsForPackedFloat64Array.ctrFromPackedFloat64ArrayFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypePackedFloat64Array, 1))
	ptrsForPackedFloat64Array.ctrFromArrayFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypePackedFloat64Array, 2))
	ptrsForPackedFloat64Array.destructorFn = ensurePtr(iface.VariantGetPtrDestructor(gdc.VariantTypePackedFloat64Array))
	{
		methodName := StringNameFromStr("size")
		defer methodName.Destroy()
		ptrsForPackedFloat64Array.methodSizeFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat64Array, methodName.AsCPtr(), 3173160232))
	}
	{
		methodName := StringNameFromStr("is_empty")
		defer methodName.Destroy()
		ptrsForPackedFloat64Array.methodIsEmptyFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat64Array, methodName.AsCPtr(), 3918633141))
	}
	{
		methodName := StringNameFromStr("set")
		defer methodName.Destroy()
		ptrsForPackedFloat64Array.methodSetFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat64Array, methodName.AsCPtr(), 1113000516))
	}
	{
		methodName := StringNameFromStr("push_back")
		defer methodName.Destroy()
		ptrsForPackedFloat64Array.methodPushBackFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat64Array, methodName.AsCPtr(), 4094791666))
	}
	{
		methodName := StringNameFromStr("append")
		defer methodName.Destroy()
		ptrsForPackedFloat64Array.methodAppendFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat64Array, methodName.AsCPtr(), 4094791666))
	}
	{
		methodName := StringNameFromStr("append_array")
		defer methodName.Destroy()
		ptrsForPackedFloat64Array.methodAppendArrayFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat64Array, methodName.AsCPtr(), 792078629))
	}
	{
		methodName := StringNameFromStr("remove_at")
		defer methodName.Destroy()
		ptrsForPackedFloat64Array.methodRemoveAtFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat64Array, methodName.AsCPtr(), 2823966027))
	}
	{
		methodName := StringNameFromStr("insert")
		defer methodName.Destroy()
		ptrsForPackedFloat64Array.methodInsertFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat64Array, methodName.AsCPtr(), 1379903876))
	}
	{
		methodName := StringNameFromStr("fill")
		defer methodName.Destroy()
		ptrsForPackedFloat64Array.methodFillFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat64Array, methodName.AsCPtr(), 833936903))
	}
	{
		methodName := StringNameFromStr("resize")
		defer methodName.Destroy()
		ptrsForPackedFloat64Array.methodResizeFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat64Array, methodName.AsCPtr(), 848867239))
	}
	{
		methodName := StringNameFromStr("clear")
		defer methodName.Destroy()
		ptrsForPackedFloat64Array.methodClearFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat64Array, methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("has")
		defer methodName.Destroy()
		ptrsForPackedFloat64Array.methodHasFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat64Array, methodName.AsCPtr(), 1296369134))
	}
	{
		methodName := StringNameFromStr("reverse")
		defer methodName.Destroy()
		ptrsForPackedFloat64Array.methodReverseFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat64Array, methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("slice")
		defer methodName.Destroy()
		ptrsForPackedFloat64Array.methodSliceFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat64Array, methodName.AsCPtr(), 2192974324))
	}
	{
		methodName := StringNameFromStr("to_byte_array")
		defer methodName.Destroy()
		ptrsForPackedFloat64Array.methodToByteArrayFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat64Array, methodName.AsCPtr(), 247621236))
	}
	{
		methodName := StringNameFromStr("sort")
		defer methodName.Destroy()
		ptrsForPackedFloat64Array.methodSortFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat64Array, methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("bsearch")
		defer methodName.Destroy()
		ptrsForPackedFloat64Array.methodBsearchFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat64Array, methodName.AsCPtr(), 1188816338))
	}
	{
		methodName := StringNameFromStr("duplicate")
		defer methodName.Destroy()
		ptrsForPackedFloat64Array.methodDuplicateFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat64Array, methodName.AsCPtr(), 949266573))
	}
	{
		methodName := StringNameFromStr("find")
		defer methodName.Destroy()
		ptrsForPackedFloat64Array.methodFindFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat64Array, methodName.AsCPtr(), 1343150241))
	}
	{
		methodName := StringNameFromStr("rfind")
		defer methodName.Destroy()
		ptrsForPackedFloat64Array.methodRfindFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat64Array, methodName.AsCPtr(), 1343150241))
	}
	{
		methodName := StringNameFromStr("count")
		defer methodName.Destroy()
		ptrsForPackedFloat64Array.methodCountFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat64Array, methodName.AsCPtr(), 2859915090))
	}
	ptrsForPackedFloat64Array.operatorNotFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, gdc.VariantTypePackedFloat64Array, gdc.VariantTypeNil))
	ptrsForPackedFloat64Array.operatorInDictionaryFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypePackedFloat64Array, gdc.VariantTypeDictionary))
	ptrsForPackedFloat64Array.operatorInArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypePackedFloat64Array, gdc.VariantTypeArray))
	ptrsForPackedFloat64Array.operatorEqualPackedFloat64ArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypePackedFloat64Array, gdc.VariantTypePackedFloat64Array))
	ptrsForPackedFloat64Array.operatorNotEqualPackedFloat64ArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypePackedFloat64Array, gdc.VariantTypePackedFloat64Array))
	ptrsForPackedFloat64Array.operatorAddPackedFloat64ArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpAdd, gdc.VariantTypePackedFloat64Array, gdc.VariantTypePackedFloat64Array))
	ptrsForPackedFloat64Array.toVariantFn = ensurePtr(iface.GetVariantToTypeConstructor(gdc.VariantTypePackedFloat64Array))
	ptrsForPackedFloat64Array.fromVariantFn = ensurePtr(iface.GetVariantFromTypeConstructor(gdc.VariantTypePackedFloat64Array))
}

type PackedFloat64Array struct {
	//data   *[classSizePackedFloat64Array]byte
	data   unsafe.Pointer
	iface  gdc.Interface
	pinner runtime.Pinner
}

// Enums

// Constructors
func newPackedFloat64Array() *PackedFloat64Array {
	me := &PackedFloat64Array{
		//data:   new([classSizePackedFloat64Array]byte),
		data:  giface.MemAlloc(classSizePackedFloat64Array),
		iface: giface,
	}
	me.pinner.Pin(me)
	me.pinner.Pin(me.AsTypePtr())
	return me
}

func NewPackedFloat64Array() *PackedFloat64Array {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	me := newPackedFloat64Array()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForPackedFloat64Array.ctrFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{}))
	return me
}

func NewPackedFloat64ArrayFromPackedFloat64Array(from PackedFloat64Array) *PackedFloat64Array {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	me := newPackedFloat64Array()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForPackedFloat64Array.ctrFromPackedFloat64ArrayFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr()}))
	return me
}

func NewPackedFloat64ArrayFromArray(from Array) *PackedFloat64Array {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	me := newPackedFloat64Array()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForPackedFloat64Array.ctrFromArrayFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr()}))
	return me
}

// Destructor
func (me *PackedFloat64Array) Destroy() {
	//me.iface.CallPtrDestructor(ensurePtr(ptrsForPackedFloat64Array.destructorFn), me.AsTypePtr())
	//me.pinner.Unpin()
}

// Variant support
func (me *Variant) AsPackedFloat64Array() (*PackedFloat64Array, error) {
	if me.Type() != gdc.VariantTypePackedFloat64Array {
		return nil, fmt.Errorf("variant is not a PackedFloat64Array")
	}
	bclass := newPackedFloat64Array()
	me.iface.CallTypeFromVariantConstructorFunc(ensurePtr(ptrsForPackedFloat64Array.toVariantFn), bclass.asUninitialized(), me.AsPtr())
	return bclass, nil
}

func (me *PackedFloat64Array) AsVariant() *Variant {
	va := newVariant()
	va.inner = me
	me.iface.CallVariantFromTypeConstructorFunc(ensurePtr(ptrsForPackedFloat64Array.fromVariantFn), va.asUninitialized(), me.AsTypePtr())
	return va
}

// Pointers
func PackedFloat64ArrayFromPtr(ptr gdc.ConstTypePtr) *PackedFloat64Array {
	me := newPackedFloat64Array()
	dataCopy(me.data, unsafe.Pointer(ptr), classSizePackedFloat64Array)
	return me
}

func (me *PackedFloat64Array) ToTypePtr(ptr gdc.TypePtr) {
	dataCopy(unsafe.Pointer(ptr), me.data, classSizePackedFloat64Array)
}

func (me *PackedFloat64Array) Type() gdc.VariantType {
	return gdc.VariantTypePackedFloat64Array
}

func (me *PackedFloat64Array) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.data)
}

func (me *PackedFloat64Array) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.AsTypePtr())
}

func (me *PackedFloat64Array) asUninitialized() gdc.UninitializedTypePtr {
	return gdc.UninitializedTypePtr(me.AsTypePtr())
}
func (me *PackedFloat64Array) Get(i int64) float64 {
	ret := me.iface.PackedFloat64ArrayOperatorIndexConst(me.AsCTypePtr(), gdc.Int(i))
	return *ret
}

// Methods

func (me *PackedFloat64Array) Size() int64 {
	ret := NewInt()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedFloat64Array.methodSizeFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *PackedFloat64Array) IsEmpty() bool {
	ret := NewBool()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedFloat64Array.methodIsEmptyFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *PackedFloat64Array) Set(index int64, value float64) {
	varg0 := NewIntFromInt(index)
	defer varg0.Destroy()
	varg1 := NewFloatFromFloat32(value)
	defer varg1.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedFloat64Array.methodSetFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedFloat64Array) PushBack(value float64) bool {
	ret := NewBool()
	defer ret.Destroy()
	varg0 := NewFloatFromFloat32(value)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedFloat64Array.methodPushBackFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *PackedFloat64Array) Append(value float64) bool {
	ret := NewBool()
	defer ret.Destroy()
	varg0 := NewFloatFromFloat32(value)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedFloat64Array.methodAppendFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *PackedFloat64Array) AppendArray(array PackedFloat64Array) {

	args := []gdc.ConstTypePtr{array.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedFloat64Array.methodAppendArrayFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedFloat64Array) RemoveAt(index int64) {
	varg0 := NewIntFromInt(index)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedFloat64Array.methodRemoveAtFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedFloat64Array) Insert(at_index int64, value float64) int64 {
	ret := NewInt()
	defer ret.Destroy()
	varg0 := NewIntFromInt(at_index)
	defer varg0.Destroy()
	varg1 := NewFloatFromFloat32(value)
	defer varg1.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedFloat64Array.methodInsertFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *PackedFloat64Array) Fill(value float64) {
	varg0 := NewFloatFromFloat32(value)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedFloat64Array.methodFillFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedFloat64Array) Resize(new_size int64) int64 {
	ret := NewInt()
	defer ret.Destroy()
	varg0 := NewIntFromInt(new_size)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedFloat64Array.methodResizeFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *PackedFloat64Array) Clear() {
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedFloat64Array.methodClearFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedFloat64Array) Has(value float64) bool {
	ret := NewBool()
	defer ret.Destroy()
	varg0 := NewFloatFromFloat32(value)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedFloat64Array.methodHasFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *PackedFloat64Array) Reverse() {
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedFloat64Array.methodReverseFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedFloat64Array) Slice(begin int64, end int64) PackedFloat64Array {
	ret := NewPackedFloat64Array()

	varg0 := NewIntFromInt(begin)
	defer varg0.Destroy()
	varg1 := NewIntFromInt(end)
	defer varg1.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedFloat64Array.methodSliceFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *PackedFloat64Array) ToByteArray() PackedByteArray {
	ret := NewPackedByteArray()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedFloat64Array.methodToByteArrayFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *PackedFloat64Array) Sort() {
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedFloat64Array.methodSortFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedFloat64Array) Bsearch(value float64, before bool) int64 {
	ret := NewInt()
	defer ret.Destroy()
	varg0 := NewFloatFromFloat32(value)
	defer varg0.Destroy()
	varg1 := NewBoolFromBool(before)
	defer varg1.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedFloat64Array.methodBsearchFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *PackedFloat64Array) Duplicate() PackedFloat64Array {
	ret := NewPackedFloat64Array()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedFloat64Array.methodDuplicateFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *PackedFloat64Array) Find(value float64, from int64) int64 {
	ret := NewInt()
	defer ret.Destroy()
	varg0 := NewFloatFromFloat32(value)
	defer varg0.Destroy()
	varg1 := NewIntFromInt(from)
	defer varg1.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedFloat64Array.methodFindFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *PackedFloat64Array) Rfind(value float64, from int64) int64 {
	ret := NewInt()
	defer ret.Destroy()
	varg0 := NewFloatFromFloat32(value)
	defer varg0.Destroy()
	varg1 := NewIntFromInt(from)
	defer varg1.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedFloat64Array.methodRfindFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *PackedFloat64Array) Count(value float64) int64 {
	ret := NewInt()
	defer ret.Destroy()
	varg0 := NewFloatFromFloat32(value)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedFloat64Array.methodCountFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

// Operators

func (me *PackedFloat64Array) EqualVariant(right Variant) bool {
	opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type())
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *PackedFloat64Array) NotEqualVariant(right Variant) bool {
	opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type())
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *PackedFloat64Array) Not() bool {
	opPtr := ptrsForPackedFloat64Array.operatorNotFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), nil, ret.AsTypePtr())
	return ret.Get()
}

func (me *PackedFloat64Array) InDictionary(right Dictionary) bool {
	opPtr := ptrsForPackedFloat64Array.operatorInDictionaryFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *PackedFloat64Array) InArray(right Array) bool {
	opPtr := ptrsForPackedFloat64Array.operatorInArrayFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *PackedFloat64Array) EqualPackedFloat64Array(right PackedFloat64Array) bool {
	opPtr := ptrsForPackedFloat64Array.operatorEqualPackedFloat64ArrayFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *PackedFloat64Array) NotEqualPackedFloat64Array(right PackedFloat64Array) bool {
	opPtr := ptrsForPackedFloat64Array.operatorNotEqualPackedFloat64ArrayFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *PackedFloat64Array) AddPackedFloat64Array(right PackedFloat64Array) PackedFloat64Array {
	opPtr := ptrsForPackedFloat64Array.operatorAddPackedFloat64ArrayFn
	ret := NewPackedFloat64Array()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

// Members
