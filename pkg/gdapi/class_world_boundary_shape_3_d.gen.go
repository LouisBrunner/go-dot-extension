// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type WorldBoundaryShape3D struct {
  obj gdc.ObjectPtr
}

func (me *WorldBoundaryShape3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *WorldBoundaryShape3D) BaseClass() string {
  return "WorldBoundaryShape3D"
}



// Enums

func (me *WorldBoundaryShape3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *WorldBoundaryShape3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *WorldBoundaryShape3D) SetPlane(plane Plane, )  {
  panic("TODO: implement")
}

func  (me *WorldBoundaryShape3D) GetPlane()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
