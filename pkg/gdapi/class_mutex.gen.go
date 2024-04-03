// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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
  classNameV := StringNameFromStr("Mutex")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("lock")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Mutex) TryLock() bool {
  classNameV := StringNameFromStr("Mutex")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("try_lock")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Mutex) Unlock()  {
  classNameV := StringNameFromStr("Mutex")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("unlock")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
