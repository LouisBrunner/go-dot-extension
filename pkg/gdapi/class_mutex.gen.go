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

type ptrsForMutexList struct {
  fnLock gdc.MethodBindPtr
  fnTryLock gdc.MethodBindPtr
  fnUnlock gdc.MethodBindPtr
}

var ptrsForMutex ptrsForMutexList

func initMutexPtrs(iface gdc.Interface) {

  className := StringNameFromStr("Mutex")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("lock")
    defer methodName.Destroy()
    ptrsForMutex.fnLock = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("try_lock")
    defer methodName.Destroy()
    ptrsForMutex.fnTryLock = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
  }
  {
    methodName := StringNameFromStr("unlock")
    defer methodName.Destroy()
    ptrsForMutex.fnUnlock = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
}

type Mutex struct {
  RefCounted
}

func (me *Mutex) BaseClass() string {
  return "Mutex"
}

func NewMutex() *Mutex {
  str := StringNameFromStr("Mutex") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &Mutex{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *Mutex) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Mutex) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Mutex) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Mutex) Lock()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMutex.fnLock), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Mutex) TryLock() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMutex.fnTryLock), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Mutex) Unlock()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMutex.fnUnlock), me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
