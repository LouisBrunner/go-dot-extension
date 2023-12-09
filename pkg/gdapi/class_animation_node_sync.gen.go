// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AnimationNodeSync struct {
  obj gdc.ObjectPtr
}

func (me *AnimationNodeSync) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AnimationNodeSync) BaseClass() string {
  return "AnimationNodeSync"
}



// Enums

func (me *AnimationNodeSync) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AnimationNodeSync) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *AnimationNodeSync) SetUseSync(enable bool, )  {
  panic("TODO: implement")
}

func  (me *AnimationNodeSync) IsUsingSync()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
