// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SceneTreeTimer struct {
  obj gdc.ObjectPtr
}

func (me *SceneTreeTimer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *SceneTreeTimer) BaseClass() string {
  return "SceneTreeTimer"
}



// Enums

func (me *SceneTreeTimer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *SceneTreeTimer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *SceneTreeTimer) SetTimeLeft(time float32, )  {
  panic("TODO: implement")
}

func  (me *SceneTreeTimer) GetTimeLeft()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
