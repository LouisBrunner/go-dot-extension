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

type ptrsForSignalList struct {
  ctrFn gdc.PtrConstructor
  ctrFromSignalFn gdc.PtrConstructor
  ctrFromObjectStringNameFn gdc.PtrConstructor
  destructorFn gdc.PtrDestructor
  methodIsNullFn gdc.PtrBuiltInMethod
  methodGetObjectFn gdc.PtrBuiltInMethod
  methodGetObjectIdFn gdc.PtrBuiltInMethod
  methodGetNameFn gdc.PtrBuiltInMethod
  methodConnectFn gdc.PtrBuiltInMethod
  methodDisconnectFn gdc.PtrBuiltInMethod
  methodIsConnectedFn gdc.PtrBuiltInMethod
  methodGetConnectionsFn gdc.PtrBuiltInMethod
  methodEmitFn gdc.PtrBuiltInMethod
  operatorNotFn gdc.PtrOperatorEvaluator
  operatorEqualSignalFn gdc.PtrOperatorEvaluator
  operatorNotEqualSignalFn gdc.PtrOperatorEvaluator
  operatorInDictionaryFn gdc.PtrOperatorEvaluator
  operatorInArrayFn gdc.PtrOperatorEvaluator
  toVariantFn gdc.TypeFromVariantConstructorFunc
  fromVariantFn gdc.VariantFromTypeConstructorFunc
}

var ptrsForSignal ptrsForSignalList

func initSignalPtrs(iface gdc.Interface) {
  ptrsForSignal.ctrFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeSignal, 0))
  ptrsForSignal.ctrFromSignalFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeSignal, 1))
  ptrsForSignal.ctrFromObjectStringNameFn = ensurePtr(iface.VariantGetPtrConstructor(gdc.VariantTypeSignal, 2))
  ptrsForSignal.destructorFn = ensurePtr(iface.VariantGetPtrDestructor(gdc.VariantTypeSignal))
  {
    methodName := StringNameFromStr("is_null")
    defer methodName.Destroy()
    ptrsForSignal.methodIsNullFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeSignal, methodName.AsCPtr(), 3918633141))
  }
  {
    methodName := StringNameFromStr("get_object")
    defer methodName.Destroy()
    ptrsForSignal.methodGetObjectFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeSignal, methodName.AsCPtr(), 4008621732))
  }
  {
    methodName := StringNameFromStr("get_object_id")
    defer methodName.Destroy()
    ptrsForSignal.methodGetObjectIdFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeSignal, methodName.AsCPtr(), 3173160232))
  }
  {
    methodName := StringNameFromStr("get_name")
    defer methodName.Destroy()
    ptrsForSignal.methodGetNameFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeSignal, methodName.AsCPtr(), 1825232092))
  }
  {
    methodName := StringNameFromStr("connect")
    defer methodName.Destroy()
    ptrsForSignal.methodConnectFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeSignal, methodName.AsCPtr(), 979702392))
  }
  {
    methodName := StringNameFromStr("disconnect")
    defer methodName.Destroy()
    ptrsForSignal.methodDisconnectFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeSignal, methodName.AsCPtr(), 3470848906))
  }
  {
    methodName := StringNameFromStr("is_connected")
    defer methodName.Destroy()
    ptrsForSignal.methodIsConnectedFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeSignal, methodName.AsCPtr(), 4129521963))
  }
  {
    methodName := StringNameFromStr("get_connections")
    defer methodName.Destroy()
    ptrsForSignal.methodGetConnectionsFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeSignal, methodName.AsCPtr(), 4144163970))
  }
  {
    methodName := StringNameFromStr("emit")
    defer methodName.Destroy()
    ptrsForSignal.methodEmitFn = ensurePtr(iface.VariantGetPtrBuiltinMethod(gdc.VariantTypeSignal, methodName.AsCPtr(), 3286317445))
  }
  ptrsForSignal.operatorNotFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, gdc.VariantTypeSignal, gdc.VariantTypeNil))
  ptrsForSignal.operatorEqualSignalFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, gdc.VariantTypeSignal, gdc.VariantTypeSignal))
  ptrsForSignal.operatorNotEqualSignalFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, gdc.VariantTypeSignal, gdc.VariantTypeSignal))
  ptrsForSignal.operatorInDictionaryFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeSignal, gdc.VariantTypeDictionary))
  ptrsForSignal.operatorInArrayFn = ensurePtr(iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, gdc.VariantTypeSignal, gdc.VariantTypeArray))
  ptrsForSignal.toVariantFn = ensurePtr(iface.GetVariantToTypeConstructor(gdc.VariantTypeSignal))
  ptrsForSignal.fromVariantFn = ensurePtr(iface.GetVariantFromTypeConstructor(gdc.VariantTypeSignal))
}

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
  me.iface.CallPtrConstructor(ensurePtr(ptrsForSignal.ctrFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{}))
  return me
}

func NewSignalFromSignal(from Signal, ) *Signal {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newSignal()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForSignal.ctrFromSignalFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return me
}

func NewSignalFromObjectStringName(object Object, signal StringName, ) *Signal {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newSignal()
  me.iface.CallPtrConstructor(ensurePtr(ptrsForSignal.ctrFromObjectStringNameFn), me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{object.AsCTypePtr(), signal.AsCTypePtr(), }))
  return me
}

// Destructor
func (me *Signal) Destroy() {
	me.iface.CallPtrDestructor(ensurePtr(ptrsForSignal.destructorFn), me.AsTypePtr())
  me.pinner.Unpin()
}

// Variant support
func (me *Variant) AsSignal() (*Signal, error) {
	if me.Type() != gdc.VariantTypeSignal {
		return nil, fmt.Errorf("variant is not a Signal")
	}
  bclass := newSignal()
	me.iface.CallTypeFromVariantConstructorFunc(ensurePtr(ptrsForSignal.toVariantFn), bclass.asUninitialized(), me.AsPtr())
	return bclass, nil
}

func (me *Signal) AsVariant() *Variant {
  va := newVariant()
  va.inner = me
  me.iface.CallVariantFromTypeConstructorFunc(ensurePtr(ptrsForSignal.fromVariantFn), va.asUninitialized(), me.AsTypePtr())
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
  ret := NewBool()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForSignal.methodIsNullFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Signal) GetObject() Object {
  ret := NewObject()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForSignal.methodGetObjectFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Signal) GetObjectId() int64 {
  ret := NewInt()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForSignal.methodGetObjectIdFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Signal) GetName() StringName {
  ret := NewStringName()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForSignal.methodGetNameFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Signal) Connect(callable Callable, flags int64, ) int64 {
  ret := NewInt()
  defer ret.Destroy()

  varg1 := NewIntFromInt(flags)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{callable.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForSignal.methodConnectFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Signal) Disconnect(callable Callable, )  {

  args := []gdc.ConstTypePtr{callable.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForSignal.methodDisconnectFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *Signal) IsConnected(callable Callable, ) bool {
  ret := NewBool()
  defer ret.Destroy()

  args := []gdc.ConstTypePtr{callable.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForSignal.methodIsConnectedFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *Signal) GetConnections() Array {
  ret := NewArray()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForSignal.methodGetConnectionsFn), me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *Signal) Emit(varargs ...Variant)  {
  args := []gdc.ConstTypePtr{}
  for _, arg := range varargs {
    args = append(args, arg.AsCTypePtr())
  }

  giface.CallPtrBuiltInMethod(ensurePtr(ptrsForSignal.methodEmitFn), me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

// Operators

func (me *Signal) EqualVariant(right Variant) bool {
  opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type())
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Signal) NotEqualVariant(right Variant) bool {
  opPtr := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type())
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Signal) Not() bool {
  opPtr := ptrsForSignal.operatorNotFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), nil, ret.AsTypePtr())
  return ret.Get()
}

func (me *Signal) EqualSignal(right Signal) bool {
  opPtr := ptrsForSignal.operatorEqualSignalFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Signal) NotEqualSignal(right Signal) bool {
  opPtr := ptrsForSignal.operatorNotEqualSignalFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Signal) InDictionary(right Dictionary) bool {
  opPtr := ptrsForSignal.operatorInDictionaryFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *Signal) InArray(right Array) bool {
  opPtr := ptrsForSignal.operatorInArrayFn
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(ensurePtr(opPtr), me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

// Members
