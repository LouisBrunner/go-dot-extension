// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Joint2D struct {
  obj gdc.ObjectPtr
}

func (me *Joint2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Joint2D) BaseClass() string {
  return "Joint2D"
}



// Enums

func (me *Joint2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Joint2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *Joint2D) SetNodeA(node NodePath, )  {
  panic("TODO: implement")
}

func  (me *Joint2D) GetNodeA()  {
  panic("TODO: implement")
}

func  (me *Joint2D) SetNodeB(node NodePath, )  {
  panic("TODO: implement")
}

func  (me *Joint2D) GetNodeB()  {
  panic("TODO: implement")
}

func  (me *Joint2D) SetBias(bias float32, )  {
  panic("TODO: implement")
}

func  (me *Joint2D) GetBias()  {
  panic("TODO: implement")
}

func  (me *Joint2D) SetExcludeNodesFromCollision(enable bool, )  {
  panic("TODO: implement")
}

func  (me *Joint2D) GetExcludeNodesFromCollision()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
