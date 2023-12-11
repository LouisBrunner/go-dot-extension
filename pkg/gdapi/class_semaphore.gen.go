// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type Semaphore struct {
  obj gdc.ObjectPtr
}

func (me *Semaphore) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Semaphore) BaseClass() string {
  return "Semaphore"
}



// Enums

func (me *Semaphore) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Semaphore) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Semaphore) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Semaphore) Wait()  {
  classNameV := StringNameFromStr("Semaphore")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("wait")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Semaphore) TryWait() bool {
  classNameV := StringNameFromStr("Semaphore")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("try_wait")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Semaphore) Post()  {
  classNameV := StringNameFromStr("Semaphore")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("post")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Signals
