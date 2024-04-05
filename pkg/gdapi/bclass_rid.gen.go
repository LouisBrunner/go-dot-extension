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

type ptrsForRIDList struct {
  ctrFn gdc.PtrConstructor
  ctrFromRIDFn gdc.PtrConstructor
  methodIsValidFn gdc.PtrBuiltInMethod
  methodGetIdFn gdc.PtrBuiltInMethod
  operatorNotFn gdc.PtrOperatorEvaluator
  operatorEqualRIDFn gdc.PtrOperatorEvaluator
  operatorNotEqualRIDFn gdc.PtrOperatorEvaluator
  operatorLessRIDFn gdc.PtrOperatorEvaluator
  operatorLessEqualRIDFn gdc.PtrOperatorEvaluator
  operatorGreaterRIDFn gdc.PtrOperatorEvaluator
  operatorGreaterEqualRIDFn gdc.PtrOperatorEvaluator
  toVariantFn gdc.TypeFromVariantConstructorFunc
  fromVariantFn gdc.VariantFromTypeConstructorFunc
}

var ptrsForRID ptrsForRIDList

func initRIDPtrs(iface gdc.Interface) {
  ptrsForRID.ctrFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeRID, 0))
  ptrsForRID.ctrFromRIDFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeRID, 1))
  {
    methodName := StringNameFromStr("is_valid")
    defer methodName.Destroy()
    ptrsForRID.methodIsValidFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeRID, methodName.AsCPtr(), 3918633141))
  }
  {
    methodName := StringNameFromStr("get_id")
    defer methodName.Destroy()
    ptrsForRID.methodGetIdFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeRID, methodName.AsCPtr(), 3173160232))
  }
  ptrsForRID.operatorNotFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, gdc.VariantTypeRID, gdc.VariantTypeNil))
  ptrsForRID.operatorEqualRIDFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeRID, gdc.VariantTypeRID))
  ptrsForRID.operatorNotEqualRIDFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeRID, gdc.VariantTypeRID))
  ptrsForRID.operatorLessRIDFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpLess, gdc.VariantTypeRID, gdc.VariantTypeRID))
  ptrsForRID.operatorLessEqualRIDFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpLessEqual, gdc.VariantTypeRID, gdc.VariantTypeRID))
  ptrsForRID.operatorGreaterRIDFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpGreater, gdc.VariantTypeRID, gdc.VariantTypeRID))
  ptrsForRID.operatorGreaterEqualRIDFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpGreaterEqual, gdc.VariantTypeRID, gdc.VariantTypeRID))
  ptrsForRID.toVariantFn = ensurePtr(iface.GetVariantToTypeConstructor(gdc.VariantTypeRID))
  ptrsForRID.fromVariantFn = ensurePtr(iface.GetVariantFromTypeConstructor(gdc.VariantTypeRID))
}

type RID struct {
  data   *[classSizeRID]byte
  iface  gdc.Interface
  pinner runtime.Pinner
}

// Enums

// Constructors
func newRID() *RID {
  me := &RID{
    data:   new([classSizeRID]byte),
    iface:  giface,
  }
  me.pinner.Pin(me)
  me.pinner.Pin(me.AsTypePtr())
  return me
}

func NewRID() *RID {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newRID()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForRID.ctrFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{}))
  return me
}

func NewRIDFromRID(from RID, ) *RID {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newRID()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForRID.ctrFromRIDFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return me
}

// Destructor
func (me *RID) Destroy() {
  me.pinner.Unpin()
}

// Variant support
func (me *Variant) AsRID() (*RID, error) {
	if me.Type() != gdc.VariantTypeRID {
		return nil, fmt.Errorf("variant is not a RID")
	}
  bclass := newRID()
	me.iface.CallTypeFromVariantConstructorFunc(ensurePtr(ptrsForRID.toVariantFn), bclass.asUninitialized(), me.AsPtr())
	return bclass, nil
}

func (me *RID) AsVariant() *Variant {
  va := newVariant()
  va.inner = me
  me.iface.CallVariantFromTypeConstructorFunc(ensurePtr(ptrsForRID.fromVariantFn), va.asUninitialized(), me.AsTypePtr())
  return va
}

// Pointers
func RIDFromPtr(ptr gdc.ConstTypePtr) *RID {
  me := newRID()
  dataFromPtr(me.data[:], ptr)
  return me
}

func (me *RID) Type() gdc.VariantType {
  return gdc.VariantTypeRID
}

func (me *RID) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(unsafe.Pointer(me.data))
}

func (me *RID) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.AsTypePtr())
}

func (me *RID) asUninitialized() gdc.UninitializedTypePtr {
  return gdc.UninitializedTypePtr(me.AsTypePtr())
}

// Methods

func (me *RID) IsValid() bool {
  ret := NewBool()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForRID.methodIsValidFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *RID) GetId() int64 {
  ret := NewInt()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForRID.methodGetIdFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

// Operators

func (me *RID) EqualVariant(right Variant) bool {
  opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type())
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *RID) NotEqualVariant(right Variant) bool {
  opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type())
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *RID) Not() bool {
  opPtr := ptrsForRID.operatorNotFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), nil, ret.AsTypePtr())
  return ret.Get()
}

func (me *RID) EqualRID(right RID) bool {
  opPtr := ptrsForRID.operatorEqualRIDFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *RID) NotEqualRID(right RID) bool {
  opPtr := ptrsForRID.operatorNotEqualRIDFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *RID) LessRID(right RID) bool {
  opPtr := ptrsForRID.operatorLessRIDFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *RID) LessEqualRID(right RID) bool {
  opPtr := ptrsForRID.operatorLessEqualRIDFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *RID) GreaterRID(right RID) bool {
  opPtr := ptrsForRID.operatorGreaterRIDFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *RID) GreaterEqualRID(right RID) bool {
  opPtr := ptrsForRID.operatorGreaterEqualRIDFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

// Members
