// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VSlider struct {
  obj gdc.ObjectPtr
}

func (me *VSlider) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VSlider) BaseClass() string {
  return "VSlider"
}



// Enums

func (me *VSlider) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VSlider) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

// TODO: properties (class)

// TODO: signals (class)
