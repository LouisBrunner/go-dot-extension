// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "fmt"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

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
  ctr := me.iface.VariantGetPtrConstructor(gdc.VariantTypeRID, 0) // FIXME: should cache?
  me.iface.CallPtrConstructor(ctr, me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{}))
  return me
}

func NewRIDFromRID(from RID, ) *RID {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newRID()
  ctr := me.iface.VariantGetPtrConstructor(gdc.VariantTypeRID, 1) // FIXME: should cache?
  me.iface.CallPtrConstructor(ctr, me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
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
	fn := me.iface.GetVariantToTypeConstructor(me.Type())
	me.iface.CallTypeFromVariantConstructorFunc(fn, bclass.asUninitialized(), me.AsPtr())
	return bclass, nil
}

func (me *RID) AsVariant() *Variant {
  va := newVariant()
  va.inner = me
  fn := me.iface.GetVariantFromTypeConstructor(me.Type())
  me.iface.CallVariantFromTypeConstructorFunc(fn, va.asUninitialized(), me.AsTypePtr())
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
  name := StringNameFromStr("is_valid")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeRID, name.AsCPtr(), 3918633141) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret bool
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *RID) GetId() int {
  name := StringNameFromStr("get_id")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeRID, name.AsCPtr(), 3173160232) // FIXME: should cache?

  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  var ret int
  pinner.Pin(&ret)
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

// Operators

func (me *RID) EqualVariant(right Variant) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *RID) NotEqualVariant(right Variant) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *RID) Not() bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, me.Type(), gdc.VariantTypeNil) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), nil, gdc.TypePtr(&ret))
  return ret
}

func (me *RID) EqualRID(right RID) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *RID) NotEqualRID(right RID) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *RID) LessRID(right RID) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpLess, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *RID) LessEqualRID(right RID) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpLessEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *RID) GreaterRID(right RID) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpGreater, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *RID) GreaterEqualRID(right RID) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpGreaterEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

// Members
