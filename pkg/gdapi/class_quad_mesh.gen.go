// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type QuadMesh struct {
  obj gdc.ObjectPtr
}

func (me *QuadMesh) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *QuadMesh) BaseClass() string {
  return "QuadMesh"
}



// Enums

func (me *QuadMesh) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *QuadMesh) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *QuadMesh) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

// TODO: properties (class)

// TODO: signals (class)
