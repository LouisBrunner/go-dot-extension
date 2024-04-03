// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "fmt"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Signal struct {
  data   *[classSizeSignal]byte
  iface  gdc.Interface
  pinner runtime.Pinner
}

// Enums

// Constructors
func newSignal() *Signal {
  me := &Signal{
    data:   new([classSizeSignal]byte),
    iface:  giface,
  }
  me.pinner.Pin(me)
  me.pinner.Pin(me.AsTypePtr())
  return me
}

func NewSignal() *Signal {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newSignal()
  ctr := me.iface.VariantGetPtrConstructor(gdc.VariantTypeSignal, 0) // FIXME: should cache?
  me.iface.CallPtrConstructor(ctr, me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{}))
  return me
}

func NewSignalFromSignal(from Signal, ) *Signal {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newSignal()
  ctr := me.iface.VariantGetPtrConstructor(gdc.VariantTypeSignal, 1) // FIXME: should cache?
  me.iface.CallPtrConstructor(ctr, me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return me
}

func NewSignalFromObjectStringName(object Object, signal StringName, ) *Signal {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newSignal()
  ctr := me.iface.VariantGetPtrConstructor(gdc.VariantTypeSignal, 2) // FIXME: should cache?
  me.iface.CallPtrConstructor(ctr, me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{object.AsCTypePtr(), signal.AsCTypePtr(), }))
  return me
}

// Destructor
func (me *Signal) Destroy() {
  dtr := me.iface.VariantGetPtrDestructor(gdc.VariantTypeSignal)
	me.iface.CallPtrDestructor(dtr, me.AsTypePtr())
  me.pinner.Unpin()
}

// Variant support
func (me *Variant) AsSignal() (*Signal, error) {
	if me.Type() != gdc.VariantTypeSignal {
		return nil, fmt.Errorf("variant is not a Signal")
	}
  bclass := newSignal()
	fn := me.iface.GetVariantToTypeConstructor(me.Type())
	me.iface.CallTypeFromVariantConstructorFunc(fn, bclass.asUninitialized(), me.AsPtr())
	return bclass, nil
}

func (me *Signal) AsVariant() *Variant {
  va := newVariant()
  va.inner = me
  fn := me.iface.GetVariantFromTypeConstructor(me.Type())
  me.iface.CallVariantFromTypeConstructorFunc(fn, va.asUninitialized(), me.AsTypePtr())
  return va
}

// Pointers
func SignalFromPtr(ptr gdc.ConstTypePtr) *Signal {
  me := newSignal()
  dataFromPtr(me.data[:], ptr)
  return me
}

func (me *Signal) Type() gdc.VariantType {
  return gdc.VariantTypeSignal
}

func (me *Signal) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(unsafe.Pointer(me.data))
}

func (me *Signal) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.AsTypePtr())
}

func (me *Signal) asUninitialized() gdc.UninitializedTypePtr {
  return gdc.UninitializedTypePtr(me.AsTypePtr())
}

// Methods

func (me *Signal) IsNull() bool {
  name := StringNameFromStr("is_null")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeSignal, name.AsCPtr(), 3918633141) // FIXME: should cache?

  ret := NewBool()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Signal) GetObject() Object {
  name := StringNameFromStr("get_object")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeSignal, name.AsCPtr(), 4008621732) // FIXME: should cache?

  ret := NewObject()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Signal) GetObjectId() int64 {
  name := StringNameFromStr("get_object_id")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeSignal, name.AsCPtr(), 3173160232) // FIXME: should cache?

  ret := NewInt()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Signal) GetName() StringName {
  name := StringNameFromStr("get_name")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeSignal, name.AsCPtr(), 1825232092) // FIXME: should cache?

  ret := NewStringName()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Signal) Connect(callable Callable, flags int64, ) int64 {
  name := StringNameFromStr("connect")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeSignal, name.AsCPtr(), 979702392) // FIXME: should cache?

  ret := NewInt()
  defer ret.Destroy()

  varg1 := NewIntFromInt(flags)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{callable.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Signal) Disconnect(callable Callable, )  {
  name := StringNameFromStr("disconnect")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeSignal, name.AsCPtr(), 3470848906) // FIXME: should cache?


  args := []gdc.ConstTypePtr{callable.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *Signal) IsConnected(callable Callable, ) bool {
  name := StringNameFromStr("is_connected")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeSignal, name.AsCPtr(), 4129521963) // FIXME: should cache?

  ret := NewBool()
  defer ret.Destroy()

  args := []gdc.ConstTypePtr{callable.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Signal) GetConnections() Array {
  name := StringNameFromStr("get_connections")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeSignal, name.AsCPtr(), 4144163970) // FIXME: should cache?

  ret := NewArray()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Signal) Emit(varargs ...Variant)  {
  name := StringNameFromStr("emit")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeSignal, name.AsCPtr(), 3286317445) // FIXME: should cache?

  args := []gdc.ConstTypePtr{}
  for _, arg := range varargs {
    args = append(args, arg.AsCTypePtr())
  }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

// Operators

func (me *Signal) EqualVariant(right Variant) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type()) // FIXME: cache
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Signal) NotEqualVariant(right Variant) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type()) // FIXME: cache
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Signal) Not() bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, me.Type(), gdc.VariantTypeNil) // FIXME: cache
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), nil, ret.AsTypePtr())
  return ret.Get()
}

func (me *Signal) EqualSignal(right Signal) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type()) // FIXME: cache
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Signal) NotEqualSignal(right Signal) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type()) // FIXME: cache
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Signal) InDictionary(right Dictionary) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, me.Type(), right.Type()) // FIXME: cache
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Signal) InArray(right Array) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, me.Type(), right.Type()) // FIXME: cache
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

// Members
