// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type StandardMaterial3D struct {
  obj gdc.ObjectPtr
}

func (me *StandardMaterial3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *StandardMaterial3D) BaseClass() string {
  return "StandardMaterial3D"
}



// Enums

func (me *StandardMaterial3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *StandardMaterial3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

// TODO: properties (class)

// TODO: signals (class)
