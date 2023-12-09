// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Path3D struct {
  obj gdc.ObjectPtr
}

func (me *Path3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Path3D) BaseClass() string {
  return "Path3D"
}



// Enums

func (me *Path3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Path3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *Path3D) SetCurve(curve Curve3D, )  {
  panic("TODO: implement")
}

func  (me *Path3D) GetCurve()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
