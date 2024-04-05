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

type ptrsForPhysicsMaterialList struct {
  fnSetFriction gdc.MethodBindPtr
  fnGetFriction gdc.MethodBindPtr
  fnSetRough gdc.MethodBindPtr
  fnIsRough gdc.MethodBindPtr
  fnSetBounce gdc.MethodBindPtr
  fnGetBounce gdc.MethodBindPtr
  fnSetAbsorbent gdc.MethodBindPtr
  fnIsAbsorbent gdc.MethodBindPtr
}

var ptrsForPhysicsMaterial ptrsForPhysicsMaterialList

func initPhysicsMaterialPtrs(iface gdc.Interface) {

  className := StringNameFromStr("PhysicsMaterial")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_friction")
    defer methodName.Destroy()
    ptrsForPhysicsMaterial.fnSetFriction = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_friction")
    defer methodName.Destroy()
    ptrsForPhysicsMaterial.fnGetFriction = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_rough")
    defer methodName.Destroy()
    ptrsForPhysicsMaterial.fnSetRough = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_rough")
    defer methodName.Destroy()
    ptrsForPhysicsMaterial.fnIsRough = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_bounce")
    defer methodName.Destroy()
    ptrsForPhysicsMaterial.fnSetBounce = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_bounce")
    defer methodName.Destroy()
    ptrsForPhysicsMaterial.fnGetBounce = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_absorbent")
    defer methodName.Destroy()
    ptrsForPhysicsMaterial.fnSetAbsorbent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_absorbent")
    defer methodName.Destroy()
    ptrsForPhysicsMaterial.fnIsAbsorbent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
}

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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&friction) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsMaterial.fnSetFriction), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PhysicsMaterial) GetFriction() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsMaterial.fnGetFriction), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PhysicsMaterial) SetRough(rough bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&rough) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsMaterial.fnSetRough), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PhysicsMaterial) IsRough() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsMaterial.fnIsRough), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PhysicsMaterial) SetBounce(bounce float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bounce) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsMaterial.fnSetBounce), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PhysicsMaterial) GetBounce() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsMaterial.fnGetBounce), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PhysicsMaterial) SetAbsorbent(absorbent bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&absorbent) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsMaterial.fnSetAbsorbent), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PhysicsMaterial) IsAbsorbent() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsMaterial.fnIsAbsorbent), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
