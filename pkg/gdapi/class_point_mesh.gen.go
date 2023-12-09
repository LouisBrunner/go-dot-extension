// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PointMesh struct {
  obj gdc.ObjectPtr
}

func (me *PointMesh) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PointMesh) BaseClass() string {
  return "PointMesh"
}



// Enums

func (me *PointMesh) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PointMesh) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PointMesh) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

// TODO: properties (class)

// TODO: signals (class)
