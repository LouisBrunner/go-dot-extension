// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type QuadOccluder3D struct {
  obj gdc.ObjectPtr
}

func (me *QuadOccluder3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *QuadOccluder3D) BaseClass() string {
  return "QuadOccluder3D"
}



// Enums

func (me *QuadOccluder3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *QuadOccluder3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *QuadOccluder3D) SetSize(size Vector2, )  {
  panic("TODO: implement")
}

func  (me *QuadOccluder3D) GetSize()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
