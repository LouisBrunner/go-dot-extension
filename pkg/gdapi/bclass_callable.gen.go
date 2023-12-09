// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Callable struct {
  iface gdc.Interface
  ptr gdc.TypePtr
}

// Enums

// Constructors

func NewCallable() Callable {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeCallable))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeCallable, 0) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{}))
  return Callable{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewCallableFromCallable(from Callable, ) Callable {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeCallable))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeCallable, 1) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return Callable{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewCallableFromObjectStringName(object Object, method StringName, ) Callable {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeCallable))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeCallable, 2) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{object.AsCTypePtr(), method.AsCTypePtr(), }))
  return Callable{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

// Destructor
func (me *Callable) Destroy() {
  if me.ptr == nil {
    return
  }
  dtr := me.iface.VariantGetPtrDestructor(gdc.VariantTypeCallable)
	me.iface.CallPtrDestructor(dtr, gdc.TypePtr(me.ptr))
	cfree(unsafe.Pointer(me.ptr))
  me.ptr = nil
}

func (me *Callable) Type() gdc.VariantType {
  return gdc.VariantTypeCallable
}

func (me *Callable) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.ptr)
}

func (me *Callable) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.ptr)
}

// Methods

func  (me *Callable) Callv(arguments Array, ) Variant {
  panic("TODO: implement")
}

func  (me *Callable) IsNull() bool {
  panic("TODO: implement")
}

func  (me *Callable) IsCustom() bool {
  panic("TODO: implement")
}

func  (me *Callable) IsStandard() bool {
  panic("TODO: implement")
}

func  (me *Callable) IsValid() bool {
  panic("TODO: implement")
}

func  (me *Callable) GetObject() Object {
  panic("TODO: implement")
}

func  (me *Callable) GetObjectId() int {
  panic("TODO: implement")
}

func  (me *Callable) GetMethod() StringName {
  panic("TODO: implement")
}

func  (me *Callable) GetBoundArgumentsCount() int {
  panic("TODO: implement")
}

func  (me *Callable) GetBoundArguments() Array {
  panic("TODO: implement")
}

func  (me *Callable) Hash() int {
  panic("TODO: implement")
}

func  (me *Callable) Bindv(arguments Array, ) Callable {
  panic("TODO: implement")
}

func  (me *Callable) Unbind(argcount int, ) Callable {
  panic("TODO: implement")
}

func  (me *Callable) Call() Variant {
  panic("TODO: implement")
}

func  (me *Callable) CallDeferred()  {
  panic("TODO: implement")
}

func  (me *Callable) Rpc()  {
  panic("TODO: implement")
}

func  (me *Callable) RpcId(peer_id int, )  {
  panic("TODO: implement")
}

func  (me *Callable) Bind() Callable {
  panic("TODO: implement")
}

// Operators

func (me *Callable) EqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *Callable) NotEqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *Callable) Not() bool {
  panic("TODO: implement")
}

func (me *Callable) EqualsCallable(right Callable) bool {
  panic("TODO: implement")
}

func (me *Callable) NotEqualsCallable(right Callable) bool {
  panic("TODO: implement")
}

func (me *Callable) InDictionary(right Dictionary) bool {
  panic("TODO: implement")
}

func (me *Callable) InArray(right Array) bool {
  panic("TODO: implement")
}

// TODO: members (bclass)
