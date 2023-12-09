// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Node3DGizmo struct {
  obj gdc.ObjectPtr
}

func (me *Node3DGizmo) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Node3DGizmo) BaseClass() string {
  return "Node3DGizmo"
}



// Enums

func (me *Node3DGizmo) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Node3DGizmo) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Node3DGizmo) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

// TODO: properties (class)

// TODO: signals (class)
