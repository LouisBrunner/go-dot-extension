// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

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
  panic("TODO: implement")
}

func  (me *VehicleBody3D) GetEngineForce()  {
  panic("TODO: implement")
}

func  (me *VehicleBody3D) SetBrake(brake float32, )  {
  panic("TODO: implement")
}

func  (me *VehicleBody3D) GetBrake()  {
  panic("TODO: implement")
}

func  (me *VehicleBody3D) SetSteering(steering float32, )  {
  panic("TODO: implement")
}

func  (me *VehicleBody3D) GetSteering()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
