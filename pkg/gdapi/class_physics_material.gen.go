// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type PhysicsMaterial struct {
  obj gdc.ObjectPtr
}

func (me *PhysicsMaterial) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicsMaterial) BaseClass() string {
  return "PhysicsMaterial"
}



// Enums

func (me *PhysicsMaterial) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PhysicsMaterial) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PhysicsMaterial) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *PhysicsMaterial) SetFriction(friction float32, )  {
  classNameV := StringNameFromStr("PhysicsMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_friction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&friction), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsMaterial) GetFriction() float32 {
  classNameV := StringNameFromStr("PhysicsMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_friction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsMaterial) SetRough(rough bool, )  {
  classNameV := StringNameFromStr("PhysicsMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_rough")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&rough), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsMaterial) IsRough() bool {
  classNameV := StringNameFromStr("PhysicsMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_rough")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsMaterial) SetBounce(bounce float32, )  {
  classNameV := StringNameFromStr("PhysicsMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bounce")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bounce), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsMaterial) GetBounce() float32 {
  classNameV := StringNameFromStr("PhysicsMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bounce")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsMaterial) SetAbsorbent(absorbent bool, )  {
  classNameV := StringNameFromStr("PhysicsMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_absorbent")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&absorbent), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsMaterial) IsAbsorbent() bool {
  classNameV := StringNameFromStr("PhysicsMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_absorbent")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *PhysicsMaterial) GetPropFriction() float32 {
  panic("TODO: implement")
}

func (me *PhysicsMaterial) SetPropFriction(value float32) {
  panic("TODO: implement")
}

func (me *PhysicsMaterial) GetPropRough() bool {
  panic("TODO: implement")
}

func (me *PhysicsMaterial) SetPropRough(value bool) {
  panic("TODO: implement")
}

func (me *PhysicsMaterial) GetPropBounce() float32 {
  panic("TODO: implement")
}

func (me *PhysicsMaterial) SetPropBounce(value float32) {
  panic("TODO: implement")
}

func (me *PhysicsMaterial) GetPropAbsorbent() bool {
  panic("TODO: implement")
}

func (me *PhysicsMaterial) SetPropAbsorbent(value bool) {
  panic("TODO: implement")
}