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

type ptrsForVehicleBody3DList struct {
  fnSetEngineForce gdc.MethodBindPtr
  fnGetEngineForce gdc.MethodBindPtr
  fnSetBrake gdc.MethodBindPtr
  fnGetBrake gdc.MethodBindPtr
  fnSetSteering gdc.MethodBindPtr
  fnGetSteering gdc.MethodBindPtr
}

var ptrsForVehicleBody3D ptrsForVehicleBody3DList

func initVehicleBody3DPtrs(iface gdc.Interface) {

  className := StringNameFromStr("VehicleBody3D")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_engine_force")
    defer methodName.Destroy()
    ptrsForVehicleBody3D.fnSetEngineForce = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_engine_force")
    defer methodName.Destroy()
    ptrsForVehicleBody3D.fnGetEngineForce = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_brake")
    defer methodName.Destroy()
    ptrsForVehicleBody3D.fnSetBrake = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_brake")
    defer methodName.Destroy()
    ptrsForVehicleBody3D.fnGetBrake = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_steering")
    defer methodName.Destroy()
    ptrsForVehicleBody3D.fnSetSteering = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_steering")
    defer methodName.Destroy()
    ptrsForVehicleBody3D.fnGetSteering = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
}

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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&engine_force) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVehicleBody3D.fnSetEngineForce), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VehicleBody3D) GetEngineForce() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVehicleBody3D.fnGetEngineForce), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *VehicleBody3D) SetBrake(brake float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&brake) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVehicleBody3D.fnSetBrake), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VehicleBody3D) GetBrake() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVehicleBody3D.fnGetBrake), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *VehicleBody3D) SetSteering(steering float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&steering) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVehicleBody3D.fnSetSteering), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VehicleBody3D) GetSteering() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVehicleBody3D.fnGetSteering), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
