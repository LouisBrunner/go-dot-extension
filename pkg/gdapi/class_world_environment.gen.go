// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type WorldEnvironment struct {
  obj gdc.ObjectPtr
}

func (me *WorldEnvironment) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *WorldEnvironment) BaseClass() string {
  return "WorldEnvironment"
}



// Enums

func (me *WorldEnvironment) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *WorldEnvironment) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *WorldEnvironment) SetEnvironment(env Environment, )  {
  panic("TODO: implement")
}

func  (me *WorldEnvironment) GetEnvironment()  {
  panic("TODO: implement")
}

func  (me *WorldEnvironment) SetCameraAttributes(camera_attributes CameraAttributes, )  {
  panic("TODO: implement")
}

func  (me *WorldEnvironment) GetCameraAttributes()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
