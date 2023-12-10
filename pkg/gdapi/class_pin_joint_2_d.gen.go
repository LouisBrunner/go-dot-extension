// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type PinJoint2D struct {
  obj gdc.ObjectPtr
}

func (me *PinJoint2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PinJoint2D) BaseClass() string {
  return "PinJoint2D"
}



// Enums

func (me *PinJoint2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PinJoint2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PinJoint2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *PinJoint2D) SetSoftness(softness float32, )  {
  classNameV := StringNameFromStr("PinJoint2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_softness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&softness), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PinJoint2D) GetSoftness() float32 {
  classNameV := StringNameFromStr("PinJoint2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_softness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *PinJoint2D) GetPropSoftness() float32 {
  panic("TODO: implement")
}

func (me *PinJoint2D) SetPropSoftness(value float32) {
  panic("TODO: implement")
}