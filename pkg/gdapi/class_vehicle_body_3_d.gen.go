// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VehicleBody3D struct {
  obj gdc.ObjectPtr
}

func (me *VehicleBody3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VehicleBody3D) BaseClass() string {
  return "VehicleBody3D"
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

func  (me *VehicleBody3D) SetEngineForce(engine_force float32, )  {
  classNameV := StringNameFromStr("VehicleBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_engine_force")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&engine_force), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VehicleBody3D) GetEngineForce() float32 {
  classNameV := StringNameFromStr("VehicleBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_engine_force")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VehicleBody3D) SetBrake(brake float32, )  {
  classNameV := StringNameFromStr("VehicleBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_brake")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&brake), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VehicleBody3D) GetBrake() float32 {
  classNameV := StringNameFromStr("VehicleBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_brake")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VehicleBody3D) SetSteering(steering float32, )  {
  classNameV := StringNameFromStr("VehicleBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_steering")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&steering), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VehicleBody3D) GetSteering() float32 {
  classNameV := StringNameFromStr("VehicleBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_steering")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *VehicleBody3D) GetPropEngineForce() float32 {
  panic("TODO: implement")
}

func (me *VehicleBody3D) SetPropEngineForce(value float32) {
  panic("TODO: implement")
}

func (me *VehicleBody3D) GetPropBrake() float32 {
  panic("TODO: implement")
}

func (me *VehicleBody3D) SetPropBrake(value float32) {
  panic("TODO: implement")
}

func (me *VehicleBody3D) GetPropSteering() float32 {
  panic("TODO: implement")
}

func (me *VehicleBody3D) SetPropSteering(value float32) {
  panic("TODO: implement")
}