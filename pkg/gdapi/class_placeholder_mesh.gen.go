// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PlaceholderMesh struct {
  obj gdc.ObjectPtr
}

func (me *PlaceholderMesh) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PlaceholderMesh) BaseClass() string {
  return "PlaceholderMesh"
}



// Enums

func (me *PlaceholderMesh) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PlaceholderMesh) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PlaceholderMesh) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *PlaceholderMesh) SetAabb(aabb AABB, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
