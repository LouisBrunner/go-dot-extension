// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type XRCamera3D struct {
  obj gdc.ObjectPtr
}

func (me *XRCamera3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *XRCamera3D) BaseClass() string {
  return "XRCamera3D"
}



// Enums

func (me *XRCamera3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *XRCamera3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

// TODO: properties (class)

// TODO: signals (class)
