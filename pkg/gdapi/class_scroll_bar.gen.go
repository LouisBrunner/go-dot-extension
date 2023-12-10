// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type ScrollBar struct {
  obj gdc.ObjectPtr
}

func (me *ScrollBar) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ScrollBar) BaseClass() string {
  return "ScrollBar"
}



// Enums

func (me *ScrollBar) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ScrollBar) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ScrollBar) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *ScrollBar) SetCustomStep(step float32, )  {
  classNameV := StringNameFromStr("ScrollBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_custom_step")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&step), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ScrollBar) GetCustomStep() float32 {
  classNameV := StringNameFromStr("ScrollBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_custom_step")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *ScrollBar) GetPropCustomStep() float32 {
  panic("TODO: implement")
}

func (me *ScrollBar) SetPropCustomStep(value float32) {
  panic("TODO: implement")
}
// Signals
// FIXME: can't seem to be able to connect them from this side of the API