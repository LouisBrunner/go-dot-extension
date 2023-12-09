// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Signal struct {
  iface gdc.Interface
  ptr gdc.TypePtr
}

// Enums

// Constructors

func NewSignal() Signal {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeSignal))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeSignal, 0) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{}))
  return Signal{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewSignalFromSignal(from Signal, ) Signal {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeSignal))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeSignal, 1) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return Signal{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewSignalFromObjectStringName(object Object, signal StringName, ) Signal {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeSignal))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeSignal, 2) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{object.AsCTypePtr(), signal.AsCTypePtr(), }))
  return Signal{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

// Destructor
func (me *Signal) Destroy() {
  if me.ptr == nil {
    return
  }
  dtr := me.iface.VariantGetPtrDestructor(gdc.VariantTypeSignal)
	me.iface.CallPtrDestructor(dtr, gdc.TypePtr(me.ptr))
	cfree(unsafe.Pointer(me.ptr))
  me.ptr = nil
}

func (me *Signal) Type() gdc.VariantType {
  return gdc.VariantTypeSignal
}

func (me *Signal) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.ptr)
}

func (me *Signal) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.ptr)
}

// Methods

func  (me *Signal) IsNull() bool {
  panic("TODO: implement")
}

func  (me *Signal) GetObject() Object {
  panic("TODO: implement")
}

func  (me *Signal) GetObjectId() int {
  panic("TODO: implement")
}

func  (me *Signal) GetName() StringName {
  panic("TODO: implement")
}

func  (me *Signal) Connect(callable Callable, flags int, ) int {
  panic("TODO: implement")
}

func  (me *Signal) Disconnect(callable Callable, )  {
  panic("TODO: implement")
}

func  (me *Signal) IsConnected(callable Callable, ) bool {
  panic("TODO: implement")
}

func  (me *Signal) GetConnections() Array {
  panic("TODO: implement")
}

func  (me *Signal) Emit()  {
  panic("TODO: implement")
}

// Operators

func (me *Signal) EqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *Signal) NotEqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *Signal) Not() bool {
  panic("TODO: implement")
}

func (me *Signal) EqualsSignal(right Signal) bool {
  panic("TODO: implement")
}

func (me *Signal) NotEqualsSignal(right Signal) bool {
  panic("TODO: implement")
}

func (me *Signal) InDictionary(right Dictionary) bool {
  panic("TODO: implement")
}

func (me *Signal) InArray(right Array) bool {
  panic("TODO: implement")
}

// TODO: members (bclass)
