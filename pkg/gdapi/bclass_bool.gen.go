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

type ptrsForBoolList struct {
  ctrFn gdc.PtrConstructor
  ctrFromBoolFn gdc.PtrConstructor
  ctrFromIntFn gdc.PtrConstructor
  ctrFromFloat32Fn gdc.PtrConstructor
  operatorNotFn gdc.PtrOperatorEvaluator
  operatorEqualBoolFn gdc.PtrOperatorEvaluator
  operatorNotEqualBoolFn gdc.PtrOperatorEvaluator
  operatorLessBoolFn gdc.PtrOperatorEvaluator
  operatorGreaterBoolFn gdc.PtrOperatorEvaluator
  operatorAndBoolFn gdc.PtrOperatorEvaluator
  operatorOrBoolFn gdc.PtrOperatorEvaluator
  operatorXorBoolFn gdc.PtrOperatorEvaluator
  operatorAndIntFn gdc.PtrOperatorEvaluator
  operatorOrIntFn gdc.PtrOperatorEvaluator
  operatorXorIntFn gdc.PtrOperatorEvaluator
  operatorAndFloat32Fn gdc.PtrOperatorEvaluator
  operatorOrFloat32Fn gdc.PtrOperatorEvaluator
  operatorXorFloat32Fn gdc.PtrOperatorEvaluator
  operatorAndObjectFn gdc.PtrOperatorEvaluator
  operatorOrObjectFn gdc.PtrOperatorEvaluator
  operatorXorObjectFn gdc.PtrOperatorEvaluator
  operatorInDictionaryFn gdc.PtrOperatorEvaluator
  operatorInArrayFn gdc.PtrOperatorEvaluator
  toVariantFn gdc.TypeFromVariantConstructorFunc
  fromVariantFn gdc.VariantFromTypeConstructorFunc
}

var ptrsForBool ptrsForBoolList

func initBoolPtrs(iface gdc.Interface) {
  ptrsForBool.ctrFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeBool, 0))
  ptrsForBool.ctrFromBoolFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeBool, 1))
  ptrsForBool.ctrFromIntFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeBool, 2))
  ptrsForBool.ctrFromFloat32Fn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeBool, 3))
  ptrsForBool.operatorNotFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, gdc.VariantTypeBool, gdc.VariantTypeNil))
  ptrsForBool.operatorEqualBoolFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeBool, gdc.VariantTypeBool))
  ptrsForBool.operatorNotEqualBoolFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeBool, gdc.VariantTypeBool))
  ptrsForBool.operatorLessBoolFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpLess, gdc.VariantTypeBool, gdc.VariantTypeBool))
  ptrsForBool.operatorGreaterBoolFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpGreater, gdc.VariantTypeBool, gdc.VariantTypeBool))
  ptrsForBool.operatorAndBoolFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpAnd, gdc.VariantTypeBool, gdc.VariantTypeBool))
  ptrsForBool.operatorOrBoolFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpOr, gdc.VariantTypeBool, gdc.VariantTypeBool))
  ptrsForBool.operatorXorBoolFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpXor, gdc.VariantTypeBool, gdc.VariantTypeBool))
  ptrsForBool.operatorAndIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpAnd, gdc.VariantTypeBool, gdc.VariantTypeInt))
  ptrsForBool.operatorOrIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpOr, gdc.VariantTypeBool, gdc.VariantTypeInt))
  ptrsForBool.operatorXorIntFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpXor, gdc.VariantTypeBool, gdc.VariantTypeInt))
  ptrsForBool.operatorAndFloat32Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpAnd, gdc.VariantTypeBool, gdc.VariantTypeFloat))
  ptrsForBool.operatorOrFloat32Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpOr, gdc.VariantTypeBool, gdc.VariantTypeFloat))
  ptrsForBool.operatorXorFloat32Fn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpXor, gdc.VariantTypeBool, gdc.VariantTypeFloat))
  ptrsForBool.operatorAndObjectFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpAnd, gdc.VariantTypeBool, gdc.VariantTypeObject))
  ptrsForBool.operatorOrObjectFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpOr, gdc.VariantTypeBool, gdc.VariantTypeObject))
  ptrsForBool.operatorXorObjectFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpXor, gdc.VariantTypeBool, gdc.VariantTypeObject))
  ptrsForBool.operatorInDictionaryFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeBool, gdc.VariantTypeDictionary))
  ptrsForBool.operatorInArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeBool, gdc.VariantTypeArray))
  ptrsForBool.toVariantFn = ensurePtr(iface.GetVariantToTypeConstructor(gdc.VariantTypeBool))
  ptrsForBool.fromVariantFn = ensurePtr(iface.GetVariantFromTypeConstructor(gdc.VariantTypeBool))
}

type Bool struct {
  data   *[classSizeBool]byte
  iface  gdc.Interface
  pinner runtime.Pinner
}

// Enums

// Constructors
func newBool() *Bool {
  me := &Bool{
    data:   new([classSizeBool]byte),
    iface:  giface,
  }
  me.pinner.Pin(me)
  me.pinner.Pin(me.AsTypePtr())
  return me
}

func NewBool() *Bool {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newBool()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForBool.ctrFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{}))
  return me
}

func NewBoolFromBool(from bool, ) *Bool {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  pinner.Pin(&from)
  me := newBool()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForBool.ctrFromBoolFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{gdc.ConstTypePtr(&from), }))
  return me
}

func NewBoolFromInt(from int64, ) *Bool {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  pinner.Pin(&from)
  me := newBool()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForBool.ctrFromIntFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{gdc.ConstTypePtr(&from), }))
  return me
}

func NewBoolFromFloat32(from float64, ) *Bool {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  pinner.Pin(&from)
  me := newBool()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForBool.ctrFromFloat32Fn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{gdc.ConstTypePtr(&from), }))
  return me
}

// Destructor
func (me *Bool) Destroy() {
  me.pinner.Unpin()
}

// Variant support
func (me *Variant) AsBool() (*Bool, error) {
	if me.Type() != gdc.VariantTypeBool {
		return nil, fmt.Errorf("variant is not a Bool")
	}
  bclass := newBool()
	me.iface.CallTypeFromVariantConstructorFunc(ensurePtr(ptrsForBool.toVariantFn), bclass.asUninitialized(), me.AsPtr())
	return bclass, nil
}

func (me *Bool) AsVariant() *Variant {
  va := newVariant()
  va.inner = me
  me.iface.CallVariantFromTypeConstructorFunc(ensurePtr(ptrsForBool.fromVariantFn), va.asUninitialized(), me.AsTypePtr())
  return va
}

// Pointers
func BoolFromPtr(ptr gdc.ConstTypePtr) *Bool {
  me := newBool()
  dataFromPtr(me.data[:], ptr)
  return me
}

func (me *Bool) Type() gdc.VariantType {
  return gdc.VariantTypeBool
}

func (me *Bool) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(unsafe.Pointer(me.data))
}

func (me *Bool) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.AsTypePtr())
}

func (me *Bool) asUninitialized() gdc.UninitializedTypePtr {
  return gdc.UninitializedTypePtr(me.AsTypePtr())
}

func (me *Bool) Get() bool {
  return *(*bool)(unsafe.Pointer(me.data))
}

// Methods

// Operators

func (me *Bool) EqualVariant(right Variant) bool {
  opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type())
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Bool) NotEqualVariant(right Variant) bool {
  opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type())
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Bool) AndVariant(right Variant) bool {
  opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpAnd, me.Type(), right.Type())
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Bool) OrVariant(right Variant) bool {
  opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpOr, me.Type(), right.Type())
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Bool) XorVariant(right Variant) bool {
  opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpXor, me.Type(), right.Type())
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Bool) Not() bool {
  opPtr := ptrsForBool.operatorNotFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), nil, ret.AsTypePtr())
  return ret.Get()
}

func (me *Bool) EqualBool(rightArg bool) bool {
  right := NewBoolFromBool(rightArg)
  defer right.Destroy()

  opPtr := ptrsForBool.operatorEqualBoolFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Bool) NotEqualBool(rightArg bool) bool {
  right := NewBoolFromBool(rightArg)
  defer right.Destroy()

  opPtr := ptrsForBool.operatorNotEqualBoolFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Bool) LessBool(rightArg bool) bool {
  right := NewBoolFromBool(rightArg)
  defer right.Destroy()

  opPtr := ptrsForBool.operatorLessBoolFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Bool) GreaterBool(rightArg bool) bool {
  right := NewBoolFromBool(rightArg)
  defer right.Destroy()

  opPtr := ptrsForBool.operatorGreaterBoolFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Bool) AndBool(rightArg bool) bool {
  right := NewBoolFromBool(rightArg)
  defer right.Destroy()

  opPtr := ptrsForBool.operatorAndBoolFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Bool) OrBool(rightArg bool) bool {
  right := NewBoolFromBool(rightArg)
  defer right.Destroy()

  opPtr := ptrsForBool.operatorOrBoolFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Bool) XorBool(rightArg bool) bool {
  right := NewBoolFromBool(rightArg)
  defer right.Destroy()

  opPtr := ptrsForBool.operatorXorBoolFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Bool) AndInt(rightArg int64) bool {
  right := NewIntFromInt(rightArg)
  defer right.Destroy()

  opPtr := ptrsForBool.operatorAndIntFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Bool) OrInt(rightArg int64) bool {
  right := NewIntFromInt(rightArg)
  defer right.Destroy()

  opPtr := ptrsForBool.operatorOrIntFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Bool) XorInt(rightArg int64) bool {
  right := NewIntFromInt(rightArg)
  defer right.Destroy()

  opPtr := ptrsForBool.operatorXorIntFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Bool) AndFloat32(rightArg float64) bool {
  right := NewFloatFromFloat32(rightArg)
  defer right.Destroy()

  opPtr := ptrsForBool.operatorAndFloat32Fn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Bool) OrFloat32(rightArg float64) bool {
  right := NewFloatFromFloat32(rightArg)
  defer right.Destroy()

  opPtr := ptrsForBool.operatorOrFloat32Fn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Bool) XorFloat32(rightArg float64) bool {
  right := NewFloatFromFloat32(rightArg)
  defer right.Destroy()

  opPtr := ptrsForBool.operatorXorFloat32Fn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Bool) AndObject(right Object) bool {
  opPtr := ptrsForBool.operatorAndObjectFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Bool) OrObject(right Object) bool {
  opPtr := ptrsForBool.operatorOrObjectFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Bool) XorObject(right Object) bool {
  opPtr := ptrsForBool.operatorXorObjectFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Bool) InDictionary(right Dictionary) bool {
  opPtr := ptrsForBool.operatorInDictionaryFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Bool) InArray(right Array) bool {
  opPtr := ptrsForBool.operatorInArrayFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

// Members
