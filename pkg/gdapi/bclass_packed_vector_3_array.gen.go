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

type ptrsForPackedVector3ArrayList struct {
	ctrFn                                gdc.PtrConstructor
	ctrFromPackedVector3ArrayFn          gdc.PtrConstructor
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
	operatorMultiplyTransform3DFn        gdc.PtrOperatorEvaluator
	operatorInDictionaryFn               gdc.PtrOperatorEvaluator
	operatorInArrayFn                    gdc.PtrOperatorEvaluator
	operatorEqualPackedVector3ArrayFn    gdc.PtrOperatorEvaluator
	operatorNotEqualPackedVector3ArrayFn gdc.PtrOperatorEvaluator
	operatorAddPackedVector3ArrayFn      gdc.PtrOperatorEvaluator
	toVariantFn                          gdc.TypeFromVariantConstructorFunc
	fromVariantFn                        gdc.VariantFromTypeConstructorFunc
}

var ptrsForPackedVector3Array ptrsForPackedVector3ArrayList

func initPackedVector3ArrayPtrs(iface gdc.Interface) {
	ptrsForPackedVector3Array.ctrFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypePackedVector3Array, 0))
	ptrsForPackedVector3Array.ctrFromPackedVector3ArrayFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypePackedVector3Array, 1))
	ptrsForPackedVector3Array.ctrFromArrayFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypePackedVector3Array, 2))
	ptrsForPackedVector3Array.destructorFn = ensurePtr(iface.VariantGetPtrDestructor(gdc.VariantTypePackedVector3Array))
	{
		methodName := StringNameFromStr("size")
		defer methodName.Destroy()
		ptrsForPackedVector3Array.methodSizeFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedVector3Array, methodName.AsCPtr(), 3173160232))
	}
	{
		methodName := StringNameFromStr("is_empty")
		defer methodName.Destroy()
		ptrsForPackedVector3Array.methodIsEmptyFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedVector3Array, methodName.AsCPtr(), 3918633141))
	}
	{
		methodName := StringNameFromStr("set")
		defer methodName.Destroy()
		ptrsForPackedVector3Array.methodSetFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedVector3Array, methodName.AsCPtr(), 3975343409))
	}
	{
		methodName := StringNameFromStr("push_back")
		defer methodName.Destroy()
		ptrsForPackedVector3Array.methodPushBackFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedVector3Array, methodName.AsCPtr(), 3295363524))
	}
	{
		methodName := StringNameFromStr("append")
		defer methodName.Destroy()
		ptrsForPackedVector3Array.methodAppendFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedVector3Array, methodName.AsCPtr(), 3295363524))
	}
	{
		methodName := StringNameFromStr("append_array")
		defer methodName.Destroy()
		ptrsForPackedVector3Array.methodAppendArrayFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedVector3Array, methodName.AsCPtr(), 203538016))
	}
	{
		methodName := StringNameFromStr("remove_at")
		defer methodName.Destroy()
		ptrsForPackedVector3Array.methodRemoveAtFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedVector3Array, methodName.AsCPtr(), 2823966027))
	}
	{
		methodName := StringNameFromStr("insert")
		defer methodName.Destroy()
		ptrsForPackedVector3Array.methodInsertFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedVector3Array, methodName.AsCPtr(), 3892262309))
	}
	{
		methodName := StringNameFromStr("fill")
		defer methodName.Destroy()
		ptrsForPackedVector3Array.methodFillFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedVector3Array, methodName.AsCPtr(), 3726392409))
	}
	{
		methodName := StringNameFromStr("resize")
		defer methodName.Destroy()
		ptrsForPackedVector3Array.methodResizeFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedVector3Array, methodName.AsCPtr(), 848867239))
	}
	{
		methodName := StringNameFromStr("clear")
		defer methodName.Destroy()
		ptrsForPackedVector3Array.methodClearFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedVector3Array, methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("has")
		defer methodName.Destroy()
		ptrsForPackedVector3Array.methodHasFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedVector3Array, methodName.AsCPtr(), 1749054343))
	}
	{
		methodName := StringNameFromStr("reverse")
		defer methodName.Destroy()
		ptrsForPackedVector3Array.methodReverseFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedVector3Array, methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("slice")
		defer methodName.Destroy()
		ptrsForPackedVector3Array.methodSliceFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedVector3Array, methodName.AsCPtr(), 2086131305))
	}
	{
		methodName := StringNameFromStr("to_byte_array")
		defer methodName.Destroy()
		ptrsForPackedVector3Array.methodToByteArrayFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedVector3Array, methodName.AsCPtr(), 247621236))
	}
	{
		methodName := StringNameFromStr("sort")
		defer methodName.Destroy()
		ptrsForPackedVector3Array.methodSortFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedVector3Array, methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("bsearch")
		defer methodName.Destroy()
		ptrsForPackedVector3Array.methodBsearchFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedVector3Array, methodName.AsCPtr(), 219263630))
	}
	{
		methodName := StringNameFromStr("duplicate")
		defer methodName.Destroy()
		ptrsForPackedVector3Array.methodDuplicateFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedVector3Array, methodName.AsCPtr(), 2754175465))
	}
	{
		methodName := StringNameFromStr("find")
		defer methodName.Destroy()
		ptrsForPackedVector3Array.methodFindFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedVector3Array, methodName.AsCPtr(), 3718155780))
	}
	{
		methodName := StringNameFromStr("rfind")
		defer methodName.Destroy()
		ptrsForPackedVector3Array.methodRfindFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedVector3Array, methodName.AsCPtr(), 3718155780))
	}
	{
		methodName := StringNameFromStr("count")
		defer methodName.Destroy()
		ptrsForPackedVector3Array.methodCountFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedVector3Array, methodName.AsCPtr(), 194580386))
	}
	ptrsForPackedVector3Array.operatorNotFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, gdc.VariantTypePackedVector3Array, gdc.VariantTypeNil))
	ptrsForPackedVector3Array.operatorMultiplyTransform3DFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypePackedVector3Array, gdc.VariantTypeTransform3D))
	ptrsForPackedVector3Array.operatorInDictionaryFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypePackedVector3Array, gdc.VariantTypeDictionary))
	ptrsForPackedVector3Array.operatorInArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypePackedVector3Array, gdc.VariantTypeArray))
	ptrsForPackedVector3Array.operatorEqualPackedVector3ArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypePackedVector3Array, gdc.VariantTypePackedVector3Array))
	ptrsForPackedVector3Array.operatorNotEqualPackedVector3ArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypePackedVector3Array, gdc.VariantTypePackedVector3Array))
	ptrsForPackedVector3Array.operatorAddPackedVector3ArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpAdd, gdc.VariantTypePackedVector3Array, gdc.VariantTypePackedVector3Array))
	ptrsForPackedVector3Array.toVariantFn = ensurePtr(iface.GetVariantToTypeConstructor(gdc.VariantTypePackedVector3Array))
	ptrsForPackedVector3Array.fromVariantFn = ensurePtr(iface.GetVariantFromTypeConstructor(gdc.VariantTypePackedVector3Array))
}

type PackedVector3Array struct {
	//data   *[classSizePackedVector3Array]byte
	data   unsafe.Pointer
	iface  gdc.Interface
	pinner runtime.Pinner
}

// Enums

// Constructors
func newPackedVector3Array() *PackedVector3Array {
	me := &PackedVector3Array{
		//data:   new([classSizePackedVector3Array]byte),
		data:  giface.MemAlloc(classSizePackedVector3Array),
		iface: giface,
	}
	me.pinner.Pin(me)
	me.pinner.Pin(me.AsTypePtr())
	return me
}

func NewPackedVector3Array() *PackedVector3Array {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	me := newPackedVector3Array()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForPackedVector3Array.ctrFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{}))
	return me
}

func NewPackedVector3ArrayFromPackedVector3Array(from PackedVector3Array) *PackedVector3Array {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	me := newPackedVector3Array()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForPackedVector3Array.ctrFromPackedVector3ArrayFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr()}))
	return me
}

func NewPackedVector3ArrayFromArray(from Array) *PackedVector3Array {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	me := newPackedVector3Array()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForPackedVector3Array.ctrFromArrayFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr()}))
	return me
}

// Destructor
func (me *PackedVector3Array) Destroy() {
	//me.iface.CallPtrDestructor(ensurePtr(ptrsForPackedVector3Array.destructorFn), me.AsTypePtr())
	//me.pinner.Unpin()
}

// Variant support
func (me *Variant) AsPackedVector3Array() (*PackedVector3Array, error) {
	if me.Type() != gdc.VariantTypePackedVector3Array {
		return nil, fmt.Errorf("variant is not a PackedVector3Array")
	}
	bclass := newPackedVector3Array()
	me.iface.CallTypeFromVariantConstructorFunc(ensurePtr(ptrsForPackedVector3Array.toVariantFn), bclass.asUninitialized(), me.AsPtr())
	return bclass, nil
}

func (me *PackedVector3Array) AsVariant() *Variant {
	va := newVariant()
	va.inner = me
	me.iface.CallVariantFromTypeConstructorFunc(ensurePtr(ptrsForPackedVector3Array.fromVariantFn), va.asUninitialized(), me.AsTypePtr())
	return va
}

// Pointers
func PackedVector3ArrayFromPtr(ptr gdc.ConstTypePtr) *PackedVector3Array {
	me := newPackedVector3Array()
	dataCopy(me.data, unsafe.Pointer(ptr), classSizePackedVector3Array)
	return me
}

func (me *PackedVector3Array) ToTypePtr(ptr gdc.TypePtr) {
	dataCopy(unsafe.Pointer(ptr), me.data, classSizePackedVector3Array)
}

func (me *PackedVector3Array) Type() gdc.VariantType {
	return gdc.VariantTypePackedVector3Array
}

func (me *PackedVector3Array) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.data)
}

func (me *PackedVector3Array) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.AsTypePtr())
}

func (me *PackedVector3Array) asUninitialized() gdc.UninitializedTypePtr {
	return gdc.UninitializedTypePtr(me.AsTypePtr())
}
func (me *PackedVector3Array) Get(i int64) Vector3 {
	ret := me.iface.PackedVector3ArrayOperatorIndexConst(me.AsCTypePtr(), gdc.Int(i))
	return *Vector3FromPtr(gdc.ConstTypePtr(ret))
}

// Methods

func (me *PackedVector3Array) Size() int64 {
	ret := NewInt()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedVector3Array.methodSizeFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *PackedVector3Array) IsEmpty() bool {
	ret := NewBool()
	defer ret.Destroy()
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedVector3Array.methodIsEmptyFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *PackedVector3Array) Set(index int64, value Vector3) {
	varg0 := NewIntFromInt(index)
	defer varg0.Destroy()

	args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), value.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedVector3Array.methodSetFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedVector3Array) PushBack(value Vector3) bool {
	ret := NewBool()
	defer ret.Destroy()

	args := []gdc.ConstTypePtr{value.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedVector3Array.methodPushBackFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *PackedVector3Array) Append(value Vector3) bool {
	ret := NewBool()
	defer ret.Destroy()

	args := []gdc.ConstTypePtr{value.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedVector3Array.methodAppendFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *PackedVector3Array) AppendArray(array PackedVector3Array) {

	args := []gdc.ConstTypePtr{array.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedVector3Array.methodAppendArrayFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedVector3Array) RemoveAt(index int64) {
	varg0 := NewIntFromInt(index)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedVector3Array.methodRemoveAtFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedVector3Array) Insert(at_index int64, value Vector3) int64 {
	ret := NewInt()
	defer ret.Destroy()
	varg0 := NewIntFromInt(at_index)
	defer varg0.Destroy()

	args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), value.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedVector3Array.methodInsertFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *PackedVector3Array) Fill(value Vector3) {

	args := []gdc.ConstTypePtr{value.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedVector3Array.methodFillFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedVector3Array) Resize(new_size int64) int64 {
	ret := NewInt()
	defer ret.Destroy()
	varg0 := NewIntFromInt(new_size)
	defer varg0.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedVector3Array.methodResizeFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *PackedVector3Array) Clear() {
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedVector3Array.methodClearFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedVector3Array) Has(value Vector3) bool {
	ret := NewBool()
	defer ret.Destroy()

	args := []gdc.ConstTypePtr{value.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedVector3Array.methodHasFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *PackedVector3Array) Reverse() {
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedVector3Array.methodReverseFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedVector3Array) Slice(begin int64, end int64) PackedVector3Array {
	ret := NewPackedVector3Array()

	varg0 := NewIntFromInt(begin)
	defer varg0.Destroy()
	varg1 := NewIntFromInt(end)
	defer varg1.Destroy()
	args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedVector3Array.methodSliceFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *PackedVector3Array) ToByteArray() PackedByteArray {
	ret := NewPackedByteArray()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedVector3Array.methodToByteArrayFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *PackedVector3Array) Sort() {
	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedVector3Array.methodSortFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedVector3Array) Bsearch(value Vector3, before bool) int64 {
	ret := NewInt()
	defer ret.Destroy()

	varg1 := NewBoolFromBool(before)
	defer varg1.Destroy()
	args := []gdc.ConstTypePtr{value.AsCTypePtr(), varg1.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedVector3Array.methodBsearchFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *PackedVector3Array) Duplicate() PackedVector3Array {
	ret := NewPackedVector3Array()

	args := []gdc.ConstTypePtr{}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedVector3Array.methodDuplicateFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return *ret
}

func (me *PackedVector3Array) Find(value Vector3, from int64) int64 {
	ret := NewInt()
	defer ret.Destroy()

	varg1 := NewIntFromInt(from)
	defer varg1.Destroy()
	args := []gdc.ConstTypePtr{value.AsCTypePtr(), varg1.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedVector3Array.methodFindFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *PackedVector3Array) Rfind(value Vector3, from int64) int64 {
	ret := NewInt()
	defer ret.Destroy()

	varg1 := NewIntFromInt(from)
	defer varg1.Destroy()
	args := []gdc.ConstTypePtr{value.AsCTypePtr(), varg1.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedVector3Array.methodRfindFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

func (me *PackedVector3Array) Count(value Vector3) int64 {
	ret := NewInt()
	defer ret.Destroy()

	args := []gdc.ConstTypePtr{value.AsCTypePtr()}

	giface.CallPtrBuiltInMethod(ensurePtr(ptrsForPackedVector3Array.methodCountFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
	return ret.Get()
}

// Operators

func (me *PackedVector3Array) EqualVariant(right Variant) bool {
	opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type())
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *PackedVector3Array) NotEqualVariant(right Variant) bool {
	opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type())
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *PackedVector3Array) Not() bool {
	opPtr := ptrsForPackedVector3Array.operatorNotFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), nil, ret.AsTypePtr())
	return ret.Get()
}

func (me *PackedVector3Array) MultiplyTransform3D(right Transform3D) PackedVector3Array {
	opPtr := ptrsForPackedVector3Array.operatorMultiplyTransform3DFn
	ret := NewPackedVector3Array()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *PackedVector3Array) InDictionary(right Dictionary) bool {
	opPtr := ptrsForPackedVector3Array.operatorInDictionaryFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *PackedVector3Array) InArray(right Array) bool {
	opPtr := ptrsForPackedVector3Array.operatorInArrayFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *PackedVector3Array) EqualPackedVector3Array(right PackedVector3Array) bool {
	opPtr := ptrsForPackedVector3Array.operatorEqualPackedVector3ArrayFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *PackedVector3Array) NotEqualPackedVector3Array(right PackedVector3Array) bool {
	opPtr := ptrsForPackedVector3Array.operatorNotEqualPackedVector3ArrayFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *PackedVector3Array) AddPackedVector3Array(right PackedVector3Array) PackedVector3Array {
	opPtr := ptrsForPackedVector3Array.operatorAddPackedVector3ArrayFn
	ret := NewPackedVector3Array()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

// Members
