// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Shape3D struct {
  obj gdc.ObjectPtr
}

func (me *Shape3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Shape3D) BaseClass() string {
  return "Shape3D"
}



// Enums

func (me *Shape3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Shape3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Shape3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *Shape3D) SetCustomSolverBias(bias float32, )  {
  panic("TODO: implement")
}

func  (me *Shape3D) GetCustomSolverBias()  {
  panic("TODO: implement")
}

func  (me *Shape3D) SetMargin(margin float32, )  {
  panic("TODO: implement")
}

func  (me *Shape3D) GetMargin()  {
  panic("TODO: implement")
}

func  (me *Shape3D) GetDebugMesh()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
