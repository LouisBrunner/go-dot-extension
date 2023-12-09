// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ORMMaterial3D struct {
  obj gdc.ObjectPtr
}

func (me *ORMMaterial3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ORMMaterial3D) BaseClass() string {
  return "ORMMaterial3D"
}



// Enums

func (me *ORMMaterial3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ORMMaterial3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

// TODO: properties (class)

// TODO: signals (class)
