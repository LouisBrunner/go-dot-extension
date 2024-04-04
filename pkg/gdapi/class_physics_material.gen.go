// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type PhysicsMaterial struct {
  Resource
}

func (me *PhysicsMaterial) BaseClass() string {
  return "PhysicsMaterial"
}

func NewPhysicsMaterial() *PhysicsMaterial {
  str := StringNameFromStr("PhysicsMaterial") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &PhysicsMaterial{}
  obj.SetBaseObject(objPtr)
  return obj
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

func  (me *PhysicsMaterial) SetFriction(friction float64, )  {
  classNameV := StringNameFromStr("PhysicsMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_friction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&friction) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PhysicsMaterial) GetFriction() float64 {
  classNameV := StringNameFromStr("PhysicsMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_friction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PhysicsMaterial) SetRough(rough bool, )  {
  classNameV := StringNameFromStr("PhysicsMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_rough")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&rough) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PhysicsMaterial) IsRough() bool {
  classNameV := StringNameFromStr("PhysicsMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_rough")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PhysicsMaterial) SetBounce(bounce float64, )  {
  classNameV := StringNameFromStr("PhysicsMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bounce")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bounce) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PhysicsMaterial) GetBounce() float64 {
  classNameV := StringNameFromStr("PhysicsMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bounce")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PhysicsMaterial) SetAbsorbent(absorbent bool, )  {
  classNameV := StringNameFromStr("PhysicsMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_absorbent")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&absorbent) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PhysicsMaterial) IsAbsorbent() bool {
  classNameV := StringNameFromStr("PhysicsMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_absorbent")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
