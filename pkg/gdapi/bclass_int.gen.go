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

type ptrsForIntList struct {
	ctrFn                          gdc.PtrConstructor
	ctrFromIntFn                   gdc.PtrConstructor
	ctrFromFloat32Fn               gdc.PtrConstructor
	ctrFromBoolFn                  gdc.PtrConstructor
	ctrFromStringFn                gdc.PtrConstructor
	operatorNegateFn               gdc.PtrOperatorEvaluator
	operatorPositiveFn             gdc.PtrOperatorEvaluator
	operatorBitNegateFn            gdc.PtrOperatorEvaluator
	operatorNotFn                  gdc.PtrOperatorEvaluator
	operatorAndBoolFn              gdc.PtrOperatorEvaluator
	operatorOrBoolFn               gdc.PtrOperatorEvaluator
	operatorXorBoolFn              gdc.PtrOperatorEvaluator
	operatorEqualIntFn             gdc.PtrOperatorEvaluator
	operatorNotEqualIntFn          gdc.PtrOperatorEvaluator
	operatorLessIntFn              gdc.PtrOperatorEvaluator
	operatorLessEqualIntFn         gdc.PtrOperatorEvaluator
	operatorGreaterIntFn           gdc.PtrOperatorEvaluator
	operatorGreaterEqualIntFn      gdc.PtrOperatorEvaluator
	operatorAddIntFn               gdc.PtrOperatorEvaluator
	operatorSubtractIntFn          gdc.PtrOperatorEvaluator
	operatorMultiplyIntFn          gdc.PtrOperatorEvaluator
	operatorDivideIntFn            gdc.PtrOperatorEvaluator
	operatorModuleIntFn            gdc.PtrOperatorEvaluator
	operatorPowerIntFn             gdc.PtrOperatorEvaluator
	operatorShiftLeftIntFn         gdc.PtrOperatorEvaluator
	operatorShiftRightIntFn        gdc.PtrOperatorEvaluator
	operatorBitAndIntFn            gdc.PtrOperatorEvaluator
	operatorBitOrIntFn             gdc.PtrOperatorEvaluator
	operatorBitXorIntFn            gdc.PtrOperatorEvaluator
	operatorAndIntFn               gdc.PtrOperatorEvaluator
	operatorOrIntFn                gdc.PtrOperatorEvaluator
	operatorXorIntFn               gdc.PtrOperatorEvaluator
	operatorEqualFloat32Fn         gdc.PtrOperatorEvaluator
	operatorNotEqualFloat32Fn      gdc.PtrOperatorEvaluator
	operatorLessFloat32Fn          gdc.PtrOperatorEvaluator
	operatorLessEqualFloat32Fn     gdc.PtrOperatorEvaluator
	operatorGreaterFloat32Fn       gdc.PtrOperatorEvaluator
	operatorGreaterEqualFloat32Fn  gdc.PtrOperatorEvaluator
	operatorAddFloat32Fn           gdc.PtrOperatorEvaluator
	operatorSubtractFloat32Fn      gdc.PtrOperatorEvaluator
	operatorMultiplyFloat32Fn      gdc.PtrOperatorEvaluator
	operatorDivideFloat32Fn        gdc.PtrOperatorEvaluator
	operatorPowerFloat32Fn         gdc.PtrOperatorEvaluator
	operatorAndFloat32Fn           gdc.PtrOperatorEvaluator
	operatorOrFloat32Fn            gdc.PtrOperatorEvaluator
	operatorXorFloat32Fn           gdc.PtrOperatorEvaluator
	operatorMultiplyVector2Fn      gdc.PtrOperatorEvaluator
	operatorMultiplyVector2IFn     gdc.PtrOperatorEvaluator
	operatorMultiplyVector3Fn      gdc.PtrOperatorEvaluator
	operatorMultiplyVector3IFn     gdc.PtrOperatorEvaluator
	operatorMultiplyVector4Fn      gdc.PtrOperatorEvaluator
	operatorMultiplyVector4IFn     gdc.PtrOperatorEvaluator
	operatorMultiplyQuaternionFn   gdc.PtrOperatorEvaluator
	operatorMultiplyColorFn        gdc.PtrOperatorEvaluator
	operatorAndObjectFn            gdc.PtrOperatorEvaluator
	operatorOrObjectFn             gdc.PtrOperatorEvaluator
	operatorXorObjectFn            gdc.PtrOperatorEvaluator
	operatorInDictionaryFn         gdc.PtrOperatorEvaluator
	operatorInArrayFn              gdc.PtrOperatorEvaluator
	operatorInPackedByteArrayFn    gdc.PtrOperatorEvaluator
	operatorInPackedInt32ArrayFn   gdc.PtrOperatorEvaluator
	operatorInPackedInt64ArrayFn   gdc.PtrOperatorEvaluator
	operatorInPackedFloat32ArrayFn gdc.PtrOperatorEvaluator
	operatorInPackedFloat64ArrayFn gdc.PtrOperatorEvaluator
	toVariantFn                    gdc.TypeFromVariantConstructorFunc
	fromVariantFn                  gdc.VariantFromTypeConstructorFunc
}

var ptrsForInt ptrsForIntList

func initIntPtrs(iface gdc.Interface) {
	ptrsForInt.ctrFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeInt, 0))
	ptrsForInt.ctrFromIntFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeInt, 1))
	ptrsForInt.ctrFromFloat32Fn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeInt, 2))
	ptrsForInt.ctrFromBoolFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeInt, 3))
	ptrsForInt.ctrFromStringFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeInt, 4))
	ptrsForInt.operatorNegateFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNegate, gdc.VariantTypeInt, gdc.VariantTypeNil))
	ptrsForInt.operatorPositiveFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpPositive, gdc.VariantTypeInt, gdc.VariantTypeNil))
	ptrsForInt.operatorBitNegateFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpBitNegate, gdc.VariantTypeInt, gdc.VariantTypeNil))
	ptrsForInt.operatorNotFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, gdc.VariantTypeInt, gdc.VariantTypeNil))
	ptrsForInt.operatorAndBoolFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpAnd, gdc.VariantTypeInt, gdc.VariantTypeBool))
	ptrsForInt.operatorOrBoolFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpOr, gdc.VariantTypeInt, gdc.VariantTypeBool))
	ptrsForInt.operatorXorBoolFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpXor, gdc.VariantTypeInt, gdc.VariantTypeBool))
	ptrsForInt.operatorEqualIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeInt, gdc.VariantTypeInt))
	ptrsForInt.operatorNotEqualIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeInt, gdc.VariantTypeInt))
	ptrsForInt.operatorLessIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpLess, gdc.VariantTypeInt, gdc.VariantTypeInt))
	ptrsForInt.operatorLessEqualIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpLessEqual, gdc.VariantTypeInt, gdc.VariantTypeInt))
	ptrsForInt.operatorGreaterIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpGreater, gdc.VariantTypeInt, gdc.VariantTypeInt))
	ptrsForInt.operatorGreaterEqualIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpGreaterEqual, gdc.VariantTypeInt, gdc.VariantTypeInt))
	ptrsForInt.operatorAddIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpAdd, gdc.VariantTypeInt, gdc.VariantTypeInt))
	ptrsForInt.operatorSubtractIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpSubtract, gdc.VariantTypeInt, gdc.VariantTypeInt))
	ptrsForInt.operatorMultiplyIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeInt, gdc.VariantTypeInt))
	ptrsForInt.operatorDivideIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpDivide, gdc.VariantTypeInt, gdc.VariantTypeInt))
	ptrsForInt.operatorModuleIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, gdc.VariantTypeInt, gdc.VariantTypeInt))
	ptrsForInt.operatorPowerIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpPower, gdc.VariantTypeInt, gdc.VariantTypeInt))
	ptrsForInt.operatorShiftLeftIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpShiftLeft, gdc.VariantTypeInt, gdc.VariantTypeInt))
	ptrsForInt.operatorShiftRightIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpShiftRight, gdc.VariantTypeInt, gdc.VariantTypeInt))
	ptrsForInt.operatorBitAndIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpBitAnd, gdc.VariantTypeInt, gdc.VariantTypeInt))
	ptrsForInt.operatorBitOrIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpBitOr, gdc.VariantTypeInt, gdc.VariantTypeInt))
	ptrsForInt.operatorBitXorIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpBitXor, gdc.VariantTypeInt, gdc.VariantTypeInt))
	ptrsForInt.operatorAndIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpAnd, gdc.VariantTypeInt, gdc.VariantTypeInt))
	ptrsForInt.operatorOrIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpOr, gdc.VariantTypeInt, gdc.VariantTypeInt))
	ptrsForInt.operatorXorIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpXor, gdc.VariantTypeInt, gdc.VariantTypeInt))
	ptrsForInt.operatorEqualFloat32Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeInt, gdc.VariantTypeFloat))
	ptrsForInt.operatorNotEqualFloat32Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeInt, gdc.VariantTypeFloat))
	ptrsForInt.operatorLessFloat32Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpLess, gdc.VariantTypeInt, gdc.VariantTypeFloat))
	ptrsForInt.operatorLessEqualFloat32Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpLessEqual, gdc.VariantTypeInt, gdc.VariantTypeFloat))
	ptrsForInt.operatorGreaterFloat32Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpGreater, gdc.VariantTypeInt, gdc.VariantTypeFloat))
	ptrsForInt.operatorGreaterEqualFloat32Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpGreaterEqual, gdc.VariantTypeInt, gdc.VariantTypeFloat))
	ptrsForInt.operatorAddFloat32Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpAdd, gdc.VariantTypeInt, gdc.VariantTypeFloat))
	ptrsForInt.operatorSubtractFloat32Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpSubtract, gdc.VariantTypeInt, gdc.VariantTypeFloat))
	ptrsForInt.operatorMultiplyFloat32Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeInt, gdc.VariantTypeFloat))
	ptrsForInt.operatorDivideFloat32Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpDivide, gdc.VariantTypeInt, gdc.VariantTypeFloat))
	ptrsForInt.operatorPowerFloat32Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpPower, gdc.VariantTypeInt, gdc.VariantTypeFloat))
	ptrsForInt.operatorAndFloat32Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpAnd, gdc.VariantTypeInt, gdc.VariantTypeFloat))
	ptrsForInt.operatorOrFloat32Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpOr, gdc.VariantTypeInt, gdc.VariantTypeFloat))
	ptrsForInt.operatorXorFloat32Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpXor, gdc.VariantTypeInt, gdc.VariantTypeFloat))
	ptrsForInt.operatorMultiplyVector2Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeInt, gdc.VariantTypeVector2))
	ptrsForInt.operatorMultiplyVector2IFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeInt, gdc.VariantTypeVector2I))
	ptrsForInt.operatorMultiplyVector3Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeInt, gdc.VariantTypeVector3))
	ptrsForInt.operatorMultiplyVector3IFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeInt, gdc.VariantTypeVector3I))
	ptrsForInt.operatorMultiplyVector4Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeInt, gdc.VariantTypeVector4))
	ptrsForInt.operatorMultiplyVector4IFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeInt, gdc.VariantTypeVector4I))
	ptrsForInt.operatorMultiplyQuaternionFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeInt, gdc.VariantTypeQuaternion))
	ptrsForInt.operatorMultiplyColorFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeInt, gdc.VariantTypeColor))
	ptrsForInt.operatorAndObjectFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpAnd, gdc.VariantTypeInt, gdc.VariantTypeObject))
	ptrsForInt.operatorOrObjectFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpOr, gdc.VariantTypeInt, gdc.VariantTypeObject))
	ptrsForInt.operatorXorObjectFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpXor, gdc.VariantTypeInt, gdc.VariantTypeObject))
	ptrsForInt.operatorInDictionaryFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeInt, gdc.VariantTypeDictionary))
	ptrsForInt.operatorInArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeInt, gdc.VariantTypeArray))
	ptrsForInt.operatorInPackedByteArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeInt, gdc.VariantTypePackedByteArray))
	ptrsForInt.operatorInPackedInt32ArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeInt, gdc.VariantTypePackedInt32Array))
	ptrsForInt.operatorInPackedInt64ArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeInt, gdc.VariantTypePackedInt64Array))
	ptrsForInt.operatorInPackedFloat32ArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeInt, gdc.VariantTypePackedFloat32Array))
	ptrsForInt.operatorInPackedFloat64ArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeInt, gdc.VariantTypePackedFloat64Array))
	ptrsForInt.toVariantFn = ensurePtr(iface.GetVariantToTypeConstructor(gdc.VariantTypeInt))
	ptrsForInt.fromVariantFn = ensurePtr(iface.GetVariantFromTypeConstructor(gdc.VariantTypeInt))
}

type Int struct {
	data   *[classSizeInt]byte
	iface  gdc.Interface
	pinner runtime.Pinner
}

// Enums

// Constructors
func newInt() *Int {
	me := &Int{
		data:  new([classSizeInt]byte),
		iface: giface,
	}
	me.pinner.Pin(me)
	me.pinner.Pin(me.AsTypePtr())
	return me
}

func NewInt() *Int {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	me := newInt()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForInt.ctrFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{}))
	return me
}

func NewIntFromInt(from int64) *Int {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	pinner.Pin(&from)
	me := newInt()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForInt.ctrFromIntFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{gdc.ConstTypePtr(&from)}))
	return me
}

func NewIntFromFloat32(from float64) *Int {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	pinner.Pin(&from)
	me := newInt()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForInt.ctrFromFloat32Fn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{gdc.ConstTypePtr(&from)}))
	return me
}

func NewIntFromBool(from bool) *Int {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	pinner.Pin(&from)
	me := newInt()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForInt.ctrFromBoolFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{gdc.ConstTypePtr(&from)}))
	return me
}

func NewIntFromString(from String) *Int {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	me := newInt()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForInt.ctrFromStringFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr()}))
	return me
}

// Destructor
func (me *Int) Destroy() {
	me.pinner.Unpin()
}

// Variant support
func (me *Variant) AsInt() (*Int, error) {
	if me.Type() != gdc.VariantTypeInt {
		return nil, fmt.Errorf("variant is not a Int")
	}
	bclass := newInt()
	me.iface.CallTypeFromVariantConstructorFunc(ensurePtr(ptrsForInt.toVariantFn), bclass.asUninitialized(), me.AsPtr())
	return bclass, nil
}

func (me *Int) AsVariant() *Variant {
	va := newVariant()
	va.inner = me
	me.iface.CallVariantFromTypeConstructorFunc(ensurePtr(ptrsForInt.fromVariantFn), va.asUninitialized(), me.AsTypePtr())
	return va
}

// Pointers
func IntFromPtr(ptr gdc.ConstTypePtr) *Int {
	me := newInt()
	dataFromPtr(me.data[:], ptr)
	return me
}

func (me *Int) ToTypePtr(ptr gdc.TypePtr) {
	dataToPtr(ptr, me.data[:])
}

func (me *Int) Type() gdc.VariantType {
	return gdc.VariantTypeInt
}

func (me *Int) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(unsafe.Pointer(me.data))
}

func (me *Int) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.AsTypePtr())
}

func (me *Int) asUninitialized() gdc.UninitializedTypePtr {
	return gdc.UninitializedTypePtr(me.AsTypePtr())
}

func (me *Int) Get() int64 {
	return *(*int64)(unsafe.Pointer(me.data))
}

// Methods

// Operators

func (me *Int) EqualVariant(right Variant) bool {
	opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type())
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Int) NotEqualVariant(right Variant) bool {
	opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type())
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Int) Negate() int64 {
	opPtr := ptrsForInt.operatorNegateFn
	ret := NewInt()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), nil, ret.AsTypePtr())
	return ret.Get()
}

func (me *Int) Positive() int64 {
	opPtr := ptrsForInt.operatorPositiveFn
	ret := NewInt()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), nil, ret.AsTypePtr())
	return ret.Get()
}

func (me *Int) BitNegate() int64 {
	opPtr := ptrsForInt.operatorBitNegateFn
	ret := NewInt()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), nil, ret.AsTypePtr())
	return ret.Get()
}

func (me *Int) AndVariant(right Variant) bool {
	opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpAnd, me.Type(), right.Type())
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Int) OrVariant(right Variant) bool {
	opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpOr, me.Type(), right.Type())
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Int) XorVariant(right Variant) bool {
	opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpXor, me.Type(), right.Type())
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Int) Not() bool {
	opPtr := ptrsForInt.operatorNotFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), nil, ret.AsTypePtr())
	return ret.Get()
}

func (me *Int) AndBool(rightArg bool) bool {
	right := NewBoolFromBool(rightArg)
	defer right.Destroy()

	opPtr := ptrsForInt.operatorAndBoolFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Int) OrBool(rightArg bool) bool {
	right := NewBoolFromBool(rightArg)
	defer right.Destroy()

	opPtr := ptrsForInt.operatorOrBoolFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Int) XorBool(rightArg bool) bool {
	right := NewBoolFromBool(rightArg)
	defer right.Destroy()

	opPtr := ptrsForInt.operatorXorBoolFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Int) EqualInt(rightArg int64) bool {
	right := NewIntFromInt(rightArg)
	defer right.Destroy()

	opPtr := ptrsForInt.operatorEqualIntFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Int) NotEqualInt(rightArg int64) bool {
	right := NewIntFromInt(rightArg)
	defer right.Destroy()

	opPtr := ptrsForInt.operatorNotEqualIntFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Int) LessInt(rightArg int64) bool {
	right := NewIntFromInt(rightArg)
	defer right.Destroy()

	opPtr := ptrsForInt.operatorLessIntFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Int) LessEqualInt(rightArg int64) bool {
	right := NewIntFromInt(rightArg)
	defer right.Destroy()

	opPtr := ptrsForInt.operatorLessEqualIntFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Int) GreaterInt(rightArg int64) bool {
	right := NewIntFromInt(rightArg)
	defer right.Destroy()

	opPtr := ptrsForInt.operatorGreaterIntFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Int) GreaterEqualInt(rightArg int64) bool {
	right := NewIntFromInt(rightArg)
	defer right.Destroy()

	opPtr := ptrsForInt.operatorGreaterEqualIntFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Int) AddInt(rightArg int64) int64 {
	right := NewIntFromInt(rightArg)
	defer right.Destroy()

	opPtr := ptrsForInt.operatorAddIntFn
	ret := NewInt()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Int) SubtractInt(rightArg int64) int64 {
	right := NewIntFromInt(rightArg)
	defer right.Destroy()

	opPtr := ptrsForInt.operatorSubtractIntFn
	ret := NewInt()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Int) MultiplyInt(rightArg int64) int64 {
	right := NewIntFromInt(rightArg)
	defer right.Destroy()

	opPtr := ptrsForInt.operatorMultiplyIntFn
	ret := NewInt()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Int) DivideInt(rightArg int64) int64 {
	right := NewIntFromInt(rightArg)
	defer right.Destroy()

	opPtr := ptrsForInt.operatorDivideIntFn
	ret := NewInt()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Int) ModuleInt(rightArg int64) int64 {
	right := NewIntFromInt(rightArg)
	defer right.Destroy()

	opPtr := ptrsForInt.operatorModuleIntFn
	ret := NewInt()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Int) PowerInt(rightArg int64) int64 {
	right := NewIntFromInt(rightArg)
	defer right.Destroy()

	opPtr := ptrsForInt.operatorPowerIntFn
	ret := NewInt()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Int) ShiftLeftInt(rightArg int64) int64 {
	right := NewIntFromInt(rightArg)
	defer right.Destroy()

	opPtr := ptrsForInt.operatorShiftLeftIntFn
	ret := NewInt()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Int) ShiftRightInt(rightArg int64) int64 {
	right := NewIntFromInt(rightArg)
	defer right.Destroy()

	opPtr := ptrsForInt.operatorShiftRightIntFn
	ret := NewInt()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Int) BitAndInt(rightArg int64) int64 {
	right := NewIntFromInt(rightArg)
	defer right.Destroy()

	opPtr := ptrsForInt.operatorBitAndIntFn
	ret := NewInt()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Int) BitOrInt(rightArg int64) int64 {
	right := NewIntFromInt(rightArg)
	defer right.Destroy()

	opPtr := ptrsForInt.operatorBitOrIntFn
	ret := NewInt()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Int) BitXorInt(rightArg int64) int64 {
	right := NewIntFromInt(rightArg)
	defer right.Destroy()

	opPtr := ptrsForInt.operatorBitXorIntFn
	ret := NewInt()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Int) AndInt(rightArg int64) bool {
	right := NewIntFromInt(rightArg)
	defer right.Destroy()

	opPtr := ptrsForInt.operatorAndIntFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Int) OrInt(rightArg int64) bool {
	right := NewIntFromInt(rightArg)
	defer right.Destroy()

	opPtr := ptrsForInt.operatorOrIntFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Int) XorInt(rightArg int64) bool {
	right := NewIntFromInt(rightArg)
	defer right.Destroy()

	opPtr := ptrsForInt.operatorXorIntFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Int) EqualFloat32(rightArg float64) bool {
	right := NewFloatFromFloat32(rightArg)
	defer right.Destroy()

	opPtr := ptrsForInt.operatorEqualFloat32Fn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Int) NotEqualFloat32(rightArg float64) bool {
	right := NewFloatFromFloat32(rightArg)
	defer right.Destroy()

	opPtr := ptrsForInt.operatorNotEqualFloat32Fn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Int) LessFloat32(rightArg float64) bool {
	right := NewFloatFromFloat32(rightArg)
	defer right.Destroy()

	opPtr := ptrsForInt.operatorLessFloat32Fn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Int) LessEqualFloat32(rightArg float64) bool {
	right := NewFloatFromFloat32(rightArg)
	defer right.Destroy()

	opPtr := ptrsForInt.operatorLessEqualFloat32Fn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Int) GreaterFloat32(rightArg float64) bool {
	right := NewFloatFromFloat32(rightArg)
	defer right.Destroy()

	opPtr := ptrsForInt.operatorGreaterFloat32Fn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Int) GreaterEqualFloat32(rightArg float64) bool {
	right := NewFloatFromFloat32(rightArg)
	defer right.Destroy()

	opPtr := ptrsForInt.operatorGreaterEqualFloat32Fn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Int) AddFloat32(rightArg float64) float64 {
	right := NewFloatFromFloat32(rightArg)
	defer right.Destroy()

	opPtr := ptrsForInt.operatorAddFloat32Fn
	ret := NewFloat()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Int) SubtractFloat32(rightArg float64) float64 {
	right := NewFloatFromFloat32(rightArg)
	defer right.Destroy()

	opPtr := ptrsForInt.operatorSubtractFloat32Fn
	ret := NewFloat()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Int) MultiplyFloat32(rightArg float64) float64 {
	right := NewFloatFromFloat32(rightArg)
	defer right.Destroy()

	opPtr := ptrsForInt.operatorMultiplyFloat32Fn
	ret := NewFloat()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Int) DivideFloat32(rightArg float64) float64 {
	right := NewFloatFromFloat32(rightArg)
	defer right.Destroy()

	opPtr := ptrsForInt.operatorDivideFloat32Fn
	ret := NewFloat()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Int) PowerFloat32(rightArg float64) float64 {
	right := NewFloatFromFloat32(rightArg)
	defer right.Destroy()

	opPtr := ptrsForInt.operatorPowerFloat32Fn
	ret := NewFloat()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Int) AndFloat32(rightArg float64) bool {
	right := NewFloatFromFloat32(rightArg)
	defer right.Destroy()

	opPtr := ptrsForInt.operatorAndFloat32Fn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Int) OrFloat32(rightArg float64) bool {
	right := NewFloatFromFloat32(rightArg)
	defer right.Destroy()

	opPtr := ptrsForInt.operatorOrFloat32Fn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Int) XorFloat32(rightArg float64) bool {
	right := NewFloatFromFloat32(rightArg)
	defer right.Destroy()

	opPtr := ptrsForInt.operatorXorFloat32Fn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Int) MultiplyVector2(right Vector2) Vector2 {
	opPtr := ptrsForInt.operatorMultiplyVector2Fn
	ret := NewVector2()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Int) MultiplyVector2I(right Vector2i) Vector2i {
	opPtr := ptrsForInt.operatorMultiplyVector2IFn
	ret := NewVector2i()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Int) MultiplyVector3(right Vector3) Vector3 {
	opPtr := ptrsForInt.operatorMultiplyVector3Fn
	ret := NewVector3()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Int) MultiplyVector3I(right Vector3i) Vector3i {
	opPtr := ptrsForInt.operatorMultiplyVector3IFn
	ret := NewVector3i()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Int) MultiplyVector4(right Vector4) Vector4 {
	opPtr := ptrsForInt.operatorMultiplyVector4Fn
	ret := NewVector4()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Int) MultiplyVector4I(right Vector4i) Vector4i {
	opPtr := ptrsForInt.operatorMultiplyVector4IFn
	ret := NewVector4i()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Int) MultiplyQuaternion(right Quaternion) Quaternion {
	opPtr := ptrsForInt.operatorMultiplyQuaternionFn
	ret := NewQuaternion()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Int) MultiplyColor(right Color) Color {
	opPtr := ptrsForInt.operatorMultiplyColorFn
	ret := NewColor()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Int) AndObject(right Object) bool {
	opPtr := ptrsForInt.operatorAndObjectFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Int) OrObject(right Object) bool {
	opPtr := ptrsForInt.operatorOrObjectFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Int) XorObject(right Object) bool {
	opPtr := ptrsForInt.operatorXorObjectFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Int) InDictionary(right Dictionary) bool {
	opPtr := ptrsForInt.operatorInDictionaryFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Int) InArray(right Array) bool {
	opPtr := ptrsForInt.operatorInArrayFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Int) InPackedByteArray(right PackedByteArray) bool {
	opPtr := ptrsForInt.operatorInPackedByteArrayFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Int) InPackedInt32Array(right PackedInt32Array) bool {
	opPtr := ptrsForInt.operatorInPackedInt32ArrayFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Int) InPackedInt64Array(right PackedInt64Array) bool {
	opPtr := ptrsForInt.operatorInPackedInt64ArrayFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Int) InPackedFloat32Array(right PackedFloat32Array) bool {
	opPtr := ptrsForInt.operatorInPackedFloat32ArrayFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Int) InPackedFloat64Array(right PackedFloat64Array) bool {
	opPtr := ptrsForInt.operatorInPackedFloat64ArrayFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

// Members
