// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type World3D struct {
  obj gdc.ObjectPtr
}

func (me *World3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *World3D) BaseClass() string {
  return "World3D"
}



// Enums

func (me *World3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *World3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *World3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *World3D) GetSpace()  {
  panic("TODO: implement")
}

func  (me *World3D) GetNavigationMap()  {
  panic("TODO: implement")
}

func  (me *World3D) GetScenario()  {
  panic("TODO: implement")
}

func  (me *World3D) SetEnvironment(env Environment, )  {
  panic("TODO: implement")
}

func  (me *World3D) GetEnvironment()  {
  panic("TODO: implement")
}

func  (me *World3D) SetFallbackEnvironment(env Environment, )  {
  panic("TODO: implement")
}

func  (me *World3D) GetFallbackEnvironment()  {
  panic("TODO: implement")
}

func  (me *World3D) SetCameraAttributes(attributes CameraAttributes, )  {
  panic("TODO: implement")
}

func  (me *World3D) GetCameraAttributes()  {
  panic("TODO: implement")
}

func  (me *World3D) GetDirectSpaceState()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
