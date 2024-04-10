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

type ptrsForFloatList struct {
	ctrFn                          gdc.PtrConstructor
	ctrFromFloat32Fn               gdc.PtrConstructor
	ctrFromIntFn                   gdc.PtrConstructor
	ctrFromBoolFn                  gdc.PtrConstructor
	ctrFromStringFn                gdc.PtrConstructor
	operatorNegateFn               gdc.PtrOperatorEvaluator
	operatorPositiveFn             gdc.PtrOperatorEvaluator
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
	operatorPowerIntFn             gdc.PtrOperatorEvaluator
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

var ptrsForFloat ptrsForFloatList

func initFloatPtrs(iface gdc.Interface) {
	ptrsForFloat.ctrFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeFloat, 0))
	ptrsForFloat.ctrFromFloat32Fn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeFloat, 1))
	ptrsForFloat.ctrFromIntFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeFloat, 2))
	ptrsForFloat.ctrFromBoolFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeFloat, 3))
	ptrsForFloat.ctrFromStringFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeFloat, 4))
	ptrsForFloat.operatorNegateFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNegate, gdc.VariantTypeFloat, gdc.VariantTypeNil))
	ptrsForFloat.operatorPositiveFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpPositive, gdc.VariantTypeFloat, gdc.VariantTypeNil))
	ptrsForFloat.operatorNotFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, gdc.VariantTypeFloat, gdc.VariantTypeNil))
	ptrsForFloat.operatorAndBoolFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpAnd, gdc.VariantTypeFloat, gdc.VariantTypeBool))
	ptrsForFloat.operatorOrBoolFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpOr, gdc.VariantTypeFloat, gdc.VariantTypeBool))
	ptrsForFloat.operatorXorBoolFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpXor, gdc.VariantTypeFloat, gdc.VariantTypeBool))
	ptrsForFloat.operatorEqualIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeFloat, gdc.VariantTypeInt))
	ptrsForFloat.operatorNotEqualIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeFloat, gdc.VariantTypeInt))
	ptrsForFloat.operatorLessIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpLess, gdc.VariantTypeFloat, gdc.VariantTypeInt))
	ptrsForFloat.operatorLessEqualIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpLessEqual, gdc.VariantTypeFloat, gdc.VariantTypeInt))
	ptrsForFloat.operatorGreaterIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpGreater, gdc.VariantTypeFloat, gdc.VariantTypeInt))
	ptrsForFloat.operatorGreaterEqualIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpGreaterEqual, gdc.VariantTypeFloat, gdc.VariantTypeInt))
	ptrsForFloat.operatorAddIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpAdd, gdc.VariantTypeFloat, gdc.VariantTypeInt))
	ptrsForFloat.operatorSubtractIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpSubtract, gdc.VariantTypeFloat, gdc.VariantTypeInt))
	ptrsForFloat.operatorMultiplyIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeFloat, gdc.VariantTypeInt))
	ptrsForFloat.operatorDivideIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpDivide, gdc.VariantTypeFloat, gdc.VariantTypeInt))
	ptrsForFloat.operatorPowerIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpPower, gdc.VariantTypeFloat, gdc.VariantTypeInt))
	ptrsForFloat.operatorAndIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpAnd, gdc.VariantTypeFloat, gdc.VariantTypeInt))
	ptrsForFloat.operatorOrIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpOr, gdc.VariantTypeFloat, gdc.VariantTypeInt))
	ptrsForFloat.operatorXorIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpXor, gdc.VariantTypeFloat, gdc.VariantTypeInt))
	ptrsForFloat.operatorEqualFloat32Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeFloat, gdc.VariantTypeFloat))
	ptrsForFloat.operatorNotEqualFloat32Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeFloat, gdc.VariantTypeFloat))
	ptrsForFloat.operatorLessFloat32Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpLess, gdc.VariantTypeFloat, gdc.VariantTypeFloat))
	ptrsForFloat.operatorLessEqualFloat32Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpLessEqual, gdc.VariantTypeFloat, gdc.VariantTypeFloat))
	ptrsForFloat.operatorGreaterFloat32Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpGreater, gdc.VariantTypeFloat, gdc.VariantTypeFloat))
	ptrsForFloat.operatorGreaterEqualFloat32Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpGreaterEqual, gdc.VariantTypeFloat, gdc.VariantTypeFloat))
	ptrsForFloat.operatorAddFloat32Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpAdd, gdc.VariantTypeFloat, gdc.VariantTypeFloat))
	ptrsForFloat.operatorSubtractFloat32Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpSubtract, gdc.VariantTypeFloat, gdc.VariantTypeFloat))
	ptrsForFloat.operatorMultiplyFloat32Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeFloat, gdc.VariantTypeFloat))
	ptrsForFloat.operatorDivideFloat32Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpDivide, gdc.VariantTypeFloat, gdc.VariantTypeFloat))
	ptrsForFloat.operatorPowerFloat32Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpPower, gdc.VariantTypeFloat, gdc.VariantTypeFloat))
	ptrsForFloat.operatorAndFloat32Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpAnd, gdc.VariantTypeFloat, gdc.VariantTypeFloat))
	ptrsForFloat.operatorOrFloat32Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpOr, gdc.VariantTypeFloat, gdc.VariantTypeFloat))
	ptrsForFloat.operatorXorFloat32Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpXor, gdc.VariantTypeFloat, gdc.VariantTypeFloat))
	ptrsForFloat.operatorMultiplyVector2Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeFloat, gdc.VariantTypeVector2))
	ptrsForFloat.operatorMultiplyVector2IFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeFloat, gdc.VariantTypeVector2I))
	ptrsForFloat.operatorMultiplyVector3Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeFloat, gdc.VariantTypeVector3))
	ptrsForFloat.operatorMultiplyVector3IFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeFloat, gdc.VariantTypeVector3I))
	ptrsForFloat.operatorMultiplyVector4Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeFloat, gdc.VariantTypeVector4))
	ptrsForFloat.operatorMultiplyVector4IFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeFloat, gdc.VariantTypeVector4I))
	ptrsForFloat.operatorMultiplyQuaternionFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeFloat, gdc.VariantTypeQuaternion))
	ptrsForFloat.operatorMultiplyColorFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpMultiply, gdc.VariantTypeFloat, gdc.VariantTypeColor))
	ptrsForFloat.operatorAndObjectFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpAnd, gdc.VariantTypeFloat, gdc.VariantTypeObject))
	ptrsForFloat.operatorOrObjectFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpOr, gdc.VariantTypeFloat, gdc.VariantTypeObject))
	ptrsForFloat.operatorXorObjectFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpXor, gdc.VariantTypeFloat, gdc.VariantTypeObject))
	ptrsForFloat.operatorInDictionaryFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeFloat, gdc.VariantTypeDictionary))
	ptrsForFloat.operatorInArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeFloat, gdc.VariantTypeArray))
	ptrsForFloat.operatorInPackedByteArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeFloat, gdc.VariantTypePackedByteArray))
	ptrsForFloat.operatorInPackedInt32ArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeFloat, gdc.VariantTypePackedInt32Array))
	ptrsForFloat.operatorInPackedInt64ArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeFloat, gdc.VariantTypePackedInt64Array))
	ptrsForFloat.operatorInPackedFloat32ArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeFloat, gdc.VariantTypePackedFloat32Array))
	ptrsForFloat.operatorInPackedFloat64ArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeFloat, gdc.VariantTypePackedFloat64Array))
	ptrsForFloat.toVariantFn = ensurePtr(iface.GetVariantToTypeConstructor(gdc.VariantTypeFloat))
	ptrsForFloat.fromVariantFn = ensurePtr(iface.GetVariantFromTypeConstructor(gdc.VariantTypeFloat))
}

type Float struct {
	data   *[classSizeFloat]byte
	iface  gdc.Interface
	pinner runtime.Pinner
}

// Enums

// Constructors
func newFloat() *Float {
	me := &Float{
		data:  new([classSizeFloat]byte),
		iface: giface,
	}
	me.pinner.Pin(me)
	me.pinner.Pin(me.AsTypePtr())
	return me
}

func NewFloat() *Float {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	me := newFloat()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForFloat.ctrFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{}))
	return me
}

func NewFloatFromFloat32(from float64) *Float {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	pinner.Pin(&from)
	me := newFloat()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForFloat.ctrFromFloat32Fn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{gdc.ConstTypePtr(&from)}))
	return me
}

func NewFloatFromInt(from int64) *Float {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	pinner.Pin(&from)
	me := newFloat()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForFloat.ctrFromIntFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{gdc.ConstTypePtr(&from)}))
	return me
}

func NewFloatFromBool(from bool) *Float {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	pinner.Pin(&from)
	me := newFloat()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForFloat.ctrFromBoolFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{gdc.ConstTypePtr(&from)}))
	return me
}

func NewFloatFromString(from String) *Float {
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	me := newFloat()
	me.iface.CallPtrConstructor(ensurePtr(ptrsForFloat.ctrFromStringFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr()}))
	return me
}

// Destructor
func (me *Float) Destroy() {
	me.pinner.Unpin()
}

// Variant support
func (me *Variant) AsFloat() (*Float, error) {
	if me.Type() != gdc.VariantTypeFloat {
		return nil, fmt.Errorf("variant is not a Float")
	}
	bclass := newFloat()
	me.iface.CallTypeFromVariantConstructorFunc(ensurePtr(ptrsForFloat.toVariantFn), bclass.asUninitialized(), me.AsPtr())
	return bclass, nil
}

func (me *Float) AsVariant() *Variant {
	va := newVariant()
	va.inner = me
	me.iface.CallVariantFromTypeConstructorFunc(ensurePtr(ptrsForFloat.fromVariantFn), va.asUninitialized(), me.AsTypePtr())
	return va
}

// Pointers
func FloatFromPtr(ptr gdc.ConstTypePtr) *Float {
	me := newFloat()
	dataFromPtr(me.data[:], ptr)
	return me
}

func (me *Float) ToTypePtr(ptr gdc.TypePtr) {
	dataToPtr(ptr, me.data[:])
}

func (me *Float) Type() gdc.VariantType {
	return gdc.VariantTypeFloat
}

func (me *Float) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(unsafe.Pointer(me.data))
}

func (me *Float) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.AsTypePtr())
}

func (me *Float) asUninitialized() gdc.UninitializedTypePtr {
	return gdc.UninitializedTypePtr(me.AsTypePtr())
}

func (me *Float) Get() float64 {
	return *(*float64)(unsafe.Pointer(me.data))
}

// Methods

// Operators

func (me *Float) EqualVariant(right Variant) bool {
	opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type())
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Float) NotEqualVariant(right Variant) bool {
	opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type())
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Float) Negate() float64 {
	opPtr := ptrsForFloat.operatorNegateFn
	ret := NewFloat()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), nil, ret.AsTypePtr())
	return ret.Get()
}

func (me *Float) Positive() float64 {
	opPtr := ptrsForFloat.operatorPositiveFn
	ret := NewFloat()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), nil, ret.AsTypePtr())
	return ret.Get()
}

func (me *Float) AndVariant(right Variant) bool {
	opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpAnd, me.Type(), right.Type())
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Float) OrVariant(right Variant) bool {
	opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpOr, me.Type(), right.Type())
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Float) XorVariant(right Variant) bool {
	opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpXor, me.Type(), right.Type())
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Float) Not() bool {
	opPtr := ptrsForFloat.operatorNotFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), nil, ret.AsTypePtr())
	return ret.Get()
}

func (me *Float) AndBool(rightArg bool) bool {
	right := NewBoolFromBool(rightArg)
	defer right.Destroy()

	opPtr := ptrsForFloat.operatorAndBoolFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Float) OrBool(rightArg bool) bool {
	right := NewBoolFromBool(rightArg)
	defer right.Destroy()

	opPtr := ptrsForFloat.operatorOrBoolFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Float) XorBool(rightArg bool) bool {
	right := NewBoolFromBool(rightArg)
	defer right.Destroy()

	opPtr := ptrsForFloat.operatorXorBoolFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Float) EqualInt(rightArg int64) bool {
	right := NewIntFromInt(rightArg)
	defer right.Destroy()

	opPtr := ptrsForFloat.operatorEqualIntFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Float) NotEqualInt(rightArg int64) bool {
	right := NewIntFromInt(rightArg)
	defer right.Destroy()

	opPtr := ptrsForFloat.operatorNotEqualIntFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Float) LessInt(rightArg int64) bool {
	right := NewIntFromInt(rightArg)
	defer right.Destroy()

	opPtr := ptrsForFloat.operatorLessIntFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Float) LessEqualInt(rightArg int64) bool {
	right := NewIntFromInt(rightArg)
	defer right.Destroy()

	opPtr := ptrsForFloat.operatorLessEqualIntFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Float) GreaterInt(rightArg int64) bool {
	right := NewIntFromInt(rightArg)
	defer right.Destroy()

	opPtr := ptrsForFloat.operatorGreaterIntFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Float) GreaterEqualInt(rightArg int64) bool {
	right := NewIntFromInt(rightArg)
	defer right.Destroy()

	opPtr := ptrsForFloat.operatorGreaterEqualIntFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Float) AddInt(rightArg int64) float64 {
	right := NewIntFromInt(rightArg)
	defer right.Destroy()

	opPtr := ptrsForFloat.operatorAddIntFn
	ret := NewFloat()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Float) SubtractInt(rightArg int64) float64 {
	right := NewIntFromInt(rightArg)
	defer right.Destroy()

	opPtr := ptrsForFloat.operatorSubtractIntFn
	ret := NewFloat()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Float) MultiplyInt(rightArg int64) float64 {
	right := NewIntFromInt(rightArg)
	defer right.Destroy()

	opPtr := ptrsForFloat.operatorMultiplyIntFn
	ret := NewFloat()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Float) DivideInt(rightArg int64) float64 {
	right := NewIntFromInt(rightArg)
	defer right.Destroy()

	opPtr := ptrsForFloat.operatorDivideIntFn
	ret := NewFloat()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Float) PowerInt(rightArg int64) float64 {
	right := NewIntFromInt(rightArg)
	defer right.Destroy()

	opPtr := ptrsForFloat.operatorPowerIntFn
	ret := NewFloat()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Float) AndInt(rightArg int64) bool {
	right := NewIntFromInt(rightArg)
	defer right.Destroy()

	opPtr := ptrsForFloat.operatorAndIntFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Float) OrInt(rightArg int64) bool {
	right := NewIntFromInt(rightArg)
	defer right.Destroy()

	opPtr := ptrsForFloat.operatorOrIntFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Float) XorInt(rightArg int64) bool {
	right := NewIntFromInt(rightArg)
	defer right.Destroy()

	opPtr := ptrsForFloat.operatorXorIntFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Float) EqualFloat32(rightArg float64) bool {
	right := NewFloatFromFloat32(rightArg)
	defer right.Destroy()

	opPtr := ptrsForFloat.operatorEqualFloat32Fn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Float) NotEqualFloat32(rightArg float64) bool {
	right := NewFloatFromFloat32(rightArg)
	defer right.Destroy()

	opPtr := ptrsForFloat.operatorNotEqualFloat32Fn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Float) LessFloat32(rightArg float64) bool {
	right := NewFloatFromFloat32(rightArg)
	defer right.Destroy()

	opPtr := ptrsForFloat.operatorLessFloat32Fn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Float) LessEqualFloat32(rightArg float64) bool {
	right := NewFloatFromFloat32(rightArg)
	defer right.Destroy()

	opPtr := ptrsForFloat.operatorLessEqualFloat32Fn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Float) GreaterFloat32(rightArg float64) bool {
	right := NewFloatFromFloat32(rightArg)
	defer right.Destroy()

	opPtr := ptrsForFloat.operatorGreaterFloat32Fn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Float) GreaterEqualFloat32(rightArg float64) bool {
	right := NewFloatFromFloat32(rightArg)
	defer right.Destroy()

	opPtr := ptrsForFloat.operatorGreaterEqualFloat32Fn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Float) AddFloat32(rightArg float64) float64 {
	right := NewFloatFromFloat32(rightArg)
	defer right.Destroy()

	opPtr := ptrsForFloat.operatorAddFloat32Fn
	ret := NewFloat()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Float) SubtractFloat32(rightArg float64) float64 {
	right := NewFloatFromFloat32(rightArg)
	defer right.Destroy()

	opPtr := ptrsForFloat.operatorSubtractFloat32Fn
	ret := NewFloat()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Float) MultiplyFloat32(rightArg float64) float64 {
	right := NewFloatFromFloat32(rightArg)
	defer right.Destroy()

	opPtr := ptrsForFloat.operatorMultiplyFloat32Fn
	ret := NewFloat()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Float) DivideFloat32(rightArg float64) float64 {
	right := NewFloatFromFloat32(rightArg)
	defer right.Destroy()

	opPtr := ptrsForFloat.operatorDivideFloat32Fn
	ret := NewFloat()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Float) PowerFloat32(rightArg float64) float64 {
	right := NewFloatFromFloat32(rightArg)
	defer right.Destroy()

	opPtr := ptrsForFloat.operatorPowerFloat32Fn
	ret := NewFloat()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Float) AndFloat32(rightArg float64) bool {
	right := NewFloatFromFloat32(rightArg)
	defer right.Destroy()

	opPtr := ptrsForFloat.operatorAndFloat32Fn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Float) OrFloat32(rightArg float64) bool {
	right := NewFloatFromFloat32(rightArg)
	defer right.Destroy()

	opPtr := ptrsForFloat.operatorOrFloat32Fn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Float) XorFloat32(rightArg float64) bool {
	right := NewFloatFromFloat32(rightArg)
	defer right.Destroy()

	opPtr := ptrsForFloat.operatorXorFloat32Fn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Float) MultiplyVector2(right Vector2) Vector2 {
	opPtr := ptrsForFloat.operatorMultiplyVector2Fn
	ret := NewVector2()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Float) MultiplyVector2I(right Vector2i) Vector2 {
	opPtr := ptrsForFloat.operatorMultiplyVector2IFn
	ret := NewVector2()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Float) MultiplyVector3(right Vector3) Vector3 {
	opPtr := ptrsForFloat.operatorMultiplyVector3Fn
	ret := NewVector3()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Float) MultiplyVector3I(right Vector3i) Vector3 {
	opPtr := ptrsForFloat.operatorMultiplyVector3IFn
	ret := NewVector3()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Float) MultiplyVector4(right Vector4) Vector4 {
	opPtr := ptrsForFloat.operatorMultiplyVector4Fn
	ret := NewVector4()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Float) MultiplyVector4I(right Vector4i) Vector4 {
	opPtr := ptrsForFloat.operatorMultiplyVector4IFn
	ret := NewVector4()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Float) MultiplyQuaternion(right Quaternion) Quaternion {
	opPtr := ptrsForFloat.operatorMultiplyQuaternionFn
	ret := NewQuaternion()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Float) MultiplyColor(right Color) Color {
	opPtr := ptrsForFloat.operatorMultiplyColorFn
	ret := NewColor()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return *ret
}

func (me *Float) AndObject(right Object) bool {
	opPtr := ptrsForFloat.operatorAndObjectFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Float) OrObject(right Object) bool {
	opPtr := ptrsForFloat.operatorOrObjectFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Float) XorObject(right Object) bool {
	opPtr := ptrsForFloat.operatorXorObjectFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Float) InDictionary(right Dictionary) bool {
	opPtr := ptrsForFloat.operatorInDictionaryFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Float) InArray(right Array) bool {
	opPtr := ptrsForFloat.operatorInArrayFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Float) InPackedByteArray(right PackedByteArray) bool {
	opPtr := ptrsForFloat.operatorInPackedByteArrayFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Float) InPackedInt32Array(right PackedInt32Array) bool {
	opPtr := ptrsForFloat.operatorInPackedInt32ArrayFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Float) InPackedInt64Array(right PackedInt64Array) bool {
	opPtr := ptrsForFloat.operatorInPackedInt64ArrayFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Float) InPackedFloat32Array(right PackedFloat32Array) bool {
	opPtr := ptrsForFloat.operatorInPackedFloat32ArrayFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

func (me *Float) InPackedFloat64Array(right PackedFloat64Array) bool {
	opPtr := ptrsForFloat.operatorInPackedFloat64ArrayFn
	ret := NewBool()
	me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
	return ret.Get()
}

// Members
