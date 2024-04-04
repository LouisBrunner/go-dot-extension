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

type VehicleBody3D struct {
  RigidBody3D
}

func (me *VehicleBody3D) BaseClass() string {
  return "VehicleBody3D"
}

func NewVehicleBody3D() *VehicleBody3D {
  str := StringNameFromStr("VehicleBody3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VehicleBody3D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *VehicleBody3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VehicleBody3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VehicleBody3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VehicleBody3D) SetEngineForce(engine_force float64, )  {
  classNameV := StringNameFromStr("VehicleBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_engine_force")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&engine_force) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VehicleBody3D) GetEngineForce() float64 {
  classNameV := StringNameFromStr("VehicleBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_engine_force")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *VehicleBody3D) SetBrake(brake float64, )  {
  classNameV := StringNameFromStr("VehicleBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_brake")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&brake) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VehicleBody3D) GetBrake() float64 {
  classNameV := StringNameFromStr("VehicleBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_brake")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *VehicleBody3D) SetSteering(steering float64, )  {
  classNameV := StringNameFromStr("VehicleBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_steering")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&steering) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VehicleBody3D) GetSteering() float64 {
  classNameV := StringNameFromStr("VehicleBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_steering")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
