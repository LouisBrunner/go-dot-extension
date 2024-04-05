// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "log"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type ptrsForRefCountedList struct {
  fnInitRef gdc.MethodBindPtr
  fnReference gdc.MethodBindPtr
  fnUnreference gdc.MethodBindPtr
  fnGetReferenceCount gdc.MethodBindPtr
}

var ptrsForRefCounted ptrsForRefCountedList

func initRefCountedPtrs(iface gdc.Interface) {

  className := StringNameFromStr("RefCounted")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("init_ref")
    defer methodName.Destroy()
    ptrsForRefCounted.fnInitRef = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
  }
  {
    methodName := StringNameFromStr("reference")
    defer methodName.Destroy()
    ptrsForRefCounted.fnReference = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
  }
  {
    methodName := StringNameFromStr("unreference")
    defer methodName.Destroy()
    ptrsForRefCounted.fnUnreference = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
  }
  {
    methodName := StringNameFromStr("get_reference_count")
    defer methodName.Destroy()
    ptrsForRefCounted.fnGetReferenceCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
}

type RefCounted struct {
  Object
}

func (me *RefCounted) BaseClass() string {
  return "RefCounted"
}

func NewRefCounted() *RefCounted {
  str := StringNameFromStr("RefCounted") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &RefCounted{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *RefCounted) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *RefCounted) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *RefCounted) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *RefCounted) InitRef() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRefCounted.fnInitRef), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RefCounted) Reference() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRefCounted.fnReference), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RefCounted) Unreference() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRefCounted.fnUnreference), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RefCounted) GetReferenceCount() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRefCounted.fnGetReferenceCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

// Signals
