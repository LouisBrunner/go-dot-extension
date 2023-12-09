// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SpotLight3D struct {
  obj gdc.ObjectPtr
}

func (me *SpotLight3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *SpotLight3D) BaseClass() string {
  return "SpotLight3D"
}



// Enums

func (me *SpotLight3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *SpotLight3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

// TODO: properties (class)

// TODO: signals (class)
