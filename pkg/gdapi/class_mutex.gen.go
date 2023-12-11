// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type Mutex struct {
  obj gdc.ObjectPtr
}

func (me *Mutex) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Mutex) BaseClass() string {
  return "Mutex"
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
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
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
