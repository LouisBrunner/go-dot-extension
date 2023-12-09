// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Joint3D struct {
  obj gdc.ObjectPtr
}

func (me *Joint3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Joint3D) BaseClass() string {
  return "Joint3D"
}



// Enums

func (me *Joint3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Joint3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *Joint3D) SetNodeA(node NodePath, )  {
  panic("TODO: implement")
}

func  (me *Joint3D) GetNodeA()  {
  panic("TODO: implement")
}

func  (me *Joint3D) SetNodeB(node NodePath, )  {
  panic("TODO: implement")
}

func  (me *Joint3D) GetNodeB()  {
  panic("TODO: implement")
}

func  (me *Joint3D) SetSolverPriority(priority int, )  {
  panic("TODO: implement")
}

func  (me *Joint3D) GetSolverPriority()  {
  panic("TODO: implement")
}

func  (me *Joint3D) SetExcludeNodesFromCollision(enable bool, )  {
  panic("TODO: implement")
}

func  (me *Joint3D) GetExcludeNodesFromCollision()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
