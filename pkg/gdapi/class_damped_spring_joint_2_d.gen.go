// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type DampedSpringJoint2D struct {
  obj gdc.ObjectPtr
}

func (me *DampedSpringJoint2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *DampedSpringJoint2D) BaseClass() string {
  return "DampedSpringJoint2D"
}



// Enums

func (me *DampedSpringJoint2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *DampedSpringJoint2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *DampedSpringJoint2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *DampedSpringJoint2D) SetLength(length float32, )  {
  classNameV := StringNameFromStr("DampedSpringJoint2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&length), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DampedSpringJoint2D) GetLength() float32 {
  classNameV := StringNameFromStr("DampedSpringJoint2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DampedSpringJoint2D) SetRestLength(rest_length float32, )  {
  classNameV := StringNameFromStr("DampedSpringJoint2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_rest_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&rest_length), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DampedSpringJoint2D) GetRestLength() float32 {
  classNameV := StringNameFromStr("DampedSpringJoint2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_rest_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DampedSpringJoint2D) SetStiffness(stiffness float32, )  {
  classNameV := StringNameFromStr("DampedSpringJoint2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_stiffness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&stiffness), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DampedSpringJoint2D) GetStiffness() float32 {
  classNameV := StringNameFromStr("DampedSpringJoint2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_stiffness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DampedSpringJoint2D) SetDamping(damping float32, )  {
  classNameV := StringNameFromStr("DampedSpringJoint2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_damping")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&damping), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DampedSpringJoint2D) GetDamping() float32 {
  classNameV := StringNameFromStr("DampedSpringJoint2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_damping")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *DampedSpringJoint2D) GetPropLength() float32 {
  panic("TODO: implement")
}

func (me *DampedSpringJoint2D) SetPropLength(value float32) {
  panic("TODO: implement")
}

func (me *DampedSpringJoint2D) GetPropRestLength() float32 {
  panic("TODO: implement")
}

func (me *DampedSpringJoint2D) SetPropRestLength(value float32) {
  panic("TODO: implement")
}

func (me *DampedSpringJoint2D) GetPropStiffness() float32 {
  panic("TODO: implement")
}

func (me *DampedSpringJoint2D) SetPropStiffness(value float32) {
  panic("TODO: implement")
}

func (me *DampedSpringJoint2D) GetPropDamping() float32 {
  panic("TODO: implement")
}

func (me *DampedSpringJoint2D) SetPropDamping(value float32) {
  panic("TODO: implement")
}